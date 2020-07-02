# API设计规范的一些思考

API状态码`200一把梭`问题历来

## 部分开源项目API

### jaeger

CNCF毕业项目，基于opentracing规范的链路跟踪系统，主要调研此系统UI界面。

```log
# 查询指定链路ID的数据
# http://localhost:16686/api/traces/7e46a082d33cacc8
# Status Code: 200
{
  "data": [
    {
      "traceID": "7e46a082d33cacc8",
      "spans": [],
      "processes": {},
      "warnings": null
    }
  ],
  "total": 0,
  "limit": 0,
  "offset": 0,
  "errors": null
}
# 传入错误ID
# http://localhost:16686/api/traces/7e46a082d33cacc8___lookhere
# Status Code: 400 Bad Request
{
  "data": null,
  "total": 0,
  "limit": 0,
  "offset": 0,
  "errors": [
    {
      "code": 400,
      "msg": "strconv.ParseUint: parsing \"cacc8___lookhere\": invalid syntax"
    }
  ]
}
# 传入超长ID
# http://localhost:16686/api/traces/7e46a082d33cacc87e46a082d33cacc87e46a082d33cacc87e46a082d33cacc87e46a082d33cacc87e46a082d33cacc87e46a082d33cacc87e46a082d33cacc87e46a082d33cacc87e46a082d33cacc87e46a082d33cacc87e46a082d33cacc87e46a082d33cacc87e46a082d33cacc87e46a082d33cacc87e46a082d33cacc87e46a082d33cacc87e46a082d33cacc8
# Status Code: 400 Bad Request
{
  "data": null,
  "total": 0,
  "limit": 0,
  "offset": 0,
  "errors": [
    {
      "code": 400,
      "msg": "TraceID cannot be longer than 32 hex characters: 7e46a082d33cacc87e46a082d33cacc87e46a082d33cacc87e46a082d33cacc87e46a082d33cacc87e46a082d33cacc87e46a082d33cacc87e46a082d33cacc87e46a082d33cacc87e46a082d33cacc87e46a082d33cacc87e46a082d33cacc87e46a082d33cacc87e46a082d33cacc87e46a082d33cacc87e46a082d33cacc87e46a082d33cacc87e46a082d33cacc8"
    }
  ]
}
```

错误相应直接输出msg到页面

### prometheus

CNCF

```log
# http://localhost:9090/api/v1/query?query=rate(prometheus_http_requests_total%5B5m%5D%20offset%205m)
# 查询正常
{
  "status": "success",
  "data": {
    "resultType": "vector",
    "result": []
  }
}
# 查询异常
{
  "status": "error",
  "errorType": "bad_data",
  "error": "invalid parameter 'query': 1:1: parse error: unknown function with name \"rate11\""
}
```

```json
{
  "status": "success" | "error",
  "data": <data>,

  // Only set if status is "error". The data field may still hold
  // additional data.
  "errorType": "<string>",
  "error": "<string>",

  // Only if there were warnings while executing the request.
  // There will still be data in the data field.
  "warnings": ["<string>"]
}
```

### k8s

CNCF毕业项目，当下最火的容器管理平台。restful风格API，官方API约定文档
https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md

#### HTTP状态码

服务器将使用与HTTP规范匹配的HTTP状态代码进行响应。有关服务器将发送的状态码类型的详细信息，请参见以下部分。

API可能会返回以下HTTP状态代码。

##### 成功代码

- 200 StatusOK
  - 指示请求已成功完成。
- 201 StatusCreated
  - 表示创建种类的请求已成功完成。
- 204 StatusNoContent
  - 表示请求已成功完成，并且响应中没有任何正文。
  - 返回以响应HTTP OPTIONS请求。

##### 错误码

- 307 StatusTemporaryRedirect
  - 指示请求资源的地址已更改。
  - 建议的客户端恢复行为：
    - 按照重定向。

- 400 StatusBadRequest
  - 指示请求的无效。
  - 建议的客户端恢复行为：
    - 不要重试。解决请求。

- 401 StatusUnauthorized
  - 表示可以访问服务器并理解该请求，但由于客户端必须提供授权，因此拒绝采取任何进一步的措施。如果客户端提供了授权，则服务器指示提供的授权不合适或无效。
  - 建议的客户端恢复行为：
    - 如果用户未提供授权信息，则提示他们输入适当的凭据。如果用户提供了授权信息，请告知他们其凭据被拒绝，并有选择地再次提示他们。

- 403 StatusForbidden
  - 表示可以访问服务器并理解该请求，但拒绝采取任何进一步的措施，因为它被配置为出于某种原因拒绝客户端对所请求资源的访问。
  - 建议的客户端恢复行为：
    - 不要重试。解决请求。

- 404 StatusNotFound
  - 表示请求的资源不存在。
  - 建议的客户端恢复行为：
    - 不要重试。解决请求。

- 405 StatusMethodNotAllowed
  - 指示代码不支持客户端尝试对资源执行的操作。
  - 建议的客户端恢复行为：
    - 不要重试。解决请求。

- 409 StatusConflict
  - 表示客户端尝试创建的资源已经存在，或者由于冲突而无法完成请求的更新操作。
  - 建议的客户端恢复行为：
    - 如果创建新资源：
      - 要么更改标识符然后重试，要么获取并比较现有对象中的字段，然后发出PUT /更新以修改现有对象。
    - 如果更新现有资源：
      - 请Conflict从status下面的响应部分中查看有关如何检索有关冲突性质的更多信息。
      - 获取并比较先前存在的对象中的字段，合并更改（如果根据前提条件仍然有效），然后使用更新的请求（包括ResourceVersion）重试。

- 410 StatusGone
  - 指示该项目在服务器上不再可用，并且未知转发地址。
  - 建议的客户端恢复行为：
    - 不要重试。解决请求。

- 422 StatusUnprocessableEntity
  - 指示由于作为请求的一部分提供的无效数据，请求的创建或更新操作无法完成。
  - 建议的客户端恢复行为：
    - 不要重试。解决请求。

- 429 StatusTooManyRequests
  - 表示已超过客户端速率限制，或者服务器接收到更多请求，然后可以处理。
  - 建议的客户端恢复行为：
    - Retry-After从响应中读取HTTP`header`，然后至少等待那么长时间再重试。

- 500 StatusInternalServerError
  - 表示可以访问服务器并理解该请求，但是发生了意外的内部错误并且调用的结果未知，或者服务器无法在合理的时间内完成操作（这可能是由于临时服务器负载或与另一台服务器的瞬时通信问题）。
  - 建议的客户端恢复行为：
    - 重试以指数补偿。

- 503 StatusServiceUnavailable
  - 表示所需的服务不可用。
  - 建议的客户端恢复行为：
    - 重试以指数补偿。

- 504 StatusServerTimeout
  - 表示请求无法在给定时间内完成。客户端仅在请求中指定超时参数时才能获得此响应。
  - 建议的客户端恢复行为：
    - 增加超时参数的值，然后使用指数退避重试。

#### 响应状态类型

发生错误时，Kubernetes将始终从任何API端点返回该`Status`种类（字段）。客户应该在适当的时候处理这​​些类型的对象。

一个`Status`种类（字段）将由API在两种情况下会返回：

- 当操作不成功时（即服务器将返回非2xx HTTP状态代码时）。
- HTTP DELETE调用成功时。

状态对象编码为JSON，并作为响应的主体提供。状态对象包含供API人员和机器使用者使用的字段，以获取有关故障原因的更多详细信息。状态对象中的信息补充但不覆盖HTTP状态代码的含义。当状态对象中的字段与通常定义的HTTP`header`具有相同的含义，并且该`header`与响应一起返回时，`header`应被认为具有更高的优先级。

示例

```sh
$ curl -v -k -H "Authorization: Bearer WhCDvq4VPpYhrcfmF6ei7V9qlbqTubUc" https://10.240.122.184:443/api/v1/namespaces/default/pods/grafana

> GET /api/v1/namespaces/default/pods/grafana HTTP/1.1
> User-Agent: curl/7.26.0
> Host: 10.240.122.184
> Accept: */*
> Authorization: Bearer WhCDvq4VPpYhrcfmF6ei7V9qlbqTubUc
>

< HTTP/1.1 404 Not Found
< Content-Type: application/json
< Date: Wed, 20 May 2015 18:10:42 GMT
< Content-Length: 232
<
{
  "kind": "Status",
  "apiVersion": "v1",
  "metadata": {},
  "status": "Failure",
  "message": "pods \"grafana\" not found",
  "reason": "NotFound",
  "details": {
    "name": "grafana",
    "kind": "pods"
  },
  "code": 404
}
```

正常返回

```json
{
  "kind": "Pod",
  "apiVersion": "v1",
  "metadata": {},
  "spec": {},
  "status": {}
```

`status` 字段包含两个可能的值之一：

- Success
- Failure

`message` 可能包含人类可读的错误描述

`reason`可能包含一个机器可读的，单字的CamelCase描述，说明此操作为何处于Failure状态。如果该值为空，则没有可用信息。该reason澄清的HTTP状态代码，但不覆盖它。

`details`可能包含与原因相关的扩展数据。每个原因都可以定义自己的扩展详细信息。该字段是可选的，除保证原因类型定义的模式外，不保证返回的数据符合任何模式。

`reason`和`details`字段的可能值：

- BadRequest
  - 指示该请求本身无效，因为该请求没有任何意义，例如删除只读对象。
  - 这与status reason Invalid上面的内容不同，后者表明API调用可能成功，但是数据无效。
  - 返回BadRequest的API调用永远不会成功。
  - Http状态码： 400 StatusBadRequest

- Unauthorized
  - 表示可以访问服务器并理解该请求，但是在客户端未提供适当授权的情况下拒绝采取任何进一步的措施。如果客户端提供了授权，则此错误表明提供的凭据不足或无效。
  - `details`（可选）：
    - kind string
      - 未授权资源的种类属性（在某些操作上可能与请求的资源不同）。
    - name string
      - 未授权资源的标识符。
  - HTTP状态码： 401 StatusUnauthorized

- Forbidden
  - 表示可以访问服务器并理解该请求，但拒绝采取任何进一步的措施，因为它被配置为出于某种原因拒绝客户端对所请求资源的访问。
  - 详细信息（可选）：
    - kind string
      - 禁止资源的种类属性（在某些操作上可能与请求的资源不同）。
    - name string
      - 禁止资源的标识符。
  - HTTP状态码： 403 StatusForbidden

- NotFound
  - 表示找不到该操作所需的一个或多个资源。
  - 详细信息（可选）：
    - kind string
      - 缺少资源的种类属性（在某些操作上可能与请求的资源不同）。
    - name string
      - 缺少资源的标识符。
  - HTTP状态码： 404 StatusNotFound

- AlreadyExists
  - 表示您正在创建的资源已经存在。
  - 详细信息（可选）：
    - kind string
      - 冲突资源的种类属性。
    - name string
      - 冲突资源的标识符。
  - HTTP状态码： 409 StatusConflict

- Conflict
  - 表示由于冲突而无法完成请求的更新操作。客户端可能需要更改请求。每个资源都可以定义指示冲突性质的自定义详细信息。
  - HTTP状态码： 409 StatusConflict

- Invalid
  - 指示由于作为请求的一部分提供的无效数据，请求的创建或更新操作无法完成。
  - 详细信息（可选）：
    - kind string
      - 无效资源的种类属性
    - name string
      - 无效资源的标识符
    - causes
      - 一个或多个StatusCause条目，指示所提供资源中的数据无效。的reason，message和field属性将被设置。
  - HTTP状态码： 422 StatusUnprocessableEntity

- Timeout
  - 表示请求无法在给定时间内完成。如果服务器已决定对客户端进行速率限制，或者服务器过载并且此时无法处理请求，则客户端可能会收到此响应。
  - Http状态码： 429 TooManyRequests
  - 服务器应设置Retry-AfterHTTP`header`，然后返回 retryAfterSeconds对象的详细信息字段。值0是默认值。

- ServerTimeout
  - 表示可以访问服务器并理解该请求，但是无法在合理的时间内完成该操作。这可能是由于临时服务器负载或与另一台服务器的瞬时通信问题引起的。
  - 详细信息（可选）：
    - kind string
      - 作用的资源的种类属性。
    - name string
      - 正在尝试的操作。
  - 服务器应设置Retry-AfterHTTP`header`，然后返回 retryAfterSeconds对象的详细信息字段。值0是默认值。
  - Http状态码： 504 StatusServerTimeout

- MethodNotAllowed
  - 指示代码不支持客户端尝试对资源执行的操作。
  - 例如，尝试删除只能创建的资源。
  - 返回MethodNotAllowed的API调用永远不会成功。
  - Http状态码： 405 StatusMethodNotAllowed

- InternalError
  - 指示发生内部错误，这是意外错误，并且调用结果未知。
  - 详细信息（可选）：
    - causes
      - 原始错误。
  - Http状态代码：500 StatusInternalServerError code可能包含此状态的建议HTTP返回代码。

##### 命名约定

以下为部分摘抄翻译（译者注）

- `Go`字段名称必须为CamelCase。JSON字段名称必须为camelCase。除了首字母大写之外，两者几乎应始终匹配。两者之间都没有下划线或破折号。
- 指定发生时间的字段名称something应称为somethingTime。请勿使用stamp（例如，creationTimestamp）。
- 除非在缩写中非常常用，例如`id`,`args`或`stdin`，否则请勿在API中使用缩写。
- 类似地，仅在极为知名的情况下才应使用首字母缩写词。首字母缩写词中的所有字母都应具有相同的大小写，并针对情况使用适当的大小写。例如，在字段名的开头，首字母缩写词应全部小写，例如“ httpGet”。当用作常量时，所有字母均应为大写，例如“ TCP”或“ UDP”。
- 表示布尔属性的字段名称称为“ fooable” Fooable，而不是IsFooable。

#### 有感

从k8s的API约定文档来看，采用了http状态码的方式。
