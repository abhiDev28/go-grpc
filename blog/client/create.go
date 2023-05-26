package main

import (
	"context"
	"log"

	pb "github.com/abhiDev28/go-grpc/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Printf("createBlog function was invoked!")

	blog := pb.Blog{
		AuthorId: "Abhishek",
		Title:    "My-first-blog",
		Content:  "Content of first blog",
	}

	res, err := c.CreateBlog(context.Background(), &blog)
	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %s\n", res.Id)
	return res.Id

}
