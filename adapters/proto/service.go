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

// GetLoad ...
func (s *Service) GetLoad(ctx context.Context, in *pb.GetLoadRequest) (*pb.GetLoadResponse, error) {
	return s.domain.GetLoad(ctx, in)
}

// CreateRifle ...
func (s *Service) CreateRifle(ctx context.Context, in *pb.CreateRifleRequest) (*pb.CreateRifleResponse, error) {
	return s.domain.CreateRifle(ctx, in)
}

// ListRifles ...
func (s *Service) ListRifles(ctx context.Context, in *pb.ListRiflesRequest) (*pb.ListRiflesResponse, error) {
	return s.domain.ListRifles(ctx, in)
}

// GetRifle ...
func (s *Service) GetRifle(ctx context.Context, in *pb.GetRifleRequest) (*pb.GetRifleResponse, error) {
	return s.domain.GetRifle(ctx, in)
}

// CreateEnvironment ...
func (s *Service) CreateEnvironment(ctx context.Context, in *pb.CreateEnvironmentRequest) (*pb.CreateEnvironmentResponse, error) {
	return s.domain.CreateEnvironment(ctx, in)
}

// ListEnvironments ...
func (s *Service) ListEnvironments(ctx context.Context, in *pb.ListEnvironmentsRequest) (*pb.ListEnvironmentsResponse, error) {
	return s.domain.ListEnvironments(ctx, in)
}

// GetEnvironment ...
func (s *Service) GetEnvironment(ctx context.Context, in *pb.GetEnvironmentRequest) (*pb.GetEnvironmentResponse, error) {
	return s.domain.GetEnvironment(ctx, in)
}

// CreateScenario ...
func (s *Service) CreateScenario(ctx context.Context, in *pb.CreateScenarioRequest) (*pb.CreateScenarioResponse, error) {
	return s.domain.CreateScenario(ctx, in)
}

// ListScenarios ...
func (s *Service) ListScenarios(ctx context.Context, in *pb.ListScenariosRequest) (*pb.ListScenariosResponse, error) {
	return s.domain.ListScenarios(ctx, in)
}

// GetScenario ...
func (s *Service) GetScenario(ctx context.Context, in *pb.GetScenarioRequest) (*pb.GetScenarioResponse, error) {
	return s.domain.GetScenario(ctx, in)
}
