package domain

import (
	"context"
	"fmt"
	"log"

	"github.com/linuxfreak003/ballistic-calculator/pb"
)

// Service ...
type Service struct{}

// CreateBullet ...
func (*Service) CreateBullet(context.Context, *pb.CreateBulletRequest) (*pb.CreateBulletResponse, error) {
	log.Printf("unimplemented in domain")
	return nil, fmt.Errorf("Unimplemented")
}
