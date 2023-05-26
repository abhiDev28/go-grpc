package main

import (
	"context"
	"log"

	pb "github.com/abhiDev28/go-grpc/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Printf("createBlog function was invoked!")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Abhishek",
		Title:    "A new title",
		Content:  "Content of the first blogm with some awesome additions!",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("Error happened while updating: %v\n", err)
	}

	log.Println("Blog was updated!")
}
