package repository

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/ek-170/loglyzer/internal/config"
	es "github.com/ek-170/loglyzer/internal/infrastructure/elasticsearch"
	"github.com/ek-170/loglyzer/internal/util"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/bulk"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/vjeantet/grok"
)

type EsParseSourceRepository struct{}

func NewEsParseSourceRepository() EsParseSourceRepository {
	return EsParseSourceRepository{}
}

func (eg EsParseSourceRepository) FindParseSources(q string) ([]*ParseSource, error) {
	client, err := es.CreateElasticsearchClient()
	if err != nil {
		return nil, err
	}
	res, err := client.Cat.Aliases().Do(context.Background())
	if err != nil {
		log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "cat Aliases")
		return nil, errors.New(es.HandleElasticsearchError(err))
	}
	var parseSources []*ParseSource = []*ParseSource{}
	for _, alias := range res {
		// if q is empty, return all parseSources
		if q != "" && q != *alias.Alias {
			continue
		}
		// alias start with "." is built-in alias
		if strings.HasPrefix(*alias.Alias, builtInAliasPrefix) {
			continue
		}
		if *alias.Alias != "" {
			parseSource := &ParseSource{
				Name: *alias.Alias,
				// parseSource
			}
			parseSources = append(parseSources, parseSource)
		}
	}
	return sortParseSource(parseSources, true), nil
}

// TODO: change "asc" to enum
func sortParseSource(arr []*ParseSource, asc bool) []*ParseSource {
	if len(arr) < 1 {
		return arr
	}
	if asc {
		sort.Slice(arr, func(i, j int) bool {
			return arr[i].Name < arr[j].Name
		})
	} else {
		// desc
		sort.Slice(arr, func(i, j int) bool {
			return arr[i].Name > arr[j].Name
		})
	}
	return arr
}

func (eg EsParseSourceRepository) GetParseSource(id string) (*ParseSource, error) {
	client, err := es.CreateElasticsearchClient()
	if err != nil {
		return nil, err
	}
	// res, err := client.Cat.Aliases().Name(id).Do(context.Background())
	// if err != nil {
	// 	log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "get Aliases")
	// 	return nil, errors.New(es.HandleElasticsearchError(err))
	// }
	// var st *ParseSource = &ParseSource{}
	// st = &ParseSource{
	// 	Name: *res[0].Alias,
	// 	// add parseSource
	// }
	// var grokPattern string
	res, err := client.Ingest.GetPipeline().Id("grok_tomcat1").Do(context.Background())
	if err != nil {
		log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "GET Pipelines")
		return nil, nil
	}
	grokPatterns := ExtractGrokPatterns(res)
	if len(grokPatterns) != 1 {
		log.Printf("Grok Pattern has not specified")
		return nil, nil
	}
	// grokPattern = grokPatterns[0].Pattern
	// text := "2023-07-26 02:10:14.335  WARN [ajp-bio-8009-exec-32] - First Multiline Log"
	text := "2014-01-09 20:03:28.269"
	conf := &grok.Config{PatternsDir: []string{filepath.Join(config.Config.Path.Base, config.Config.Path.Patterns)}}
	g, err := grok.NewWithConfig(conf)
	if err != nil {
		return nil, err
	}
	if len(grokPatterns[0].PatternDefs) > 0 {
		for k, v := range grokPatterns[0].PatternDefs {
			g.AddPattern(k, v)
		}
	}
	// pattern := "%{TOMCAT_DATESTAMP:timestamp}  %{LOGLEVEL:log.level} \\[%{DATA:java.log.origin.thread.name}\\] - %{MESSAGE:message}"
	pattern := "%{TOMCAT_DATESTAMP:timestamp}"
	val, err := g.Parse(pattern, text)
	if len(val) > 0 {
		for k, v := range val {
			log.Println(k +": "+v)
		}
	}
if err != nil {
	log.Print("test")
}

	return nil, nil
}

type Log struct {
	LineNum int    `json:"lineNumber"`
	File string    `json:"file"`
	Message string `json:"message"`
}

func (eg EsParseSourceRepository) CreateParseSource(
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
	searchRes, err := client.Search().Index(psInfoIndexName).Fields(es.BuildParseSourceFields()).Sort(es.BuildParseSourceSort()).Do(context.Background())
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
		err = parseMultiLineLog(client, grokId, psIndexName, file)
	} else {
		err = parseLog(client, grokId, psIndexName, file)
	}
	
	if err != nil {
		return err
	}

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

	log.Printf("Time of Creating ParseSource is: %s\n", elapsedTime)

	return nil
}

func parseLog(client *elasticsearch.TypedClient, grokId string, psIndexName string, file *os.File) error {
	scanner := bufio.NewScanner(file)
	var lineNumber int

  bulk := client.Bulk()
  for scanner.Scan() {
		line := scanner.Text()
		lineNumber++

		if strings.TrimSpace(line) == "" {
			// ignore brank line
			continue
		}
		addBulk(bulk, lineNumber, line, file.Name())
	}

  _, err := bulk.Index(psIndexName).Pipeline(grokId).Do(context.Background())
	if err != nil {
		log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "Bulk Index")
		return errors.New(es.HandleElasticsearchError(err))
	}
	// if res != nil {
	// 	for _, i := range res.Items {
	// 		for _, v := range i {
	// 			log.Print(*v.Result)
	// 			log.Print(v.Status)
	// 			if v.Error != nil {
	// 				log.Print(v.Error.Reason)
	// 				log.Print(v.Error.StackTrace)
	// 			}
	// 		}
	// 	}
	// }

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func parseMultiLineLog(client *elasticsearch.TypedClient, grokId string, psIndexName string, file *os.File) error {
	var g *grok.Grok
  var grokPattern string
	res, err := client.Ingest.GetPipeline().Id(grokId).Do(context.Background())
	if err != nil {
		log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "GET Pipelines")
		return errors.New(es.HandleElasticsearchError(err))
	}
	grokPatterns := ExtractGrokPatterns(res)
	if len(grokPatterns) != 1 {
		log.Printf("Grok Pattern has not specified")
		return errors.New("specified Grok Pattern is invalid")
	}
	grokPattern = grokPatterns[0].Pattern
	conf := &grok.Config{PatternsDir: []string{filepath.Join(config.Config.Path.Base, config.Config.Path.Patterns)}}
	// Cosidering of duplicate additional patterns key, always generate New grok parser.
	g, err = grok.NewWithConfig(conf)

	if err != nil {
		log.Printf("failed to init grok library")
		return err
	}
	if len(grokPatterns[0].PatternDefs) > 0 {
		for k, v := range grokPatterns[0].PatternDefs {
			g.AddPattern(k, v)
		}
	}

	scanner := bufio.NewScanner(file)
	var lineNumber int = 1
	// if multi line, this slice's len() is over two
	var multiLineLog []string

  bulk := client.Bulk()
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		count++

		if strings.TrimSpace(line) == "" {
			// ignore brank line
			continue
		}

		isMatched, err := g.Match(grokPattern, line)
		if err != nil {
			log.Printf("couldn't check whether or not multi line. due to grok library error")
		}
		if isMatched {
			if len(multiLineLog) == 1 {
				// send previous log to bulk request
				addBulk(bulk, lineNumber, multiLineLog[0], file.Name())
				// initialize multiLineLog
				multiLineLog = []string{}
				lineNumber = count
			} else if len(multiLineLog) >= 2 {
				// send previous log to bulk request
				addBulk(bulk, lineNumber, strings.Join(multiLineLog, newLineCode), file.Name())
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
		addBulk(bulk, lineNumber, multiLineLog[0], file.Name())
	} else if len(multiLineLog) >= 2 {
		addBulk(bulk, lineNumber, strings.Join(multiLineLog, newLineCode), file.Name())
	}

  _, err = bulk.Index(psIndexName).Pipeline(grokId).Do(context.Background())
	if err != nil {
		log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "Bulk Index")
		return errors.New(es.HandleElasticsearchError(err))
	}
	
	// if res != nil {
	// 	for _, i := range res.Items {
	// 		for _, v := range i {
	// 			log.Print(*v.Result)
	// 			log.Print(v.Status)
	// 			if v.Error != nil {
	// 				log.Print(v.Error.CausedBy.RootCause)
	// 				log.Print(v.Error.StackTrace)
	// 				log.Print(v.Error.Reason)
	// 			}
	// 		}
	// 	}
	// }

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func addBulk(bulk *bulk.Bulk, lineNumber int, message string, file string) {
	log := Log{LineNum: lineNumber, Message: message, File: file}
	op := types.NewIndexOperation()
	bulk.IndexOp(*op, log)
}

func (eg EsParseSourceRepository) DeleteParseSource(name string) error {
	client, err := es.CreateElasticsearchClient()
	if err != nil {
		return err
	}
	// get all Indices name
	res, err := client.Indices.GetAlias().Name(name).Do(context.Background())
	if err != nil {
		log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "get all indices")
		return errors.New(es.HandleElasticsearchError(err))
	}
	// delete all Indices
	for key := range res {
		_, err = client.Indices.Delete(key).Do(context.Background())
		if err != nil {
			log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "delete all indices")
			return errors.New(es.HandleElasticsearchError(err))
		}
	}
	// after delete all Indices, automatically Alias has deleted
	return nil
}
