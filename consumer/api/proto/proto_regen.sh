#!/bin/bash

#install -> go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
#install -> go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

export PATH="$PATH:$(go env GOPATH)/bin"

# add file name to generate the new proto

PROTO_PATH="../api/proto/v1/"
PROTO_FILE="payments"

protoc --go_out=${PROTO_PATH} --go_opt=paths=source_relative \
    --go-grpc_out=${PROTO_PATH} --go-grpc_opt=paths=source_relative \
    --proto_path=${PROTO_PATH} ${PROTO_PATH}${PROTO_FILE}.proto
