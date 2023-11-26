package repository

type ParseSourceRepository interface {
	FindParseSources(q string) ([]*ParseSource, error)
	GetParseSource(name string) (*ParseSource, error)
	CreateParseSource(searchTarget string, multiLine bool, fileName string, grokId string) (error)
	DeleteParseSource(name string) error
}

// ParseSource is information of parsed log's index
type ParseSource struct {
	Id string    `json:"id"`    // Doc ID
	Name string  `json:"name"`  // parse target file name
	Index string `json:"index"` // save target Index name
	Order int16  `json:"order"` // use to name index
}