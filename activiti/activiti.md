# activiti

## 表结构

表设计原则：流程数据和业务数据相分离。

### 用户角色相关表

一套成熟的管理系统必不可少的包含用户、角色等权限管理设计，activiti提供了极简的实现，分别是以下四张表

- ACT_ID_USER 用户表
- ACT_ID_GROUP 用户组表（对应部门 或 职位 或 角色）
- ACT_ID_MEMBERSHIP 用户和用户组关系 只有USER和GROP表两个主键组成联合主键
- ACT_ID_INFO 用户信息扩展表

- ACT_GE_BYTEARRAY 各种二进制资源，bpm xml png等
- ACT_GE_PROPERTY

## 官方三大组件

### activiti-app

流程引擎用户控制台。使用此工具可以启动新流程，分配任务，查看和声明任务等

配置文件目录 /usr/local/tomcat/webapps/activiti-app/WEB-INF/classes/META-INF/activiti-app/
主要配置是 activiti-app.properties

默认账密 admin/test

### activiti-admin

配置文件目录 /usr/local/tomcat/webapps/activiti-rest/WEB-INF/classes
主要配置是 db.properties engine.properties

默认账密 admin/admin

### activiti-rest

默认账密 kermit/kermit

## 参考

[官方文档](https://www.activiti.org/userguide/#databaseConfiguration)

[Activiti6详细教程](https://blog.csdn.net/babylovewei/article/details/85166182)

[flowable6,activit6,activit7的对比中文翻译列表](https://my.oschina.net/u/2464371/blog/3027732)
