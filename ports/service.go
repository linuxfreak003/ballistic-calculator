package ports

import (
	"context"

	"github.com/linuxfreak003/ballistic-calculator/pb"
)

// Service ...
type Service interface {
	CreateLoad(context.Context, *pb.CreateLoadRequest) (*pb.CreateLoadResponse, error)
}
