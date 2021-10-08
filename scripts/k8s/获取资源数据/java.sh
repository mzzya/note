#! /bin/bash

pod="tr-test-sso-api-fb79c6d88-5nmcj"
# 获取最新的两条dump文件

kubectl exec $pod -- ls --sort time ./dump_file/ | grep hprof | awk 'NR<3{print "./dump_file/"$0}' | xargs kubectl exec $pod tar czvf dump_file.tar.gz

sleep 1s

kubectl cp $pod:dump_file.tar.gz ./dump_file.tar.gz

sleep 1s

tar -xzvf ./dump_file.tar.gz -C ~

rm ./dump_file.tar.gz
