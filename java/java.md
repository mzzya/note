# java学习笔记

## 常见名词

- `JAVA SE`
  - Java Standard Edition Java标准版
- `JIT`
  - just-in-time compilation 即时编译（动态编译）。
  - 相应的概念有`AOT` ahead-of-time compilation 事前编译（静态编译）。
- `JDBC`
  - Java Database Connectivity Java数据库连接
- `AOP`
  - Aspect Oriented Programming 面向切面编程
  - 面向动词领域
  - 将通用需求功能从不相关的类当中分离出来，能够使得很多类共享一个行为，一旦发生变化，不必修改很多类，而只需要修改这个行为即可
  - 框架
    - Spring AOP
    - AspectJ
      - Joinpoint（连接点）指那些被拦截到的点，在 Spring 中，可以被动态代理拦截目标类的方法。
      - Pointcut（切入点）指要对哪些 Joinpoint 进行拦截，即被拦截的连接点。
      - Advice（通知）指拦截到 Joinpoint 之后要做的事情，即对切入点增强的内容。
      - Target（目标）指代理的目标对象。
      - Weaving（植入）指把增强代码应用到目标上，生成代理对象的过程。
      - Proxy（代理）指生成的代理对象。
      - Aspect（切面）切入点和通知的结合。
- `OOP`
  - Object Oriented Programming 面向对象编程
  - 面向名词领域
  - 将需求功能划分为不同的并且相对独立，封装良好的类，并让它们有着属于自己的行为，依靠继承和多态等来定义彼此的关系的话
  - 3大特性
    - 封装
    - 继承
    - 多态
  - 对象+类+继承+多态+消息，核心概念是类和对象
- `GoF`
  - Gang of Four 四人组 泛指 设计模式


## 注解

- org.springframework.stereotype
  - @Controller 控制器
  - @Service 业务逻辑
  - @Repository  数据访问
    - @Target({ElementType.TYPE})
    - @Retention(RetentionPolicy.RUNTIME)
    - @Documented
    - @Component
  - @Component 组件
    - @Target(ElementType.TYPE)
    - @Retention(RetentionPolicy.RUNTIME)
    - @Documented
    - @Indexed
  - @Indexed
    - @Target(ElementType.TYPE)
    - @Retention(RetentionPolicy.RUNTIME)
    - @Documented

- java.lang.annotation
  - @Retention 注解保留策略
    - SOURCE 编译时丢弃
    - CLASS 编译到class文件中，但vm运行时不保留 默认
    - RUNTIME 编译到class文件中，vm运行时保留，可以通过反射获取。
  - @Documented 将注解信息生成到java doc中
  - @Target - 标记这个注解应该是哪种 Java 成员。
    - TYPE class、interface、annotation type、enum上
    - FIELD 字段上
    - METHOD 方法上
    - PARAMETER 参数上
    - CONSTRUCTOR 构造方法上
    - LOCAL_VARIABLE
    - ANNOTATION_TYPE 注解上 @interface
    - PACKAGE
    - TYPE_PARAMETER 1.8
    - TYPE_USE 1.8
    - MODULE 9
  - @Inherited - 标记这个注解是继承于哪个注解类(默认 注解并没有继承于任何子类)
  - @Repeatable 可重复的
- java.lang
  - @Override 重载
  - @Deprecated 废弃的方法
  - @SuppressWarnings 抑制警告，使用了`@Deprecated`注解方法法后，编译器会警告提醒，此注解是忽略警告用的。


```java
package java.lang.annotation;

public interface Annotation {
    boolean equals(Object obj);
    int hashCode();
    String toString();
    Class<? extends Annotation> annotationType();
}

```

## spring-boot


### 热更新、热交换、hotswapping

- Command/Ctl + , `Build、Execution、Deployment` -> `Compiler` ->勾选 `Build project automatically`

- Command/Ctl + Shift + A或者导航栏`Help`->`Find Action`->`Registry...`->勾选 `compiler.automake.allow.when.app.running`

```xml
<dependencies>
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-devtools</artifactId>
        <optional>true</optional>
    </dependency>
</dependencies>
```


### 集成redis缓存

- @Cacheable：触发​​缓存填充。
- @CacheEvict：触发​​缓存逐出。
- @CachePut：更新缓存，而不会干扰方法的执行。
- @Caching：重新组合要在一个方法上应用的多个缓存操作。
- @CacheConfig：在类级别共享一些与缓存相关的常见设置。

https://docs.spring.io/spring-boot/docs/current/reference/html/spring-boot-features.html#boot-features-caching

https://docs.spring.io/spring/docs/5.2.8.RELEASE/spring-framework-reference/integration.html#cache

###

- [官-热更新](https://docs.spring.io/spring-boot/docs/current-SNAPSHOT/reference/htmlsingle/#howto-hotswapping)
- [管-开发工具配置](https://docs.spring.io/spring-boot/docs/current-SNAPSHOT/reference/htmlsingle/#using-boot-devtools)


### 相关资料

- [spring官方文档](https://docs.spring.io/spring/docs/current/spring-framework-reference/)
- [spring-boot官方文档](https://docs.spring.io/spring-boot/docs/current-SNAPSHOT/reference/htmlsingle/)
- [java SE规范](https://docs.oracle.com/javase/specs/index.html)

### JMX

JMX（Java Management Extensions，即Java管理扩展）是一个为应用程序、设备、系统等植入管理功能的框架。JMX可以跨越一系列异构操作系统平台、系统体系结构和网络传输协议，灵活的开发无缝集成的系统、网络和服务管理应用。

### proxy

```sh
-Dhttp.proxyHost=127.0.0.1 -Dhttp.proxyPort=8899 -Dhttps.proxyHost=127.0.0.1 -Dhttps.proxyPort=8899
```


## java nacos sdk

### nacos-config-spring-boot-starter

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

#### 若要实现自动刷新有两种注解 @NacosValue @NacosConfigurationProperties

- `@NacosValue` 字段级别注解
- `@NacosConfigurationProperties` 类级别注解

##### 方法一 @NacosValue

```java
@Configuration
public class Config {

    @NacosValue(value = "${key}", autoRefreshed = true)
    private String key;
}
```

##### 方法二 @NacosConfigurationProperties

```java
@Component
@NacosConfigurationProperties(dataId = "nacos", groupId = "TR", autoRefreshed = true)
public class NcpConfig {

    private String key;
}
```


#### 版本说明

`2021-06-21`实测`0.2.8`对`spring-boot-starter-parent`最高支持到`2.3.*`，不支持`2.4.*`,`2.5.*`

见 https://github.com/nacos-group/nacos-spring-boot-project/issues/169

### spring-cloud-starter-alibaba-nacos-config

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

#### 若要实现自动刷新 @Value&@RefreshScope @ConfigurationProperties

- @Value字段级别注解 @RefreshScope 类级别 组合使用才可以
- @ConfigurationProperties 类级别注解

##### @Value&@RefreshScope

```java
@Configuration
@RefreshScope
public class GlobalConfig {

    @Value("${key}")
    private String key;
}
```

@ConfigurationProperties

```java
@Component
@ConfigurationProperties(prefix = "prefix")
public class NcpConfig {

    private String key;
}
```

#### 版本说明*

`2021-06-21`实测`2.2.5.RELEASE`对`spring-boot-starter-parent`最高支持到`2.3.*`

### 总结

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

### 参考资料


- nacos spring boot 版本
  - https://github.com/nacos-group/nacos-spring-boot-project
  -  https://mvnrepository.com/artifact/com.alibaba.boot/nacos-config-spring-boot-starter
- nacos spring cloud 版本
  - https://github.com/alibaba/spring-cloud-alibaba
  - https://mvnrepository.com/artifact/com.alibaba.cloud/spring-cloud-starter-alibaba-nacos-config

## 配置文件application.yml和bootstrap.yml

### 分环境

#### application.yml

默认配置`application.yml`不同环境可增加`application.{env}.yml`文件配置

#### bootstrap.yml

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
