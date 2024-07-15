package repository

type GrokRepository interface {
	FindGrokPatterns(q string) ([]*GrokPattern, error)
	CreateGrokPattern(id string, pattern string, patternDefs map[string]string, description string) error
	DeleteGrokPattern(id string) error
}

type GrokPattern struct {
	Id string       `json:"id"`
	Pattern string  `json:"pattern"`
	PatternDefs map[string]string `json:"pattern_definitions"`
  Description string `json:"description"`
}