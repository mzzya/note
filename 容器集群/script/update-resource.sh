#! /bin/bash

# # 查看 deploy 的资源分配情况
# kubectl get deploy -o=custom-columns=name:.metadata.name,ns:.metadata.namespace,replicas:.spec.replicas,request-cpu:.spec.template.spec.containers[0].resources.requests.cpu,limit-cpu:.spec.template.spec.containers[0].resources.limits.cpu,request-memory:.spec.template.spec.containers[0].resources.requests.memory,limit-memory:.spec.template.spec.containers[0].resources.limits.memory,affinity:.spec.template.spec.affinity

# # 查看 pod 的资源分配情况
# kubectl get pod -o=custom-columns=name:.metadata.name,ns:.metadata.namespace,request-cpu:.spec.containers[0].resources.requests.cpu,limit-cpu:.spec.containers[0].resources.limits.cpu,request-memory:.spec.containers[0].resources.requests.memory,limit-memory:.spec.containers[0].resources.limits.memory,node:.spec.nodeName,affinity:.spec.affinity

# kubectl resource-capacity --kubeconfig ~/.kube/uat.yaml

# for deployName in $(kubectl get deploy | awk 'NR>1{print $1}'); do
#     echo $deployName
#     deployJson=$(kubectl get deploy $deployName -o json)
#     # echo $deployJson
#     if [[ $deployJson =~ "actuator/health" ]]; then
#         # echo "{\"spec\":{\"template\":{\"spec\":{\"containers\":[{\"name\":\"${deployName}\",\"resources\":{\"requests\":{\"cpu\":\"10m\",\"memory\":\"256Mi\"},\"limits\":{\"cpu\":\"1\",\"memory\":\"2Gi\"}},\"livenessProbe\":{\"failureThreshold\":10,\"httpGet\":{\"path\":\"/actuator/health/liveness\",\"port\":8888,\"scheme\":\"HTTP\"},\"initialDelaySeconds\":180,\"periodSeconds\":10,\"successThreshold\":1,\"timeoutSeconds\":5},\"readinessProbe\":{\"failureThreshold\":10,\"httpGet\":{\"path\":\"/actuator/health/readiness\",\"port\":8888,\"scheme\":\"HTTP\"},\"initialDelaySeconds\":30,\"periodSeconds\":10,\"successThreshold\":1,\"timeoutSeconds\":5}}]}}}}"
#         kubectl patch deploy $deployName -p "{\"spec\":{\"template\":{\"spec\":{\"containers\":[{\"name\":\"${deployName}\",\"resources\":{\"requests\":{\"cpu\":\"10m\",\"memory\":\"256Mi\"},\"limits\":{\"cpu\":\"1\",\"memory\":\"2Gi\"}},\"livenessProbe\":{\"failureThreshold\":10,\"httpGet\":{\"path\":\"/actuator/health/liveness\",\"port\":8888,\"scheme\":\"HTTP\"},\"initialDelaySeconds\":180,\"periodSeconds\":10,\"successThreshold\":1,\"timeoutSeconds\":5},\"readinessProbe\":{\"failureThreshold\":10,\"httpGet\":{\"path\":\"/actuator/health/readiness\",\"port\":8888,\"scheme\":\"HTTP\"},\"initialDelaySeconds\":30,\"periodSeconds\":10,\"successThreshold\":1,\"timeoutSeconds\":5}}]}}}}"
#     fi
# done

for deployName in $(kubectl --kubeconfig ~/.kube/test.yaml get deploy -o=custom-columns=name:.metadata.name,ns:.metadata.namespace,replicas:.spec.replicas,containerName:.spec.template.spec.containers[0].name,request-cpu:.spec.template.spec.containers[0].resources.requests.cpu,limit-cpu:.spec.template.spec.containers[0].resources.limits.cpu,request-memory:.spec.template.spec.containers[0].resources.requests.memory,limit-memory:.spec.template.spec.containers[0].resources.limits.memory -A | awk '{print $1","$2","$4}'); do
    echo $deployName
    OLD_IFS="$IFS"
    IFS=","
    deployArray=($deployName)
    IFS="$OLD_IFS"
    # echo "${deployArray[0]}=${deployArray[1]}"
    kubectl --kubeconfig ~/.kube/test.yaml -n ${deployArray[1]} patch deploy ${deployArray[0]} -p "{\"spec\":{\"template\":{\"spec\":{\"containers\":[{\"name\":\"${deployArray[2]}\",\"resources\":{\"limits\":{\"cpu\":\"2\",\"memory\":\"2Gi\"}}}]}}}}"
done

# curl http://localhost:8888/actuator/health/liveness
# curl http://localhost:8888/actuator/health/readiness
