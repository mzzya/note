# #要监控的pods列表
# for pod in $(kubectl get pods -o wide|grep prd-api|awk '{print $1}')
# do
#     # echo $pod
#     #给他们打上debug标签
#     kubectl label pods $pod debug=$pod --overwrite
# done
# echo "输出所有标签"
# kubectl get pods -o wide --show-labels|grep prd-api|awk '{print $1 "\t" $10}'

# #创建pod的 svc 和 ing
# for pod in $(kubectl get pods -o wide|grep prd-api|awk '{print $1}')
# do
#     cp pprof.yaml debug.yaml
#     sed -i "s/{POD}/${pod}/g" debug.yaml
#     # cat debug.yaml
#     kubectl apply -f debug.yaml
# done

# #循环所有的pods 初始化时执行
# for pod in $(kubectl get pods -o wide|grep prd-api|awk '{print $1}')
# do
#     memory=$(kubectl top pod|grep prd-api|grep $pod|awk '{print $3}')
#     echo "http://$pod.cab7b4da4476242bc9ac8743854ce1325.cn-shanghai.alicontainer.com/debug/pprof"
#     fileName=${pod}-$(date "+%Y-%m-%d|%H:%M")-${memory}.heap
#     curl -s http://$pod.cab7b4da4476242bc9ac8743854ce1325.cn-shanghai.alicontainer.com/debug/pprof/heap > ${fileName}
#     echo -e "复制下边一行命令执行比较,-base参数可替换成早期执行的记录\n"
#     echo "go tool pprof -http :9090 -base \"${fileName}\" http://$pod.cab7b4da4476242bc9ac8743854ce1325.cn-shanghai.alicontainer.com/debug/pprof/heap"
#     echo -e "\n\n"
# done

# for pod in $(kubectl get pods -o wide|grep prd-api|awk '{print $1}')
# do
#     memory=$(kubectl top pod|grep prd-api|grep $pod|awk '{print $3}')
#     echo "http://$pod.cab7b4da4476242bc9ac8743854ce1325.cn-shanghai.alicontainer.com/debug/pprof"
#     fileName=${pod}-$(date "+%Y-%m-%d|%H:%M")-${memory}.allocs
#     curl -s http://$pod.cab7b4da4476242bc9ac8743854ce1325.cn-shanghai.alicontainer.com/debug/pprof/allocs > ${fileName}
#     echo -e "复制下边一行命令执行比较,-base参数可替换成早期执行的记录\n"
#     echo "go tool pprof -http :9090 -base \"${fileName}\" http://$pod.cab7b4da4476242bc9ac8743854ce1325.cn-shanghai.alicontainer.com/debug/pprof/allocs"
#     echo -e "\n\n"
# done
