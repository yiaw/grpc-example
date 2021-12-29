package app

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	v1user "grpc-example/internal/app/v1/user"
)

func NewGRPCServer() *grpc.Server {
	s := grpc.NewServer()
	v1user.NewUserServer(s)
	reflection.Register(s)
	return s
}
