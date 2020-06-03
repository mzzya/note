#! /bin/bash

cp ./prometheus.yml /tmp/prometheus.yml
docker run --restart always -d -p 9090:9090 -v "/tmp/prometheus.yml:/etc/prometheus/prometheus.yml" prom/prometheus:v2.18.1 --name prometheus