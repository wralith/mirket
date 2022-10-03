#!bin/bash

PATH=$PATH:$GOPATH/bin
protodir=../../pb

protoc --go_out=./pb \
  --go-grpc_out=./pb \
  --go_opt=paths=source_relative \
  --go-grpc_opt=paths=source_relative \
  -I=$protodir \
    $protodir/user.proto