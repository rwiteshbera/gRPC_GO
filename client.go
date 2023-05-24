package main

import (
	"context"
	"log"
	"rpcProject/pb"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	var connection *grpc.ClientConn

	connection, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("couldn't connect: %s", err)
	}
	defer connection.Close()

	c := pb.NewLaptopGuideClient(connection)

	resp, err := c.CreateNew(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error when calling: %s", err)
	}

	log.Printf("Response from server: %s", resp)
}
