#! /bin/bash

docker rm -f activiti

# docker rmi -f hellojqk/activiti:6

# docker build -t hellojqk/activiti:6 .

docker run -it -p 8080:8080 \
    -v $(pwd)/activiti-app.properties:/usr/local/tomcat/webapps/activiti-app/WEB-INF/classes/META-INF/activiti-app/activiti-app.properties \
    -v $(pwd)/db.properties:/usr/local/tomcat/webapps/activiti-rest/WEB-INF/classes/db.properties \
    -v $(pwd)/activiti-admin.properties:/usr/local/tomcat/webapps/activiti-admin/WEB-INF/classes/META-INF/activiti-admin/activiti-admin.properties \
    --link mysql --name activiti hellojqk/activiti:6
