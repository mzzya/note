#! /bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=jsoniter -ldflags "-s -w" -o ./app ./main.go

docker build -t hellojqk/httptest:latest .

docker login -u ${REGISTRY_USERNAME} -p ${REGISTRY_PASSWORD} ${REGISTRY_HOST}
docker tag hellojqk/httptest:latest registry.cn-shanghai.aliyuncs.com/clp-test/tr-tracer:httptest

docker push registry.cn-shanghai.aliyuncs.com/clp-test/tr-tracer:httptest

rm ./app
