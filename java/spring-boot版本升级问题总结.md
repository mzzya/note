# spring-boot版本升级问题总结

由`2.3.4`到`2.6.3`, 各版本生命周期见[spring-boot官方说明](https://spring.io/projects/spring-boot#support)

## 升级历程

直接升级项目spring-boot版本到2.6.3，问题如下

### java.lang.ClassNotFoundException: org.springframework.boot.context.properties.ConfigurationBeanFactoryMetadata

在搜索引擎或者github中能找到的答案大部分与nacos有关，在nacos中也有相关的issues。

项目中的确使用了

```xml
<dependency>
    <groupId>com.alibaba.cloud</groupId>
    <artifactId>spring-cloud-starter-alibaba-nacos-config</artifactId>
    <version>2.2.5.RELEASE</version>
</dependency>
```

一般情况下直接升级这个包即可，与此相关的项目有如下三个：

- alibaba/nacos
- alibaba/spring-cloud-alibaba
- nacos-group/nacos-spring-boot-project

在[mvn官方仓库](https://mvnrepository.com/)最新的版本是`2021.1`，github指向的是`alibaba/spring-cloud-alibaba`仓库，但`具体实现`是在`nacos-group/nacos-spring-boot-project`仓库，版本说明在`alibaba/spring-cloud-alibaba`的`Wiki`->[版本说明](https://github.com/alibaba/spring-cloud-alibaba/wiki/%E7%89%88%E6%9C%AC%E8%AF%B4%E6%98%8E)中。该问题在3个仓库issues中经常被误导为 至今不支持高版本的spring-boot。

实际在Wiki中标明了能支持到2.4.2

| Spring Cloud Alibaba Version | Spring Cloud Version  | Spring Boot Version |
| ---------------------------- | --------------------- | ------------------- |
| 2021.1                       | Spring Cloud 2020.0.1 | 2.4.2               |

解决方法：`spring-cloud-starter-alibaba-nacos-config`升至`2021.1`

### Failed to configure a DataSource: 'url' attribute is not specified and no embedded datasource could be configured

异常时配置如下

```xml
<parent>
  <groupId>org.springframework.boot</groupId>
  <artifactId>spring-boot-starter-parent</artifactId>
  <version>2.4.2</version>
  <relativePath/>
</parent>
<dependencies>
  <dependency>
      <groupId>com.alibaba.cloud</groupId>
      <artifactId>spring-cloud-starter-alibaba-nacos-config</artifactId>
      <version>2021.1</version>
  </dependency>
</dependencies>
```

碰到这个错误第一反应是配置文件没有加载进来，可见[issue-2021.1版本无法读取Nacos配置中心配置](https://github.com/alibaba/spring-cloud-alibaba/issues/2060)

通过以下配置可以观察到异常情况下并没有加载配置中心的日志（正常情况下启动时会输出配置文件内容）

```yaml
logging.level:
  com.alibaba: debug
  org.apache.http: debug
```

原因是因为新的包基于spring-boot版本是2.4.2，这个版本的spring-boot用的spring-cloud-commons用的是3.0.1版本，对应的spring cloud是2020.0。

以下摘自[官方版本说明](https://github.com/spring-cloud/spring-cloud-release/wiki/Spring-Cloud-2020.0-Release-Notes)：

由 spring-cloud-commons 提供的 Bootstrap 默认不再启用。如果您的项目需要它，可以通过属性或新的启动器重新启用它。
- 通过属性集spring.cloud.bootstrap.enabled=true或重新启用spring.config.use-legacy-processing=true。这些需要设置为环境变量、java 系统属性或命令行参数。
- 另一种选择是包含新的spring-cloud-starter-bootstrap.

解决方法：

```xml
<dependency>
    <groupId>org.springframework.cloud</groupId>
    <artifactId>spring-cloud-starter-bootstrap</artifactId>
    <version>3.1.0</version>
</dependency>
```

### java.lang.ClassNotFoundException: org.springframework.boot.Bootstrapper

异常时配置如下

```xml
<parent>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-parent</artifactId>
        <version>2.6.2</version>
        <relativePath/>
</parent>
<dependencies>
  <dependency>
      <groupId>com.alibaba.cloud</groupId>
      <artifactId>spring-cloud-starter-alibaba-nacos-config</artifactId>
      <version>2021.1</version>
  </dependency>
</dependencies>
```

解决方法同样是：

```xml
<dependency>
    <groupId>org.springframework.cloud</groupId>
    <artifactId>spring-cloud-starter-bootstrap</artifactId>
    <version>3.1.0</version>
</dependency>
```

### The dependencies of some of the beans in the application context form a cycle: com.github.pagehelper.autoconfigure.PageHelperAutoConfiguration

原因见[issue-现有的版本好像与 SpringBoot 2.6.0 都不兼容](https://github.com/pagehelper/pagehelper-spring-boot/issues/126)

解决方法：

```xml
<dependency>
    <groupId>com.github.pagehelper</groupId>
    <artifactId>pagehelper-spring-boot-starter</artifactId>
    <version>1.4.1</version>
</dependency>
```

### Unable to locate the default servlet for serving static content. Please set the 'defaultServletName' property explicitly

原因：Spring Boot 2.4 将不再注册DefaultServlet您的 servlet 容器提供的内容。在大多数应用程序中，它没有被使用，因为 Spring MVCDispatcherServlet是唯一需要的 servlet。

见：[spring-boot-2.4文档](https://github.com/spring-projects/spring-boot/wiki/Spring-Boot-2.4-Release-Notes#default-servlet-registration)

解决方法：

```yaml
server:
  servlet:
    register-default-servlet: true
```

### Spring Boot [2.6.3] is not compatible with this Spring Cloud release train

详细错误内容

```log
Your project setup is incompatible with our requirements due to following reasons:

- Spring Boot [2.6.3] is not compatible with this Spring Cloud release train


Action:

Consider applying the following actions:

- Change Spring Boot version to one of the following versions [2.3.x, 2.4.x] .
You can find the latest Spring Boot versions here [https://spring.io/projects/spring-boot#learn].
If you want to learn more about the Spring Cloud Release train compatibility, you can visit this page [https://spring.io/projects/spring-cloud#overview] and check the [Release Trains] section.
If you want to disable this check, just set the property [spring.cloud.compatibility-verifier.enabled=false]
```

通过上方的日志可以看出来是是因为spring-cloud有多版本的依赖导致检测没过，该错误官方也有描述，
见[spring-cloud-common](https://cloud.spring.io/spring-cloud-commons/multi/multi__spring_cloud_commons_common_abstractions.html#_spring_cloud_compatibility_verification)

解决方法一：

```xml
<dependency>
    <groupId>org.springframework.cloud</groupId>
    <artifactId>spring-cloud-commons</artifactId>
    <version>3.1.0</version>
</dependency>
```

解决方法二：

忽略警告，个人不推荐

```yaml
spring:
  cloud:
    compatibility-verifier.enabled: false
```

## 总结

spring-boot版本由`2.3.4`到`2.6.3`，会出现问题的组件包已知的有：

- spring-cloud-starter-alibaba-nacos-config
- pagehelper-spring-boot-starter

核心原因是spring组件版本升级带来的兼容性问题，就需要出问题的第三方组件同步升级解决。出现问题时，最好跟着错误中的关键字去官方文档中找找答案。

已知问题解决方案汇总如下:

```yaml
server:
  servlet:
    register-default-servlet: true
```

```xml
<!-- 解决spring-cloud-starter-alibaba-nacos-config -->
<dependency>
    <groupId>com.alibaba.cloud</groupId>
    <artifactId>spring-cloud-starter-alibaba-nacos-config</artifactId>
    <version>2021.1</version>
</dependency>
<dependency>
    <groupId>org.springframework.cloud</groupId>
    <artifactId>spring-cloud-starter-bootstrap</artifactId>
    <version>3.1.0</version>
</dependency>
<!-- pagehelper-spring-boot-starter -->
<dependency>
    <groupId>com.github.pagehelper</groupId>
    <artifactId>pagehelper-spring-boot-starter</artifactId>
    <version>1.4.1</version>
</dependency>

<dependency>
    <groupId>org.springframework.cloud</groupId>
    <artifactId>spring-cloud-commons</artifactId>
    <version>3.1.0</version>
</dependency>
```

## 多个项目如何统一管理这些依赖？

基于maven父子模块，父项目`dependencyManagement`管理公共依赖和版本，子项目在引用时就无需配置。