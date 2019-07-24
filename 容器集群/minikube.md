# minikube安装

## 拉取镜像

mac或win10(git bash)执行pull_image.sh从阿里云拉取k8s相关镜像（当前文档k8s版本v1.15.0，时间2019.7.19）并重新打标签为k8s镜像

## 安装虚拟机软件

### mac 启动命令

官方推荐 virtualbox

```shell
minikube start
```

### win10 启动命令

控制面板->程序和功能->启动hyper-v和虚拟机相关选项，安装完成重启

由于安装docker时需要启用hyper-v 所以默认minikube start 会报错无法启动virtualbox所以制定虚拟机驱动为hyperv来执行

```shell
minikube start --vm-driver=hyperv --hyperv-virtual-switch="Default Switch"
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

### 参考文章

- https://github.com/kubernetes/minikube/blob/master/docs/offline.md
