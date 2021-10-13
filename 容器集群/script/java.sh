#! /bin/bash

pod="tr-prd-sso-api-fdb79c996-gqxs4"
# 获取最新的两条dump文件

# kubectl exec $pod -- ls --sort time ./dump_file/ | grep hprof | awk 'NR<3{print "./dump_file/"$0}' | xargs kubectl exec $pod tar czvf dump_file.tar.gz

kubectl cp $pod:dump_file.tar.gz ./dump_file.tar.gz

tar -xzvf ./dump_file.tar.gz -C ~

rm ./dump_file.tar.gz
