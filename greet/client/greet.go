package main

import (
	"context"
	"log"
	"time"

	"github.com/joseMarciano/grpc/greet/protogen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreet(client protogen.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	clientResponse, err := client.Greet(ctx, &protogen.GreetRequest{
		FirstName: "John",
	})

	if err != nil {
		if e, ok := status.FromError(err); ok {
			if e.Code() == codes.InvalidArgument {
				log.Printf("Client received an error: %v with message: %v", e.Code(), e.Message())
			}

			if e.Code() == codes.DeadlineExceeded {
				log.Printf("Client deadline exceeded: %v with message: %v", e.Code(), e.Message())
			}

			return
		}

		log.Fatalf("Error while calling Greet RPC: %v", err)
	}

	log.Printf("Response from Greet: %v", clientResponse.Result)
}
