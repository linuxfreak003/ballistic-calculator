package ports

import (
	"context"

	"github.com/linuxfreak003/ballistic-calculator/pb"
)

// Service ...
type Service interface {
	// Rifle
	CreateRifle(context.Context, *pb.CreateRifleRequest) (*pb.CreateRifleResponse, error)
	ListRifles(context.Context, *pb.ListRiflesRequest) (*pb.ListRiflesResponse, error)

	// Load
	CreateLoad(context.Context, *pb.CreateLoadRequest) (*pb.CreateLoadResponse, error)
	ListLoads(context.Context, *pb.ListLoadsRequest) (*pb.ListLoadsResponse, error)
}
