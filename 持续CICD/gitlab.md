
# 获取当前时间 docker镜像

```shell
    # 方案一 限基于linux的镜像(支持apk) 镜像存在时区的无需执行第一行
    - sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && apk add tzdata
    - cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo "Asia/Shanghai" > /etc/timezone
    - date "+%Y%m%d%H%M"
    # 方案二
    - date -d @"$(($(date +%s)+8*60*60))" "+%Y%m%d%H%M"
```

## 配置文件

/Users/Shared/gitlab-runner-smb/config/config.toml

确保并发数量等于环境数量

```toml
concurrent = 4
```

## 检查rnner容器启动方式

SMB-docker-Linux-Shared

- --privileged 特权模式
- volume 映射配置文件到本地

```shell
RUNNER_IMAGE="gitlab/gitlab-runner:alpine-v12.9.0"
GROUP_NAME="smb" #组名
docker run -d --name gitlab-runner-${GROUP_NAME} --privileged --restart always \
        -v /var/run/docker.sock:/var/run/docker.sock \
        -v /Users/Shared/gitlab-runner-${GROUP_NAME}/config:/etc/gitlab-runner \
        ${RUNNER_IMAGE}
```

## 注册4个runner

```shell
RUNNER_IMAGE="gitlab/gitlab-runner:alpine-v12.9.0"
GROUP_NAME="smb" #组名
#标准构建分支
branchs=('dev' 'test' 'uat' 'prd')
for branch in ${branchs[@]}; do
    echo ${branch} #分支名
    #注册
    docker run --rm \
    -v /Users/Shared/gitlab-runner-${GROUP_NAME}/config:/etc/gitlab-runner \
    ${RUNNER_IMAGE} register \
    --tag-list "${GROUP_NAME}-${branch}" \
    --non-interactive \
    --registration-token "your gitlab token" \
    --run-untagged="false" \
    --limit=1 \
    --locked="false" \
    --custom_build_dir-enabled="true" \
    --access-level "not_protected" \
    --name "smb-runner" \
    --url "https://gitlab.example.com/" \
    --executor "docker" \
    --docker-tlsverify="true" \
    --docker-image "docker:19.03.4" \
    --docker-privileged="true" \
    --docker-volumes '/certs/client' \
    --docker-volumes '/cache'
done
```
