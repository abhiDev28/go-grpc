package main

import (
	"context"
	"log"
	"time"

	pb "github.com/abhiDev28/go-grpc/calculator/proto"
)

func doAverage(c pb.CalculatorServiceClient) {
	log.Println("doAverage function invoked!")

	reqs := []*pb.AverageRequest{
		{Number: 10},
		{Number: 20},
		{Number: 30},
	}

	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Average: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending number: %v\n", req.Number)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from Average: %v\n", err)
	}

	log.Printf("Average: %v\n", res.Result)
}
