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

// CreateLoad ...
func (s *Service) CreateLoad(ctx context.Context, in *pb.CreateLoadRequest) (*pb.CreateLoadResponse, error) {
	return s.domain.CreateLoad(ctx, in)
}
