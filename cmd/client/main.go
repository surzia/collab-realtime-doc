package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/surzia/grpc-starter/conf"
	"github.com/surzia/grpc-starter/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	// load config
	conf, err := conf.Load("../../conf/grpc.conf")
	if err != nil {
		panic(err)
	}

	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := protobuf.NewNewServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetHotTopNews(ctx, &protobuf.Request{
		Type:     conf.NewsConf.Type,
		Page:     int64(conf.NewsConf.Page),
		Size:     int64(conf.NewsConf.Size),
		IsFilter: int64(conf.NewsConf.IsFilter),
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	for _, v := range r.GetNews() {
		fmt.Println(v)
	}
}
