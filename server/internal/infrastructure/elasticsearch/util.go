package elasticsearch

import (
	types "github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
)

func BuildGrokPipeline(pattern string, patternDefs map[string]string, description string) *types.ProcessorContainer {
	grokProcessor := types.NewGrokProcessor()
	grokProcessor.Field = "message"
	grokProcessor.PatternDefinitions = patternDefs
	grokProcessor.Patterns = []string{pattern}
	grokProcessor.Description = &description
	pipeline := types.NewProcessorContainer()
	pipeline.Grok = grokProcessor
	return pipeline
}

func BuildAlias(name string) map[string]types.Alias {
	alias := make(map[string]types.Alias, 1)
	alias[name] = types.Alias{}
	return alias
}

func BuildParseSourceMapping() *types.TypeMapping {
	properties := make(map[string]types.Property, 2)
	properties["name"] = types.KeywordProperty{}
	properties["index"] = types.KeywordProperty{}
	properties["order"] = types.ShortNumberProperty{}
	mapping := types.NewTypeMapping()
	mapping.Properties = properties
	return mapping
}

func BuildParseSourceSort(f string) *types.SortCombinations {
	sort := make(map[string]*types.FieldSort, 1)
	fSort := types.NewFieldSort()
	fSort.Order = &sortorder.Desc
	sort[f] = fSort
	var sortComb types.SortCombinations = sort
	return &sortComb
}

func BuildParseSourceFields() types.FieldAndFormat {
	field := types.NewFieldAndFormat()
	field.Field = "*"
	return *field
}

func BuildParseSourceWildcardQuery(v string, f string) map[string]types.WildcardQuery {
	wq := types.NewWildcardQuery()
	b := true
	wq.CaseInsensitive = &b
	wq.Value = &v
	qMap := make(map[string]types.WildcardQuery, 1)
	qMap[f] = *wq
	return qMap
}