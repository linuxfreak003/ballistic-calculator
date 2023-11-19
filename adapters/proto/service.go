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
