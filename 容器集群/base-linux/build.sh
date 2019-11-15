# alpine 加速 https://blog.csdn.net/freeking101/article/details/80795752
version="3.10.3"
docker build -t hellojqk/alpine:${version} .
docker push hellojqk/alpine:${version}