package main

import (
	"log"
	"net"

	"github.com/linuxfreak003/ballistic-calculator/adapters/proto"
	"github.com/linuxfreak003/ballistic-calculator/domain"
	"github.com/linuxfreak003/ballistic-calculator/pb"
	"google.golang.org/grpc"
)

func main() {
	s := &domain.Service{}
	protoService := proto.NewService(s)

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("could not listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterBallisticServiceServer(grpcServer, protoService)

	log.Printf("server listening at %v", lis.Addr())
	grpcServer.Serve(lis)
}
