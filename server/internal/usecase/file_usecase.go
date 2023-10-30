package usecase

import (
	"github.com/ek-170/loglyzer/internal/util"
)

type FileUsecase struct {
}

func NewFileUsecase() *FileUsecase {
  return &FileUsecase{}
}

func (gu FileUsecase) FindFiles(q string) (map[string][]string, error) {
  files, err := util.ScanFiles("/logs")
	if err != nil {
	  return nil, err
	}
	if q == "" {
		return map[string][]string{"files": files}, nil
	}
	matchingFiles := util.FindMatchingStrings(files, q)
	return map[string][]string{"files": matchingFiles}, nil
}