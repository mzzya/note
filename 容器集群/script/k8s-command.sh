# 查看 deploy 的资源分配情况
kubectl get deploy -o=custom-columns=name:.metadata.name,ns:.metadata.namespace,replicas:.spec.replicas,request-cpu:.spec.template.spec.containers[0].resources.requests.cpu,limit-cpu:.spec.template.spec.containers[0].resources.limits.cpu,request-memory:.spec.template.spec.containers[0].resources.requests.memory,limit-memory:.spec.template.spec.containers[0].resources.limits.memory

# 查看 pod 的资源分配情况
kubectl get pod -o=custom-columns=name:.metadata.name,request-cpu:.spec.containers[0].resources.requests.cpu,limit-cpu:.spec.containers[0].resources.limits.cpu,request-memory:.spec.containers[0].resources.requests.memory,limit-memory:.spec.containers[0].resources.limits.memory

# 修改 deploy 的资源分配

kubectl patch deploy smb-test-api -p '{"spec":{"template":{"spec":{"containers":[{"name":"smb-test-api","resources":{"requests":{"cpu":"100m","memory":"128Mi"},"limits":{"cpu":"200m","memory":"256Mi"}}}]}}}}'
