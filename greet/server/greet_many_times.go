package main

import (
	"fmt"
	"log"

	pb "github.com/abhiDev28/go-grpc/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with: %v", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.FirstName, i)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}
	return nil
}
