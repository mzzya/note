# 参考文章 https://github.com/kubernetes/minikube/blob/master/docs/offline.md

mkdir -p ~/.minikube/cache/images/k8s.gcr.io/
mkdir -p ~/.minikube/cache/images/gcr.io/k8s-minikube/

# minikube k8s.grc.io 缓存目录
minikube_cache_k8s_grc_io_dir=~/.minikube/cache/images/k8s.gcr.io/

# minikube grc.io 缓存目录
minikube_cache_grc_io_dir=~/.minikube/cache/images/gcr.io/k8s-minikube/

# 阿里云杭州谷歌镜像映射地址
source=registry.cn-hangzhou.aliyuncs.com/google_containers/

# k8s.gcr.io地址
target_k8s_grc_io=k8s.gcr.io/


k8s_grc_io_images=('kubernetes-dashboard-amd64:v1.10.1'
'kube-controller-manager:v1.15.0'
'kube-proxy:v1.15.0'
'kube-apiserver:v1.15.0'
'kube-scheduler:v1.15.0'
'kube-addon-manager:v9.0'
'coredns:1.3.1'
'etcd:3.3.10'
'pause:3.1'
'k8s-dns-kube-dns-amd64:1.14.13'
'k8s-dns-sidecar-amd64:1.14.13'
'k8s-dns-dnsmasq-nanny-amd64:1.14.13')

for image in ${k8s_grc_io_images[@]}; do
    docker pull ${source}${image}
    docker tag ${source}${image} ${target_k8s_grc_io}${image}
    docker rmi ${source}${image}
    docker save -o ${minikube_cache_k8s_grc_io_dir}${image//:/_} ${target_k8s_grc_io}${image}
done

grc_io_images=('storage-provisioner:v1.8.1')

target_grc_io=gcr.io/

for image in ${grc_io_images[@]}; do
    docker pull ${source}${image}
    docker tag ${source}${image} ${target_grc_io}k8s-minikube/${image}
    docker rmi ${source}${image}
    docker save -o ${minikube_cache_grc_io_dir}${image//:/_} ${target_grc_io}${image}
done

docker images | grep ${source}
docker images | grep ${target_k8s_grc_io}
docker images | grep ${target_grc_io}

ls -l ${minikube_cache_k8s_grc_io_dir}
ls -l ${minikube_cache_grc_io_dir}

echo "已知问题！！！！！！"
echo "已知问题！！！！！！"
echo "已知问题！！！！！！"

echo "gcr.io/k8s-minikube/storage-provisioner:v1.8.1这个镜像无法顺利的提取到minikube中，但不影响K8S启动"