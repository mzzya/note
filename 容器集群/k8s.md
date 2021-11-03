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


## ingress controller

通过进入nginx controller pod可以查看nginx 配置文件。

```nginx
proxy_connect_timeout                   10s;
proxy_send_timeout                      60s;
proxy_read_timeout                      60s;

proxy_next_upstream                     error timeout;
proxy_next_upstream_timeout             0;
proxy_next_upstream_tries               3;
```

比较重要的几个参数
- proxy_connect_timeout 定义与代理服务器建立连接的超时时间。应该注意，这个超时通常不能超过75秒。
- proxy_send_timeout 设置将请求传输到代理服务器的超时时间。如果代理服务器在这段时间内没有收到任何信息，则关闭连接。
- proxy_read_timeout 定义从代理服务器读取响应的超时。如果代理服务器在这段时间内没有传输任何信息，连接将被关闭。

- proxy_next_upstream 指定在什么情况下请求应该传递给下一个服务器
- proxy_next_upstream_timeout 限制将请求传递给下一个服务器的时间。0值关闭此限制。
- proxy_next_upstream_tries 限制将请求传递给下一个服务器的可能尝试次数。0值关闭此限制。

以上的配置可以在具体业务的ingress配置annotation中重写

```yml
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  annotations:
    nginx.ingress.kubernetes.io/proxy-next-upstream-tries: '2'
    nginx.ingress.kubernetes.io/proxy-read-timeout: '600'
  name: tr-test-sso-api-ing
  namespace: tr
```


```yml
apiVersion: apps/v1
kind: Deployment
spec:
  template:
    metadata:
      labels:
        app: tr-dev-sso-api
    spec:
      terminationGracePeriodSeconds: 30
```

### 参考资料

- https://zhuanlan.zhihu.com/p/127959800