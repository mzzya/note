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
