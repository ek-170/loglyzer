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

func (psu ParseSourceUsecase) FindParseSources(q string, searchTarget string) ([]*repository.ParseSource, error) {
	result, err := psu.parseSourceRepository.FindParseSources(q, searchTarget)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (psu ParseSourceUsecase) CreateParseSource(
	searchTarget string, multiLine bool, filePath string, grokId string) error {
	err := psu.parseSourceRepository.CreateParseSource(searchTarget, multiLine, filePath, grokId)
	if err != nil {
		return err
	}
	return nil
}

func (psu ParseSourceUsecase) DeleteParseSource(id string, searchTarget string) error {
	err := psu.parseSourceRepository.DeleteParseSource(id, searchTarget)
	if err != nil {
		return err
	}
	return nil
}
