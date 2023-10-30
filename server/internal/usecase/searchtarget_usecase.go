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

func (stu SearchTargetUsecase) FindSearchTargets(q string) ([]*repository.SearchTarget, error) {
  result, err := stu.searchTargetRepository.FindSearchTargets(q)
  if err != nil {
    return nil, err
  }
  return result, nil
}

func (stu SearchTargetUsecase) GetSearchTarget(name string) (*repository.SearchTarget, error) {
  result, err := stu.searchTargetRepository.GetSearchTarget(name)
  if err != nil {
    return nil, err
  }
  return result, nil
}

func (stu SearchTargetUsecase) CreateSearchTarget(name string) error {
  err := stu.searchTargetRepository.CreateSearchTarget(name)
  if err != nil {
    return err
  }
  return nil
}

func (stu SearchTargetUsecase) DeleteSearchTarget(name string) error {
  err := stu.searchTargetRepository.DeleteSearchTarget(name)
  if err != nil {
    return err
  }
  return nil
}