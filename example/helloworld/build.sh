#! /bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./hello ./main.go
docker build -t hellojqk/helloworld:latest .
docker push hellojqk/helloworld:latest
rm ./hello