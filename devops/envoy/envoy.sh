CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=jsoniter -ldflags "-s -w" -o ./app ./main.go

docker rm -f envoy_app envoy_app_1 envoy_app_2 envoy
APP_IMAGE_TAG=BBBB
docker build -t hellojqk/envoy_app:v2 --build-arg APP_IMAGE_TAG=${APP_IMAGE_TAG} -f ./dockerfile_app .
docker run -it --name envoy_app_2 hellojqk/envoy_app:v2
# docker push hellojqk/envoy_app:v2
# docker run -itd --name envoy_app -p 1234:1234 --network net envoy_app:v1 --port 1234
# docker run -itd --name envoy_app_1 -p 1235:1235 --network net envoy_app:v1 --port 1235

# docker rm -f envoy
# docker rmi envoy:v1

# docker build -t envoy:v1 .
# docker run -it --name envoy -p 9901:9901 -p 10000:10000 -p 10001:10001 -v `pwd`/envoy.yaml:/etc/envoy/envoy.yaml --network net envoy:v1 -c /etc/envoy/envoy.yaml
