package repository

type GrokRepository interface {
	FindGrokPatterns(q string) ([]*GrokPattern, error)
}

type GrokPattern struct {
	Name string     `json:"name"`
	Pattern string  `json:"pattern"`
}