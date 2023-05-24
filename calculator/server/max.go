package main

import (
	"io"
	"log"

	pb "github.com/abhiDev28/go-grpc/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Printf("Max function was invoked!")
	var maxNumber int32 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}

		if number := req.Number; number > maxNumber {
			maxNumber = number
			err = stream.Send(&pb.MaxResponse{
				Result: maxNumber,
			})

			if err != nil {
				log.Fatalf("Error while sending data to client: %v\n", err)
			}
		}

	}
}
