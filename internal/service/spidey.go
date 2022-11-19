package service

import (
	"github.com/micktor/spidey/internal/config"
	"github.com/micktor/spidey/internal/repository"
)

type SpideyService struct {
}

func NewSpideyService(config *config.Config, repo *repository.SpideyRepository) *SpideyService {
	return &SpideyService{}
}

func (p SpideyService) Test() (err error) {
	panic("implement me")
}
