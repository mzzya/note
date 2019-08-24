#! /bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=jsoniter -ldflags "-s -w" -o ./app ./main.go

docker build -t hellojqk/grpc-api:0.0.1 .
docker push hellojqk/grpc-api:0.0.1

rm ./app