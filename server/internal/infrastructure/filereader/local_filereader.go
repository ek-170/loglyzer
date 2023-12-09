package filereader

import (
	"io"
	"log"
	"os"

	"github.com/ek-170/loglyzer/internal/config"
	"github.com/ek-170/loglyzer/internal/util"
)

type LocalFileReader struct {
	path string
}

func (lfr *LocalFileReader) ReadFile() (io.Reader, error) {
	// at first, check existence of file
	searchedFile, err := util.SearchFile(config.Config.Server.LogDir, lfr.path)
	if err != nil {
		return nil, err
	}
	fullPath := config.Config.Server.LogDir + "/" + searchedFile
	file, err := os.Open(fullPath)
	if err != nil {
		log.Println("failed to open file")
		return nil, err
	}
	return file , nil
}

func NewLocalFileReader(conf FileReaderConfig) LocalFileReader{
	lfr := LocalFileReader{}
	lfr.path =conf.Path
	return lfr
}