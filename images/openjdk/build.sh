#!/bin/bash

#基础镜像
# base_image="hellojqk/alpine:3.12.3"
base_image="adoptopenjdk/openjdk8:jdk8u252-b09-alpine"

cp ./Template Dockerfile

sed -i "s!\$BASE_IMAGE!${base_image}!g" Dockerfile

version="jdk8u252-b09-alpine"

#改造后镜像TAG
image_full_name="hellojqk/adoptopenjdk8:${version}"

# alpine 加速 https://blog.csdn.net/freeking101/article/details/80795752
docker build --no-cache --build-arg kubectl_version=${version} -t ${image_full_name} .

docker push ${image_full_name}

rm Dockerfile
