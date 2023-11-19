package ports

import (
	"context"

	"github.com/linuxfreak003/ballistic-calculator/pb"
)

// Service ...
type Service interface {
	Solve(context.Context, *pb.SolveRequest) (*pb.SolveResponse, error)

	// Rifle
	CreateRifle(context.Context, *pb.CreateRifleRequest) (*pb.CreateRifleResponse, error)
	ListRifles(context.Context, *pb.ListRiflesRequest) (*pb.ListRiflesResponse, error)

	// Load
	CreateLoad(context.Context, *pb.CreateLoadRequest) (*pb.CreateLoadResponse, error)
	ListLoads(context.Context, *pb.ListLoadsRequest) (*pb.ListLoadsResponse, error)

	// Environment
	CreateEnvironment(context.Context, *pb.CreateEnvironmentRequest) (*pb.CreateEnvironmentResponse, error)
	ListEnvironments(context.Context, *pb.ListEnvironmentsRequest) (*pb.ListEnvironmentsResponse, error)

	// Scenario
	CreateScenario(context.Context, *pb.CreateScenarioRequest) (*pb.CreateScenarioResponse, error)
	ListScenarios(context.Context, *pb.ListScenariosRequest) (*pb.ListScenariosResponse, error)
}
