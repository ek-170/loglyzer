package repository

import (
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
  res, err := client.Cat.Aliases().Do(context.TODO())
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
  res, err := client.Cat.Aliases().Name(name).Do(context.TODO())
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
  searchTarget string, parseSource string, multiLine bool, fileName string, grokId string) error {
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
  res, err := client.Search().Index(psIndex).Fields(es.BuildParseSourceFields()).Sort(es.BuildParseSourceSort()).Do(context.TODO())
  if err != nil {
    log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "search Docs")
    return errors.New(es.HandleElasticsearchError(err))
  }
  hitSize := len(res.Hits.Hits)

  var psName string
  var order int16
  if hitSize == 0 {
    order = 1
    psName = "ps_" + searchTarget + "_" + strconv.Itoa(int(order))
  } else {
    // TODO ここでJSONの扱いから
    head := res.Hits.Hits[0]
    source := head.Source_
    ps := ParseSource{}
    json.Unmarshal(source, order)
    psName = "ps_" + searchTarget + "_" + strconv.Itoa(int(order))
  }
  // parsesource index -> "ps_"+{searchSourceName}+"_"+連番
  // ParseSource.Indexが重複していないか確認&末尾の連番を確認
  ps := ParseSource{
    Name: fileName,
    Index: psName,
    Order: order,
  }
  // ParseSourceInfo作成(ここに引数のparseSource)
  res2, _ := client.Index(psIndex).Document(ps).Do(context.TODO())
  if res2 != nil {
    log.Print("bbb")
  }

  // ファイルを読み込み、ログをElasticsearchに送信
  // scanner := bufio.NewScanner(file)
  // var lineNumber int
  // // var multiLineLog []string

  // for scanner.Scan() {
  //   line := scanner.Text()
  //   lineNumber++

  //   if strings.TrimSpace(line) == "" {
  //       // ignore brank line
  //       continue
  //   }

  //   // multi line logの判定をGrokライブラリで実施
  //   // multi lineならcontinue
  //   // log本文を取得し終えたらドキュメント生成してBulkへ（ここをチャネル使いたい）
  //   // 処理にあたっては先にタスクIDだけ返却し、フロントはそれを元に処理の進捗をポーリングする予定
  //   client.Bulk().Index(searchTarget).Do(context.TODO())

  // }


  // if err := scanner.Err(); err != nil {
  //   return err
  // }

  return nil
}

func (eg EsParseSourceRepository) DeleteParseSource(name string) error {
  client, err := es.CreateElasticsearchClient()
  if err != nil {
    return err
  }
  // get all Indices name
  res, err := client.Indices.GetAlias().Name(name).Do(context.TODO())
  if err != nil {
    log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "get all indices")
    return errors.New(es.HandleElasticsearchError(err))
  }
  // delete all Indices
  for key := range res {
    _, err = client.Indices.Delete(key).Do(context.TODO())
    if err != nil {
      log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "delete all indices")
      return errors.New(es.HandleElasticsearchError(err))
    }
  }
  // after delete all Indices, automatically Alias has deleted
  return nil
}