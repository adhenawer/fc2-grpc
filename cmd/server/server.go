package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	pb "github.com/adhenawer/fc2-grpc/pb"
	services "github.com/adhenawer/fc2-grpc/services"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
}