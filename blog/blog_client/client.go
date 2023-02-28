package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mohdjishin/gRPC/blog/blogpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello,I'm a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	c := blogpb.NewBlogServiceClient(conn)

	blog := &blogpb.Blog{
		AuthorId: "Jishin",
		Title:    "My first blog",
		Content:  "Content of the first blog",
	}
	fmt.Println("Creating the blog")
	createBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{
		Blog: blog,
	})

	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}

	fmt.Printf("Blog has been created %v", createBlogRes)

}
