package elasticsearch

import (
	"fmt"
	"log"

	"github.com/ek-170/loglyzer/internal/config"
	es "github.com/elastic/go-elasticsearch/v8"
)

func CreateElasticsearchClient() (*es.TypedClient, error) {
  esUrl := fmt.Sprintf("%s://%s:%s",
    config.Config.FullTextSearch.Schme,
    config.Config.FullTextSearch.Host,
    config.Config.FullTextSearch.Port,
  )

  cfg := es.Config{
		Addresses: []string{
			esUrl,
		},
	}
  es, err := es.NewTypedClient(cfg)
  if err != nil {
    log.Printf("Failed to create Elasticsearch Client.")
    return nil, err
  }
  return es, err
}
