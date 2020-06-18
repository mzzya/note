# xxl-job

```sh
docker run -e PARAMS="--spring.datasource.url=jdbc:mysql://docker.for.mac.host.internal:3306/xxl_job?useUnicode=true&characterEncoding=UTF-8&autoReconnect=true&serverTimezone=Asia/Shanghai --spring.datasource.username=root --spring.datasource.password=123123" -p 8080:8080 -v /tmp:/data/applogs --name xxl-job-admin  -d xuxueli/xxl-job-admin:2.2.0
```
