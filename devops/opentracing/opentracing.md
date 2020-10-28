# opentracing 相关项目学习

## 为什么要用opentracing

[官方解释](https://opentracing.io/docs/overview/what-is-tracing/)


## 相关项目

- [opentracing-go](#opentracing-go)
- [jaeger-client-go](#jaeger-client-go)
- [jaeger](#jaeger)

## opentracing-go

github: [https://github.com/opentracing/opentracing-go](https://github.com/opentracing/opentracing-go)


### 源码模块

- ext/
- log/
- ext.go
- globaltracer.go
- gocontext.go
- noop.go
- propagation.go
- span.go
- tracer.go

## jaeger-client-go



### 与日志系统

- 每个业务系统独立创建各自的request_id
- 日志记录同时记录request_id和tracer_id
- id的解析由业务系统代码拦截器处理解析，并通过上下文侵入式传递