#! /bin/bash

# curl -L https://istio.io/downloadIstio | sh -

# cd istio-1.4.3

# export PATH=$PWD/bin:$PATH

rm -rf ./istio.tar.gz && curl -L -C - -o istio.tar.gz -O https://github.com/istio/istio/releases/download/1.5.1/istio-1.5.1-osx.tar.gz
rm -rf ~/istio && mkdir ~/istio && tar -zxf istio.tar.gz -C ~/istio --strip-components 1
# 以上步骤下载不动的话可使用NDM+8888 8844DNS配合下载

istioctl manifest apply --set profile=demo

kubectl apply -f ~/istio/samples/bookinfo/platform/kube/bookinfo.yaml
kubectl apply -f ~/istio/samples/bookinfo/networking/bookinfo-gateway.yaml
kubectl apply -f ~/istio/samples/httpbin/httpbin.yaml
export INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].nodePort}')
export SECURE_INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="https")].nodePort}')
export INGRESS_HOST=$(minikube ip)
export GATEWAY_URL=$INGRESS_HOST:$INGRESS_PORT
curl -s http://${GATEWAY_URL}/productpage | grep -o "<title>.*</title>"
