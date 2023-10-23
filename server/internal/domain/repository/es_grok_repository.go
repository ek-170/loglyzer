package repository

import (
	"context"
	"log"

  "github.com/ek-170/loglyzer/internal/consts"
	"github.com/ek-170/loglyzer/internal/infrastructure/elasticsearch"
)

type EsGrokRepository struct {}

func NewEsGrokRepository() EsGrokRepository{
  return EsGrokRepository{}
}

func (eg EsGrokRepository) GetGrokPatterns(q string) ([]*GrokPattern, error){
  client, err := elasticsearch.CreateElasticsearchClient()
  if err != nil {
    return nil, err
  }
  res, err := client.Ingest.GetPipeline().Do(context.TODO())
  if err != nil {
    log.Printf(consts.FAIL_REQUEST_ELASTIC_SEARCH, "GET Pipelines")
    return nil, err
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
  log.Printf("air test...")
  return grokPatterns, nil
}