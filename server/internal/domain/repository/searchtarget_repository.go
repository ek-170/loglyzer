package repository

type SearchTargetRepository interface {
	FindSearchTargets(q string) ([]*SearchTarget, error)
	GetSearchTarget(id string) (*SearchTarget, error)
	CreateSearchTarget(id string) error
	DeleteSearchTarget(id string) error
}

type SearchTarget struct {
	Id string     `json:"id"`
	// ParseSources []*ParseSource
}