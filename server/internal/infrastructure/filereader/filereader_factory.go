package filereader

import (
	fr "github.com/ek-170/loglyzer/internal/domain/filereader"
)


type FileReaderConfig struct {
	// common conf
	FileReadMode string
	Path string
	// ssh conf
	SshKeyPath   string
	UserName     string
	Password     string
	Host         string
	Port         int
}

type FileReadMode string

func InitFileReader(conf FileReaderConfig) fr.FileReader{
	switch conf.FileReadMode {
	case "ssh":
		return NewSshFileReader(conf)
	case "local":
		return NewLocalFileReader(conf)
	default:
		return NewLocalFileReader(conf)
	}

}