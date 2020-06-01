# helm笔记

## helm repo

```shell
# 添加官方charts源
helm repo add stable https://kubernetes-charts.storage.googleapis.com
# 获取repo最新
helm repo update
```

- stable http://mirror.azure.cn/kubernetes/charts/ 推荐
- incubator http://mirror.azure.cn/kubernetes/charts-incubator/

## 基本概念

- Chart 图表
  - 包含在Kubernetes集群内运行应用程序、工具或服务所需的所有资源定义
  - 作用类似 docker image
- Repository 仓库
  - Chart 仓库
  - 作用类似 hub.docker.com 提供镜像托管服务的仓库
- Release 发行版
  - helm部署到k8s中Chart的实例


## Chart说明

### 标准目录结构

- wordpress/
  - Chart.yaml 包含有关图表信息的YAML文件
  - LICENSE 许可证信息
  - README.md 描述文件
  - values.yaml 图标默认配置
  - values.schema.json  # OPTIONAL: A JSON Schema for imposing a structure on the values.yaml file
  - charts/ 包含此图表所依赖的任何图表的目录
  - crds/ Custom Resource Definitions
  - templates/  模板与values结合生成k8s文件
  - templates/NOTES.txt 模板用法说明


```yaml
apiVersion: v2 # The chart API version (required) chart依赖helm的api版本 helm3是v2
name: myenvoy # The name of the chart (required) 图表名称
version: 1.14.1 # A SemVer 2 version (required) 图表版本
kubeVersion: ">= 1.17.*" # A SemVer range of compatible Kubernetes versions (optional) k8s 版本
description: A single-sentence description of this project (optional)
type: It is the type of chart (optional)
keywords:
  - A list of keywords about this project (optional)
home: The URL of this projects home page (optional)
sources:
  - A list of URLs to source code for this project (optional)
dependencies: # A list of the chart requirements (optional)
  - name: The name of the chart (nginx)
    version: The version of the chart ("1.2.3")
    repository: The repository URL ("https://example.com/charts") or alias ("@repo-name")
    condition: (optional) A yaml path that resolves to a boolean, used for enabling/disabling charts (e.g. subchart1.enabled )
    tags: # (optional)
      - Tags can be used to group charts for enabling/disabling together
    enabled: (optional) Enabled bool determines if chart should be loaded
    import-values: # (optional)
      - ImportValues holds the mapping of source values to parent key to be imported. Each item can be a string or pair of child/parent sublist items.
    alias: (optional) Alias usable alias to be used for the chart. Useful when you have to add the same chart multiple times
maintainers: # (optional)
  - name: The maintainers name (required for each maintainer)
    email: The maintainers email (optional for each maintainer)
    url: A URL for the maintainer (optional for each maintainer)
icon: A URL to an SVG or PNG image to be used as an icon (optional).
appVersion: The version of the app that this contains (optional). This needn't be SemVer.
deprecated: Whether this chart is deprecated (optional, boolean) 是否弃用
annotations:
  example: A list of annotations keyed by name (optional).
```

[官方文档](https://helm.sh/docs/topics/charts/)