package grpc

import (
	"context"
	"log"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func customMiddleware() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		log.Print("Requested at:", time.Now())
		resp, err := handler(ctx, req)
		log.Print("Response at:", time.Now())
		return resp, err
	}
}

func NewChainMiddleware() grpc.UnaryServerInterceptor {
	logrus.ErrorKey = "grpc.error"
	logrusEntry := logrus.NewEntry(logrus.StandardLogger())
	middleware := grpc_middleware.ChainUnaryServer(
		grpc_ctxtags.UnaryServerInterceptor(),
		grpc_logrus.UnaryServerInterceptor(logrusEntry),
		customMiddleware(),
		grpc_recovery.UnaryServerInterceptor(),
	)

	return middleware
}
