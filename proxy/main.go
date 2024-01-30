package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/linuxfreak003/ballistic-calculator/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

var (
	ballisticEndpoint = flag.String("ballistic-endpoint", "localhost:50051", "endpoint of ballistic server")
	port              = flag.String("port", "8080", "port to listen on")
)

func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

// preflightHandler adds the necessary headers in order to serve
// CORS from any origin using the methods "GET", "HEAD", "POST", "PUT", "DELETE"
// We insist, don't do this without consideration in production systems.
func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept", "Authorization"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	log.Infof("preflight request for %s", r.URL.Path)
}

// RunEndPoint ...
func RunEndPoint(address string, opts ...runtime.ServeMuxOption) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithTimeout(time.Second * 10),
	}
	err := pb.RegisterBallisticServiceHandlerFromEndpoint(ctx, mux, *ballisticEndpoint, dialOpts)
	if err != nil {
		return err
	}

	log.Printf("Listening on %v", address)
	http.ListenAndServe(address, allowCORS(mux))
	return nil
}

func main() {
	flag.Parse()

	if err := RunEndPoint(":" + *port); err != nil {
		log.Fatal(err)
	}
}
