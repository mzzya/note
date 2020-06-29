# prometheus实践

- 采集接口 http://localhost:9090/metrics
- 展示界面 http://localhost:9090/graph

## 指标类型

- Counter 只增计数器
- Gauge  可增可减计量器
- Histogram 直方统计图
- Summary 摘要

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

官方解释 https://prometheus.io/docs/practices/histograms/