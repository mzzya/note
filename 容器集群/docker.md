# docker笔记

## docker镜像与docker in docker（dind）镜像区别

- docker镜像：启动时是不包含`服务端进程`，包含客户端程序。
- dind镜像：需要使用`--privileged`特权模式启动，包含`服务端进程`和客户端程序。相当于你在mac、win或linux上安装的docker镜像，正常运行docker命令。CI/CD过程中使用不会缓存基础镜像，每次都会获取。

如果映射 /var/run/docker.sock 则是与主机通信

## overlay2

## 启动 mysql

```shell script
# 第一次启动
docker run --restart=always --name mysql -p 3306:3306 -p 33060:33060 -e MYSQL_ROOT_PASSWORD=123123 -v ~/DockerVolumes/mysql:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123123 -d mysql:8.0.20 --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci
# 再次启动
docker run --restart=always --name mysql -p 3306:3306 -p 33060:33060 -e MYSQL_ROOT_PASSWORD=123123 -v ~/DockerVolumes/mysql:/var/lib/mysql -d mysql:8.0.20
```

```shell
docker run --restart=always --name mongo -p 27017:27017 -v ~/DockerVolumes/mongo:/data/db -d mongo:4.2.3
```
