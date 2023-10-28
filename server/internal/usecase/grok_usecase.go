package usecase

import (
	"log"

	"github.com/ek-170/loglyzer/internal/consts"
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
    log.Printf(consts.FAIL_GET, "Grok Pattern")
    return nil, err
  }
  return result, err
}