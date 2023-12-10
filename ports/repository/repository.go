package repository

import (
	"context"
)

// Repository holds all the related entities
type Repository interface {
	// Rifle
	CreateRifle(context.Context, *Rifle) (int64, error)
	ListRifles(context.Context, *ListRequest) ([]*Rifle, *ListResponse, error)
	GetRifle(context.Context, int64) (*Rifle, error)

	// Load
	CreateLoad(context.Context, *Load) (int64, error)
	ListLoads(context.Context, *ListRequest) ([]*Load, *ListResponse, error)
	GetLoad(context.Context, int64) (*Load, error)

	// Environment
	CreateEnvironment(context.Context, *Environment) (int64, error)
	ListEnvironments(context.Context, *ListRequest) ([]*Environment, *ListResponse, error)
	GetEnvironment(context.Context, int64) (*Environment, error)

	// Scenario
	CreateScenario(context.Context, *Scenario) (int64, error)
	ListScenarios(context.Context, *ListRequest) ([]*Scenario, *ListResponse, error)
	GetScenario(context.Context, int64) (*Scenario, error)
}
