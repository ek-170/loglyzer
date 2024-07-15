package repository

import fr "github.com/ek-170/loglyzer/internal/domain/filereader"

type ParseSourceRepository interface {
	FindParseSources(q string, analysis string) ([]*ParseSource, error)
	CreateParseSource(analysis string, multiLine bool, frConf fr.FileReaderConfig, grokId string) error
	DeleteParseSource(id string, analysis string) error
}

// ParseSource is information of parsed log's index
type ParseSource struct {
	Id    string `json:"id"`    // ParseSource Info Doc ID(only fill as response)
	Name  string `json:"name"`  // parse target file name
	Index string `json:"index"` // save target Index name
	Order int16  `json:"order"` // use to name index
}
