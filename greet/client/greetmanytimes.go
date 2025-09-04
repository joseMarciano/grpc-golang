package main

import (
	"context"
	"errors"
	"io"
	"log"

	"github.com/joseMarciano/grpc/greet/protogen"
)

func doGreetManyTimes(client protogen.GreetServiceClient) {
	responseStream, err := client.GreetManyTimes(context.Background(), &protogen.GreetRequest{
		FirstName: "John",
	})

	if err != nil {
		log.Fatalf("Greet(_) = _, %v", err)
	}

	for {
		var msg *protogen.GreetResponse
		msg, err = responseStream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			log.Fatalf("GreetManyTimes(_) = _, %v", err)
		}

		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}

}
