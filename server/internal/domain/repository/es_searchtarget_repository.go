package repository

import (
	"context"
	"errors"
	"log"
	"sort"
	"strings"

	es "github.com/ek-170/loglyzer/internal/infrastructure/elasticsearch"
)

type EsAnalysisRepository struct{}

func NewEsAnalysisRepository() EsAnalysisRepository {
	return EsAnalysisRepository{}
}

// Not allow multiple "q" separated by space
func (est EsAnalysisRepository) FindAnalysiss(q string) ([]*Analysis, error) {
	client, err := es.CreateElasticsearchClient()
	if err != nil {
		return nil, err
	}
	if q == "" {
		// retrieve all aliases
		q = "*"
	} else {
		q = "*" + q + "*"
	}
	res, err := client.Cat.Aliases().Name(q).Do(context.Background())
	if err != nil {
		log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "cat Aliases")
		return nil, errors.New(es.HandleElasticsearchError(err))
	}
	var analyses []*Analysis = []*Analysis{}
	for _, alias := range res {
		// alias start with "." is built-in alias
		if strings.HasPrefix(*alias.Alias, builtInAliasPrefix) {
			continue
		}
		// duplicate aliases may be returned
		if containsAnalysis(analyses, *alias.Alias) {
			continue
		}
		// parsesource info index name
		infoIndexName := "ps_" + *alias.Alias + "_info"
		parseSources, err := SearchParseSources(client, "", infoIndexName)
		if err != nil {
			return nil, err
		}
		if *alias.Alias != "" {
			analysis := &Analysis{
				Id:           *alias.Alias,
				ParseSources: parseSources,
			}
			analyses = append(analyses, analysis)
		}
	}
	return sortAnalysis(analyses, true), nil
}

func containsAnalysis(analyses []*Analysis, new string) bool {
	for _, s := range analyses {
		if s.Id == new {
			return true
		}
	}
	return false
}

// TODO: change "asc" to enum
func sortAnalysis(arr []*Analysis, asc bool) []*Analysis {
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

func (est EsAnalysisRepository) GetAnalysis(id string) (*Analysis, error) {
	client, err := es.CreateElasticsearchClient()
	if err != nil {
		return nil, err
	}
	res, err := client.Cat.Aliases().Name(id).Do(context.Background())
	if err != nil {
		log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "get Aliases")
		return nil, errors.New(es.HandleElasticsearchError(err))
	}
	// parsesource info index name
	infoIndexName := "ps_" + *res[0].Alias + "_info"
	parseSources, err := SearchParseSources(client, "", infoIndexName)
	if err != nil {
		return nil, err
	}
	st := &Analysis{
		Id:           *res[0].Alias,
		ParseSources: parseSources,
	}
	return st, nil
}

func (est EsAnalysisRepository) CreateAnalysis(id string) error {
	err := validateAnalysisId(id)
	if err != nil {
		return err
	}
	client, err := es.CreateElasticsearchClient()
	if err != nil {
		return err
	}
	// to create Alias, also need to exist Index
	// so create placeholder Index which is not used
	indexName := id + "_placeholder"
	_, err = client.Indices.Create(indexName).Do(context.Background())
	if err != nil {
		log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "create placeholder Index")
		return errors.New(es.HandleElasticsearchError(err))
	}
	_, err = client.Indices.PutAlias(indexName, id).Do(context.Background())
	if err != nil {
		log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "create Alias")
		return errors.New(es.HandleElasticsearchError(err))
	}
	psInfoIndexName := "ps_" + id + "_info"
	_, err = client.Indices.Create(psInfoIndexName).
		Mappings(es.BuildParseSourceMapping()).
		Do(context.Background())
	if err != nil {
		log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "create ParseSource Index")
		return errors.New(es.HandleElasticsearchError(err))
	}
	return nil
}

func validateAnalysisId(name string) error {
	if strings.HasPrefix(name, builtInAliasPrefix) {
		return errors.New(es.ES_EM00002)
	}
	return nil
}

func (est EsAnalysisRepository) DeleteAnalysis(id string) error {
	client, err := es.CreateElasticsearchClient()
	if err != nil {
		return err
	}
	// get all Indices name
	res, err := client.Indices.GetAlias().Name(id).Do(context.Background())
	if err != nil {
		log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "get all indices")
		return errors.New(es.HandleElasticsearchError(err))
	}
	// delete all Indices
	// after delete all Indices, automatically Alias has deleted
	for key := range res {
		_, err = client.Indices.Delete(key).Do(context.Background())
		if err != nil {
			log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "delete all indices")
			return errors.New(es.HandleElasticsearchError(err))
		}
	}
	psInfoIndexName := "ps_" + id + "_info"
	_, err = client.Indices.Delete(psInfoIndexName).Do(context.Background())
	if err != nil {
		log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "delete Index")
		return errors.New(es.HandleElasticsearchError(err))
	}

	return nil
}
