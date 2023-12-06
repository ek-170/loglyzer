package usecase

import (
	"github.com/ek-170/loglyzer/internal/domain/repository"
)

type GrokUsecase struct {
  grokRepository repository.GrokRepository
}

func NewGrokUsecase(grokRepository repository.GrokRepository) *GrokUsecase {
  return &GrokUsecase{
    grokRepository: grokRepository,
  }
}

func (gu GrokUsecase) FindGrokPatterns(q string) ([]*repository.GrokPattern, error) {
  result, err := gu.grokRepository.FindGrokPatterns(q)
  if err != nil {
    return nil, err
  }
  return result, nil
}

func (gu GrokUsecase) CreateGrokPatterns(id string, pattern string, patternDefs map[string]string, description string) error {
  err := gu.grokRepository.CreateGrokPattern(id, pattern, patternDefs, description)
  if err != nil {
    return err
  }
  return nil
}

func (gu GrokUsecase) DeleteGrokPatterns(id string) error {
  err := gu.grokRepository.DeleteGrokPattern(id)
  if err != nil {
    return err
  }
  return nil
}