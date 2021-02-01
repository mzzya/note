#! /bin/bash

docker rm -f jbpm-workbench

docker build -t hellojqk/jbpm-workbench .

docker run -p 8080:8080 -p 8001:8001 -p 9999:9999 -v ~/git/drools:/opt/jboss/wildfly/bin/.niogit:Z --name jbpm-workbench --link kie-server hellojqk/jbpm-workbench
