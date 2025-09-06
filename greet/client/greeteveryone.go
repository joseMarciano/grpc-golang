package main

import (
	"context"
	"errors"
	"io"
	"log"

	"github.com/joseMarciano/grpc/greet/protogen"
	"google.golang.org/grpc"
)

func doGreetEveryone(client protogen.GreetServiceClient) {
	stream, err := client.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("%v.GreetEveryone(_) = _, %v", client, err)
	}

	go sendDoGreetEveryone(stream)

	var response *protogen.GreetResponse
	for {
		response, err = stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			log.Fatalf("Failed to receive a response: %v", err)
		}

		log.Printf("Response from GreetEveryone: %v", response.GetResult())
	}
}

func sendDoGreetEveryone(stream grpc.BidiStreamingClient[protogen.GreetRequest, protogen.GreetResponse]) {
	requests := []*protogen.GreetRequest{
		{FirstName: "John"},
		{FirstName: "Jane"},
		{FirstName: "Doe"},
		{FirstName: "Ann"},
	}

	for _, req := range requests {
		err := stream.Send(req)
		if err != nil {
			log.Fatalf("Failed to send a request: %v", err)
		}
	}

	err := stream.CloseSend()
	if err != nil {
		log.Fatalf("Failed to close the stream: %v", err)
	}
}
