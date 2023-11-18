package domain

import (
	"context"

	"github.com/linuxfreak003/ballistic-calculator/pb"
	"github.com/linuxfreak003/ballistic-calculator/ports/repository"
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
	load := &repository.Load{}
	load.FromProto(in.GetLoad())
	id, err := s.repo.CreateLoad(ctx, load)
	if err != nil {
		return nil, err
	}
	load.LoadId = id
	return &pb.CreateLoadResponse{
		Load: load.ToProto(),
	}, nil
}
