package repository

import (
	"context"
	"log"
	"sort"
  "errors"

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
  res, err := client.Ingest.GetPipeline().Do(context.TODO())
  if err != nil {
    log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "GET Pipelines")
    return nil, errors.New(es.HandleElasticsearchError(err))
  }
  var grokPatterns []*GrokPattern = []*GrokPattern{}
  for name, pipeline := range res {
    // if q is empty, return all grok patterns
    if(q != "" && q != name){
      continue
    }
    for _, processor := range pipeline.Processors {
      if processor.Grok != nil {
        grokPattern := &GrokPattern{
          Name: name,
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