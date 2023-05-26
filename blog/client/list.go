package main

import (
	"context"
	"io"
	"log"

	pb "github.com/abhiDev28/go-grpc/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlogs(c pb.BlogServiceClient) {
	log.Printf("listBlogs function was invoked!!")

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error while reading stream: %v\n", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error in listBlogs: %v\n", err)
		}

		log.Println(res)
	}

}
