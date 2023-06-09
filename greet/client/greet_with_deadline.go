package main

import (
	"log"
	"time"

	pb "github.com/abhiDev28/go-grpc/greet/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Printf("doGreetWithDeadline function was invoked!")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := pb.GreetRequest{
		FirstName: "Abhishek",
	}

	res, err := c.GreetWithDeadline(ctx, &req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded !")
				return
			} else {
				log.Fatalf("Unexpected gRPC error: %v\n", err)
			}
		} else {
			log.Fatalf("A non gRPC error: %v\n", err)
		}
	}
	log.Printf("GreetWithDeadline: %s\n", res.Result)
}
