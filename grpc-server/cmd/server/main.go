package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

const portNumber = "9000"

func main() {
	lis, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		log.Fatalf("failed to listen; %s\n", err.Error())
	}

	grpcServer := grpc.NewServer()
	log.Printf("start gRPC Server on %s port\n", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server :%s\n", err.Error())
	}
}
