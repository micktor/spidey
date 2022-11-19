package internal

import "github.com/micktor/spidey/internal/dto"

type SpideyService interface {
	Test(dto.SpideyRequest) (dto.SpideyResponse, error)
}
