# [我的收藏](https://hub.docker.com/u/smgqk/starred)

- [grafana 标准](https://hub.docker.com/r/grafana/grafana)
- [prometheus 标准](https://hub.docker.com/r/prom/prometheus)
  - docker run -p 9090:9090 -v /tmp/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus:v2.8.2




- opentracing 测试 https://github.com/opentracing/opentracing-go
  - 这个项目是一个规范定义了跟踪信息的处理规范，注入与提取，开启方式
- jaeger-client-go 测试 https://github.com/jaegertracing/jaeger-client-go
  - 这个项目实现了opentracing的跟踪接口，对数据进行处理
- opentracing 测试 https://hub.docker.com/r/jaegertracing/all-in-one
  - docker run -d -p 6831:6831/udp -p 16686:16686 --name jaegertracing jaegertracing/all-in-one:1.12.0

- drool
  - https://hub.docker.com/r/jboss/drools-workbench-showcase
  - https://hub.docker.com/r/jboss/kie-server-showcase