package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/lattots/julius/proto"
)

const port = 8080

func main() {
	conn, err := grpc.NewClient(fmt.Sprintf(":%d", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error connecting to port %d: %v\n", port, err)
	}
	defer conn.Close()

	fmt.Println("Connected!")

	c := pb.NewEventServiceClient(conn)

	fmt.Println("Calling GetEvent...")

	e, err := c.GetEvent(context.Background(), &pb.SingleEventRequest{EventID: 4})
	if err != nil {
		log.Fatalf("error getting event: %v\n", err)
	}
	fmt.Println(e.StartTime.AsTime())
}
