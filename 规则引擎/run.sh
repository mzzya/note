#! /bin/bash

docker rm -f jbpm-wb

docker build -t hellojqk/jbpm-wb .

docker run -p 8080:8080 -p 8001:8001 -p 9990:9990 -e WILDFLY_PASS="admin" -v ~/git/drools:/opt/jboss/wildfly/bin/.niogit:Z --name jbpm-wb --link kie-server hellojqk/jbpm-wb
