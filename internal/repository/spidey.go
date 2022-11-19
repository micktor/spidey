package repository

import (
	"context"
	"github.com/micktor/spidey/internal/ent"
)

type SpideyRepository struct {
	db *ent.Client
}

func NewSpideyRepository(db *ent.Client) *SpideyRepository {
	return &SpideyRepository{
		db: db,
	}
}

func (p SpideyRepository) GetByID(ctx context.Context, id int) error {
	_, err := p.db.User.Get(ctx, id)
	return err
}
