#！/bin/bash

docker build -t hellojqk/nginx:1.19.8-alpine -f ./Dockerfile_base .
docker push hellojqk/nginx:1.19.8-alpine
