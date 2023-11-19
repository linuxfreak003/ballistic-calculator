package proto

import (
	"context"

	"github.com/linuxfreak003/ballistic-calculator/pb"
	"github.com/linuxfreak003/ballistic-calculator/ports"
)

// NewService ...
func NewService(s ports.Service) *Service {
	return &Service{
		domain: s,
	}
}

// Service ...
type Service struct {
	domain ports.Service
	unimplementedBallisticServiceServer
}

type unimplementedBallisticServiceServer struct {
	pb.UnimplementedBallisticServiceServer
}

// Solve ...
func (s *Service) Solve(ctx context.Context, in *pb.SolveRequest) (*pb.SolveResponse, error) {
	return s.domain.Solve(ctx, in)
}

// CreateLoad ...
func (s *Service) CreateLoad(ctx context.Context, in *pb.CreateLoadRequest) (*pb.CreateLoadResponse, error) {
	return s.domain.CreateLoad(ctx, in)
}

// ListLoads ...
func (s *Service) ListLoads(ctx context.Context, in *pb.ListLoadsRequest) (*pb.ListLoadsResponse, error) {
	return s.domain.ListLoads(ctx, in)
}

// CreateRifle ...
func (s *Service) CreateRifle(ctx context.Context, in *pb.CreateRifleRequest) (*pb.CreateRifleResponse, error) {
	return s.domain.CreateRifle(ctx, in)
}

// ListRifles ...
func (s *Service) ListRifles(ctx context.Context, in *pb.ListRiflesRequest) (*pb.ListRiflesResponse, error) {
	return s.domain.ListRifles(ctx, in)
}

// CreateEnvironment ...
func (s *Service) CreateEnvironment(ctx context.Context, in *pb.CreateEnvironmentRequest) (*pb.CreateEnvironmentResponse, error) {
	return s.domain.CreateEnvironment(ctx, in)
}

// ListEnvironments ...
func (s *Service) ListEnvironments(ctx context.Context, in *pb.ListEnvironmentsRequest) (*pb.ListEnvironmentsResponse, error) {
	return s.domain.ListEnvironments(ctx, in)
}

// CreateScenario ...
func (s *Service) CreateScenario(ctx context.Context, in *pb.CreateScenarioRequest) (*pb.CreateScenarioResponse, error) {
	return s.domain.CreateScenario(ctx, in)
}

// ListScenarios ...
func (s *Service) ListScenarios(ctx context.Context, in *pb.ListScenariosRequest) (*pb.ListScenariosResponse, error) {
	return s.domain.ListScenarios(ctx, in)
}
