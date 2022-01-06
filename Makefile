# Makefile

.PHONY: v1user
v1user:
	protoc -I=./protos \
	    --go_out ./protos \
	    --go_opt paths=source_relative \
	    --go-grpc_out ./protos \
		--go-grpc_opt paths=source_relative \
	    protos/v1/user/user.proto

.PHONY: v2user
v2user:
	protoc -I=./protos \
	    --go_out ./protos \
        --go_opt paths=source_relative \
        --go-grpc_out ./protos \
        --go-grpc_opt paths=source_relative \
	    protos/v2/user/user.proto

.PHONY: v2user-gateway
v2user-gateway:
	protoc -I ./protos \
	    --grpc-gateway_out ./protos \
	    --grpc-gateway_opt logtostderr=true \
	    --grpc-gateway_opt paths=source_relative \
	    protos/v2/user/user.proto

.PHONY: v2user-all
v2user-all:
	protoc -I=./protos \
		--go_out ./protos \
    	--go_opt paths=source_relative \
    	--go-grpc_out ./protos \
		--go-grpc_opt paths=source_relative \
	    --grpc-gateway_out ./protos \
	    --grpc-gateway_opt logtostderr=true \
	    --grpc-gateway_opt paths=source_relative \
	    protos/v2/user/user.proto