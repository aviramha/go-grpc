package main

import (
	pb "bookshop/client/pb/inventory"
	"context"
	"log"
       "time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	
        for {
          conn, err := grpc.Dial("bookserver-svc:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewInventoryClient(conn)
	bookList, err := client.GetBookList(context.Background(), &pb.GetBookListRequest{})
	if err != nil {
		log.Fatalf("failed to get book list: %v", err)
	}
	log.Printf("book list: %v", bookList)
        time.Sleep(1 * time.Second)
        }
}
