package main

import (
	"context"
	"log"

	pb "github.com/abhiDev28/go-grpc/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Printf("doSum function was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  10,
		SecondNumber: 20,
	})
	if err != nil {
		log.Fatalf("Could not do sum: %v\n", err)
	}
	log.Printf("Sum is :%v\n", res.Result)
}
