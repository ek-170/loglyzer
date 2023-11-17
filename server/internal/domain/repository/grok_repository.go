package repository

type GrokRepository interface {
	FindGrokPatterns(q string) ([]*GrokPattern, error)
}

type GrokPattern struct {
	Id string       `json:"id"`
  // Name string     `json:"name"`
	Pattern string  `json:"pattern"`
  // Discription string `json:"discription`
}