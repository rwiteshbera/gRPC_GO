package main

import (
	"context"
	"log"
	"net"
	"rpcProject/pb"
	"rpcProject/utils"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type LaptopGuide struct {
	pb.LaptopGuideServer
}

func (l *LaptopGuide) CreateNew(ctx context.Context, in *emptypb.Empty) (*pb.Laptop, error) {
	newLaptop := utils.NewLaptop()
	log.Println(newLaptop)
	return newLaptop, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterLaptopGuideServer(grpcServer, &LaptopGuide{})

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start gRPC server: %v", err)
	}
}
