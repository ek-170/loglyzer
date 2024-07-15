package usecase

import (
	"github.com/ek-170/loglyzer/internal/domain/repository"
)

type AnalysisUsecase struct {
	analysisRepository repository.AnalysisRepository
}

func NewAnalysisUsecase(analysisRepository repository.AnalysisRepository) *AnalysisUsecase {
	return &AnalysisUsecase{
		analysisRepository: analysisRepository,
	}
}

func (stu AnalysisUsecase) FindAnalysiss(q string) ([]*repository.Analysis, error) {
	result, err := stu.analysisRepository.FindAnalysiss(q)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (stu AnalysisUsecase) GetAnalysis(name string) (*repository.Analysis, error) {
	result, err := stu.analysisRepository.GetAnalysis(name)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (stu AnalysisUsecase) CreateAnalysis(name string) error {
	err := stu.analysisRepository.CreateAnalysis(name)
	if err != nil {
		return err
	}
	return nil
}

func (stu AnalysisUsecase) DeleteAnalysis(name string) error {
	err := stu.analysisRepository.DeleteAnalysis(name)
	if err != nil {
		return err
	}
	return nil
}
