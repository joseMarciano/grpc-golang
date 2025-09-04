package main

import (
	"fmt"

	"github.com/joseMarciano/grpc/greet/protogen"
	"google.golang.org/grpc"
)

func (s *Server) GreetManyTimes(in *protogen.GreetRequest, stream grpc.ServerStreamingServer[protogen.GreetResponse]) error {
	var err error
	for i := 0; i < 2; i++ {
		res := fmt.Sprintf("Hello %s number %d", in.GetFirstName(), i)
		if err = stream.Send(&protogen.GreetResponse{Result: res}); err != nil {
			return err
		}
	}
	return err
}
