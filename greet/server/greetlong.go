package main

import (
	"errors"
	"io"
	"strings"

	"github.com/joseMarciano/grpc/greet/protogen"
	"google.golang.org/grpc"
)

func (*Server) LongGreet(stream grpc.ClientStreamingServer[protogen.GreetRequest, protogen.GreetResponse]) error {
	var names []string
	for {
		res, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return err
		}

		names = append(names, res.FirstName)
	}

	return stream.SendAndClose(&protogen.GreetResponse{Result: "Hello " + strings.Join(names, ", ")})
}
