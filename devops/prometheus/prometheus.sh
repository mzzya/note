#! /bin/bash

cp ./prometheus.yml /tmp/prometheus.yml
docker rm -f prometheus
docker run --restart always -d -p 9090:9090 -v "/tmp/prometheus.yml:/etc/prometheus/prometheus.yml" --name prometheus prom/prometheus:v2.19.2