package domain

import (
	"context"

	"github.com/linuxfreak003/ballistic-calculator/pb"
	"github.com/linuxfreak003/ballistic-calculator/ports/repository"
	"github.com/linuxfreak003/util/slice"
	"gitlab.com/linuxfreak003/ballistic"
)

// Service ...
type Service struct {
	repo repository.Repository
}

// NewService returns a new Service
func NewService(repo repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// Solve ...
func (s *Service) Solve(ctx context.Context, in *pb.SolveRequest) (*pb.SolveResponse, error) {
	scenario, err := s.repo.GetScenario(ctx, in.ScenarioId)
	if err != nil {
		return nil, err
	}

	r, err := s.repo.GetRifle(ctx, scenario.RifleId)
	if err != nil {
		return nil, err
	}

	e, err := s.repo.GetEnvironment(ctx, scenario.EnvironmentId)
	if err != nil {
		return nil, err
	}

	l, err := s.repo.GetLoad(ctx, scenario.LoadId)
	if err != nil {
		return nil, err
	}

	solver := ballistic.NewSolver(r.ToBallistic(), e.ToBallistic(), l.ToBallistic(), true)
	solutionSet := solver.Generate(0, 1400, 1)

	solutions := slice.Map(solutionSet, func(s *ballistic.Solution) *pb.Solution {
		return &pb.Solution{
			Range:      int64(s.Range),
			RawRange:   s.RawRange,
			Drop:       s.Drop,
			DropMoa:    s.DropMOA,
			Time:       s.Time,
			Windage:    s.Windage,
			WindageMoa: s.WindageMOA,
			Energy:     s.Energy,
			Velocity:   s.Velocity,
			VelocityX:  s.VelocityX,
			VelocityY:  s.VelocityY,
		}
	})
	return &pb.SolveResponse{
		Solutions: solutions,
	}, nil
}

// CreateLoad ...
func (s *Service) CreateLoad(ctx context.Context, in *pb.CreateLoadRequest) (*pb.CreateLoadResponse, error) {
	var load *repository.Load
	load = load.FromProto(in.GetLoad())
	id, err := s.repo.CreateLoad(ctx, load)
	if err != nil {
		return nil, err
	}
	load.LoadId = id
	return &pb.CreateLoadResponse{
		Load: load.ToProto(),
	}, nil
}

// ListLoads ...
func (s *Service) ListLoads(ctx context.Context, _ *pb.ListLoadsRequest) (*pb.ListLoadsResponse, error) {
	loads, err := s.repo.ListLoads(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.ListLoadsResponse{
		Loads: slice.Map(loads, func(l *repository.Load) *pb.Load { return l.ToProto() }),
	}, nil
}

// CreateRifle ...
func (s *Service) CreateRifle(ctx context.Context, in *pb.CreateRifleRequest) (*pb.CreateRifleResponse, error) {
	var r *repository.Rifle
	r = r.FromProto(in.GetRifle())

	id, err := s.repo.CreateRifle(ctx, r)
	if err != nil {
		return nil, err
	}
	r.RifleId = id
	return &pb.CreateRifleResponse{
		Rifle: r.ToProto(),
	}, nil
}

// ListRifles ...
func (s *Service) ListRifles(ctx context.Context, _ *pb.ListRiflesRequest) (*pb.ListRiflesResponse, error) {
	rifles, err := s.repo.ListRifles(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.ListRiflesResponse{
		Rifles: slice.Map(rifles, func(r *repository.Rifle) *pb.Rifle { return r.ToProto() }),
	}, nil
}

// CreateEnvironment ...
func (s *Service) CreateEnvironment(ctx context.Context, in *pb.CreateEnvironmentRequest) (*pb.CreateEnvironmentResponse, error) {
	var r *repository.Environment
	r = r.FromProto(in.GetEnvironment())

	id, err := s.repo.CreateEnvironment(ctx, r)
	if err != nil {
		return nil, err
	}
	r.EnvironmentId = id
	return &pb.CreateEnvironmentResponse{
		Environment: r.ToProto(),
	}, nil
}

// ListEnvironments ...
func (s *Service) ListEnvironments(ctx context.Context, _ *pb.ListEnvironmentsRequest) (*pb.ListEnvironmentsResponse, error) {
	environments, err := s.repo.ListEnvironments(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.ListEnvironmentsResponse{
		Environments: slice.Map(environments, func(r *repository.Environment) *pb.Environment { return r.ToProto() }),
	}, nil
}

// CreateScenario ...
func (s *Service) CreateScenario(ctx context.Context, in *pb.CreateScenarioRequest) (*pb.CreateScenarioResponse, error) {
	var r *repository.Scenario
	r = r.FromProto(in.GetScenario())

	id, err := s.repo.CreateScenario(ctx, r)
	if err != nil {
		return nil, err
	}
	r.ScenarioId = id
	return &pb.CreateScenarioResponse{
		Scenario: r.ToProto(),
	}, nil
}

// ListScenarios ...
func (s *Service) ListScenarios(ctx context.Context, _ *pb.ListScenariosRequest) (*pb.ListScenariosResponse, error) {
	scenarios, err := s.repo.ListScenarios(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.ListScenariosResponse{
		Scenarios: slice.Map(scenarios, func(r *repository.Scenario) *pb.Scenario { return r.ToProto() }),
	}, nil
}
