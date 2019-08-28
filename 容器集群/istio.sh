#! /bin/bash

curl -L https://git.io/getLatestIstio | ISTIO_VERSION=1.2.4 sh -

cd istio-1.2.4

export PATH=$PWD/bin:$PATH

for i in install/kubernetes/helm/istio-init/files/crd*yaml; do kubectl apply -f $i; done


 kubectl apply -f install/kubernetes/istio-demo.yaml