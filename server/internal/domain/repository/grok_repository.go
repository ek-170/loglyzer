package repository

type GrokRepository interface {
	GetGrokPatterns(q string) ([]*GrokPattern, error) // 必要なパラメータの調査
}

type GrokPattern struct {
	Name string
  Pattern string
}