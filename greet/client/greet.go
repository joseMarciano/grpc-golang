package main

import (
	"context"
	"log"

	"github.com/joseMarciano/grpc/greet/protogen"
)

func doGreet(client protogen.GreetServiceClient) {
	clientResponse, err := client.Greet(context.Background(), &protogen.GreetRequest{
		FirstName: "John",
	})

	if err != nil {
		log.Fatalf("Greet(_) = _, %v", err)
	}

	log.Printf("Response from Greet: %v", clientResponse.Result)
}
