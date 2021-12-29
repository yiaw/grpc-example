package main

import (
	"grpc-example/cmd/app"
	"log"
	"net"
)

const portNumber = "12345"

func main() {
	lis, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		log.Fatalf("failed to listen; %s\n", err.Error())
	}

	s := app.NewUserServer()
	log.Printf("start gRPC Server on %s port\n", portNumber)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server :%s\n", err.Error())
	}
}