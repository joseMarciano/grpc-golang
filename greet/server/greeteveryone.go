package main

import (
	"errors"
	"io"
	"time"

	"github.com/joseMarciano/grpc/greet/protogen"
	"google.golang.org/grpc"
)

func (*Server) GreetEveryone(stream grpc.BidiStreamingServer[protogen.GreetRequest, protogen.GreetResponse]) error {
	var requests []*protogen.GreetRequest
	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return nil
		}

		if err != nil {
			return err
		}

		time.Sleep(time.Second)
		requests = append(requests, req)

		res := &protogen.GreetResponse{Result: "Hello " + req.FirstName}
		if err = stream.Send(res); err != nil {
			return err
		}
	}

}
