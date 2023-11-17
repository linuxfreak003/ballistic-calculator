package ports

import (
	"context"

	"github.com/linuxfreak003/ballistic-calculator/pb"
)

type Service interface {
	CreateBullet(context.Context, *pb.CreateBulletRequest) (*pb.CreateBulletResponse, error)
}
