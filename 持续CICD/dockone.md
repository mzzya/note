# gitlab-ci在晨光科力普中应用

## 项目部署方式演进

### 背景介绍

科力普省心购是晨光文具集团在19年初为了拓展综合办公物资采购业务成立的B2B电商平台，面向中小企业客群，具有”轻快“的特点。项目启动之前，公司项目多为企业、政府、事业单位等提供办公用品采购服务，采用定期发版的方式保证系统的稳健运行，一个小的需求也可能要等上一周才会发布。部署过程中，人工打包偶发文件不正确导致系统可用性的降低。

### 主要编程语言：golang+nodejs

### 运行环境

- dev 开发环境，极不稳定
- test 集成测试环境，BUG率高
- uat 验收环境，基于线上备份数据搭建
- pre 预发布环境 等同于生产环境，”金丝雀“
- prd 生产环境

#### 容器集群

- test集群 对应 test 环境
- uat集群 对应 uat环境 dev环境（早期仅有uat集群）
- prd集群 对应 pre prd 环境

### git分支

- dev 开发分支
- test 集成测试分支
- uat 验收测试分支
- prd 生产分支
- feature-*需求分支

### 开发流程

开发同学基于`teambition`认领新的需求，创建对应`feature-*分支`。

涉及到前后端联调的话，会先合并到`dev分支`进行冒烟测试。

开发完毕，合并`feature-*分支`到`test分支`并通知测试团队测试。

测试团队测试通过后，合并`feature-*分支`到`uat分支`进行验收测试。

验收通过后按需使用`feature-*分支`或者`uat分支`合并到`prd`。


### gitlab-ci

每个环境部署基本上可以简化为3个步骤`stage`:`complie`编译，`docker-build`镜像构建，`deployment`部署。

```yaml
stages:
  - compile
  - docker-build
  - deployment

job-compile-dev:
  stage: compile
  script:
      - npm run ci #或go build 等
  artifacts:
    paths:
      - bin/


job-docker-build-dev:
  stage: docker-build
  script:
    - docker build -t registry.*.com/clp-dev/project:CI_COMMIT_SHORT_SHA-YYYYMMDDHHmm

job-deployment-dev:
  stage: deployment
  script:
    - kubectl patch deploy K8S_DEPLOYMENT_NAME -p '更新镜像json字符串'
```

其中每个stage必须项：

- image 运行镜像名称
- script 执行的命令 `go build`、`docker build`、`kubectl patch`
- stage 属于哪个阶段 `compile`、`docker-build`、`deployment`

## 多项目CI/CD配置管理

项目初始情况

- A项目基于go语言，`compile`阶段 `image: go 1.12.8`
- B项目基与nodejs，`compile`阶段 `node: v10.8`

虽则需求的增多

- C项目基于go语言，`compile`阶段 `image: go 1.13.1`
- D项目...
- E项目...
- ...

每个项目下都独立维护着各自的`.gitlab-ci.yml`

遇到的问题:

问题1：`docker-build`阶段和`deployment`阶段都是冗余代码，每个项目下要维护。

问题2：go1.13.1~go1.13.4版本 http server存在内存泄漏问题，想升级项目到go编译镜像到1.13.5需要对所有go项目通过合并代码方式走一遍全套环境发布流程。

如何解决：

针对问题1 `gitlab 11.*`引入了`include`和`extends`两个关键字，可以将冗余的代码提取到公共的文件中，可以是`local`本git仓库内的文件、`file`其他git仓库、`template`官方提供的模板、`remote`远程：例如OSS内

我们将`docker-build`阶段和`deployment`阶段提取出来放在`common`组下`cicd`项目内。

```yaml
# common/cicd 下 prepared-docker-build.yml
docker-build-dev:
    stage: docker-build
    script:
        - docker build -t registry.*.com/clp-dev/project:CI_COMMIT_SHORT_SHA-YYYYMMDDHHmm

# common/cicd 下 prepared-deploy.yml
deployment-dev:
    stage: deployment
    script:
        - kubectl patch deploy K8S_DEPLOYMENT_NAME -p '更新镜像json字符串'

#相应项目内的 .gitlab-ci.yml

include:
  - project: "common/cicd"
    file: "/prepared-docker-build.yml"
  - project: "common/cicd"
    file: "/prepared-deploy.yml"

stages:
    - compile
    - docker-build
    - deployment

job-compile-dev:
    stage: compile
    script:
        - npm run ci #或go build 等
```

## 缓存与并发控制

我们注册了4个group runner分别对应我们的4个代码分支。push到对应分支时，触发相应的构建，每个runner同时只允许一个项目构建。提取部分配置如下

```toml
# config/config.toml 仅展示 实际未注册命令生成的配置文件

concurrent = 4

[[runners]]
  name = "group-runner-dev"   #runner名称
  url = "https://gitlab.example.com/"
  token = "***"
  executor = "docker"
  [runners.custom_build_dir]
    enabled = true
  [runners.docker]
    tls_verify = true
    image = "docker:19.03.8"
    privileged = true
    pull_policy = "if-not-present"
    disable_entrypoint_overwrite = false
    oom_kill_disable = false
    disable_cache = false
    volumes = ["/certs/client", "/cache", "/etc/docker/daemon.json:/etc/docker/daemon.json"]
    shm_size = 0
  [runners.cache]
    [runners.cache.s3]
    [runners.cache.gcs]
```

## 功能探索