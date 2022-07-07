package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/surzia/grpc-starter/conf"
	"github.com/surzia/grpc-starter/protobuf"
	"github.com/surzia/grpc-starter/service"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
	srv  *service.NewService
)

type Server struct {
	protobuf.UnimplementedNewServiceServer
}

func (s *Server) GetHotTopNews(ctx context.Context, req *protobuf.Request) (*protobuf.News, error) {
	ret := srv.RequestPublishAPI()
	return &protobuf.News{
		News: ret,
	}, nil
}

func main() {
	// load config
	conf, err := conf.Load("../../conf/grpc.conf")
	if err != nil {
		panic(err)
	}

	// register new service for http request
	srv = service.NewNewService(conf)

	// register grpc service
	s := grpc.NewServer()
	protobuf.RegisterNewServiceServer(s, &Server{})

	// listen tcp connection
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// start grpc server
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
