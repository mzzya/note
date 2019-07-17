# docker

## mac

### mac宿主机 路由表 localhost->docker.for.mac.host.internal

### prometheus

docker run -e TZ=Asia/Shanghai -p 9090:9090 -d -v /tmp/prometheus.yml:/etc/prometheus/prometheus.yml --name=prometheus prom/prometheus:v2.9.2

### grafana

docker run -d -p 3000:3000 --name=grafana grafana/grafana:6.1.6

# k8s

## dashboard

文档:https://kubernetes.io/docs/tasks/access-application-cluster/web-ui-dashboard/  
启动命令:kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/master/aio/deploy/recommended/kubernetes-dashboard.yaml
登录界面:http://localhost:8001/api/v1/namespaces/kube-system/services/https:kubernetes-dashboard:/proxy/  
授权文档：https://github.com/kubernetes/dashboard/wiki/Creating-sample-user  

获取令牌命令:
kubectl create serviceaccount admin-user  
kubectl -n kube-system describe secret $(kubectl -n kube-system get secret | grep admin-user | awk '{print $1}')  
随便复制一个token就行

## 学习笔记

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

## 实战

### namespace

#### 新增 kubectl create -f ./my-namespace.yaml

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: <insert-namespace-name-here>
```

kubectl config set-context dev --namespace=development --cluster=docker-desktop --user=docker-desktop

kubectl config set-context prod --namespace=production --cluster=docker-desktop --user=docker-desktop