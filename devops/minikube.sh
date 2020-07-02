#网络问题时 关闭网络共享 WIFI 使用有线连接试试

minikube start \
    --alsologtostderr \
    --memory=16384 --cpus=4 --nodes=1 \
    --extra-config=kubelet.authentication-token-webhook=true \
    --image-mirror-country=cn \
    --registry-mirror="https://o40mvhma.mirror.aliyuncs.com" \
    --kubernetes-version='1.18.2'


#win10失败情况一 hyper未启动 解决办法
Enable-WindowsOptionalFeature -Online -FeatureName Microsoft-Hyper-V -All  443434434343