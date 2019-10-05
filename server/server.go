package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/eschizoid/go-grpc/handler"
	pb "github.com/eschizoid/go-grpc/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func serveSwagger(response http.ResponseWriter, request *http.Request) {
	if !strings.HasSuffix(request.URL.Path, ".swagger.json") {
		log.Printf("Not found: %s", request.URL.Path)
		http.NotFound(response, request)
		return
	}
	http.ServeFile(response, request, "/go/src/github.com/eschizoid/go-grpc/proto/ingestgw.swagger.json")
}

func startRESTServer(address, grpcAddress string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	var opts = []grpc.DialOption{
		grpc.WithInsecure(),
	}
	if err := pb.RegisterIngestHandlerFromEndpoint(ctx, mux, grpcAddress, opts); err != nil {
		return fmt.Errorf("could not register service Ingest: %s", err)
	}
	r := http.NewServeMux()
	r.HandleFunc("/swagger/", serveSwagger)
	r.Handle("/", mux)
	log.Printf("starting HTTP/1.1 REST server on %s", address)
	if err := http.ListenAndServe(address, r); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
	return nil
}

func startGRPCServer() error {
	listener, err := net.Listen("tcp", ":5300")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterIngestServer(grpcServer, &handler.IngestServer{})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
	return nil
}

func main() {
	grpcAddress := fmt.Sprintf("%s:%d", "go-grpc", 5300)
	restAddress := fmt.Sprintf("%s:%d", "go-grpc", 8081)
	go func() {
		if err := startGRPCServer(); err != nil {
			log.Fatalf("failed to start gRPC server: %s", err)
		}
	}()
	go func() {
		if err := startRESTServer(restAddress, grpcAddress); err != nil {
			log.Fatalf("failed to start gRPC server: %s", err)
		}
	}()
	select {}
}
