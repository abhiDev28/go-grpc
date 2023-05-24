package main

import (
	"io"
	"log"

	pb "github.com/abhiDev28/go-grpc/calculator/proto"
)

func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {
	log.Println("Average function was invoked!")

	var sum int32 = 0
	count := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{
				Result: float64(sum) / float64(count),
			})
		}
		if err != nil {
			log.Panicf("Error while reading client stream: %v\n", err)
		}
		log.Printf("Receving number: %v\n", req.Number)
		sum += req.Number
		count++
	}

}
