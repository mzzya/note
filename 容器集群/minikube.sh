#网络问题时 关闭网络共享 WIFI 使用有线连接试试

minikube start \
    --alsologtostderr \
    --memory=16384 --cpus=4 --nodes=3 \
    --extra-config=kubelet.authentication-token-webhook=true \
    --image-mirror-country=cn \
    --registry-mirror="https://o40mvhma.mirror.aliyuncs.com" \
    --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers \
    --kubernetes-version='1.18.0'
