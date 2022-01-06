package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/yiaw/grpc-example/cmd/app"
)

func main() {

	tlsenable := flag.Bool("tls", false, "enable SSL/TLS, def false")
	grpcport := flag.Int("port", 8090, "GRPC Server Port Num, def 8090")
	httpport := flag.Int("http", 8080, "HTTP Gateway Port Num, def 8080")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcport))
	if err != nil {
		log.Fatalf("failed to listen; %s\n", err.Error())
	}

	s, err := app.NewGRPCServer(*tlsenable)
	if err != nil {
		log.Fatalf("failed New GRPCServer() .. err=%s\n", err.Error())
	}

	log.Printf("start gRPC Server on %d port, enableTLS=%t\n", *grpcport, *tlsenable)

	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	gw, err := app.NewGateWay(*grpcport)
	if err != nil {
		log.Fatalf("failed NewGateWay() .. err=%s\n", err.Error())
	}

	httpServ := &http.Server{
		Addr:    fmt.Sprintf(":%d", *httpport),
		Handler: gw,
	}

	log.Printf("HTTP Server GRPC Gateway on http://0.0.0.0:%d", *httpport)

	log.Fatalln(httpServ.ListenAndServe())
}
