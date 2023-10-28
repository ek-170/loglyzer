package usecase

import (
	"log"

	"github.com/ek-170/loglyzer/internal/consts"
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
    log.Printf(consts.FAIL_FIND, "Search Target")
    return nil, err
  }
  return result, err
}

func (gu SearchTargetUsecase) GetSearchTarget(q string) (*repository.SearchTarget, error) {
  result, err := gu.searchTargetRepository.GetSearchTarget(q)
  if err != nil {
    log.Printf(consts.FAIL_GET, "Search Target")
    return nil, err
  }
  return result, err
}

func (gu SearchTargetUsecase) CreateSearchTarget(name string) (error) {
  // err := validateName(name)
  // if err != nil {
  //   log.Printf(consts.FAIL_CREATE, "Search Target")
  //   return err
  // }
  err := gu.searchTargetRepository.CreateSearchTarget(name)
  if err != nil {
    log.Printf(consts.FAIL_CREATE, "Search Target")
    return err
  }
  return err
}

// func validateName(name string) error {
// 
// }

func (gu SearchTargetUsecase) DeleteSearchTarget(name string) (error) {
  err := gu.searchTargetRepository.DeleteSearchTarget(name)
  if err != nil {
    log.Printf(consts.FAIL_DELETE, "Search Target")
    return err
  }
  return err
}