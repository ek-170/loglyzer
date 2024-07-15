package repository

type AnalysisRepository interface {
	FindAnalysiss(q string) ([]*Analysis, error)
	GetAnalysis(id string) (*Analysis, error)
	CreateAnalysis(id string) error
	DeleteAnalysis(id string) error
}

type Analysis struct {
	Id           string         `json:"id"`
	DataViewId   string         `json:"dataViewId"`
	ParseSources []*ParseSource `json:"parseSources"`
}
