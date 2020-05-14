#!/bin/bash
cp ./logstash-kafka.conf /usr/local/etc/logstash/
sudo logstash ./logstash-kafka.conf