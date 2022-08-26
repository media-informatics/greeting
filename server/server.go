package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/media-informatics/greeting/service"

	"google.golang.org/grpc"
)

const port = 54321

type server struct {
	service.UnimplementedGreetingServer
}

// SayHello implements GreetingServer
func (s *server) SayHello(ctx context.Context, in *service.HelloRequest) (*service.HelloResponse, error) {
	log.Printf("Received: %v", in.GetRequest())
	return &service.HelloResponse{Response: "Hello " + in.GetRequest()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	service.RegisterGreetingServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
