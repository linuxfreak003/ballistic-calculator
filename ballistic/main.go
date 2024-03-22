package main

import (
	"log"
	"net"

	"github.com/linuxfreak003/ballistic-calculator/adapters/proto"
	"github.com/linuxfreak003/ballistic-calculator/adapters/sqlite"
	"github.com/linuxfreak003/ballistic-calculator/domain"
	"github.com/linuxfreak003/ballistic-calculator/pb"
	"google.golang.org/grpc"
)

func main() {
	r, err := sqlite.NewRepository("sql.db")
	if err != nil {
		log.Fatalf("could not create repository: %v", err)
	}
	d := domain.NewService(r)
	protoService := proto.NewService(d)

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("could not listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterBallisticServiceServer(grpcServer, protoService)

	log.Printf("server listening at %v", lis.Addr())
	grpcServer.Serve(lis)
}
