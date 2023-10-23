package repository

type GrokRepository interface {
	GetGrokPatterns(q string) ([]*GrokPattern, error)
}

type GrokPattern struct {
	Name string     `json:"name"`
	Pattern string  `json:"pattern"`
}