#! /bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./app ./main.go
docker build --no-cache -t hellojqk/opentracing:latest .
# docker login -u hellojqk
docker push hellojqk/opentracing:latest
rm ./app
docker rmi $(docker image ls --filter="dangling=true" --format="{{.ID}}")
sleep 10