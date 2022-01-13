# Makefile

all:
	protoc -I=. \
        -I=$(GOPATH)/src/grpc-example/protos \
        --go_out $(GOPATH)/src/grpc-example/protos/v1/chat \
        --go_opt paths=source_relative \
        --go-grpc_out $(GOPATH)/src/grpc-example/protos/v1/chat \
        --go-grpc_opt paths=source_relative \
        chat.proto;
clean:
	rm -rf *.go
