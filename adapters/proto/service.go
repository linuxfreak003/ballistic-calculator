package proto

import (
	"context"
	"log"

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

// CreateBullet ...
func (s *Service) CreateBullet(ctx context.Context, in *pb.CreateBulletRequest) (*pb.CreateBulletResponse, error) {
	log.Printf("create_bullet called")
	return s.domain.CreateBullet(ctx, in)
}
