#! /bin/bash

kubectl --kubeconfig ~/.kube/test.yaml top pod -A>top.txt
kubectl --kubeconfig ~/.kube/test.yaml get pod -A -o=custom-columns=name:.metadata.name,ns:.metadata.namespace,request-cpu:.spec.containers[0].resources.requests.cpu,limit-cpu:.spec.containers[0].resources.limits.cpu,request-memory:.spec.containers[0].resources.requests.memory,limit-memory:.spec.containers[0].resources.limits.memory>resource.txt


# paste top.txt resource.txt
awk 'NR==FNR{a[i]=$0;i++}NR>FNR{print a[j]" "$0;j++}' top.txt resource.txt