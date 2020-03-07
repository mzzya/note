docker run -d -p 3000:3000 --name=grafana grafana/grafana:6.5.2
docker run -d -p 6831:6831/udp -p 16686:16686 --name jaegertracing jaegertracing/all-in-one:1.16.0
docker run -d -p 9090:9090 -v /Users/wangyang/blog/scripts/etc/prometheus/:/etc/prometheus/ --name=prometheus prom/prometheus:v2.15.2
