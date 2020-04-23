
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
- daemon.json 请提前配置镜像源加速地址

```shell
RUNNER_IMAGE="gitlab/gitlab-runner:alpine-v12.9.0"
GROUP_NAME="smb" #组名
docker run -d --name gitlab-runner-${GROUP_NAME} --privileged --restart always \
        -v /var/run/docker.sock:/var/run/docker.sock \
        -v /Users/Shared/gitlab-runner-${GROUP_NAME}/config:/etc/gitlab-runner \
        -v /Users/Shared/gitlab-runner-${GROUP_NAME}/cache:/cache \
        -v /Users/Shared/gitlab-runner-${GROUP_NAME}/config/daemon.json:/etc/docker/daemon.json \
        ${RUNNER_IMAGE}
```

## 注册4个runner worker

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
    --docker-image "docker:19.03.8" \
    --docker-privileged="true" \
    --docker-volumes '/certs/client' \
    --docker-volumes '/cache' \
    --docker-volumes '/etc/docker/daemon.json:/etc/docker/daemon.json'
done
```

### 常见问题和优化

大原则：并发进行时 cache目录 GIT_CLONE_PATH目录要带上特定情况下唯一的标识，如环境字段f

#### 多项目多分支并发时

extends和include组合起来使用非常强大

extends支持多级继承，但是不建议使用三个以上级别。支持的最大嵌套级别为10。

#### docker模式问题

docker in docker privileged模式

- 好处：创建的是属于自己的独立docker服务，执行docker rm也不会影响宿主机镜像。隔离的。
- 坏处：CI/CD的dockerfile引用镜像每次都要重新拉取，折中方案，docker pull yourimage:latest docker build --cache-from yourimage:latest .

/var/run/docker.sock 卷映射

- 好处：CI/CD的dockerfile引用镜像不需要每次都拉取。
- 坏处：没有隔离，假设多个团队在进行开发，如果执行了docker rm，会影响到别的团队，因为是公用的宿主机docker.sock。

#### cache问题

concurrent不等于1时需要指定不同的key 否则可能还会出现严重问题

```yaml
cache:
    key: ${CI_COMMIT_REF_NAME}-h5-node_modules
    paths:
      - node_modules/
```

宿主机执行 find / -name '*cache.zip' 观察前后缓存文件变化

#### 自动运行问题

```yaml
rules:
    - if: '$CI_PIPELINE_SOURCE == "merge_request_event"'
      when: never
    - if: '$CI_COMMIT_REF_NAME == "dev" || $CI_COMMIT_REF_NAME == "test" || $CI_COMMIT_REF_NAME == "prd" || $CI_COMMIT_REF_NAME == "uat"'
      when: always
```

如果不配置 $CI_PIPELINE_SOURCE == "merge_request_event" never

则merge_request点击时就会执行，此时代码尚未合并到分支中

#### docker构建步骤时docker仓库代理问题

docker宿主机配置 /etc/docker/daemon.json

然后worker映射 --docker-volumes '/etc/docker/daemon.json:/etc/docker/daemon.json'

官方暂时不支持直接配置

#### 多环境构建问题

假设有 dev test uat prd 4个分支 有 api和web两个项目

想法 各环境可以并行运行 即运行 dev-api时可运行 test-web或test-api

gitlab-ci.yaml 中 steg tag 用环境标识（dev test ...）

limit=1 concurrent=4 会有一个问题

构建目录会交叉 被同时使用！！！

默认构建目录 /builds/组名/项目名

则需要配置
config.toml
custom_build_dir-enabled =true

gitlab-ci.yaml

GIT_CLONE_PATH: /builds{环境}/组名/项目名

以下针对于go语言不使用代理构建

GOPATH: "/builds{环境}"

GOBIN: "/builds{环境}/bin"

### docker in docker 问题

全局配置

```yaml
variables:
  DOCKER_DRIVER: overlay2
  DOCKER_TLS_CERTDIR: "/certs"
```
