package main

import (
	"log"

	"github.com/joseMarciano/grpc/greet/protogen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := protogen.NewGreetServiceClient(conn)
	doGreet(c)
	doGreetManyTimes(c)
	doGreetLong(c)
}
