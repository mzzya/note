#! /bin/bash

rm -rf ./*.yaml
rm -rf ./*.json

# 获取所有命名空间下镜像和镜像拉取密钥配置
# kubectl --kubeconfig ~/.kube/test.yaml get deploy -n smb -o custom-columns=name:.metadata.name,image:.spec.template.spec.containers[0].image,imagePullSecrets:.spec.template.spec.imagePullSecrets | grep -v none

deploy=$(kubectl --kubeconfig ~/.kube/test.yaml get deploy -n tr -o custom-columns=name:.metadata.name,image:.spec.template.spec.containers[0].image,imagePullSecrets:.spec.template.spec.imagePullSecrets | awk 'NR>1 {print $1}')

for d in $deploy; do
    # echo $d
    # kubectl --kubeconfig ~/.kube/test.yaml get deploy $d -n tr -o custom-columns=name:.metadata.name,image:.spec.template.spec.containers[0].image,imagePullSecrets:.spec.template.spec.imagePullSecrets | awk 'NR>1'
    kubectl --kubeconfig ~/.kube/test.yaml -n tr patch deploy $d -p '{"spec":{"template":{"spec":{"imagePullSecrets":null}}}}'
    # exit 0
    # kubectl --kubeconfig ~/.kube/test.yaml get deploy -n tr $d -o json >$d.json
done
