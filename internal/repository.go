package internal

import "context"

type SpideyRepository interface {
	GetByID(ctx context.Context, id int) error
}
