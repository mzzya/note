#!/bin/bash

#基础镜像
# base_image="hellojqk/alpine:3.12.3"
base_image="openjdk:8u282-buster"

cp ./Template Dockerfile

sed -i "s!\$BASE_IMAGE!${base_image}!g" Dockerfile

version="8u282-buster"

#改造后镜像TAG
image_full_name="hellojqk/openjdk:${version}"

# alpine 加速 https://blog.csdn.net/freeking101/article/details/80795752
docker build --no-cache --build-arg kubectl_version=${version} -t ${image_full_name} .

docker push ${image_full_name}

rm Dockerfile
