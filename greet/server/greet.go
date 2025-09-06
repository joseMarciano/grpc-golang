package main

import (
	"context"

	"github.com/joseMarciano/grpc/greet/protogen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Greet(_ context.Context, in *protogen.GreetRequest) (*protogen.GreetResponse, error) {
	if in.FirstName != "John" {
		return nil, status.Error(codes.InvalidArgument, "First name must be John")
	}

	return &protogen.GreetResponse{Result: "Hello " + in.FirstName}, nil
}
