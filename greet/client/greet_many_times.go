package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/abhiDev28/go-grpc/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	req := pb.GreetRequest{
		FirstName: "Abhishek",
	}

	stream, err := c.GreetManyTimes(context.Background(), &req)

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		fmt.Printf("GreetManyTimes: %s\n", msg.Result)
	}
}
