#! /bin/bash

# k8s 环境标识
k8s_env=test
# k8s 命名空间标识 可换成 -A 所有命名空间
k8s_ns="-n tr"

# 删除pod指定目录
kubectl --kubeconfig ~/.kube/$k8s_env.yaml get pod -o=custom-columns=name:.metadata.name,ns:.metadata.namespace $k8s_ns | while read line; do
    lineAry=($line)
    echo $line
    # 输出所有有dump_file的目录信息
    # kubectl --kubeconfig ~/.kube/$k8s_env.yaml -n ${lineAry[1]} exec ${lineAry[0]} -- sh -c "du -h ./dump_file 2>/dev/null"
    # 删除所有dump_file的目录
    # kubectl --kubeconfig ~/.kube/$k8s_env.yaml -n ${lineAry[1]} exec ${lineAry[0]} -- sh -c "du -h ./dump_file && rm -rf ./dump_file && mkdir ./dump_file"
    # 查找所有有+HeapDumpBeforeFullGC参数的容器列表
    # kubectl --kubeconfig ~/.kube/$k8s_env.yaml -n ${lineAry[1]} exec ${lineAry[0]} -- sh -c "du -h ./app.jar 1> /dev/null && jps|grep app|awk '{print \$1}'|xargs -t jinfo -flags" | grep "+HeapDumpBeforeFullGC" && echo $line >>vm-test.list
    kubectl --kubeconfig ~/.kube/$k8s_env.yaml -n ${lineAry[1]} exec ${lineAry[0]} -- sh -c "ls -lh|grep log"
done
