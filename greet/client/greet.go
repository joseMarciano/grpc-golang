package main

import (
	"context"
	"log"

	"github.com/joseMarciano/grpc/greet/protogen"
	"google.golang.org/grpc/status"
)

func doGreet(client protogen.GreetServiceClient) {
	clientResponse, err := client.Greet(context.Background(), &protogen.GreetRequest{
		FirstName: "Johna",
	})

	if err != nil {
		if e, ok := status.FromError(err); ok {
			log.Printf("Error message from server: %v", e.Message())
			log.Printf("Error code from server: %v", e.Code())
			return
		}

		log.Fatalf("Error while calling Greet RPC: %v", err)
	}

	log.Printf("Response from Greet: %v", clientResponse.Result)
}
