minikube start \
--extra-config=kubelet.authentication-token-webhook=true \
--vm-driver=hyperkit \
--memory=16384 --cpus=4 \
--iso-url='https://kubernetes.oss-cn-hangzhou.aliyuncs.com/minikube/iso/minikube-v1.8.0.iso' \
--registry-mirror=https://o40mvhma.mirror.aliyuncs.com \
--image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers \
--kubernetes-version='1.17.4' \
--alsologtostderr
