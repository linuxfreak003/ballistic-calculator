package domain

import (
	"context"

	"github.com/linuxfreak003/ballistic-calculator/pb"
	"github.com/linuxfreak003/ballistic-calculator/ports/repository"
	"github.com/linuxfreak003/util/slice"
)

// Service ...
type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
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
func (s *Service) ListLoads(ctx context.Context, in *pb.ListLoadsRequest) (*pb.ListLoadsResponse, error) {
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
	var rifle *repository.Rifle
	rifle = rifle.FromProto(in.GetRifle())

	id, err := s.repo.CreateRifle(ctx, rifle)
	if err != nil {
		return nil, err
	}
	rifle.RifleId = id
	return &pb.CreateRifleResponse{
		Rifle: rifle.ToProto(),
	}, nil
}

// ListRifles ...
func (s *Service) ListRifles(ctx context.Context, in *pb.ListRiflesRequest) (*pb.ListRiflesResponse, error) {
	rifles, err := s.repo.ListRifles(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.ListRiflesResponse{
		Rifles: slice.Map(rifles, func(r *repository.Rifle) *pb.Rifle { return r.ToProto() }),
	}, nil
}
