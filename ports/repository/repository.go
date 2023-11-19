package repository

import (
	"context"
)

// Repository holds all the related entities
type Repository interface {
	// Rifle
	CreateRifle(context.Context, *Rifle) (int64, error)
	ListRifles(context.Context) ([]*Rifle, error)
	GetRifle(context.Context, int64) (*Rifle, error)

	// Load
	CreateLoad(context.Context, *Load) (int64, error)
	ListLoads(context.Context) ([]*Load, error)
	GetLoad(context.Context, int64) (*Load, error)

	// Environment
	CreateEnvironment(context.Context, *Environment) (int64, error)
	ListEnvironments(context.Context) ([]*Environment, error)
	GetEnvironment(context.Context, int64) (*Environment, error)

	// Scenario
	CreateScenario(context.Context, *Scenario) (int64, error)
	ListScenarios(context.Context) ([]*Scenario, error)
	GetScenario(context.Context, int64) (*Scenario, error)
}
