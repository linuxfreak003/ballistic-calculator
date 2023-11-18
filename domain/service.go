package domain

import (
	"context"

	"github.com/linuxfreak003/ballistic-calculator/pb"
	"github.com/linuxfreak003/ballistic-calculator/ports"
)

// Service ...
type Service struct {
	repo ports.Repository
}

// CreateLoad ...
func (s *Service) CreateLoad(ctx context.Context, in *pb.CreateLoadRequest) (*pb.CreateLoadResponse, error) {
	id, err := s.repo.CreateLoad(ctx, in)
	if err != nil {
		return nil, err
	}
	in.LoadId = id
	return in
}
