package repository

import (
	"context"
	"github.com/micktor/spidey/internal/dto"
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

func (p SpideyRepository) GetByID(ctx context.Context, id int) (dto.SpideyResponse, error) {
	r, err := p.db.User.Get(ctx, id)

	return toOurStruct(*r), err
}

func toOurStruct(user ent.User) dto.SpideyResponse {
	return dto.SpideyResponse{
		Test: "123",
	}
}
