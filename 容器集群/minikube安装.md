# minikube安装

1. 使用地址 https://kubernetes.io/docs/tasks/tools/install-minikube/ 安装minikube
2. 使用迅雷下载 https://storage.googleapis.com/minikube/iso/minikube-v1.3.0.iso 

```shell
mv ~/Downloads/minikube-v1.3.0.iso ~/.minikube/cache/iso/
```

## 拉取镜像

笨方法：mac或win10(git bash)执行pull_image.sh从阿里云拉取k8s相关镜像（当前文档k8s版本v1.15.0，时间2019.7.19）并重新打标签为k8s镜像

新方法：`minikube start`时指定`--registry-mirror=https://registry.docker-cn.com --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers`无需手动拉去镜像

还有：https://hub.docker.com/u/mirrorgooglecontainers

## 安装虚拟机软件

### mac 启动命令

官方推荐 virtualbox

```shell
minikube start --registry-mirror=https://registry.docker-cn.com --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers --alsologtostderr
```

### win10 启动命令

控制面板->程序和功能->启动hyper-v和虚拟机相关选项，安装完成重启

由于安装docker时需要启用hyper-v 所以默认minikube start 会报错无法启动virtualbox所以制定虚拟机驱动为hyperv来执行

```shell
minikube start --vm-driver=hyperv --hyperv-virtual-switch="Default Switch"
# 这种方法启动无需自己去拉镜像
minikube start --registry-mirror=https://registry.docker-cn.com --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers --memory=8192 --vm-driver="hyperv" --hyperv-virtual-switch="Default Switch" --alsologtostderr
```

### linux 启动命令

```shell
#不使用虚拟机
minikube start --registry-mirror=https://registry.docker-cn.com --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers --memory=4096 --vm-driver=none --alsologtostderr
```

注意：`Default Switch`是`hyper-v管理器`中默认的虚拟交换机，网上有些启动示例使用的是`minikube`虚拟交换机【这个会导致无法启动，无法绑定到apiserver上】

错误调试方法

```shell
minikube start --vm-driver=hyperv --hyperv-virtual-switch="Default Switch" --v=7 --alsologtostderr
```

异常解决方法`参考`[文档](https://www.freecodecamp.org/news/get-started-with-docker-and-kubernetes-on-windows-10-73c328c6f89a/)

### 验证

```shell
kubectl version
```

发现版本不一致是因为 我们首先装的是docker,客户端显示的是docker集成的kuberctl版本

```shell
Client Version: version.Info{Major:"1", Minor:"14", GitVersion:"v1.14.3", GitCommit:"5e53fd6bc17c0dec8434817e69b04a25d8ae0ff0", GitTreeState:"clean", BuildDate:"2019-06-06T01:44:30Z", GoVersion:"go1.12.5", Compiler:"gc", Platform:"windows/amd64"}
Server Version: version.Info{Major:"1", Minor:"15", GitVersion:"v1.15.0", GitCommit:"e8462b5b5dc2584fdcd18e6bcfe9f1e4d970a529", GitTreeState:"clean", BuildDate:"2019-06-19T16:32:14Z", GoVersion:"go1.12.5", Compiler:"gc", Platform:"linux/amd64"}
```

运行`minikube kubectl version`版本一致

也可单独安装最新版kubectl 会导致电脑上最多有三个版本的kubectl

- minikube 附带
- docker 附带
- 自己安装
  
按需求创建软连接 或 配置 PATH优先级

至此 minibube方式安装完毕

`也可采用docker自带的方式安装，同样前提需拉docker支持的版本镜像`

### 已知问题

如果`minikube status`查看未正常启动,`minikube delete`删除重新执行启动命令

```shell
kubectl get pods -A
#部分结果
# kube-system   storage-provisioner                     0/1     ErrImagePull   0          20s

minikube docker-env

eval $(minikube docker-env)

docker ps
#部分结果
# gcr.io/storage-provisioner                v1.8.1              4689081edb10        20 months ago       80.8MB
```

K8S启动后执行查看所有pods会发现`storage-provisioner`启动失败，原因是`gcr.io/k8s-minikube/storage-provisioner:v1.8.1`被错误的从`cache`中提取成了`gcr.io/storage-provisioner:v1.8.1`

需要手动补充执行

```shell
docker tag gcr.io/storage-provisioner:v1.8.1 gcr.io/k8s-minikube/storage-provisioner:v1.8.1
docker rmi gcr.io/storage-provisioner:v1.8.1
```

### 启动控制台

启动命令

```shell
minikube dashboard
```

常见问题

```shell
#启动失败先查看pod状态
kubectl get pods -A
#kube-system   kubernetes-dashboard-7b8ddcb5d6-jl296   0/1     CrashLoopBackOff   8          17m
#说明这个pods启动时崩溃了

#查看pods详细描述
kubectl describe pod kubernetes-dashboard-7b8ddcb5d6-jl296 -n kube-system
#看最后的Events没看出来什么问题

#查看pods运行log
kubectl logs -f kubernetes-dashboard-7b8ddcb5d6-jl296 -n kube-system
#panic: secrets is forbidden: User "system:serviceaccount:kube-system:default" cannot create resource "secrets" in API group "" in the namespace "kube-system"
#说的是当前也就是minikube使用的用户没有权限创建资源
#搜索结果 github上
# kubectl create clusterrolebinding add-on-cluster-admin --clusterrole=cluster-admin --serviceaccount=kube-system:default
#没卵用
#mac正常启动 win10 无法启动 祭大招 重新安装
minikube delete
minikube start ......
minikube dashobard
#这次成功了
```

启动衡量指标服务

```shell
minikube addons enable metrics-server
```

或者使用下方链接创建 dashboard 不推荐
https://kubernetes.io/docs/tasks/access-application-cluster/web-ui-dashboard/


## dashboard

文档

https://kubernetes.io/docs/tasks/access-application-cluster/web-ui-dashboard/  

启动命令

```shell
kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/master/aio/deploy/recommended/kubernetes-dashboard.yaml
```

登录界面

http://localhost:8001/api/v1/namespaces/kube-system/services/https:kubernetes-dashboard:/proxy/  

授权文档

https://github.com/kubernetes/dashboard/wiki/Creating-sample-user  

获取令牌命令:
kubectl create serviceaccount admin-user  
kubectl -n kube-system describe secret $(kubectl -n kube-system get secret | grep admin-user | awk '{print $1}')  
随便复制一个token就行

### 相关概念

- Kubernetes Master 主节点
  - kube-apiserver
  - kube-controller-manager 
  - and kube-scheduler
- Kubernetes Node 子节点
  - kubelet，与Kubernetes Master进行通信。
  - kube-proxy，一个网络代理，反映每个节点上的Kubernetes网络服务。
- Kubernetes Objects 
  - Pod `豆荚` 封装应用程序容器（或者，在某些情况下，多个容器），存储资源，唯一的网络IP以及控制容器应如何运行的选项 代表者 docker
  - Service `服务` 它定义了一个逻辑集Pods 和一个访问它们的策略 - 有时称为微服务。Podsa 的目标集合Service（通常）由a确定.  
    - 外部访问web应用通过service转交给具体的Pod执行。
    - 
  - Volume
  - Namespace

### 参考文章

- https://github.com/kubernetes/minikube/blob/master/docs/offline.md
