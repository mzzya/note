# k8s 集群内安装helm

此方法暂时行不通

## 添加helm chart 仓库

helm repo add gitlab https://charts.gitlab.io

## 列出k8s tr命名空间下是否有tiller

helm list --tiller-namespace tr

## 安装tiller

helm init --history-max 200 --tiller-namespace tr --debug --kube-context=aliyun-colipu-tr
报错镜像无法获取

## 查看tiller使用的镜像版本

kubectl get deployment tiller-deploy -o custom-columns=name:.spec.template.spec.containers[0].name,image:.spec.template.spec.containers[0].image

## 替换tiller镜像源

kubectl set image deployment/tiller-deploy tiller=registry.cn-hangzhou.aliyuncs.com/google_containers/tiller:v2.14.3


## 列出helm 版本

helm version --tiller-namespace tr

helm install --namespace tr --name gitlab-runner -f ./values.yaml gitlab/gitlab-runner --tiller-namespace tr
