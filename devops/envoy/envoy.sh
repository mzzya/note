CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=jsoniter -ldflags "-s -w" -o ./app ./main.go

docker network create envoy

docker rm -f envoy_app envoy_app_1 envoy
docker build -t envoy_app:v1 -f ./dockerfile_app .
docker run -itd --name envoy_app -p 1234:1234 --network envoy envoy_app:v1 --port 1234
docker run -itd --name envoy_app_1 -p 1235:1235 --network envoy envoy_app:v1 --port 1235

docker rm -f envoy
docker rmi envoy:v1

docker build -t envoy:v1 .
docker run -it --name envoy -p 9901:9901 -p 10000:10000 -p 10001:10001 -v /mnt/c/git/note/devops/envoy/envoy.yaml:/etc/envoy/envoy.yaml --network envoy envoy:v1 -c /etc/envoy/envoy.yaml
docker run -it --name envoy_1 envoy:v1