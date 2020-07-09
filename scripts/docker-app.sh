docker run -d -p 3000:3000 --name=grafana grafana/grafana:6.5.2
docker run -d --restart always --network net --name jaeger \
  -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 14250:14250 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.18
docker run -d -p 9090:9090 -v /Users/wangyang/blog/scripts/etc/prometheus/:/etc/prometheus/ --name=prometheus prom/prometheus:v2.15.2
