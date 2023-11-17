package repository

import (
	"context"
	"errors"
	"log"
	"sort"

	es "github.com/ek-170/loglyzer/internal/infrastructure/elasticsearch"
)

type EsGrokRepository struct {}

func NewEsGrokRepository() EsGrokRepository{
  return EsGrokRepository{}
}

func (eg EsGrokRepository) FindGrokPatterns(q string) ([]*GrokPattern, error){
  client, err := es.CreateElasticsearchClient()
  if err != nil {
    return nil, err
  }
  if q == "" {
    // retrieve all aliases
    q = "*"
  }else {
    q = "*" + q + "*"
  }
  res, err := client.Ingest.GetPipeline().Id(q).Do(context.Background())
  if err != nil {
    log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "GET Pipelines")
    return nil, errors.New(es.HandleElasticsearchError(err))
  }
  var grokPatterns []*GrokPattern = []*GrokPattern{}
  for id, pipeline := range res {
    for _, processor := range pipeline.Processors {
      if processor.Grok != nil {
        grokPattern := &GrokPattern{
          Id: id,
          Pattern: processor.Grok.Patterns[0],
        }
        grokPatterns = append(grokPatterns, grokPattern)
      }
    }
  }
  return sortGrokPatterns(grokPatterns, true), nil
}

// TODO: change "asc" to enum
func sortGrokPatterns(arr []*GrokPattern, asc bool) ([]*GrokPattern){
	if len(arr) < 1 {
		return arr
	}
  if asc {
    sort.Slice(arr, func(i, j int) bool {
      return arr[i].Id < arr[j].Id
    })
  } else {
    // desc
    sort.Slice(arr, func(i, j int) bool {
      return arr[i].Id > arr[j].Id
    })
  }
	return arr
}