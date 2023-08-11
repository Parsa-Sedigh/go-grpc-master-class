package main

import (
	"context"
	"fmt"
	"github.com/Parsa-Sedigh/go-grpc-master-class/blog/blogpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"log/slog"
	"os"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

func main() {
	fmt.Println("Blog Client")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	// create Blog
	fmt.Println("Creating the blog...")

	blog := &blogpb.Blog{
		AuthorId: "Parsa",
		Title:    "First blog",
		Content:  "content of the first blog",
	}

	createBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected err: %v", err)
	}

	fmt.Printf("Blog has been created: %v\n", createBlogRes)

	blogID := createBlogRes.Blog.GetId()

	// read Blog
	fmt.Println("Reading the blog")

	_, err = c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "sdsds"})
	if err != nil {
		fmt.Printf("Error happened while reading: %v \n", err)
	}

	// we're gonna read exactly what we just created, yeah it doesn't make sense but this is just for learning
	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogID}

	readBlogRes, err := c.ReadBlog(context.Background(), readBlogReq)
	if err != nil {
		fmt.Printf("Error happened while reading: %v \n", err)
	}

	fmt.Printf("Blog was read: %v", readBlogRes)

	// update Blog
	blog = &blogpb.Blog{
		Id:       blogID,
		AuthorId: "Parsa 2",
		Title:    "Second blog",
		Content:  "content of the second blog",
	}
	updateBlogRes, err := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: blog})
	if err != nil {
		logger.Error("Error happened while updating: ", err, "\n")
	}

	logger.Info("Blog was read: ", updateBlogRes)

	// delete Blog
	deleteRes, err := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{BlogId: blogID})
	if err != nil {
		logger.Error("Error happened while deleting: ", err, "\n")
	}

	logger.Info("Blog was deleted: ", deleteRes, "\n")

	// list Blogs

	// the request is empty(it's ok)
	stream, err := c.ListBlog(context.Background(), &blogpb.ListBlogRequest{})
	if err != nil {
		logger.Error("error while calling ListBlog RPC: ", err, "\n")
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Error("error while calling ListBlog RPC: ", err, "\n")
		}

		fmt.Println(res.GetBlog())
	}
}
