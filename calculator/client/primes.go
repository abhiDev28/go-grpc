package main

import (
	"context"
	"io"
	"log"

	pb "github.com/abhiDev28/go-grpc/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Printf("doPrimes function was invoked")
	req := &pb.PrimeRequest{
		Number: 12390392840,
	}
	res, err := c.Primes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Primes: %v\n", err)
	}

	for {
		msg, err := res.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}
		log.Println(msg.Result)
	}

}
