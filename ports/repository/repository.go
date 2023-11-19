package repository

import (
	"context"
)

// Repository holds all the related entities
type Repository interface {
	// Rifle
	CreateRifle(context.Context, *Rifle) (int64, error)
	ListRifles(context.Context) ([]*Rifle, error)

	// Load
	CreateLoad(context.Context, *Load) (int64, error)
	ListLoads(context.Context) ([]*Load, error)
}
