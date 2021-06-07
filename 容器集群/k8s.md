# k8s

1. k8s nginx ingress 添加自定义头部
参考资料：https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/annotations/


2. curl 参数 排查网络问题
https://curl.se/docs/manpage.html

```sh
curl 'your_url'   -H 'Content-Type: application/json'   -H 'your_cookie'  -o /dev/null -w '%{size_header}\n'
```


```yaml
# 示例一
# http response信息中添加公共头部
apiVersion: v1
data:
  strict-transport-security: max-age=0; includeSubDomains;
  x-company-group-modify: strict-transport-security;
kind: ConfigMap
metadata:
  name: nginx-configuration-headers
  namespace: kube-system
---
apiVersion: v1
data:
  add-headers: kube-system/nginx-configuration-headers
kind: ConfigMap
metadata:
  labels:
    app: ingress-nginx
  name: nginx-configuration
  namespace: kube-system
```
