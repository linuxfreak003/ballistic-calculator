package repository

import (
	"context"
)

// Repository holds all the related entities
type Repository interface {
	CreateLoad(context.Context, *Load) (int64, error)
}
