package elasticsearch

import (
	types "github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func BuildAlias(name string) map[string]types.Alias {
	alias := make(map[string]types.Alias, 1)
	alias[name] = types.Alias{}
	return alias
}

func BuildParseSourceInfoMapping() *types.TypeMapping {
	properties := make(map[string]types.Property, 2)
	properties["name"] = types.KeywordProperty{}
	properties["index"] = types.KeywordProperty{}
	mapping := types.NewTypeMapping()
	mapping.Properties = properties
	return mapping
}