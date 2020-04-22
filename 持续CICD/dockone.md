# gitlab-ci/cd在晨光科力普中应用

## 项目部署方式

### 背景介绍

科力普省心购是晨光文具集团在19年初为了拓展综合办公物资采购业务成立的B2B电商平台，面向中小企业客群。省心购项目启动之前，公司其他项目多为企业、政府、事业单位等提供办公用品采购服务，采用定期发版的方式保证系统的稳健运行，一个小的需求也可能要等上一周才会发布。多达五套的运行环境使得我们需要一款能够保证省心购项目顺利快速准确迭代的CI/CD工具。

### 为什么选择gitlab-ci

首先是公司选择了gitlab作为代码仓库。gitlab-ci运行状况界面深度集成在了gitlab仓库页面左侧导航栏，查看运行状况非常方便，只需配置ci/cd文件`.gitlab-ci.yml`和`dockerfile`文件即可满足我们的自动化需求。使用了大半年的时间再看，官方对于gitlab和gitlab-ci的迭代速度也是非常快，基本上每个月都会有新特性的加入。

### 主要编程语言：golang+nodejs

### 运行环境

- dev 开发环境，极不稳定
- test 集成测试环境，BUG率高
- uat 验收环境，基于线上备份数据搭建
- pre 预发布环境 等同于生产环境，”金丝雀“
- prd 生产环境

#### K8S集群

- test集群 对应 test 环境
- uat集群 对应 uat环境 dev环境（早期仅有uat集群）
- prd集群 对应 pre prd 环境

### git分支

我们采用合并到对于分支，部署对应环境的方式来进行迭代。

- dev 开发分支
- test 集成测试分支
- uat 验收测试分支
- prd 生产分支
- feature-* 需求分支

### 开发流程

开发同学基于`teambition`认领新的需求，创建对应`feature-*分支`。

涉及到前后端联调的话，会先合并到`dev分支`进行冒烟测试。

开发完毕，合并`feature-*分支`到`test分支`并通知测试团队测试。

测试团队测试通过后，合并`feature-*分支`到`uat分支`进行验收测试。

验收通过后按需使用`feature-*分支`或`uat分支`合并到`prd`。

### gitlab-runner

gitlab包含了用于协调作业的开源持续集成服务`gitlab-ci/cd`。

`gitlab-runner`是`gitlab-ci/cd`的处理程序，采用轮询的方式获取gitlab项目变更，执行相应的作业并将结果返回给gitlab。支持二进制，docker,docker-machine,k8s等方式部署。

我们采用docker作为gitlab-runner的运行环境，为每个团队启动一个runner容器，容器内按分支注册了4个runner分别处理各个分支的CI/CD任务。

#### 常用菜单介绍

- CI/CD
  - Pipelines 管道页面，展示所有的CI/CD过程，运行状态，结果。
  - Schedules 可以配置定时触发管道运行。
- Setting
  - CI/CD
    - Runners 选择项目使用的runner，一般使用`Group Runners`组级别runner。
    - Variables 一般用于存放镜像仓库的地址，账号，密码，kubectl配置文件等信息。

ci/cd的过程可以简化为3个`stage`阶段:`complie`编译，`docker-build`镜像构建，`deployment`部署。这三个阶段组合在一起就是`pipeline`。每个阶段都会启动各自的运行容器,容器运行时生成的缓存文件和编译文件需指定保存，否则下个阶段无法获取到。

阶段必须包含的关键参数：

- image 运行镜像 例如：golang:1.14.2
- script 执行的命令 `go build`、`docker build`、`kubectl patch`
- stage 属于哪个阶段 `compile`、`docker-build`、`deployment`

比较实用的可选参数：

- artifacts 工件 例如：编译阶段生成的二进制文件应提供给镜像构建阶段使用。

一个简单的示例

```yaml
# .gitlab-ci.yml
stages:
  - compile
  - docker-build
  - deployment

compile:
  stage: compile
  image: golang:1.14.2
  script:
      - go build #或npm ci 等
  artifacts:
    paths:
      - bin/

docker-build:
  stage: docker-build
  image: docker:19.03.8
  services:
    - docker:19.03.8-dind
  script:
    - docker build -t registry.*.com/clp-dev/project:CI_COMMIT_SHORT_SHA-YYYYMMDDHHmm

deployment:
  stage: deployment
  image: registry.*.com/kubectl:v1.17.3 #需要自己构建
  script:
    - kubectl patch deploy K8S_DEPLOYMENT_NAME -p '更新镜像json字符串'

```

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

### 缓存

#### image services 镜像缓存

`pull_policy`有三种拉取策略

- always 总是会拉取镜像
- if-not-present 如果本地存在则直接使用本地镜像 不存在则拉取，推荐使用
- never 仅从本地获取镜像

```toml
#config.toml

[[runners]]
  pull_policy = "if-not-present"
```

#### 构建阶段缓存 node_modules vendor

每个阶段的job都会在初始化阶段获取

编译阶段，node项目需要执行`npm ci`或`npm i`拉取`node_modules`信息，如果不配置缓存每次都要拉取，拖慢编译时间。

```yaml
## 编译阶段
#方式一
cache:
    #字符串不支持路径 CI_COMMIT_REF_NAME 分支名 CI_PROJECT_NAMESPACE 组名 CI_PROJECT_NAME 项目名
    #例如: dev-mygroup-myproject-node_modules 最终node_modules会被压缩成cache.zip放在此目录下
    key: ${CI_COMMIT_REF_NAME}-${CI_PROJECT_NAMESPACE}-${CI_PROJECT_NAME}-node_modules
    paths:
      - node_modules
#方式二 依赖gitlab 12.5
cache:
  key:
    files:
      - package.json
    prefix: ${CI_PROJECT_NAMESPACE}-${CI_PROJECT_NAME}-node_modules
## 编译阶段之后的镜像构建和部署阶段
cache:
  policy: pull #pull-push

```

- `key` 存放路径
- `paths` 配置当前构建容器中的那些路径需要被缓存起来。

因目前gitlab版本较低，我们目前采用的是`方式一`的配置方式，每个项目的每个构建环境分支维护一个属于自己的缓存地址。

`方式二`依赖较新的gitlab版本，对`key`进行了扩展。我们更倾向于这种方式，原因是依node项目为例，node_modules文件中的内容是根据`package.json`或`package-lock.json`生成的，如果这两个文件没有发生变化，那么就无需更新缓存信息。

`key`=`prefix`+`-`+`SHA(files)`

- `prefix` 生成目录的前缀，可以不定义
- `files` 判定缓存是否需要更新的文件，最多2个，最终生成的路径是根据这两个文件计算出SHA码

### 并发控制

早期项目较少 多个项目多个分支在构建时会按提交顺序依次执行。

```toml
concurrent = 1
[[runners]]
  name = "runner-dev"
  limit = 1
```

改进配置 随着项目的推进，我们基于分支环境拆分了各自独立的runner，每个环境最多只有一个构建任务执行。

```toml
concurrent = 4
[[runners]]
  name = "runner-dev"
  limit = 1
[[runners]]
  name = "runner-test"
  limit = 1
[[runners]]
  name = "runner-uat"
  limit = 1
[[runners]]
  name = "runner-prd"
  limit = 1
```

现在 多个项目多分支支持并发构建。假设A项目的dev分支连续有5次提交，在以往的配置中会触发5个构建任务。第一个在执行，2~5或者B项目dev分支的构建任务需要等待。

```toml
concurrent = 15
[[runners]]
  name = "runner-dev"
  limit = 5
[[runners]]
  name = "runner-test"
  limit = 5
[[runners]]
  name = "runner-uat"
  limit = 5
[[runners]]
  name = "runner-prd"
  limit = 5
```

## 思考与探索

### trigger 触发器的应用

### resource_group