package main

import (
	"context"

	"github.com/joseMarciano/grpc/greet/protogen"
)

func (s *Server) Greet(_ context.Context, in *protogen.GreetRequest) (*protogen.GreetResponse, error) {
	return &protogen.GreetResponse{Result: "Hello " + in.FirstName}, nil
}
