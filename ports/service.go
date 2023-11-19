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
	GetRifle(context.Context, *pb.GetRifleRequest) (*pb.GetRifleResponse, error)

	// Load
	CreateLoad(context.Context, *pb.CreateLoadRequest) (*pb.CreateLoadResponse, error)
	ListLoads(context.Context, *pb.ListLoadsRequest) (*pb.ListLoadsResponse, error)
	GetLoad(context.Context, *pb.GetLoadRequest) (*pb.GetLoadResponse, error)

	// Environment
	CreateEnvironment(context.Context, *pb.CreateEnvironmentRequest) (*pb.CreateEnvironmentResponse, error)
	ListEnvironments(context.Context, *pb.ListEnvironmentsRequest) (*pb.ListEnvironmentsResponse, error)
	GetEnvironment(context.Context, *pb.GetEnvironmentRequest) (*pb.GetEnvironmentResponse, error)

	// Scenario
	CreateScenario(context.Context, *pb.CreateScenarioRequest) (*pb.CreateScenarioResponse, error)
	ListScenarios(context.Context, *pb.ListScenariosRequest) (*pb.ListScenariosResponse, error)
	GetScenario(context.Context, *pb.GetScenarioRequest) (*pb.GetScenarioResponse, error)
}
