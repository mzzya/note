minikube start --extra-config=kubelet.authentication-token-webhook=true --vm-driver=hyperkit --memory=16384 --cpus=4 --registry-mirror=https://o40mvhma.mirror.aliyuncs.com --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers --alsologtostderr
minikube start --memory=16384 --cpus=4 --registry-mirror=https://o40mvhma.mirror.aliyuncs.com --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers --alsologtostderr
minikube start --memory=16384 --cpus=4 --iso-url='https://kubernetes.oss-cn-hangzhou.aliyuncs.com/minikube/iso/minikube-v1.7.3.iso' --image-mirror-country cn --alsologtostderr

# https://storage.googleapis.com/minikube-preloaded-volume-tarballs/preloaded-images-k8s-v1-v1.17.3-docker-overlay2.tar.lz4
