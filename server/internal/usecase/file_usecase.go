package usecase

import (
	"github.com/ek-170/loglyzer/internal/util"
	"github.com/ek-170/loglyzer/internal/config"
)

type FileUsecase struct {
}

const LOG_DIR = "/log"

func NewFileUsecase() *FileUsecase {
  return &FileUsecase{}
}

func (gu FileUsecase) FindFiles(q string) (map[string][]string, error) {
  files, err := util.ScanFiles(config.Config.Server.LogDir)
	if err != nil {
	  return nil, err
	}
	if q == "" {
		return map[string][]string{"files": files}, nil
	}
	matchingFiles := util.FindMatchingStrings(files, q)
	return map[string][]string{"files": matchingFiles}, nil
}