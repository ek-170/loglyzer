package usecase

import (
	"github.com/ek-170/loglyzer/internal/domain/repository"
)

type SearchTargetUsecase struct {
  searchTargetRepository repository.SearchTargetRepository
}

func NewSearchTargetUsecase(searchTargetRepository repository.SearchTargetRepository) *SearchTargetUsecase {
  return &SearchTargetUsecase{
    searchTargetRepository: searchTargetRepository,
  }
}

func (gu SearchTargetUsecase) FindSearchTargets(q string) ([]*repository.SearchTarget, error) {
  result, err := gu.searchTargetRepository.FindSearchTargets(q)
  if err != nil {
    return nil, err
  }
  return result, err
}

func (gu SearchTargetUsecase) GetSearchTarget(q string) (*repository.SearchTarget, error) {
  result, err := gu.searchTargetRepository.GetSearchTarget(q)
  if err != nil {
    return nil, err
  }
  return result, err
}

func (gu SearchTargetUsecase) CreateSearchTarget(name string) error {
  err := gu.searchTargetRepository.CreateSearchTarget(name)
  if err != nil {
    return err
  }
  return err
}

func (gu SearchTargetUsecase) DeleteSearchTarget(name string) error {
  err := gu.searchTargetRepository.DeleteSearchTarget(name)
  if err != nil {
    return err
  }
  return err
}