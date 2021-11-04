# ingress

## 请求响应耗时

通过进入nginx controller pod可以查看nginx 配置文件。

```nginx
	## start server testssoapi.colipu.com
	server {
		server_name testssoapi.colipu.com ;

		listen 80  ;

		location / {

			# Custom headers to proxied server

			proxy_connect_timeout                   10s;
			proxy_send_timeout                      60s;
			proxy_read_timeout                      60s;

			# In case of errors try the next upstream server before returning an error
			proxy_next_upstream                     error timeout;
			proxy_next_upstream_timeout             0;
			proxy_next_upstream_tries               3;
		}

	}
	## end server testssoapi.colipu.com
```

比较重要的几个参数
- proxy_connect_timeout 定义与代理服务器建立连接的超时时间。应该注意，这个超时通常不能超过75秒。默认10秒。
- proxy_send_timeout 设置将请求传输到代理服务器的超时时间。如果代理服务器在这段时间内没有收到任何信息，则关闭连接。默认60秒。
- proxy_read_timeout 定义从代理服务器读取响应的超时。如果代理服务器在这段时间内没有传输任何信息，连接将被关闭。默认60秒。

- proxy_next_upstream 指定在什么情况下请求应该传递给下一个服务器
- proxy_next_upstream_timeout 限制将请求传递给下一个服务器的时间。0值关闭此限制。
- proxy_next_upstream_tries 限制将请求传递给下一个服务器的可能尝试次数。0值关闭此限制。

以上的配置可以在具体业务的ingress配置annotation中重写


```yml
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  annotations:
    nginx.ingress.kubernetes.io/proxy-read-timeout: '20'
  name: tr-test-sso-api-ing
  namespace: tr
```

测试案例

| 配置情况               | 请求类型 | 业务处理耗时 | 网关相应 | 客户端响应总时长 | 日志                     | k8s日志                          |
| ---------------------- | -------- | ------------ | -------- | ---------------- | ------------------------ | -------------------------------- |
| proxy_read_timeout:60s | GET      | 70s          | 504      | 3m               | 记录了3条请求且返回了200 | 只能从nginx pod看，有3条超时日志 |
| proxy_read_timeout:60s | POST     | 70s          | 504      | 1m               | 记录了1条请求且返回了200 | 只能从nginx pod看，有1条超时日志 |
| proxy_read_timeout:20s | GET      | 70s          | 504      | 1m               | 记录了3条请求且返回了200 | 只能从nginx pod看，有3条超时日志 |
| proxy_read_timeout:20s | POST     | 70s          | 504      | 20.4s            | 记录了1条请求且返回了200 | 只能从nginx pod看，有1条超时日志 |

- 幂等 任意多次执行所产生的影响均与一次执行的影响相同，
  - GET,PUT,DELEET,OPTIONS等
- 非幂等
  - POST, LOCK, PATCH

所以我们从上面的表格中可以看到，幂等的方法在后端服务器发生超时时会转发到下一个可用的后端服务器（只有1台的后端服务器的话就还是对自身再次调用）。

nginx文档 http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_next_upstream



prd-k8s-log 查看域名的请求相应耗时分布。

```sql
* and host: "ssoapi.colipu.com" |select ceil(request_time) latency,count(1) qty group by latency order by latency asc
```



### 参考资料

- https://zhuanlan.zhihu.com/p/127959800