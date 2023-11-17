package repository

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/ek-170/loglyzer/internal/config"
	es "github.com/ek-170/loglyzer/internal/infrastructure/elasticsearch"
	"github.com/ek-170/loglyzer/internal/util"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	// "github.com/vjeantet/grok"
)

type EsParseSourceRepository struct {}

func NewEsParseSourceRepository() EsParseSourceRepository{
  return EsParseSourceRepository{}
}

func (eg EsParseSourceRepository) FindParseSources(q string) ([]*ParseSource, error){
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
    if strings.HasPrefix(*alias.Alias, builtInAliasPrefix){
      continue
    }
    if *alias.Alias != "" {
      parseSource := &ParseSource{
        Name:  *alias.Alias,
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

func (eg EsParseSourceRepository) GetParseSource(name string) (*ParseSource, error){
  client, err := es.CreateElasticsearchClient()
  if err != nil {
    return nil, err
  }
  res, err := client.Cat.Aliases().Name(name).Do(context.Background())
  if err != nil {
    log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "get Aliases")
    return nil, errors.New(es.HandleElasticsearchError(err))
  }
  var st *ParseSource = &ParseSource{}
  st = &ParseSource{
    Name: *res[0].Alias,
    // add parseSource
  }
  return st, nil
}

type LogEntry struct {
  Line int    `json:"line"`
  Log  string `json:"log"`
}

func (eg EsParseSourceRepository) CreateParseSource(
  searchTarget string, multiLine bool, fileName string, grokId string) error {
  searchedFile, err := util.SearchFile(config.Config.Server.LogDir, fileName)
  if err != nil {
    return err
  }
  fullPath := config.Config.Server.LogDir+"/"+searchedFile
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

  psIndex := searchTarget+"_parsesource"
  // Get all already existing ParseSource in descending order of Order
  searchRes, err := client.Search().Index(psIndex).Fields(es.BuildParseSourceFields()).Sort(es.BuildParseSourceSort()).Do(context.Background())
  if err != nil {
    log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "search Docs")
    return errors.New(es.HandleElasticsearchError(err))
  }
  hitSize := len(searchRes.Hits.Hits)

  var psName string
  var order int16
  if hitSize == 0 {
    order = 1
    psName = "ps_" + searchTarget + "_" + strconv.Itoa(int(order))
  } else {
    source := searchRes.Hits.Hits[0].Source_
    headPs := ParseSource{}
    json.Unmarshal(source, &headPs)
    order = int16(headPs.Order) + 1
    psName = "ps_" + searchTarget + "_" + strconv.Itoa(int(order))
  }

  ps := ParseSource{
    Name: fileName,
    Index: psName,
    Order: order,
  }

  _, err = client.Index(psIndex).Document(ps).Do(context.Background())
  if err != nil {
    log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "create Doc")
    return errors.New(es.HandleElasticsearchError(err))
  }

  // ファイルを読み込み、ログをElasticsearchに送信
  scanner := bufio.NewScanner(file)
  var lineNumber int
  // var multiLineLog []string
  type Log struct {
    LineNum int    `json:"lineNumber"`
    Message string `json:"message"`
  }

  logs := []Log{}
  for scanner.Scan() {
    line := scanner.Text()
    lineNumber++

    if strings.TrimSpace(line) == "" {
        // ignore brank line
        continue
    }

    // multi line logの判定をGrokライブラリで実施
    // multi lineならcontinue
    // log本文を取得し終えたらドキュメント生成してBulkへ（ここをチャネル使いたい）


    log := Log{LineNum: lineNumber, Message: line}
    logs = append(logs, log)
    // 処理にあたっては先にタスクIDだけ返却し、フロントはそれを元に処理の進捗をポーリングする予定

  }
  op := types.NewIndexOperation()
  err = client.Bulk().Index(psName).Pipeline(grokId).IndexOp(*op, logs)
  if err != nil {
    log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "Bulk Index")
    return errors.New(es.HandleElasticsearchError(err))
  }

  if err := scanner.Err(); err != nil {
    return err
  }

  return nil
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