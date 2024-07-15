package usecase

import (
	fr "github.com/ek-170/loglyzer/internal/domain/filereader"
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

func (psu ParseSourceUsecase) FindParseSources(q string, analysis string) ([]*repository.ParseSource, error) {
	result, err := psu.parseSourceRepository.FindParseSources(q, analysis)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (psu ParseSourceUsecase) CreateParseSource(
	analysis string, multiLine bool, frConf fr.FileReaderConfig, grokId string) error {
	err := psu.parseSourceRepository.CreateParseSource(analysis, multiLine, frConf, grokId)
	if err != nil {
		return err
	}
	return nil
}

func (psu ParseSourceUsecase) DeleteParseSource(id string, analysis string) error {
	err := psu.parseSourceRepository.DeleteParseSource(id, analysis)
	if err != nil {
		return err
	}
	return nil
}
