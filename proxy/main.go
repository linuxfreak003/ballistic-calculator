package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/linuxfreak003/ballistic-calculator/pb"
	"google.golang.org/grpc"
)

var (
	ballisticEndpoint = flag.String("ballistic-endpoint", "localhost:50051", "endpoint of ballsitic")
)

func RunEndPoint(address string, opts ...runtime.ServeMuxOption) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterBallisticServiceHandlerFromEndpoint(ctx, mux, *ballisticEndpoint, dialOpts)
	if err != nil {
		return err
	}

	log.Printf("Listening on %v", address)
	http.ListenAndServe(address, mux)
	return nil
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := RunEndPoint(":8080"); err != nil {
		glog.Fatal(err)
	}
}
