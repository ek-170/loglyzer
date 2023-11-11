package repository

type GrokRepository interface {
	FindGrokPatterns(q string) ([]*GrokPattern, error)
}

type GrokPattern struct {
	Id string     `json:"id"`
	Pattern string  `json:"pattern"`
}