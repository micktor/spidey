package service

import (
	"context"
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

func (p SpideyService) Get(ctx context.Context, spideyRequest dto.SpideyRequest) (resp dto.SpideyResponse, err error) {
	resp, err = p.repository.GetByID(ctx, 1)
	return resp, err
}
