# prometheus实践

Prometheus是最初在SoundCloud上构建的开源系统监视和警报工具包 。自2012年成立以来，许多公司和组织都采用了Prometheus，该项目拥有非常活跃的开发人员和用户社区。现在，它是一个独立的开源项目，并且独立于任何公司进行维护。为了强调这一点并阐明项目的治理结构，Prometheus 于2016年加入了 Cloud Native Computing Foundation，成为继Kubernetes之后的第二个托管项目。

Prometheus非常适合记录任何纯数字时间序列。它既适合以机器为中心的监视，也适合于高度动态的面向服务的体系结构的监视。在微服务世界中，它对多维数据收集和查询的支持是一种特别的优势。

## 特征

- 一个多维数据模型，其中包含通过度量标准名称和键/值对标识的时间序列数据
- PromQL，一种灵活的查询语言 ，可利用此维度
- 不依赖分布式存储；单服务器节点是自治的
- 指标收集可由使用方暴露http接口，由服务端主动拉取收集
- 指标收集也可通过中间网关主动拉取后推送到服务端
- 通过服务发现或静态配置发现目标
- 多种图形和仪表板支持模式，官方提供的有一个简单地展示界面，更多时候是跟granfa搭配食用

## 组件

- Prometheus主服务器，它会收集并存储时间序列数据
- 客户端库，用于检测应用程序代码
- 一个支持短暂工作的推送网关
- 诸如HAProxy，StatsD，Graphite等服务的专用出口商
- 一个alertmanager处理警报
- 各种支持工具

大多数Prometheus组件都是用Go编写的，因此易于构建和部署为静态二进制文件。

- 采集接口 http://localhost:9090/metrics
- 展示界面 http://localhost:9090/graph

## OpenMetrics与Prometheus

约等于 Opentracing与jaegertracing

## 指标类型

- Counter 只增计数器
- Gauge  可增可减计量器
- Histogram 直方统计图
- Summary 摘要

```sh
<metric name>{<label name>=<label value>, ...}
api_http_requests_total{method="POST", handler="/messages"}
```

### Counter

单调递增计数器,例如：已服务请求，已完成任务或错误的数量。

```sh
# http各请求总数。
# HELP prometheus_http_requests_total Counter of HTTP requests.
# TYPE prometheus_http_requests_total counter
prometheus_http_requests_total{code="200",handler="/alerts"} 2
prometheus_http_requests_total{code="200",handler="/api/v1/label/:name/values"} 12
prometheus_http_requests_total{code="200",handler="/api/v1/query"} 58
```

### Gauge

可增可减计量器，例如：线程数，并发请求数或内存使用量。

```sh
# 当前存在的goroutines（可以当协程理解）数量
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 45
# 已分配且仍在使用的字节数
# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
# TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 3.3142464e+07
```

### Histogram

直方图，例如：请求持续时间或响应大小。`le` upper inclusive bound，当前统计上限。
所有指标数值都是累积的。

`<basename>_count`==`<basename>_bucket{le="+Inf"}`

```sh
# HTTP请求相应大小直方统计图
# HELP prometheus_http_response_size_bytes Histogram of response size for HTTP requests.
# TYPE prometheus_http_response_size_bytes histogram
prometheus_http_response_size_bytes_bucket{handler="/metrics",le="100"} 0
prometheus_http_response_size_bytes_bucket{handler="/metrics",le="1000"} 0
prometheus_http_response_size_bytes_bucket{handler="/metrics",le="10000"} 2
prometheus_http_response_size_bytes_bucket{handler="/metrics",le="100000"} 2
prometheus_http_response_size_bytes_bucket{handler="/metrics",le="1e+06"} 2
prometheus_http_response_size_bytes_bucket{handler="/metrics",le="1e+07"} 2
prometheus_http_response_size_bytes_bucket{handler="/metrics",le="1e+08"} 2
prometheus_http_response_size_bytes_bucket{handler="/metrics",le="1e+09"} 2
prometheus_http_response_size_bytes_bucket{handler="/metrics",le="+Inf"} 2
prometheus_http_response_size_bytes_sum{handler="/metrics"} 12583
prometheus_http_response_size_bytes_count{handler="/metrics"} 2
```

### Summary

摘要

```sh
# 垃圾回收周期的暂停持续时间的摘要。
# HELP go_gc_duration_seconds A summary of the pause duration of garbage collection cycles.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 0.0002343
go_gc_duration_seconds{quantile="0.25"} 0.0002634
go_gc_duration_seconds{quantile="0.5"} 0.0003184
go_gc_duration_seconds{quantile="0.75"} 0.0006648
go_gc_duration_seconds{quantile="1"} 0.0006918
go_gc_duration_seconds_sum 0.0030561
go_gc_duration_seconds_count 8
```

### 直方图和摘要区别

- 推荐知乎 https://zhuanlan.zhihu.com/p/76904793
- 官方解释 https://prometheus.io/docs/practices/histograms/

## 查询

### 表达式

- =：选择与提供的字符串完全相同的标签。
- !=：选择不等于提供的字符串的标签。
- =~：选择与提供的字符串进行正则表达式匹配的标签。
- !~：选择不与提供的字符串进行正则表达式匹配的标签。

`http_requests_total{environment=~"staging|testing|development",method!="GET"}`

- s -秒
- m - 分钟
- h - 小时
- d - 天
- w -周
- y -年


获取近5分钟内的所有采集时记录明细

`prometheus_http_requests_total[5m]`

```log
prometheus_http_requests_total{code="200",handler="/metrics",instance="localhost:9090",job="prometheus"}

2220 @1593700880.995
2221 @1593700885.995
2222 @1593700890.995
2223 @1593700895.995
2224 @1593700900.995
2225 @1593700905.995
2226 @1593700910.995
2227 @1593700915.995
2228 @1593700920.995
2229 @1593700925.995
```


获取5分钟前的采集记录汇总情况
`prometheus_http_requests_total offset 5m`

```log
prometheus_http_requests_total{code="200",handler="/metrics",instance="localhost:9090",job="prometheus"}	2235
prometheus_http_requests_total{code="200",handler="/api/v1/label/:name/values",instance="localhost:9090",job="prometheus"}	8

```

获取5分钟前的5分钟明细，也就是第6~10分钟的采集记录明细

`prometheus_http_requests_total[5m] offset 5m`