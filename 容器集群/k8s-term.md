# pod term

通常容器启动后，服务进程需要一段时间加载。例如java容器通常需要10来秒启动时间。默认情况下，k8s中容器1号进程启动后就被认为是服务已经准备完毕，就会导致服务出现问题，因此需要配置k8s健康检查的探针确保服务真正可用时才会提供服务。

容器更新或主动停止。

```log
2021-11-03 16:13:57.643  INFO 48330 --- [  restartedMain] com.colipu.sso.SSOApplication            : Started SSOApplication in 9.418 seconds (JVM running for 10.164)
```

## 容器启动

三种探针

| 探针类型 | 作用                                   | 启动时机                                                            |
| -------- | -------------------------------------- | ------------------------------------------------------------------- |
| 启动探测 | 容器内应用是否已启动                   | 容器拉起后                                                          |
| 就绪检查 | 容器内应用是否已经准备好了接收流量     | 有启动探测会在启动探测成功后启动<br/>没有启动探测将在容器拉起后启动 |
| 存货检查 | 容器内的应用是否运行正常，不正常则重启 | 同就绪检查                                                          |

```yaml
livenessProbe:
  failureThreshold: 6 #连续失败多少次触发重启
  httpGet:
    path: /actuator/health/liveness
    port: 8888
    scheme: HTTP
  initialDelaySeconds: 10 #初始延迟多少秒
  periodSeconds: 10 #间隔多久
  successThreshold: 1 #连续成功多少次算正常
  timeoutSeconds: 5 #请求超时时间

readinessProbe:
  failureThreshold: 3

startupProbe:
  failureThreshold: 60 #连续检查60次都是失败才认为是容器无法启动
```

## 容器终止

容器停止时，k8s会同时去做2件事情：1.在后端服务列表里移除当前容器（作用是停止新的流量进来），此操作类似从nginx中移除upstream。2.向容器中的1号进程发送SIGTERM指令，等30秒，如果容器1号进程没有停止则发送SIGKILL强制关闭。因为这两件事是同步进行的，就会出现一种情况发送`SIGTERM`指令后，应用立马中断了新的服务请求连接建立，但此时后端服务列表里并没有将当前容器移除（有延迟），就会出现502的情况。所以要确保`SIGTERM`指令发送时，不再有新的流量进来。所以最简单的处理方式是使用`lifecycle preStop`停止前处理，他会阻塞`SIGTERM`信号的发出。阻塞期间能够让后端服务列表里移除掉当前容器，没有流量后在发送信号让应用处理停止事件。

`terminationGracePeriodSeconds`这个是发送`SIGTERM`容忍应用还有多长时间处理停止事件，超过这个时间后将被直接杀掉。也可以说是容器级别的优雅退出最大等待时间。

```yml
apiVersion: apps/v1
kind: Deployment
spec:
  template:
    metadata:
      labels:
        app: tr-dev-sso-api
    spec:
      containers:
        - lifecycle:
            preStop:
              exec:
                command:
                  - sleep
                  - '20'
      terminationGracePeriodSeconds: 30
```

参考资料：https://kubernetes.io/zh/docs/concepts/workloads/pods/pod-lifecycle/#pod-termination