package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/surzia/grpc-starter/protobuf"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type Server struct {
	protobuf.UnimplementedAttractionServiceServer
}

func (s *Server) GetAttractions(ctx context.Context, req *protobuf.Request) (*protobuf.Response, error) {
	log.Printf("Request attractions that contains %s", req.GetName())
	return &protobuf.Response{Atrtactions: []*protobuf.Attraction{}}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	protobuf.RegisterAttractionServiceServer(s, &Server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
