package main

import (
	"log"
	"net"

	"github.com/joseMarciano/grpc/greet/protogen"
	"google.golang.org/grpc"
)

type Server struct {
	protogen.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	protogen.RegisterGreetServiceServer(s, &Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
