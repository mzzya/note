# istio尝试问题记录

## 问题

### 第一次开通istio后发现启动了自动注入的namespace中的pod重启后无法访问外部服务。

原因：istio开启后默认屏蔽了外部流量，需要手动开启。

## yaml中配置的dnsConfig失效，pod启动失败或退出

原因: istio为pod注入了`side car`即 `istio proxy` 导致dns解析失效，无法访问私有服务。

## pod中访问阿里云opensearch报header Authorization参数丢失

原因：阿里云opensearch严格校验header中Authorization参数大小写，标准http 1.1协议规定的是不区分大小写，http 2协议规定的是全部小写。
Istio Proxy组件在接管流量时会将所有的header参数转换成小写。opensearch没有计划兼容处理。故临时解决方案是通过注解关掉proxy。

```yaml
# deployment等
spec:
  template:
    metadata:
      annotations:
        sidecar.istio.io/inject: 'false'
```