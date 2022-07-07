package main

import (
	"log"
	"net"

	"github.com/wenealves10/golang-grpc/pb"
	"github.com/wenealves10/golang-grpc/servers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, &servers.User{})
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", ":4005")

	if err != nil {
		log.Fatalf("error %v", err)
	}

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatalf("error %v", err)
	}

}
