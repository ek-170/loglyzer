package repository

import (
	"context"
	"errors"
	"log"
	"sort"
	"strings"

	es "github.com/ek-170/loglyzer/internal/infrastructure/elasticsearch"
)

type EsSearchTargetRepository struct {}

func NewEsSearchTargetRepository() EsSearchTargetRepository{
  return EsSearchTargetRepository{}
}

// Not allow multiple "q" separated by space
func (eg EsSearchTargetRepository) FindSearchTargets(q string) ([]*SearchTarget, error){
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
  res, err := client.Cat.Aliases().Name(q).Do(context.Background())
  if err != nil {
    log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "cat Aliases")
    return nil, errors.New(es.HandleElasticsearchError(err))
  }
  var searchTargets []*SearchTarget = []*SearchTarget{}
  for _, alias := range res {
    // alias start with "." is built-in alias
    if strings.HasPrefix(*alias.Alias, builtInAliasPrefix){
      continue
    }
    // duplicate aliases may be returned
    if containsSearchTarget(searchTargets, *alias.Alias){
      continue
    }
    if *alias.Alias != "" {
      searchTarget := &SearchTarget{
        Id:  *alias.Alias,
        // parseSource
      }
      searchTargets = append(searchTargets, searchTarget)
    }
  }
  return sortSearchTarget(searchTargets, true), nil
}

func containsSearchTarget(searchTargets []*SearchTarget, new string) bool {
  for _, s := range searchTargets {
    if s.Id == new {
      return true
    }
  }
  return false
}

// TODO: change "asc" to enum
func sortSearchTarget(arr []*SearchTarget, asc bool) []*SearchTarget {
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

func (eg EsSearchTargetRepository) GetSearchTarget(id string) (*SearchTarget, error){
  client, err := es.CreateElasticsearchClient()
  if err != nil {
    return nil, err
  }
  res, err := client.Cat.Aliases().Name(id).Do(context.Background())
  if err != nil {
    log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "get Aliases")
    return nil, errors.New(es.HandleElasticsearchError(err))
  }
  var st *SearchTarget = &SearchTarget{}
  st = &SearchTarget{
    Id: *res[0].Alias,
    // add parseSource
  }
  return st, nil
}

func (eg EsSearchTargetRepository) CreateSearchTarget(id string) error {
  err := validateSearchTargetId(id)
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
  indexName = id + "_parsesource"
  _, err = client.Indices.Create(indexName).
    Mappings(es.BuildParseSourceMapping()).
    Do(context.Background())
  if err != nil {
    log.Printf(FAIL_REQUEST_ELASTIC_SEARCH, "create ParseSource Index")
    return errors.New(es.HandleElasticsearchError(err))
  }
  return nil
}

func validateSearchTargetId(name string) error {
  if strings.HasPrefix(name, builtInAliasPrefix) {
    return errors.New(es.ES_EM00002)
  }
  return nil
}

func (eg EsSearchTargetRepository) DeleteSearchTarget(id string) error {
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