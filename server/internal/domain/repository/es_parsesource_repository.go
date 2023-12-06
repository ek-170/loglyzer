package repository

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ek-170/loglyzer/internal/config"
	es "github.com/ek-170/loglyzer/internal/infrastructure/elasticsearch"
	"github.com/ek-170/loglyzer/internal/util"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/vjeantet/grok"
)

type EsParseSourceRepository struct{}

func NewEsParseSourceRepository() EsParseSourceRepository {
	return EsParseSourceRepository{}
}

// Not allow multiple "q" separated by space
func (ep EsParseSourceRepository) FindParseSources(q string, searchTarget string) ([]*ParseSource, error) {
	client, err := es.CreateElasticsearchClient()
	if err != nil {
		return nil, err
	}
	infoIndexName := "ps_" + searchTarget + "_info"
	parseSources, err := SearchParseSources(client, q, infoIndexName)
	if err != nil {
		return nil, err
	}

	return parseSources, nil
}

func SearchParseSources(client *elasticsearch.TypedClient, q string, infoIndexName string) ([]*ParseSource, error) {
	var s *search.Search
	// if q is empty, return all parseSources
	if q != "" {
		v := "*" + q + "*"
		wq := es.BuildParseSourceWildcardQuery(v, "name")
		q := types.NewQuery()
		q.Wildcard = wq
		s = client.Search().Index(infoIndexName).Query(q).Fields(es.BuildParseSourceFields()).Sort(es.BuildParseSourceSort("name"))
	}else {
		s = client.Search().Index(infoIndexName).Fields(es.BuildParseSourceFields()).Sort(es.BuildParseSourceSort("name"))
	}

	res, err := s.Do(context.Background())
	if err != nil {
		log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "search Docs")
		return nil, errors.New(es.HandleElasticsearchError(err))
	}
	var parseSources []*ParseSource = []*ParseSource{}
	
	for _, h := range res.Hits.Hits {
		s := h.Source_
		p := &ParseSource{}
		json.Unmarshal(s, p)
		p.Id = h.Id_
		parseSources = append(parseSources, p)
	}

	return parseSources, nil
}

type Log struct {
	LineNum int    `json:"lineNumber"`
	File string    `json:"file"`
	Message string `json:"message"`
}

func (ep EsParseSourceRepository) CreateParseSource(
	searchTarget string, multiLine bool, fileName string, grokId string) error {
	
	startTime := time.Now()

	searchedFile, err := util.SearchFile(config.Config.Server.LogDir, fileName)
	if err != nil {
		return err
	}
	fullPath := config.Config.Server.LogDir + "/" + searchedFile
	file, err := os.Open(fullPath)
	if err != nil {
		log.Println("failed to open file")
		return err
	}
	defer file.Close()

	client, err := es.CreateElasticsearchClient()
	if err != nil {
		return err
	}

	psInfoIndexName := "ps_" + searchTarget + "_info"
	// Get all already existing ParseSource in descending order of Order
	searchRes, err := client.Search().Index(psInfoIndexName).Fields(es.BuildParseSourceFields()).Sort(es.BuildParseSourceSort("order")).Do(context.Background())
	if err != nil {
		log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "search Docs")
		return errors.New(es.HandleElasticsearchError(err))
	}
	hitSize := len(searchRes.Hits.Hits)

	var psIndexName string
	var order int16
	if hitSize == 0 {
		order = 1
		psIndexName = "ps_" + searchTarget + "_" + strconv.Itoa(int(order))
	} else {
		source := searchRes.Hits.Hits[0].Source_
		headPs := ParseSource{}
		json.Unmarshal(source, &headPs)
		order = int16(headPs.Order) + 1
		psIndexName = "ps_" + searchTarget + "_" + strconv.Itoa(int(order))
	}

	ps := ParseSource{
		Name:  fileName,
		Index: psIndexName,
		Order: order,
	}

	// indexing new ParseSource information
	_, err = client.Index(psInfoIndexName).Document(ps).Do(context.Background())
	if err != nil {
		log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "create Doc")
		return errors.New(es.HandleElasticsearchError(err))
	}

  _, err = client.Indices.Create(psIndexName).Do(context.Background())
  if err != nil {
    log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "create Index")
    return errors.New(es.HandleElasticsearchError(err))
  }

  _, err = client.Indices.PutAlias(psIndexName, searchTarget).Do(context.Background())
  if err != nil {
    log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "create Alias")
    return errors.New(es.HandleElasticsearchError(err))
  }

	if multiLine {
		err = indexMultiLineLog(client, grokId, psIndexName, file)
	} else {
		err = indexLog(client, grokId, psIndexName, file)
	}
	
	if err != nil {
		return err
	}

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

	log.Printf("Time of Creating ParseSource is: %s\n", elapsedTime)

	return nil
}

func indexLog(client *elasticsearch.TypedClient, grokId string, psIndexName string, file *os.File) error {
	var lineNumber int
	
	done := make(chan struct{})
	log.Printf("Bulk unit is: %d", config.Config.FullTextSearch.BulkUnit)
	parsedLog := make(chan Log, config.Config.FullTextSearch.BulkUnit)
	defer close(parsedLog)
	
	wg := &sync.WaitGroup{}
	log.Printf("Number of worker: %d", config.Config.Parser.Worker)
	// set up bulk request worker
	for i := 0; i < config.Config.Parser.Worker; i++ {
		wg.Add(1)
		go sendBulkRequest(client, done, parsedLog, psIndexName, grokId, wg)
	}
	
	scanner := bufio.NewScanner(file)
	log.Println("Start scannnig log file")
  for scanner.Scan() {
		line := scanner.Text()
		lineNumber++

		if strings.TrimSpace(line) == "" {
			// ignore brank line
			continue
		}
		l := Log{LineNum: lineNumber, Message: line, File: filepath.Base(file.Name())}
		parsedLog <- l
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	close(done)
	wg.Wait()

	return nil
}

func indexMultiLineLog(client *elasticsearch.TypedClient, grokId string, psIndexName string, file *os.File) error {
	var g *grok.Grok
  var grokPattern string
	log.Println("Fetch Grok Pattern from Elasticsearch")
	res, err := client.Ingest.GetPipeline().Id(grokId).Do(context.Background())
	if err != nil {
		log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "GET Pipelines")
		return errors.New(es.HandleElasticsearchError(err))
	}
	grokPatterns := ExtractGrokPatterns(res)
	if len(grokPatterns) != 1 {
		log.Printf("Grok Pattern has not specified")
		return errors.New("pecified Grok Pattern is invalid")
	}
	grokPattern = grokPatterns[0].Pattern
	log.Println("Initialize Grok parser")
	conf := &grok.Config{PatternsDir: []string{filepath.Join(config.Config.Path.Base, config.Config.Path.Patterns)}}
	// Cosidering of duplicate additional patterns key, always generate New grok parser.
	g, err = grok.NewWithConfig(conf)

	if err != nil {
		log.Printf("Failed to init Grok parser")
		return err
	}
	if len(grokPatterns[0].PatternDefs) > 0 {
		for k, v := range grokPatterns[0].PatternDefs {
			g.AddPattern(k, v)
		}
	}
	log.Println("Initialize Grok parser has done")

	var lineNumber int = 1
	// if multi line, this slice's len() is over two
	var multiLineLog []string
	
	log.Printf("Bulk unit is: %d", config.Config.FullTextSearch.BulkUnit)
	done := make(chan struct{})
	parsedLog := make(chan Log, config.Config.FullTextSearch.BulkUnit)
	defer close(parsedLog)
	
	wg := &sync.WaitGroup{}
	log.Printf("Number of multiline worker: %d", config.Config.Parser.MultilineWorker)
	// set up bulk request worker
	for i := 0; i < config.Config.Parser.MultilineWorker; i++ {
		wg.Add(1)
		go sendBulkRequest(client, done, parsedLog, psIndexName, grokId, wg)
	}
	
	count := 0
	scanner := bufio.NewScanner(file)
	log.Println("Start scannnig log file")
	for scanner.Scan() {
		line := scanner.Text()
		count++

		if strings.TrimSpace(line) == "" {
			// ignore brank line
			continue
		}

		isMatched, err := g.Match(grokPattern, line)
		if err != nil {
			log.Println("couldn't check whether or not multi line. due to grok library error")
		}
		if isMatched {
			if len(multiLineLog) == 1 {
				// send previous log to bulk request
				l := Log{LineNum: lineNumber, Message: multiLineLog[0], File: filepath.Base(file.Name())}
				parsedLog <- l
				// initialize multiLineLog
				multiLineLog = []string{}
				lineNumber = count
			} else if len(multiLineLog) >= 2 {
				// send previous log to bulk request
				l := Log{LineNum: lineNumber, Message: strings.Join(multiLineLog, newLineCode), File: filepath.Base(file.Name())}
				parsedLog <- l
				// initialize multiLineLog
				multiLineLog = []string{}
				lineNumber = count
			}
			multiLineLog = append(multiLineLog, line)
			continue
		}
		// not match grok, treat as multi line log by adding to previous log
		multiLineLog = append(multiLineLog, line)
	}

	// last log is left
	if len(multiLineLog) == 1 {
		l := Log{LineNum: lineNumber, Message: multiLineLog[0], File: filepath.Base(file.Name())}
		parsedLog <- l
	} else if len(multiLineLog) >= 2 {
		l := Log{LineNum: lineNumber, Message: strings.Join(multiLineLog, newLineCode), File: filepath.Base(file.Name())}
		parsedLog <- l
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	close(done)
	wg.Wait()

	return nil
}

// TODO error handling
func sendBulkRequest(client *elasticsearch.TypedClient, done chan struct{}, parsedLog <-chan Log, psIndexName string, grokId string, wg *sync.WaitGroup) {
	for {
		bulk := client.Bulk()
		bulkSize := 0
		BULK:
		for {
			select {
			case l, ok := <-parsedLog:
				if !ok {
					log.Println("<-chan parsedLog is closed")
					return
				}
				op := types.NewIndexOperation()
				bulk.IndexOp(*op, l)
				bulkSize++
				if(bulkSize == config.Config.FullTextSearch.BulkUnit){
					_, err := bulk.Index(psIndexName).Pipeline(grokId).Do(context.Background())
					if err != nil {
						log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "Bulk Index")
						log.Println(es.HandleElasticsearchError(err))
					}
					break BULK
				}
			case <-done:
				// At this point, sending Log to parseLog is completed
				log.Println("Finish scanning log file")
				log.Printf("Remaining parsed log is: %d", len(parsedLog))
				existsRemainingLog := false
				if len(parsedLog) > 0 {
					existsRemainingLog = true
					for l := range parsedLog {
						op := types.NewIndexOperation()
						bulk.IndexOp(*op, l)
						if len(parsedLog) == 0{
							break
						} 
					}
				}
				if existsRemainingLog || bulkSize > 0 {
					log.Println("Excute bulk request of remaining logs")
					_, err := bulk.Index(psIndexName).Pipeline(grokId).Do(context.Background())
					if err != nil {
						log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "Bulk Index")
						log.Println(es.HandleElasticsearchError(err))
						wg.Done()
						return
					}
				}
				wg.Done()
				return
			}
		}
	}
}

func (ep EsParseSourceRepository) DeleteParseSource(id string, searchTarget string) error {
	client, err := es.CreateElasticsearchClient()
	if err != nil {
		return err
	}
	psInfoIndexName := "ps_" + searchTarget + "_info"
	// get delete target parsesource info
	res, err := client.Get(psInfoIndexName, id).Source_("true").Do(context.Background())
  if err != nil {
    log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "get Doc")
    return errors.New(es.HandleElasticsearchError(err))
  }
	s := res.Source_
	p := &ParseSource{}
	json.Unmarshal(s, p)

	// delete parsesource
	_, err = client.Indices.Delete(p.Index).Do(context.Background())
	if err != nil {
    log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "delete Index")
    return errors.New(es.HandleElasticsearchError(err))
  }

	// delete parsesource info
	_, err = client.Delete(psInfoIndexName, id).Do(context.Background())
	if err != nil {
    log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "delete Doc")
    return errors.New(es.HandleElasticsearchError(err))
  }

	return nil
}
