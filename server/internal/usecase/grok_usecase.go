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
  return result, err
}