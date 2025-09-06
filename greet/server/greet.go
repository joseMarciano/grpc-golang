package main

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/joseMarciano/grpc/greet/protogen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Greet(ctx context.Context, in *protogen.GreetRequest) (*protogen.GreetResponse, error) {
	if in.FirstName != "John" {
		return nil, status.Error(codes.InvalidArgument, "First name must be John")
	}

	for i := 0; i <= 5; i++ {
		if errors.Is(ctx.Err(), context.Canceled) {
			log.Println("context deadline exceeded 2")
			return nil, status.Error(codes.DeadlineExceeded, "deadline was exceeded test")
		}

		time.Sleep(time.Second)
	}

	return &protogen.GreetResponse{Result: "Hello " + in.FirstName}, nil
}
