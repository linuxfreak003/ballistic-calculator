package domain

import (
	"context"
	"os"

	"github.com/linuxfreak003/ballistic-calculator/pb"
	"github.com/linuxfreak003/ballistic-calculator/ports/repository"
	"github.com/linuxfreak003/util/slice"
	log "github.com/sirupsen/logrus"
	"gitlab.com/linuxfreak003/ballistic"
)

// Service ...
type Service struct {
	repo repository.Repository
}

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

// NewService returns a new Service
func NewService(repo repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// BallisticSolutionToProto ...
func BallisticSolutionToProto(s *ballistic.Solution) *pb.Solution {
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
}

// Solve ...
func (s *Service) Solve(ctx context.Context, in *pb.SolveRequest) (*pb.SolveResponse, error) {
	log.Infof("Solve called")
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
	solution, err := solver.SolveFor(int(in.GetRange()))
	if err != nil {
		return nil, err
	}
	return &pb.SolveResponse{
		Solution: BallisticSolutionToProto(solution),
	}, nil
}

// SolveTable ...
func (s *Service) SolveTable(ctx context.Context, in *pb.SolveTableRequest) (*pb.SolveTableResponse, error) {
	log.Infof("SolveTable called")
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
	solutionSet := solver.Generate(0, int(in.GetMaxRange()), int(in.GetIncrement()))

	solutions := slice.Map(solutionSet, BallisticSolutionToProto)
	return &pb.SolveTableResponse{
		Solutions: solutions,
	}, nil
}

// CreateLoad ...
func (s *Service) CreateLoad(ctx context.Context, in *pb.CreateLoadRequest) (*pb.CreateLoadResponse, error) {
	log.Infof("CreateLoad called")
	var load *repository.Load
	load = load.FromProto(in.GetLoad())
	id, err := s.repo.CreateLoad(ctx, load)
	if err != nil {
		log.Errorf("could not create load: %v", err)
		return nil, err
	}
	load.LoadId = id
	return &pb.CreateLoadResponse{
		Load: load.ToProto(),
	}, nil
}

// ListLoads ...
func (s *Service) ListLoads(ctx context.Context, _ *pb.ListLoadsRequest) (*pb.ListLoadsResponse, error) {
	log.Infof("ListLoads called")
	loads, err := s.repo.ListLoads(ctx)
	if err != nil {
		log.Errorf("could not list loads: %v", err)
		return nil, err
	}
	return &pb.ListLoadsResponse{
		Loads: slice.Map(loads, func(l *repository.Load) *pb.Load { return l.ToProto() }),
	}, nil
}

// GetLoad ...
func (s *Service) GetLoad(ctx context.Context, in *pb.GetLoadRequest) (*pb.GetLoadResponse, error) {
	log.Infof("GetLoad called")
	load, err := s.repo.GetLoad(ctx, in.LoadId)
	if err != nil {
		log.Errorf("could not get load: %v", err)
		return nil, err
	}
	return &pb.GetLoadResponse{
		Load: load.ToProto(),
	}, nil
}

// CreateRifle ...
func (s *Service) CreateRifle(ctx context.Context, in *pb.CreateRifleRequest) (*pb.CreateRifleResponse, error) {
	log.Infof("CreateRifle called")
	var r *repository.Rifle
	r = r.FromProto(in.GetRifle())

	id, err := s.repo.CreateRifle(ctx, r)
	if err != nil {
		log.Errorf("could not create rifle: %v", err)
		return nil, err
	}
	r.RifleId = id
	return &pb.CreateRifleResponse{
		Rifle: r.ToProto(),
	}, nil
}

// ListRifles ...
func (s *Service) ListRifles(ctx context.Context, _ *pb.ListRiflesRequest) (*pb.ListRiflesResponse, error) {
	log.Infof("ListRifles called")
	rifles, err := s.repo.ListRifles(ctx)
	if err != nil {
		log.Errorf("could not list rifles: %v", err)
		return nil, err
	}
	return &pb.ListRiflesResponse{
		Rifles: slice.Map(rifles, func(r *repository.Rifle) *pb.Rifle { return r.ToProto() }),
	}, nil
}

// GetRifle ...
func (s *Service) GetRifle(ctx context.Context, in *pb.GetRifleRequest) (*pb.GetRifleResponse, error) {
	log.Infof("GetRifle called")
	rifle, err := s.repo.GetRifle(ctx, in.RifleId)
	if err != nil {
		log.Errorf("could not get rifle: %v", err)
		return nil, err
	}
	return &pb.GetRifleResponse{
		Rifle: rifle.ToProto(),
	}, nil
}

// CreateEnvironment ...
func (s *Service) CreateEnvironment(ctx context.Context, in *pb.CreateEnvironmentRequest) (*pb.CreateEnvironmentResponse, error) {
	log.Infof("CreateEnvironment called")
	var r *repository.Environment
	r = r.FromProto(in.GetEnvironment())

	id, err := s.repo.CreateEnvironment(ctx, r)
	if err != nil {
		log.Errorf("could not create environment: %v", err)
		return nil, err
	}
	r.EnvironmentId = id
	return &pb.CreateEnvironmentResponse{
		Environment: r.ToProto(),
	}, nil
}

// ListEnvironments ...
func (s *Service) ListEnvironments(ctx context.Context, _ *pb.ListEnvironmentsRequest) (*pb.ListEnvironmentsResponse, error) {
	log.Infof("ListEnvironments called")
	environments, err := s.repo.ListEnvironments(ctx)
	if err != nil {
		log.Errorf("could not list environments: %v", err)
		return nil, err
	}
	return &pb.ListEnvironmentsResponse{
		Environments: slice.Map(environments, func(r *repository.Environment) *pb.Environment { return r.ToProto() }),
	}, nil
}

// GetEnvironment ...
func (s *Service) GetEnvironment(ctx context.Context, in *pb.GetEnvironmentRequest) (*pb.GetEnvironmentResponse, error) {
	log.Infof("GetEnvironment called")
	environment, err := s.repo.GetEnvironment(ctx, in.EnvironmentId)
	if err != nil {
		log.Errorf("could not get environment: %v", err)
		return nil, err
	}
	return &pb.GetEnvironmentResponse{
		Environment: environment.ToProto(),
	}, nil
}

// CreateScenario ...
func (s *Service) CreateScenario(ctx context.Context, in *pb.CreateScenarioRequest) (*pb.CreateScenarioResponse, error) {
	log.Infof("CreateScenario called")
	var r *repository.Scenario
	r = r.FromProto(in.GetScenario())

	id, err := s.repo.CreateScenario(ctx, r)
	if err != nil {
		log.Errorf("could not create scenario: %v", err)
		return nil, err
	}
	r.ScenarioId = id
	return &pb.CreateScenarioResponse{
		Scenario: r.ToProto(),
	}, nil
}

// ListScenarios ...
func (s *Service) ListScenarios(ctx context.Context, _ *pb.ListScenariosRequest) (*pb.ListScenariosResponse, error) {
	log.Infof("ListScenarios called")
	scenarios, err := s.repo.ListScenarios(ctx)
	if err != nil {
		log.Errorf("could not list scenarios: %v", err)
		return nil, err
	}
	return &pb.ListScenariosResponse{
		Scenarios: slice.Map(scenarios, func(r *repository.Scenario) *pb.Scenario { return r.ToProto() }),
	}, nil
}

// GetScenario ...
func (s *Service) GetScenario(ctx context.Context, in *pb.GetScenarioRequest) (*pb.GetScenarioResponse, error) {
	log.Infof("GetScenario called")
	scenario, err := s.repo.GetScenario(ctx, in.ScenarioId)
	if err != nil {
		log.Errorf("could not get scenario: %v", err)
		return nil, err
	}
	return &pb.GetScenarioResponse{
		Scenario: scenario.ToProto(),
	}, nil
}
