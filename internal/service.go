package internal

import (
	"context"
	"github.com/micktor/spidey/internal/dto"
)

type SpideyService interface {
	Get(context.Context, dto.SpideyRequest) (dto.SpideyResponse, error)
}
