# 启动 mysql

```shell script
# 第一次启动
docker run --restart=always --name mysql -p 3306:3306 -p 33060:33060 -v ~/DockerVolumes/mysql:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123123 -d mysql:8.0.17 --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci
# 再次启动
docker run --restart=always --name mysql -p 3306:3306 -p 33060:33060 -v ~/DockerVolumes/mysql:/var/lib/mysql -d mysql:8.0.17
```

```shell
docker run --restart=always --name mongo -p 27017:27017 -v ~/DockerVolumes/mongo:/data/db -d mongo:4.2.3
```
