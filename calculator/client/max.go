package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/abhiDev28/go-grpc/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Printf("doMax function was invoked!")

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	numbers := []int32{1, 5, 3, 6, 2, 20}

	waitc := make(chan struct{})

	go func() {
		for _, number := range numbers {
			log.Printf("Sending number: %v\n", number)
			stream.Send(&pb.MaxRequest{
				Number: number,
			})
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}

			log.Printf("Received: %v\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
