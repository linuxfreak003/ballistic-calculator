package domain

import (
	"context"

	"github.com/linuxfreak003/ballistic-calculator/pb"
)

// Service ...
type Service struct{}

// CreateBullet ...
func (*Service) CreateBullet(context.Context, *pb.CreateBulletRequest) (*pb.CreateBulletResponse, error) {
	return nil, nil
}
