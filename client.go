package main

import (
	"context"
	"log"
	"rpcProject/chat"

	"google.golang.org/grpc"
)

func main() {
	var connection *grpc.ClientConn

	connection, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("couldn't connect: %s", err)
	}
	defer connection.Close()

	c := chat.NewChatServiceClient(connection)

	message := chat.Message{
		Body: "Hi, I am client",
	}

	resp, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from server: %s", resp.Body)
}
