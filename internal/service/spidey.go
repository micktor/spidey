package service

import (
	"github.com/micktor/spidey/internal/config"
	"github.com/micktor/spidey/internal/dto"
	"github.com/micktor/spidey/internal/repository"
)

type SpideyService struct {
	config     *config.Config
	repository *repository.SpideyRepository
}

func NewSpideyService(config *config.Config, repo *repository.SpideyRepository) *SpideyService {
	return &SpideyService{
		config:     config,
		repository: repo,
	}
}

func (p SpideyService) Test(spideyRequest dto.SpideyRequest) (resp dto.SpideyResponse, err error) {
	panic("implement me")
}
