FROM jboss/business-central-workbench-showcase:7.48.0.Final




ENV  JAVA_OPTS="-Xms256m -Xmx2048m -XX:MetaspaceSize=96M -XX:MaxMetaspaceSize=512m -Djava.net.preferIPv4Stack=true -Dfile.encoding=UTF-8 \
    -Djboss.modules.system.pkgs=org.jboss.logmanager \
    -Djava.util.logging.manager=org.jboss.logmanager.LogManager \
    -Djava.rmi.server.hostname=localhost -Dcom.sun.management.jmxremote.port=9999 -Dcom.sun.management.jmxremote.ssl=false -Dcom.sun.management.jmxremote.authenticate=false"