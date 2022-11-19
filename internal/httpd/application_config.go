package httpd

import (
	"github.com/micktor/spidey/internal"
	"github.com/micktor/spidey/internal/config"
	"github.com/micktor/spidey/internal/repository"
	"github.com/micktor/spidey/internal/service"
)

type applicationConfig struct {
	config           *config.Config
	db               *ent.Client
	spideyHandler    *SpideyHandler
	spideyService    internal.SpideyService
	spideyRepository internal.SpideyRepository
}

func newApplicationConfig(config *config.Config, db *ent.Client) *applicationConfig {
	applicationConfig := &applicationConfig{
		config: config,
		db:     db,
	}

	applicationConfig.addSpideyDependencies()
	return applicationConfig
}

func (a *applicationConfig) addSpideyDependencies() *applicationConfig {
	repo := repository.NewSpideyRepository(a.db)
	a.spideyService = service.NewSpideyService(a.config, repo)
	a.spideyHandler = NewSpideyHandler(a)

	return a
}
