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

## [我的收藏](https://hub.docker.com/u/hellojqk/starred)

- [grafana](https://hub.docker.com/r/grafana/grafana)
  - docker run -d -p 3000:3000 --name=grafana grafana/grafana:6.1.6
- [prometheus](https://hub.docker.com/r/prom/prometheus)
  - docker run -p 9090:9090 -v /tmp/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus:v2.8.2
- [opentracing](https://github.com/opentracing/opentracing-go)
  - 这个项目是一个规范定义了跟踪信息的处理规范，注入与提取，开启方式
- [jaeger-client-go](https://github.com/jaegertracing/jaeger-client-go)
  - 这个项目实现了opentracing的跟踪接口，对数据进行处理
- [opentracing](https://hub.docker.com/r/jaegertracing/all-in-one)
  - docker run -d -p 6831:6831/udp -p 16686:16686 --name jaegertracing jaegertracing/all-in-one:1.12.0

- drool
  - https://hub.docker.com/r/jboss/drools-workbench-showcase
  - https://hub.docker.com/r/jboss/kie-server-showcase

## 基本操作

### format与filter组合过滤

```shell
docker container ls --format "{{.ID}}" --filter "Name=git*"
#输出 none 镜像 awk方式
docker image ls|grep none|awk '{print $3}'
#清空 none 镜像 官方支持
docker rmi $(docker image ls --filter "dangling=true" --format "{{.ID}}")
#配合docker rm $(cmd)
```

### mac宿主机 路由表 localhost->docker.for.mac.host.internal