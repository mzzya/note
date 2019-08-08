1.format与filter组合过滤

```shell
docker container ls --format "{{.ID}}" --filter "Name=git*"
#配合docker rm $(cmd)
```

2. 清理镜像

```shell
# 过滤出ID
docker image ls --filter "dangling=true" --format "{{.ID}}"
docker rmi $(docker image ls --filter "dangling=true" --format "{{.ID}}")
docker image ls|grep none|awk '{print $3}'


docker rmi $(docker image ls|grep none|awk '{print $3}') -f

docker rm $(docker ps -a|grep 'Exited'|awk '{print $1}')
```
