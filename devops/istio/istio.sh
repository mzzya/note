#! /bin/bash

# 安装istio相关组件
istioctl install --set profile=demo

# default 命名空间开启自动注入 Envoy sidecar
kubectl label namespace default istio-injection=enabled

