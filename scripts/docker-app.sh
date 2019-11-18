docker run -d -p 3000:3000 --name=grafana grafana/grafana:6.4.4
docker run -d -p 6831:6831/udp -p 16686:16686 --name jaegertracing jaegertracing/all-in-one:1.15.1
docker run -d -p 9090:9090 -v /Users/wangyang/blog/scripts/etc/prometheus/prometheus.yaml:/etc/prometheus/prometheus.yaml --name=prometheus prom/prometheus:v2.14.0
