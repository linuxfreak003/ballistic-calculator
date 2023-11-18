package ports

import (
	"context"
)

// Repository holds all the related entities
type Repository interface {
	CreateLoad(context.Context, *Load) (int64, error)
}

// Load ...
type Load struct {
	LoadId         int64
	Bullet         Bullet
	MuzzleVelocity float64
}

// Bullet ...
type Bullet struct {
	Caliber float64
	Weight  float64
	BC      Coefficient
	Length  float64
}

// Coefficient ...
type Coefficient struct {
	Value    float64
	DragFunc int
}
