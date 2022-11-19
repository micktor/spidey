package repository

import "github.com/micktor/spidey/internal/ent"

type SpideyRepository struct {
	db *ent.Client
}

func NewSpideyRepository(db *ent.Client) *SpideyRepository {
	return &SpideyRepository{
		db: db,
	}
}

func (p SpideyRepository) Create() error {
	panic("implement me")
}
