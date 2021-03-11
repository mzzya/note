FROM tomcat:jdk8-adoptopenjdk-openj9

COPY wars/activiti-admin.war webapps/activiti-admin/activiti-admin.war
COPY wars/activiti-app.war webapps/activiti-app/activiti-app.war
COPY wars/activiti-rest.war webapps/activiti-rest/activiti-rest.war
COPY sqljdbc4-4.0.jar lib/
COPY mysql-connector-java-8.0.21.jar lib/
COPY mysql-connector-java-8.0.21.jar /usr/local/tomcat/webapps/activiti-rest/WEB-INF/lib/
COPY mysql-connector-java-8.0.21.jar /usr/local/tomcat/webapps/activiti-admin/WEB-INF/lib/

RUN cd /usr/local/tomcat/webapps/activiti-admin && jar -xf activiti-admin.war && rm activiti-admin.war \
    && cd /usr/local/tomcat/webapps/activiti-app && jar -xf activiti-app.war && rm activiti-app.war \
    && cd /usr/local/tomcat/webapps/activiti-rest && jar -xf activiti-rest.war && rm activiti-rest.war

# COPY activiti-app.properties /usr/local/tomcat/webapps/activiti-app/WEB-INF/classes/META-INF/activiti-app/activiti-app.properties
