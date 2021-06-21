# nacos sdk

- [nacos sdk](#nacos-sdk)
  - [nacos-config-spring-boot-starter](#nacos-config-spring-boot-starter)
    - [若要实现自动刷新有两种注解](#若要实现自动刷新有两种注解)
      - [@NacosValue](#nacosvalue)
      - [@NacosConfigurationProperties](#nacosconfigurationproperties)
    - [版本说明](#版本说明)
  - [spring-cloud-starter-alibaba-nacos-config](#spring-cloud-starter-alibaba-nacos-config)
    - [可以实现自动刷新两种注解](#可以实现自动刷新两种注解)
      - [@Value&@RefreshScope](#valuerefreshscope)
      - [@ConfigurationProperties](#configurationproperties)
    - [版本说明*](#版本说明-1)
  - [总结](#总结)
  - [参考资料](#参考资料)
  - [分环境](#分环境)
    - [application.yml](#applicationyml)
    - [bootstrap.yml](#bootstrapyml)

## nacos-config-spring-boot-starter

版本

- spring boot 2.3.12.RELEASE
- sdk 0.2.8

```yml
# application.yml 中配置
nacos:
  config:
    access-key: ********
    secret-key: ********
    namespace: ********
    type: yml #指定配置文件类型，与ACM里配置的类型对应
    bootstrap:
      enable: true #引导启动
    data-id: nacos #对应Data ID
    group: TR #对应Group
    endpoint: acm.aliyun.com #固定配置
    auto-refresh: true #自动刷新，按需开启，需要搭配@NacosValue(value = "${key}", autoRefreshed = true)使用
    remote-first: true #远程配置覆盖本地配置
```

```xml
<dependency>
  <groupId>com.alibaba.boot</groupId>
  <artifactId>nacos-config-spring-boot-starter</artifactId>
  <version>0.2.8</version>
</dependency>
```

### 若要实现自动刷新有两种注解

- `@NacosValue` 字段级别注解
- `@NacosConfigurationProperties` 类级别注解

#### @NacosValue

```java
@Configuration
public class Config {

    @NacosValue(value = "${key}", autoRefreshed = true)
    private String key;
}
```

#### @NacosConfigurationProperties

```java
@Component
@NacosConfigurationProperties(dataId = "nacos", groupId = "TR", autoRefreshed = true)
public class NcpConfig {

    private String key;
}
```


### 版本说明

`2021-06-21`实测`0.2.8`对`spring-boot-starter-parent`最高支持到`2.3.*`，不支持`2.4.*`,`2.5.*`

见 https://github.com/nacos-group/nacos-spring-boot-project/issues/169

## spring-cloud-starter-alibaba-nacos-config

版本

- spring boot 2.3.12.RELEASE
- sdk 2.2.5.RELEASE

```yml
# bootstrap.yml 中配置
spring:
  cloud:
    nacos:
      config:
        access-key: *******
        secret-key: *******
        namespace: *******
        group: TR #对应Group
        endpoint: acm.aliyun.com #固定配置
        file-extension: "yml" #指定配置文件类型
        refresh-enabled: true #自动刷新
        name: nacos #对应Data ID
```

```xml
<dependency>
    <groupId>com.alibaba.cloud</groupId>
    <artifactId>spring-cloud-starter-alibaba-nacos-config</artifactId>
    <version>2.2.5.RELEASE</version>
</dependency>
```

### 可以实现自动刷新两种注解

- @Value字段级别注解 @RefreshScope 类级别 组合使用才可以
- @ConfigurationProperties 类级别注解

#### @Value&@RefreshScope

```java
@Configuration
@RefreshScope
public class GlobalConfig {

    @Value("${key}")
    private String key;
}
```

#### @ConfigurationProperties

```java
@Component
@ConfigurationProperties(prefix = "prefix")
public class NcpConfig {

    private String key;
}
```

### 版本说明*

`2021-06-21`实测`2.2.5.RELEASE`对`spring-boot-starter-parent`最高支持到`2.3.*`

## 总结

- nacos-config-spring-boot-starter
  - 配置在application.yml中
  - 如果想实现自动刷新需要使用它提供的注解
    - @NacosValue
      - 必须配置 autoRefreshed = true
    - @NacosConfigurationProperties 比较傻还要配置dataID和group
      - 必须加 autoRefreshed = true
- spring-cloud-starter-alibaba-nacos-config
  - 必须配置在bootstrap.yml中
  - 自动刷新直接使用spring注解即可，写法相对简洁
    - @Value&@RefreshScope
    - @ConfigurationProperties

综合来看，更推荐使用`spring-cloud-starter-alibaba-nacos-config`这个包，用起来更原生和简洁，分环境配置的问题可以看`分环境`所述

## 参考资料

- nacos spring boot 版本
  - https://github.com/nacos-group/nacos-spring-boot-project
  -  https://mvnrepository.com/artifact/com.alibaba.boot/nacos-config-spring-boot-starter
- nacos spring cloud 版本
  - https://github.com/alibaba/spring-cloud-alibaba
  - https://mvnrepository.com/artifact/com.alibaba.cloud/spring-cloud-starter-alibaba-nacos-config

## 分环境

### application.yml

默认配置`application.yml`不同环境可增加`application.{env}.yml`文件配置

### bootstrap.yml

只能写在一个文件里，但是可以通过`---`隔开不同环境的配置，例如

```yaml
spring:
  profiles: dev
  cloud:
    nacos:
      config:
        group: TR #对应Group
        endpoint: acm.aliyun.com #固定配置
        file-extension: "yml" #指定配置文件类型，与ACM里配置的类型对应
        refresh-enabled: true #自动刷新
        name: nacos #对应Data ID
logging:
  level:
    com.alibaba.*: debug
---
spring:
  profiles: dev
  cloud:
    nacos:
      config:
        access-key: ******
        secret-key: ******
        namespace: ************
---
spring:
  profiles: test
  cloud:
    nacos:
      config:
        access-key: ******
        secret-key: ******
        namespace: ******
```
