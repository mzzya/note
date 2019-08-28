#! /bin/bash

minikube start --memory=16384 --cpus=4 --registry-mirror=https://registry.docker-cn.com --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers --alsologtostderr