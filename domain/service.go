package domain

import (
	"context"
	"fmt"
	"log"

	"github.com/linuxfreak003/ballistic-calculator/pb"
	"github.com/linuxfreak003/ballistic-calculator/ports"
)

// Service ...
type Service struct {
	repo ports.Repository
}

// CreateBullet ...
func (s *Service) CreateBullet(ctx context.Context, in *pb.CreateBulletRequest) (*pb.CreateBulletResponse, error) {
	s.repo.CreateBullet(ctx, in)
	log.Printf("unimplemented in domain")
	return nil, fmt.Errorf("Unimplemented")
}
