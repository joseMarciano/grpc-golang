package main

import (
	"context"
	"log"

	"github.com/joseMarciano/grpc/greet/protogen"
)

func doGreetLong(client protogen.GreetServiceClient) {
	requests := []*protogen.GreetRequest{
		{FirstName: "John"},
		{FirstName: "Jane"},
		{FirstName: "Doe"},
		{FirstName: "Ann"},
	}

	stream, err := client.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("LongGreet error: %v", err)
	}

	for _, req := range requests {
		err = stream.Send(req)

		if err != nil {
			log.Fatalf("Failed to send a request: %v", err)
		}
	}

	var res *protogen.GreetResponse
	if res, err = stream.CloseAndRecv(); err != nil {
		log.Fatalf("CloseAndRecv error: %v", err)
	}

	log.Printf("GreetResponses: %v", res.Result)
}
