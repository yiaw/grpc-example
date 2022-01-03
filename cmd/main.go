package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"grpc-example/cmd/app"
)

func main() {

	tlsenable := flag.Bool("tls", false, "enable SSL/TLS, def false")
	port := flag.Int("port", 50001, "Server Port Num, def 50001")

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen; %s\n", err.Error())
	}

	s, err := app.NewGRPCServer(*tlsenable)
	if err != nil {
		log.Fatalf("failed New GRPCServer() .. err=%s\n", err.Error())
	}

	log.Printf("start gRPC Server on %d port, enableTLS=%t\n", *port, *tlsenable)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server :%s\n", err.Error())
	}
}
