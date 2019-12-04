
# 获取当前时间 docker镜像

```shell
    # 方案一 限基于linux的镜像(支持apk) 镜像存在时区的无需执行第一行
    - sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && apk add tzdata
    - cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo "Asia/Shanghai" > /etc/timezone
    - date "+%Y%m%d%H%M"
    # 方案二
    - date -d @"$(($(date +%s)+8*60*60))" "+%Y%m%d%H%M"
```