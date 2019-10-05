package handler

import (
	"log"

	pb "github.com/eschizoid/go-grpc/proto"
	"golang.org/x/net/context"
)

type IngestServer struct{}

func (s *IngestServer) Do(c context.Context, request *pb.Request) (response *pb.Response, err error) {
	log.Printf("Payload: %s", request.GetMessage())
	response = &pb.Response{Message: "pong",}
	return response, nil
}
