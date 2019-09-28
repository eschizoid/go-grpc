package main

import (
	"context"
	"fmt"

	pb "github.com/eschizoid/go-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial("go-grpc:5300", opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewIngestClient(conn)
	request := &pb.Request{Message: "ping",}
	response, err := client.Do(context.Background(), request)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	fmt.Println(response.Message)
}
