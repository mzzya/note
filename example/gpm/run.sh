#! /bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./app ./main.go

docker build -t gpm .

rm ./app

docker run -it -c 2 --cpus 0.1 gpm
