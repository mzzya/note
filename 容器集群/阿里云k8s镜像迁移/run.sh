#! /bin/bash

rm -rf ./*.yaml
rm -rf ./*.json

kubectl config use-context tr-uat

# smb-prd-job 这个暂时不运行

deploys=(smb-pre-api smb-pre-srv smb-pre-web smb-pre-config-web smb-prd-api smb-prd-srv smb-prd-web smb-prd-config-web smb-nginx-lb smb-prd-acm-listener)

for deploy in "${deploys[@]}"; do
    kubectl get deploy/${deploy} -o json >deploy-${deploy}.json
done

svcs=(smb-pre-api smb-pre-srv smb-pre-web smb-pre-config-web smb-prd-api smb-prd-srv smb-prd-web smb-prd-config-web)

for svc in "${svcs[@]}"; do
    kubectl get svc/${svc}-svc -o json >svc-${svc}-svc.json
done

dss=(jaeger-agent)

for ds in "${dss[@]}"; do
    kubectl get ds/${ds} -o json >ds-${ds}.json
done

secretss=(registry.cn-shanghai.aliyuncs.com)

for secrets in "${secretss[@]}"; do
    kubectl get secrets/${secrets} -o json >secrets-${secrets}.json
done

cms=(opentracing smb-prd smb-nginx smb-weixin)

for cm in "${cms[@]}"; do
    kubectl get cm/${cm} -o json >cm-${cm}.json
done

#kubectl config use-context 切换集群
# kubectl config use-context

/usr/local/bin/python3 run.py

kubectl config use-context minikube

kubectl apply -f . -n smb
