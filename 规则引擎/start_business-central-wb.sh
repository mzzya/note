#!/usr/bin/env bash

# Start Wildfly with the given arguments.
echo "Running business-central workbench on JBoss Wildfly..."
exec ./standalone.sh -b $JBOSS_BIND_ADDRESS -c $KIE_SERVER_PROFILE.xml -Dorg.kie.demo=$KIE_DEMO -Dorg.kie.example=$KIE_DEMO -Djava.net.preferIPv4Stack=true -Djava.net.preferIPv4Addresses=true -Djava.rmi.server.hostname=localhost -Dcom.sun.management.jmxremote.rmi.port=9999 -Dcom.sun.management.jmxremote.port=9999 -Dcom.sun.management.jmxremote.authenticate=false -Dcom.sun.management.jmxremote.ssl=false
exit $?
