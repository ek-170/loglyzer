package repository

type SearchTargetRepository interface {
	FindSearchTargets(q string) ([]*SearchTarget, error)
	GetSearchTarget(name string) (*SearchTarget, error)
	CreateSearchTarget(name string) error
	DeleteSearchTarget(name string) error
}

type SearchTarget struct {
	Name string     `json:"name"`
	// ParseSources []*ParseSource
}