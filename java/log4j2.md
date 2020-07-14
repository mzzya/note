# log4j2

配置文件格式，优先级`properties`->`yaml`->`json`->`xml`->`默认设置（控制台输出）`

- Xml
  - 较为常见，搜索引擎出现最多的方式
- Yaml
- Properties
- Json

输出方式

- Console
- File

日志级别

ALL，TRACE，DEBUG，INFO，WARN，ERROR，OFF


亮点

- org.apache.logging.log4j.ThreadContext 可存储请求ID和用户ID等信息，并通过配置文件直接配置到日志中，无需额外处理。

常见问题

- `includeLocation`捕获位置信息（类名，文件名，方法名和调用者的行号）可能很慢

```xml
<!-- 标准xml示例 -->
<?xml version="1.0" encoding="UTF-8"?>
<Configuration>
  <Properties>
    <Property name="name1">value</property>
    <Property name="name2" value="value2"/>
  </Properties>
  <filter  ... />
  <Appenders>
    <appender ... >
      <filter  ... />
    </appender>
    ...
  </Appenders>
  <Loggers>
    <Logger name="name1">
      <filter  ... />
    </Logger>
    ...
    <Root level="level">
      <AppenderRef ref="name"/>
    </Root>
  </Loggers>
</Configuration>
```

# Configuration

- status log4j2自身的打印日志级别
- monitorInterval 检测配置文件更改间隔，默认30秒，最小间隔5秒

## Property Substitution

语法 `${base64:Base64_encoded_data}`、`${env:ENV_NAME}`、`${sys:some.property}`

- base64
- ctx
  - org.apache.logging.log4j.ThreadContext中保存的变量，通常用于`requestId`透传
- env
  - 从环境变量中读取，例如hostName等
- sys
  - 从系统配置文件中读取，例如log4j2.yaml中，`Properties`->`Propertie`
- main
  - 从main函数参数中读取
- ...

## Layout

## JsonLayout

- `charset` string
  - 转换为字节数组时要使用的字符集。该值必须是有效的Charset。如果未指定，将使用UTF-8。
- `compact` bool
  - 如果为true，则附加程序不使用行尾和缩进。默认为false。
- `eventEol` bool
  - 如果为true，则追加器在每个记录之后追加一个行尾。默认为false。与eventEol = true和compact = true一起使用时，每行获得一条记录。
- `endOfLine` string
  - 如果设置，将覆盖默认的行尾字符串。例如，将其设置为“ \ n”，并与eventEol = true和compact = true配合使用，使每行有一条记录，并由“ \ n”而不是“ \ r \ n”分隔。默认为null（即未设置）。
- `complete` bool
  - 如果为true，则附加程序包括JSON标头和页脚以及记录之间的逗号。默认为false。
- `properties` bool
  - 如果为true，则附加程序将线程上下文映射包含在生成的JSON中。默认为false。
- `propertiesAsList` bool
  - 如果为true，则将线程上下文映射作为映射条目对象的列表包括在内，其中每个条目都具有“键”属性（其值为键）和“值”属性（其值为值）。默认为false，在这种情况下，线程上下文映射作为键-值对的简单映射包含在内。
- `locationInfo` bool
  - 如果为true，则附加程序将位置信息包括在生成的JSON中。默认为false。生成位置信息 是一项昂贵的操作，可能会影响性能。请谨慎使用。
- `includeStacktrace`
  - 如果为true，则包括所有已记录的Throwable的完整stacktrace （可选，默认为true）。
- `includeTimeMillis`
  - 如果为true，则timeMillis属性将包含在Json有效负载中，而不是包含在即时中。timeMillis将包含自UTC 1970年1月1日午夜以来的毫秒数。
- `stacktraceAsString`
  - 是否将stacktrace格式化为字符串而不是嵌套对象（可选，默认为false）。
- `includeNullDelimiter`
  - 在每个事件之后是否包括NULL字节作为定界符（可选，默认为false）。
- `objectMessageAsJsonObject`
  - 如果为true，则ObjectMessage被序列化为JSON对象到输出日志的“ message”字段。默认为false。

[layout 官方文档](https://logging.apache.org/log4j/2.x/manual/layouts.html)


## Appender

Appender负责将LogEvents传递到其目的地。

- AsyncAppender
- ConsoleAppender
- FileAppender
- JDBCAppender
- HttpAppender
- KafkaAppender
- RollingFileAppender
- ...10+种

## ConsoleAppender

常用PartternLayout

格式
`%d{HH:mm:ss.SSS} [%thread] %-5level %logger{36} - %msg%n`

## RollingFileAppender

- append
- bufferedIO
- bufferSize
- createOnDemand
- filter
- fileName
- immediateFlush
- layout
- locking
- name
- ignoreExceptions
- filePermissions
- fileOwner
- fileGroup


## Layouts

日志格式，通常在控制台

- Json
- Pattern
- Syslog
- Csv
- ...

## Json

## Pattern


## 相关说明

log4j2作为使用量最多的日志组件，有着丰富且强大的配置，提供了很多可选项。

- [appender说明](https://logging.apache.org/log4j/2.x/manual/appenders.html)
- [异步日志说明](https://logging.apache.org/log4j/2.x/manual/async.html)
- [官方文档](https://logging.apache.org/log4j/2.x/manual/configuration.html)