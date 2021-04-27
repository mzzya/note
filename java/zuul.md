# zuul踩坑记录

## 问题

线上服务隔断时间服务就开始无响应，查看cpu,memory都没有发现异常，重启后才能正常访问，但隔不了多久就会再次挂掉。

## 复现问题

### 创建两个http服务

服务一：http://localhost:8899/timeout/0
服务二: http://localhost:9988/timeout/0

### springboot+zuul提供服务代理

#### 代理配置文件

```yaml
zuul:
  routes:
    proxy1:
      path: /proxy1/**
      url: http://127.0.0.1:8899/
    proxy2:
      path: /proxy2/**
      url: http://127.0.0.1:9988/
logging:
  level:
    org.apache.http.impl.conn.PoolingHttpClientConnectionManager: debug
```

访问 http://localhost:8080/proxy1/timeout/0 转发到服务一

访问 http://localhost:8080/proxy2/timeout/0 转发到服务二

#### 代理过滤器

```java
@Component
public class ResponseFilter extends ZuulFilter {
    @Override
    public String filterType() {
        return FilterConstants.POST_TYPE;
    }

    @Override
    public int filterOrder() {
        return 0;
    }

    @Override
    public boolean shouldFilter() {
        return true;
    }

    @SneakyThrows
    @Override
    public Object run() throws ZuulException {
        HttpServletRequest request = RequestContext.getCurrentContext().getRequest();
        HttpServletResponse response = RequestContext.getCurrentContext().getResponse();

        String requestUrl = request.getRequestURL().toString();
        //proxy1 代理正常转发什么也不做
        //proxy2 故意引发一个异常
        if (requestUrl.contains("/proxy2/")) {
           //注意，这里会抛异常
           final PrintWriter writer;
           writer = response.getWriter();
           writer.print("what are you looking at");
            // throw new Exception("呵呵额和");
        }
        return null;
    }
}
```

说明：
通过zuul代理，如果是访问proxy1则正常通过，如果是访问proxy2则抛出异常（异常原因和线上相同），然后开启http连接池日志查看连接池变化。

### 观察日志

#### proxy1正常请求日志

```log
# 第1次
2021-04-27 16:15:06.229 DEBUG 49339 --- [nio-8080-exec-1] h.i.c.PoolingHttpClientConnectionManager : Connection request: [route: {}->http://127.0.0.1:8899][total available: 0; route allocated: 0 of 20; total allocated: 0 of 200]
2021-04-27 16:15:06.236 DEBUG 49339 --- [nio-8080-exec-1] h.i.c.PoolingHttpClientConnectionManager : Connection leased: [id: 0][route: {}->http://127.0.0.1:8899][total available: 0; route allocated: 1 of 20; total allocated: 1 of 200]
2021-04-27 16:15:06.248 DEBUG 49339 --- [nio-8080-exec-1] h.i.c.PoolingHttpClientConnectionManager : Connection [id: 0][route: {}->http://127.0.0.1:8899] can be kept alive indefinitely
2021-04-27 16:15:06.248 DEBUG 49339 --- [nio-8080-exec-1] h.i.c.PoolingHttpClientConnectionManager : Connection released: [id: 0][route: {}->http://127.0.0.1:8899][total available: 1; route allocated: 1 of 20; total allocated: 1 of 200]


# 第2次
2021-04-27 16:15:48.269 DEBUG 49339 --- [nio-8080-exec-2] h.i.c.PoolingHttpClientConnectionManager : Connection request: [route: {}->http://127.0.0.1:8899][total available: 1; route allocated: 1 of 20; total allocated: 1 of 200]
2021-04-27 16:15:48.272 DEBUG 49339 --- [nio-8080-exec-2] h.i.c.PoolingHttpClientConnectionManager : Connection leased: [id: 0][route: {}->http://127.0.0.1:8899][total available: 0; route allocated: 1 of 20; total allocated: 1 of 200]
2021-04-27 16:15:48.273 DEBUG 49339 --- [nio-8080-exec-2] h.i.c.PoolingHttpClientConnectionManager : Connection [id: 0][route: {}->http://127.0.0.1:8899] can be kept alive indefinitely
2021-04-27 16:15:48.273 DEBUG 49339 --- [nio-8080-exec-2] h.i.c.PoolingHttpClientConnectionManager : Connection released: [id: 0][route: {}->http://127.0.0.1:8899][total available: 1; route allocated: 1 of 20; total allocated: 1 of 200]

# 第3次
2021-04-27 16:15:49.292 DEBUG 49339 --- [nio-8080-exec-3] h.i.c.PoolingHttpClientConnectionManager : Connection request: [route: {}->http://127.0.0.1:8899][total available: 1; route allocated: 1 of 20; total allocated: 1 of 200]
2021-04-27 16:15:49.292 DEBUG 49339 --- [nio-8080-exec-3] h.i.c.PoolingHttpClientConnectionManager : Connection leased: [id: 0][route: {}->http://127.0.0.1:8899][total available: 0; route allocated: 1 of 20; total allocated: 1 of 200]
2021-04-27 16:15:49.293 DEBUG 49339 --- [nio-8080-exec-3] h.i.c.PoolingHttpClientConnectionManager : Connection [id: 0][route: {}->http://127.0.0.1:8899] can be kept alive indefinitely
2021-04-27 16:15:49.294 DEBUG 49339 --- [nio-8080-exec-3] h.i.c.PoolingHttpClientConnectionManager : Connection released: [id: 0][route: {}->http://127.0.0.1:8899][total available: 1; route allocated: 1 of 20; total allocated: 1 of 200]
```

手动刷新可见 第一次请求时连接池没有任何连接可用（total available: 0），接着创建了一个连接（Connection leased），连接释放（Connection released）后再次放入连接池（total available: 1），多次手动刷新，因为请求量不大，所以直接复用了连接池的连接，所以观测到`连接可用数量`一直是`1`.

#### proxy2异常请求日志

```log
# 第1次
2021-04-27 16:02:14.216 DEBUG 47731 --- [nio-8080-exec-1] h.i.c.PoolingHttpClientConnectionManager : Connection request: [route: {}->http://127.0.0.1:9988][total available: 0; route allocated: 0 of 20; total allocated: 0 of 200]
2021-04-27 16:02:14.224 DEBUG 47731 --- [nio-8080-exec-1] h.i.c.PoolingHttpClientConnectionManager : Connection leased: [id: 0][route: {}->http://127.0.0.1:9988][total available: 0; route allocated: 1 of 20; total allocated: 1 of 200]
2021-04-27 16:02:14.241  WARN 47731 --- [nio-8080-exec-1] o.s.c.n.z.filters.post.SendErrorFilter   : Error during filtering

# 第2次
2021-04-27 16:02:30.230 DEBUG 47731 --- [nio-8080-exec-2] h.i.c.PoolingHttpClientConnectionManager : Connection request: [route: {}->http://127.0.0.1:9988][total available: 0; route allocated: 1 of 20; total allocated: 1 of 200]
2021-04-27 16:02:30.230 DEBUG 47731 --- [nio-8080-exec-2] h.i.c.PoolingHttpClientConnectionManager : Connection leased: [id: 1][route: {}->http://127.0.0.1:9988][total available: 0; route allocated: 2 of 20; total allocated: 2 of 200]
2021-04-27 16:02:30.233  WARN 47731 --- [nio-8080-exec-2] o.s.c.n.z.filters.post.SendErrorFilter   : Error during filtering

# 第3次
2021-04-27 16:02:31.396 DEBUG 47731 --- [nio-8080-exec-3] h.i.c.PoolingHttpClientConnectionManager : Connection request: [route: {}->http://127.0.0.1:9988][total available: 0; route allocated: 2 of 20; total allocated: 2 of 200]
2021-04-27 16:02:31.396 DEBUG 47731 --- [nio-8080-exec-3] h.i.c.PoolingHttpClientConnectionManager : Connection leased: [id: 2][route: {}->http://127.0.0.1:9988][total available: 0; route allocated: 3 of 20; total allocated: 3 of 200]
2021-04-27 16:02:31.398  WARN 47731 --- [nio-8080-exec-3] o.s.c.n.z.filters.post.SendErrorFilter   : Error during filtering


# 第20次
2021-04-27 16:04:10.172 DEBUG 47731 --- [nio-8080-exec-1] h.i.c.PoolingHttpClientConnectionManager : Connection request: [route: {}->http://127.0.0.1:9988][total available: 0; route allocated: 19 of 20; total allocated: 19 of 200]
2021-04-27 16:04:10.172 DEBUG 47731 --- [nio-8080-exec-1] h.i.c.PoolingHttpClientConnectionManager : Connection leased: [id: 19][route: {}->http://127.0.0.1:9988][total available: 0; route allocated: 20 of 20; total allocated: 20 of 200]
2021-04-27 16:04:10.174  WARN 47731 --- [nio-8080-exec-1] o.s.c.n.z.filters.post.SendErrorFilter   : Error during filtering



# 说明开始阻塞 开始阻塞

2021-04-27 16:04:10.920 DEBUG 47731 --- [nio-8080-exec-3] h.i.c.PoolingHttpClientConnectionManager : Connection request: [route: {}->http://127.0.0.1:9988][total available: 0; route allocated: 20 of 20; total allocated: 20 of 200]
```

手动刷新可见 第一次请求时连接池没有任何连接可用（total available: 0），接着创建了一个连接（Connection leased），对比正常日志缺少 连接释放（Connection released）,手动刷新20次后，达到连接池限制最大数量，无法创建新的连接，导致后续的请求全部阻塞。`Connection request: [route: {}->http://127.0.0.1:9988][total available: 0; route allocated: 20 of 20;`


异常时比正常时缺少了连接释放（Connection released）日志，说明Error会导致连接不能正常释放。

而Error的产生是因为错误的重写response导致的。

因此修复Error即可解决连接被占用问题。

再者 tomcat连接池和httpclient连接池识具体情况进行优化（待补充）