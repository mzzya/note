# java学习笔记

## 常见名词

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