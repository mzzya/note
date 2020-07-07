CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=jsoniter -ldflags "-s -w" -o ./app ./main.go
docker rm -f envoy_app
docker build -t envoy_app:v1 -f ./dockerfile_app .
docker run -itd --name envoy_app -p 1234:1234 envoy_app:v1

docker rm -f envoy
docker rmi envoy:v1

docker build -t envoy:v1 .
docker run -it --name envoy -p 9901:9901 -p 10000:10000 -p 10001:10001 envoy:v1
