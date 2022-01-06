package app

import (
	"crypto/tls"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	v1user "github.com/yiaw/grpc-example/internal/app/v1/user"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair("server-cert.pem", "server-key.pem")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
		//ClientAuth:   tls.RequireAndVerifyClientCert,
		//ClientCAs:    certPool,
	}

	return credentials.NewTLS(config), nil
}

func NewGRPCServer(tlsenable bool) (*grpc.Server, error) {

	/*
		serverOptions := []grpc.ServerOption{
			grpc.UnaryInterceptor(interceptor.Unary()),
			grpc.StreamInterceptor(interceptor.Stream()),
		}
	*/

	var s *grpc.Server
	if tlsenable {
		tlsCredentials, err := loadTLSCredentials()
		if err != nil {
			return nil, fmt.Errorf("cannot load TLS credentials: %w", err)
		}
		s = grpc.NewServer(grpc.Creds(tlsCredentials))
	} else {
		s = grpc.NewServer()
	}

	if s == nil {
		return nil, fmt.Errorf("grpcNewServer Fail..\n")
	}

	v1user.NewUserServer(s)
	// v2user.NewUserServer(s)

	reflection.Register(s)
	return s, nil
}
