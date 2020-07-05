docker rm -f envoy
docker rmi envoy:v1
docker build -t envoy:v1 .
docker run -it --name envoy -p 9901:9901 -p 10000:10000 -p 10001:10001 envoy:v1
