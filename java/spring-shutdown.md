# spring shutdown

基于spring-boot 2.3.4

## 默认配置-没有优雅退出

首先我们先看一下一个默认的sping项目在idea中点击停止时的日志

```log
2021-11-02 17:41:44.419  WARN 38752 --- [      Thread-10] c.a.nacos.common.notify.NotifyCenter     : [NotifyCenter] Start destroying Publisher
2021-11-02 17:41:44.419  WARN 38752 --- [       Thread-6] c.a.n.common.http.HttpClientBeanHolder   : [HttpClientBeanHolder] Start destroying common HttpClient
2021-11-02 17:41:44.420  WARN 38752 --- [      Thread-10] c.a.nacos.common.notify.NotifyCenter     : [NotifyCenter] Destruction of the end
2021-11-02 17:41:44.420  WARN 38752 --- [       Thread-6] c.a.n.common.http.HttpClientBeanHolder   : [HttpClientBeanHolder] Destruction of the end
2021-11-02 17:41:44.542  INFO 38752 --- [extShutdownHook] o.s.s.concurrent.ThreadPoolTaskExecutor  : Shutting down ExecutorService 'applicationTaskExecutor'
2021-11-02 17:41:44.545  INFO 38752 --- [      Thread-37] com.xxl.job.core.server.EmbedServer      : >>>>>>>>>>> xxl-job remoting server stop.
2021-11-02 17:41:44.562  INFO 38752 --- [rRegistryThread] c.x.j.c.thread.ExecutorRegistryThread    : >>>>>>>>>>> xxl-job registry-remove success, registryParam:RegistryParam{registryGroup='EXECUTOR', registryKey='tr-ssoapi-executor-agent', registryValue='http://10.10.31.80:9999/'}, registryResult:ReturnT [code=200, msg=null, content=null]
2021-11-02 17:41:44.562  INFO 38752 --- [rRegistryThread] c.x.j.c.thread.ExecutorRegistryThread    : >>>>>>>>>>> xxl-job, executor registry thread destory.
2021-11-02 17:41:44.562  INFO 38752 --- [extShutdownHook] com.xxl.job.core.server.EmbedServer      : >>>>>>>>>>> xxl-job remoting server destroy success.
2021-11-02 17:41:44.563  INFO 38752 --- [FileCleanThread] c.x.j.core.thread.JobLogFileCleanThread  : >>>>>>>>>>> xxl-job, executor JobLogFileCleanThread thread destory.
2021-11-02 17:41:44.563  INFO 38752 --- [rCallbackThread] c.x.j.core.thread.TriggerCallbackThread  : >>>>>>>>>>> xxl-job, executor callback thread destory.
2021-11-02 17:41:44.564  INFO 38752 --- [      Thread-36] c.x.j.core.thread.TriggerCallbackThread  : >>>>>>>>>>> xxl-job, executor retry callback thread destory.
2021-11-02 17:41:44.572  INFO 38752 --- [extShutdownHook] com.zaxxer.hikari.HikariDataSource       : HikariPool-4 - Shutdown initiated...
2021-11-02 17:41:44.575  INFO 38752 --- [extShutdownHook] com.zaxxer.hikari.HikariDataSource       : HikariPool-4 - Shutdown completed.
2021-11-02 17:41:44.575  INFO 38752 --- [extShutdownHook] com.zaxxer.hikari.HikariDataSource       : HikariPool-3 - Shutdown initiated...
2021-11-02 17:41:44.577  INFO 38752 --- [extShutdownHook] com.zaxxer.hikari.HikariDataSource       : HikariPool-3 - Shutdown completed.
2021-11-02 17:41:44.578  INFO 38752 --- [extShutdownHook] com.zaxxer.hikari.HikariDataSource       : HikariPool-2 - Shutdown initiated...
2021-11-02 17:41:44.578  INFO 38752 --- [extShutdownHook] com.zaxxer.hikari.HikariDataSource       : HikariPool-2 - Shutdown completed.
2021-11-02 17:41:44.579  INFO 38752 --- [extShutdownHook] com.zaxxer.hikari.HikariDataSource       : HikariPool-1 - Shutdown initiated...
2021-11-02 17:41:44.580  INFO 38752 --- [extShutdownHook] com.zaxxer.hikari.HikariDataSource       : HikariPool-1 - Shutdown completed.

Process finished with exit code 130 (interrupted by signal 2: SIGINT)
```

程序关闭动作触发后，可以识别的关键信息

- 一些组件`nacos`,`xxljob`,`hikari`等会执行一些销毁事件，xxljob基于`DisposableBean`实现的。
- idea中程序中断的指令是使用了kill -2的信号量`SIGINT`
- 同时观测到如果有执行中的http请求，会被立即中断，本地环境没有网关的情况下是不会有http状态码的。
  - chrome中可见的错误是`net::ERR_EMPTY_RESPONSE`
  - k8s上有nginx网关，网关如果发现业务服务器无响应后有套特殊处理机制，会返回http状态码。后面讲k8s ingress时细讲。

#### 简单介绍kill信号量

- 2	SIGINT 相当于用键盘输入[ctrl]-c 来中断一个程序的进行
- 9	SIGKILL	代表强制中断一个程序的进行，如果该程序进行到一半， 那么尚未完成的部分可能会有『半产品』产生，类似vim会有.filename.swp 保留下来。
- 15 SIGTERM	以正常的结束程序来终止该程序。由于是正常的终止， 所以后续的动作会将他完成。不过，如果该程序已经发生问题，就是无法使用正常的方法终止时， 输入这个signal 也是没有用的。

参考 鸟哥私饭菜第 16.3.2 程序的管理 http://linux.vbird.org/linux_basic/0440processcontrol.php

## 配置优雅退出

```yaml

spring:
 lifecycle:
   # 优雅退出最大等待时间 默认值:30s
   timeout-per-shutdown-phase: 30s

server:
  shutdown: graceful #开启优雅退出
```

```log
2021-11-03 11:05:16.614  WARN 15654 --- [       Thread-9] c.a.nacos.common.notify.NotifyCenter     : [NotifyCenter] Start destroying Publisher
2021-11-03 11:05:16.614  WARN 15654 --- [       Thread-5] c.a.n.common.http.HttpClientBeanHolder   : [HttpClientBeanHolder] Start destroying common HttpClient
2021-11-03 11:05:16.615  WARN 15654 --- [       Thread-9] c.a.nacos.common.notify.NotifyCenter     : [NotifyCenter] Destruction of the end
2021-11-03 11:05:16.616  WARN 15654 --- [       Thread-5] c.a.n.common.http.HttpClientBeanHolder   : [HttpClientBeanHolder] Destruction of the end
2021-11-03 11:05:16.617  INFO 15654 --- [extShutdownHook] o.s.b.w.e.tomcat.GracefulShutdown        : Commencing graceful shutdown. Waiting for active requests to complete
2021-11-03 11:05:16.625  INFO 15654 --- [tomcat-shutdown] o.s.b.w.e.tomcat.GracefulShutdown        : Graceful shutdown complete
2021-11-03 11:05:16.680  INFO 15654 --- [extShutdownHook] o.s.b.w.e.tomcat.GracefulShutdown        : Commencing graceful shutdown. Waiting for active requests to complete
2021-11-03 11:05:46.684  INFO 15654 --- [extShutdownHook] o.s.c.support.DefaultLifecycleProcessor  : Failed to shut down 1 bean with phase value 2147483647 within timeout of 30000ms: [webServerGracefulShutdown]
2021-11-03 11:05:46.708  INFO 15654 --- [tomcat-shutdown] o.s.b.w.e.tomcat.GracefulShutdown        : Graceful shutdown aborted with one or more requests still active
2021-11-03 11:05:53.885  INFO 15654 --- [extShutdownHook] o.s.s.concurrent.ThreadPoolTaskExecutor  : Shutting down ExecutorService 'applicationTaskExecutor'
2021-11-03 11:05:53.888  INFO 15654 --- [      Thread-37] com.xxl.job.core.server.EmbedServer      : >>>>>>>>>>> xxl-job remoting server stop.
2021-11-03 11:05:53.908  INFO 15654 --- [rRegistryThread] c.x.j.c.thread.ExecutorRegistryThread    : >>>>>>>>>>> xxl-job registry-remove success, registryParam:RegistryParam{registryGroup='EXECUTOR', registryKey='tr-ssoapi-executor-agent', registryValue='http://10.10.31.80:9999/'}, registryResult:ReturnT [code=200, msg=null, content=null]
2021-11-03 11:05:53.909  INFO 15654 --- [rRegistryThread] c.x.j.c.thread.ExecutorRegistryThread    : >>>>>>>>>>> xxl-job, executor registry thread destory.
2021-11-03 11:05:53.909  INFO 15654 --- [extShutdownHook] com.xxl.job.core.server.EmbedServer      : >>>>>>>>>>> xxl-job remoting server destroy success.
2021-11-03 11:05:53.909  INFO 15654 --- [FileCleanThread] c.x.j.core.thread.JobLogFileCleanThread  : >>>>>>>>>>> xxl-job, executor JobLogFileCleanThread thread destory.
2021-11-03 11:05:53.910  INFO 15654 --- [rCallbackThread] c.x.j.core.thread.TriggerCallbackThread  : >>>>>>>>>>> xxl-job, executor callback thread destory.
2021-11-03 11:05:53.910  INFO 15654 --- [      Thread-36] c.x.j.core.thread.TriggerCallbackThread  : >>>>>>>>>>> xxl-job, executor retry callback thread destory.
2021-11-03 11:05:53.919  INFO 15654 --- [extShutdownHook] com.zaxxer.hikari.HikariDataSource       : HikariPool-4 - Shutdown initiated...
2021-11-03 11:05:53.923  INFO 15654 --- [extShutdownHook] com.zaxxer.hikari.HikariDataSource       : HikariPool-4 - Shutdown completed.
2021-11-03 11:05:53.924  INFO 15654 --- [extShutdownHook] com.zaxxer.hikari.HikariDataSource       : HikariPool-3 - Shutdown initiated...
2021-11-03 11:05:53.927  INFO 15654 --- [extShutdownHook] com.zaxxer.hikari.HikariDataSource       : HikariPool-3 - Shutdown completed.
2021-11-03 11:05:53.927  INFO 15654 --- [extShutdownHook] com.zaxxer.hikari.HikariDataSource       : HikariPool-2 - Shutdown initiated...
2021-11-03 11:05:53.928  INFO 15654 --- [extShutdownHook] com.zaxxer.hikari.HikariDataSource       : HikariPool-2 - Shutdown completed.
2021-11-03 11:05:53.929  INFO 15654 --- [extShutdownHook] com.zaxxer.hikari.HikariDataSource       : HikariPool-1 - Shutdown initiated...
2021-11-03 11:05:53.929  INFO 15654 --- [extShutdownHook] com.zaxxer.hikari.HikariDataSource       : HikariPool-1 - Shutdown completed.

Process finished with exit code 130 (interrupted by signal 2: SIGINT)
```

通过日志可以观察到的信息有

- `[tomcat-shutdown] o.s.b.w.e.tomcat.GracefulShutdown` 优雅退出时由内置的tomcat支持的。
- 从接收到终止指令后有30秒的时间可供继续处理请求。
  - 如果有正在执行中的耗时请求，30秒到期后将被强制停止。
  - 此时如果有新的请求，会直接`ERR_CONNECTION_REFUSED`拒绝请求。

### 注意：如果你模拟接口超时用的是`thread.sleep()`会在退出时发送发生线程异常，导致服务报错

## 总结

以下是基于直连本地服务的结论，如果使用了nginx等中间网关，网关的重试机制和超时控制会有不同的相应，后面讲k8s ingress再介绍。

- 默认情况下没有优雅退出机制，程序关闭时会立即中断正在执行中的请求导致请求失败。
- 通过tomcat和spring共同提供的优雅退出机制配置，程序关闭时会挡掉新的请求，正在处理中的请求有30秒可以继续执行，可能会有部分请求失败。
- 请求失败没有http状态码。
