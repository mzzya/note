# 坑

## dns问题

因为某些原因 集群的dns不能改，只能从pod入手。设置为dnsPolicy: None可直接使用，但是会导致集群的用不了。
综合方案 设置重试和轮询解决

```yaml
dnsConfig:
  nameservers:
    - 10.10.10.1
    - 10.10.10.2
  searches:
    - *.hellojqk.com
    - *.hellojqk.com
  options:
    - name: attempts
      value: "3"
    - name: rotate
```
