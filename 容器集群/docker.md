# [我的收藏](https://hub.docker.com/u/hellojqk/starred)

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