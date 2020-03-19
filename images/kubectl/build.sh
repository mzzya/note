#!/bin/bash

#基础镜像
base_image="alpine:3.10.4"
kubectl_version="v1.17.3"

cp ./Template Dockerfile

sed -i "s!\$BASE_IMAGE!hellojqk/${base_image}!g" Dockerfile

#改造后镜像TAG
image_full_name=hellojqk/${base_image}

# alpine 加速 https://blog.csdn.net/freeking101/article/details/80795752
docker build --no-cache --build-arg kubectl_version=${kubectl_version} -t ${image_full_name} .

docker push ${image_full_name}

rm Dockerfile