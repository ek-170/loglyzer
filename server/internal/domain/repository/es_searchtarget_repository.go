package repository

import (
	"context"
	"log"
	"sort"
	"strings"

	"github.com/ek-170/loglyzer/internal/consts"
	"github.com/ek-170/loglyzer/internal/infrastructure/elasticsearch"
)

type EsSearchTargetRepository struct {}

func NewEsSearchTargetRepository() EsSearchTargetRepository{
  return EsSearchTargetRepository{}
}

func (eg EsSearchTargetRepository) FindSearchTargets(q string) ([]*SearchTarget, error){
  client, err := elasticsearch.CreateElasticsearchClient()
  if err != nil {
    return nil, err
  }
  res, err := client.Cat.Aliases().Do(context.TODO())
  if err != nil {
    log.Printf(consts.FAIL_REQUEST_ELASTIC_SEARCH, "cat Aliases")
    return nil, err
  }
  var searchTargets []*SearchTarget = []*SearchTarget{}
  builtInAliasPrefix := "."
  for _, alias := range res {
    // if q is empty, return all searchTargets
    if q != "" && q != *alias.Alias {
      continue
    }
    // alias start with "." is built-in alias
    if strings.HasPrefix(*alias.Alias, builtInAliasPrefix){
      continue
    }
    // duplicate aliases may be returned
    if contains(searchTargets, *alias.Alias){
      continue
    }
    if *alias.Alias != "" {
      searchTarget := &SearchTarget{
        Name:  *alias.Alias,
        // parseSource
      }
      searchTargets = append(searchTargets, searchTarget)
    }
  }
  return sortSearchTarget(searchTargets, true), nil
}

func contains(searchTargets []*SearchTarget, new string) bool {
  for _, s := range searchTargets {
    if s.Name == new {
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

func (eg EsSearchTargetRepository) GetSearchTarget(name string) (*SearchTarget, error){
  client, err := elasticsearch.CreateElasticsearchClient()
  if err != nil {
    return nil, err
  }
  res, err := client.Cat.Aliases().Name(name).Do(context.TODO())
  if err != nil {
    log.Printf(consts.FAIL_REQUEST_ELASTIC_SEARCH, "get Aliases")
    return nil, err
  }
  var st *SearchTarget = &SearchTarget{}
  st = &SearchTarget{
    Name: *res[0].Alias,
    // add parseSource
  }
  return st, nil
}

func (eg EsSearchTargetRepository) CreateSearchTarget(name string) error {
  client, err := elasticsearch.CreateElasticsearchClient()
  if err != nil {
    return err
  }
  // to create Alias, also need to exist Index
  // so create placeholder Index which is not used
  indexName := name + "_placeholder"
  _, err = client.Indices.Create(indexName).Do(context.TODO())
  if err != nil {
    log.Printf(consts.FAIL_REQUEST_ELASTIC_SEARCH, "create placeholder Index")
    return err
  }
  _, err = client.Indices.PutAlias(indexName, name).Do(context.TODO())
  if err != nil {
    log.Printf(consts.FAIL_REQUEST_ELASTIC_SEARCH, "create Alias")
    return err
  }
  return nil
}

func (eg EsSearchTargetRepository) DeleteSearchTarget(name string) error {
  client, err := elasticsearch.CreateElasticsearchClient()
  if err != nil {
    return err
  }
  // get all Indices name
  res, err := client.Indices.GetAlias().Name(name).Do(context.TODO())
  if err != nil {
    log.Printf(consts.FAIL_REQUEST_ELASTIC_SEARCH, "get all indices")
    return err
  }
  // delete all Indices 
  for key := range res {
    _, err = client.Indices.Delete(key).Do(context.TODO())
    if err != nil {
      log.Printf(consts.FAIL_REQUEST_ELASTIC_SEARCH, "delete all indices")
      return err
    }
  }
  // after delete all Indices, automatically Alias has deleted
  return nil
}