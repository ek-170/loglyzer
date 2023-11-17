package usecase

import (
	"github.com/ek-170/loglyzer/internal/domain/repository"
)

type ParseSourceUsecase struct {
  parseSourceRepository repository.ParseSourceRepository
}

func NewParseSourceUsecase(parseSourceRepository repository.ParseSourceRepository) *ParseSourceUsecase {
  return &ParseSourceUsecase{
    parseSourceRepository: parseSourceRepository,
  }
}

func (psu ParseSourceUsecase) FindParseSources(q string) ([]*repository.ParseSource, error) {
  result, err := psu.parseSourceRepository.FindParseSources(q)
  if err != nil {
    return nil, err
  }
  return result, nil
}

func (psu ParseSourceUsecase) GetParseSource(name string) (*repository.ParseSource, error) {
  result, err := psu.parseSourceRepository.GetParseSource(name)
  if err != nil {
    return nil, err
  }
  return result, nil
}

func (psu ParseSourceUsecase) CreateParseSource(
  searchTarget string, multiLine bool, fileName string, grokId string) error {
  err := psu.parseSourceRepository.CreateParseSource(searchTarget, multiLine, fileName, grokId)
  if err != nil {
    return err
  }
  return nil
}

func (psu ParseSourceUsecase) DeleteParseSource(name string) error {
  err := psu.parseSourceRepository.DeleteParseSource(name)
  if err != nil {
    return err
  }
  return nil
}