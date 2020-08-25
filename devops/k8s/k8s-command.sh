# 查看 deploy 的资源分配情况
kubectl get deploy -o=custom-columns=name:.metadata.name,replicas:.spec.replicas,\
request-cpu:.spec.template.spec.containers[0].resources.requests.cpu,\
limit-cpu:.spec.template.spec.containers[0].resources.limits.cpu,\
request-memory:.spec.template.spec.containers[0].resources.requests.memory,\
limit-memory:.spec.template.spec.containers[0].resources.limits.memory


# 查看 pod 的资源分配情况
kubectl get pod -o=custom-columns=name:.metadata.name,\
request-cpu:.spec.containers[0].resources.requests.cpu,\
limit-cpu:.spec.containers[0].resources.limits.cpu,\
request-memory:.spec.containers[0].resources.requests.memory,\
limit-memory:.spec.containers[0].resources.limits.memory

# 修改 deploy 的资源分配

kubectl patch deploy smb-test-api -p '{"spec":{"template":{"spec":{"containers":[{"name":"smb-test-api","resources":{"requests":{"cpu":"100m","memory":"128Mi"},"limits":{"cpu":"200m","memory":"256Mi"}}}]}}}}'

kubectl rollout restart deploy/smb-test-api deploy/smb-test-job deploy/smb-test-srv deploy/smb-test-web
kubectl rollout restart deploy/smb-uat-api deploy/smb-uat-job deploy/smb-uat-srv deploy/smb-uat-web
kubectl rollout restart deploy/smb-pre-api deploy/smb-pre-job deploy/smb-pre-srv deploy/smb-pre-web


kubectl patch deploy smb-test-api -p '{"spec":{"template":{"spec":{"containers":[{"name":"smb-test-api","resources":{"requests":{"cpu":"100m","memory":"128Mi"},"limits":{"cpu":"250m","memory":"512Mi"}}}]}}}}'
kubectl patch deploy smb-test-job -p '{"spec":{"template":{"spec":{"containers":[{"name":"smb-test-job","resources":{"requests":{"cpu":"100m","memory":"128Mi"},"limits":{"cpu":"250m","memory":"512Mi"}}}]}}}}'
kubectl patch deploy smb-test-srv -p '{"spec":{"template":{"spec":{"containers":[{"name":"smb-test-srv","resources":{"requests":{"cpu":"100m","memory":"128Mi"},"limits":{"cpu":"250m","memory":"512Mi"}}}]}}}}'
kubectl patch deploy smb-test-web -p '{"spec":{"template":{"spec":{"containers":[{"name":"smb-test-web","resources":{"requests":{"cpu":"100m","memory":"128Mi"},"limits":{"cpu":"250m","memory":"512Mi"}}}]}}}}'


kubectl patch deploy smb-uat-api -p '{"spec":{"template":{"spec":{"containers":[{"name":"smb-uat-api","resources":{"requests":{"cpu":"100m","memory":"128Mi"},"limits":{"cpu":"250m","memory":"512Mi"}}}]}}}}'
kubectl patch deploy smb-uat-job -p '{"spec":{"template":{"spec":{"containers":[{"name":"smb-uat-job","resources":{"requests":{"cpu":"100m","memory":"128Mi"},"limits":{"cpu":"250m","memory":"512Mi"}}}]}}}}'
kubectl patch deploy smb-uat-srv -p '{"spec":{"template":{"spec":{"containers":[{"name":"smb-uat-srv","resources":{"requests":{"cpu":"100m","memory":"128Mi"},"limits":{"cpu":"250m","memory":"512Mi"}}}]}}}}'
kubectl patch deploy smb-uat-web -p '{"spec":{"template":{"spec":{"containers":[{"name":"smb-uat-web","resources":{"requests":{"cpu":"100m","memory":"128Mi"},"limits":{"cpu":"250m","memory":"512Mi"}}}]}}}}'

kubectl patch deploy smb-prd-api -p '{"spec":{"template":{"spec":{"containers":[{"name":"smb-prd-api","resources":{"requests":{"cpu":"100m","memory":"128Mi"},"limits":{"cpu":"250m","memory":"512Mi"}}}]}}}}'
kubectl patch deploy smb-prd-srv -p '{"spec":{"template":{"spec":{"containers":[{"name":"smb-prd-srv","resources":{"requests":{"cpu":"100m","memory":"128Mi"},"limits":{"cpu":"250m","memory":"512Mi"}}}]}}}}'
kubectl patch deploy smb-prd-web -p '{"spec":{"template":{"spec":{"containers":[{"name":"smb-prd-web","resources":{"requests":{"cpu":"100m","memory":"128Mi"},"limits":{"cpu":"250m","memory":"512Mi"}}}]}}}}'
kubectl patch deploy smb-prd-config-web -p '{"spec":{"template":{"spec":{"containers":[{"name":"smb-prd-config-web","resources":{"requests":{"cpu":"100m","memory":"128Mi"},"limits":{"cpu":"250m","memory":"512Mi"}}}]}}}}'
kubectl patch deploy smb-pre-api -p '{"spec":{"template":{"spec":{"containers":[{"name":"smb-pre-api","resources":{"requests":{"cpu":"100m","memory":"128Mi"},"limits":{"cpu":"250m","memory":"512Mi"}}}]}}}}'
kubectl patch deploy smb-pre-srv -p '{"spec":{"template":{"spec":{"containers":[{"name":"smb-pre-srv","resources":{"requests":{"cpu":"100m","memory":"128Mi"},"limits":{"cpu":"250m","memory":"512Mi"}}}]}}}}'
kubectl patch deploy smb-pre-web -p '{"spec":{"template":{"spec":{"containers":[{"name":"smb-pre-web","resources":{"requests":{"cpu":"100m","memory":"128Mi"},"limits":{"cpu":"250m","memory":"512Mi"}}}]}}}}'
kubectl patch deploy smb-pre-config-web -p '{"spec":{"template":{"spec":{"containers":[{"name":"smb-pre-config-web","resources":{"requests":{"cpu":"100m","memory":"128Mi"},"limits":{"cpu":"250m","memory":"512Mi"}}}]}}}}'
kubectl patch deploy smb-nginx-lb -p '{"spec":{"template":{"spec":{"containers":[{"name":"smb-nginx-lb","resources":{"requests":{"cpu":"100m","memory":"128Mi"},"limits":{"cpu":"250m","memory":"512Mi"}}}]}}}}'