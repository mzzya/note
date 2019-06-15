任务需求
==
设计一套灵活的促销规则
---
- 假设 闹着玩的
  - 某商品限时促销 2020-2-2 2:2:2 ~ 2020-2-2 3:3:3 这段时间内8折
  - 某商品限时满二减一 2020-2-2 2:22:22 ~ 2020-2-2 3:33:33 最多减免一件
  - 活动页商品 优惠券 满100减10 满200减25 (要求 订单中包含活动页商品即可满减，多个活动可同时使用各自活动的优惠券)
  - 通用优惠券形式 可叠加活动优惠券（红包也可通过这样发放）
  - 积分直接抵消金额
  - 某商品限购10件

实际需求分析
- 
- `优惠券规则`单独定义逻辑 有适用商品范围和类目限制
- `关于促销` 有适用商品范围和类目限制
  - `满减` 满130-40 满200-70
  - `打折` 10件9折 20件8折 30件7折
  - `每满100-10`
- 京东为例
  - 加入购物车后 会按照促销分组 实时显示各促销下的折扣价
  - 每个商品可以参加多个促销，但只能选择一个结算(多个的话存在叠加折扣风险，导致商品超低价，不允许)
- 我们要做的
  - 促销信息在我们后台维护
  - 促销勾选后类目和商品Id后计算出促销对应的Id集合 和商品对应的促销ID集合
  - 

- 促销信息定义
  - ID 促销ID
  - Type 促销类型 大促（大部分商品参加的活动）中促（部分商品参加的活动）小促（单个商品或含3个商品以下组合促销）
  - Name 名称
  - Desc 简称
  - URL 页面地址
  - CategoryIds 涉及类目集合
  - GoodsIds 涉及商品集合
  - SkuIds 涉及商品集合
  - BeginTime 开始时间
  - EndTime 结束时间
  - SaleType 促销类型 满减 打折 组合 每满减
    - Condition 满减 100-20;200-40;300-80;
    - Condition 打折 1-0.9;2-0.8;3-0.7;
    - Condition 没满100-10块
  - 

规则引擎介绍
=
摘自[百度百科](https://baike.baidu.com/item/%E8%A7%84%E5%88%99%E5%BC%95%E6%93%8E/3076955?fr=aladdin)

规则引擎由推理引擎发展而来，是一种嵌入在应用程序中的组件，实现了将业务决策从应用程序代码中分离出来，并使用预定义的语义模块编写业务决策。接受数据输入，解释业务规则，并根据业务规则做出业务决策。

应用背景编辑
-
企业级管理者对企业IT系统的开发有着如下的要求：
- 为提高效率，管理流程必须自动化，即使现代商业规则异常复杂。
- 市场要求业务规则经常变化，IT系统必须依据业务规则的变化快速、低成本的更新。
- 为了快速、低成本的更新，业务人员应能直接管理IT系统中的规则，不需要程序开发人员参与。

产品优点
-
使用规则引擎可以通过降低实现复杂业务逻辑的组件的复杂性，降低应用程序的维护和可扩展性成本，其优点如下：
- 分离商业决策者的商业决策逻辑和应用开发者的技术决策；
- 能有效的提高实现复杂逻辑的代码的可维护性；
- 在开发期间或部署后修复代码缺陷；
- 应付特殊状况，即客户一开始没有提到要将业务逻辑考虑在内；
- 符合组织对敏捷或迭代开发过程的使用；

我的总结
-
简而言之就是为了将重的，易变化的业务规则与代码逻辑分离。

相关概念
-
打开drools官网 https://www.drools.org/


给出的docker镜像有
- Drools Workbench    
- Drools Workbench Showcase   ---->即 `Business Central` 设计决策和其他工件


- KIE Execution Server
- KIE Execution Server Showcase  ---->即`KIE Server` 执行和测试您创建的工件

区别在于带showcase的默认创建了默认权限和用户


- Author
    - Authoring of knowledge using a UI metaphor, such as: DRL, BPMN2, decision table, class models.
    - 使用UI隐喻创作知识，例如：DRL，BPMN2，决策表，类模型。
- Build
    - Builds the authored knowledge into deployable units.
    - 将创作的知识构建到可部署的单元中。
    - For KIE this unit is a JAR.
    - 对于KIE，这个单位是JAR。
- Test

    - Test KIE knowledge before it’s deployed to the application.
    - 在将KIE知识部署到应用程序之前对其进行测试。
- Deploy
    - Deploys the unit to a location where applications may utilize (consume) them.
    - 将单元部署到应用程序可以利用（使用）它们的位置。
    - KIE uses Maven style repository.
    - KIE使用Maven风格的存储库。
- Utilize
    - The loading of a JAR to provide a KIE session (KieSession), for which the application can interact with.
    - 加载JAR以提供应用程序可与之交互的KIE会话（KieSession）。
    - KIE exposes the JAR at runtime via a KIE container (KieContainer).
    - KIE在运行时通过KIE容器（KieContainer）公开JAR。
    - KieSessions, for the runtime’s to interact with, are created from the KieContainer.
    -KieSessions，用于与之交互的运行时，是从KieContainer创建的。
- Run
    - System interaction with the KieSession, via API.
    - 通过API与KieSession进行系统交互。
- Work
    - User interaction with the KieSession, via command line or UI.
    - 用户通过命令行或UI与KieSession进行交互。
- Manage
    - Manage any KieSession or KieContainer.
    - 管理任何KieSession或KieContainer。

部署
--
- 开发环境：通常由一个Business Central安装和至少一个KIE Server安装组成。
- 运行时环境：由一个或多个具有或不具有Business Central的KIE服务器实例组成。Business Central有一个嵌入式Drools控制器。如果安装Business Central，请使用菜单 → 部署 → 执行服务器**`Menu → Deploy → Execution servers`**页面来创建和维护容器。

Assert菜单选项说明
- Decision Model and Notation (DMN) models `决策模型和符号（DMN）模型`
- Guided decision table `指导决策表`
    - 基于UI的表设计器中创建的规则表
- Spreadsheet decision tables `电子表格决策表`
    - 将XLS或XLSX电子表格决策表上载到Business Central
- Guided rules `引导规则`
    - 基于UI的规则设计器中创建的单个规则
- Guided rule templates `引导规则模板`
    - 基于UI的模板设计器中创建的可重用规则结构
- DRL rules `DRL规则`
    - 在.drl文本文件中定义的单个规则
- Predictive Model Markup Language (PMML) models `预测模型标记语言（PMML）模型`
    - 基于数据挖掘组（DMG）定义的符号标准的预测数据分析模型


起个名字
--
- Knowledge Representation and Reasoning `知识表示与推理`
  
  KRR是关于我们如何以符号形式表示我们的知识，即我们如何描述某些东西。推理是关于我们如何利用这些知识进行思考。基于系统的面向对象语言（如C ++，Java和C＃）具有称为类的数据定义，用于描述建模实体的组成和行为。
  
  您可能已经听过讨论将正向链接（反应性和数据驱动）的优点与反向链接（被动和查询驱动）进行比较。存在许多其他类型的推理技术，每种推理技术都扩大了我们可以声明性地解决的问题的范围。仅举几例：不完美推理（模糊逻辑，确定性因素），可废止逻辑，信念系统，时间推理和相关性。您无需了解所有这些术语即可理解和使用Drools。他们只是想知道研究主题的范围，实际上更广泛，并随着研究人员推动新的界限而不断发展。
  
  KRR通常被称为人工智能的核心。即使使用像神经网络这样的生物学方法来模拟大脑并且更多地关注模式识别而不是思考，它们仍然建立在KRR理论的基础之上。我对Drools的第一次努力是以工程为导向，因为我没有正式的培训或对KRR的理解。学习KRR让我获得了更广泛的理论背景，更好地理解我所做的和我要去的地方，因为它为我们的Drools R＆D提供了几乎所有理论方面的支持。它确实是一个巨大而引人入胜的主题，它将为那些花时间学习的人带来红利。我知道它确实存在并且仍然适合我。Bracham和Levesque写了一篇开创性的作品，名为“知识表示和推理” 对于想要建立坚实基础的人来说，这是必读的。我还会推荐Russel和Norvig的书“人工智能，一种现代方法”，它也涵盖了KRR。
- Rules Engines and Production Rule Systems (PRS) `规则引擎和生产规则系统（PRS）`

  Drools引擎是为开发人员提供KRR功能的计算机程序。从高层次来看，它有三个组成部分：
    - 本体论
    - 规则
    - 数据
- Hybrid Reasoning Systems (HRS) `混合推理系统（HRS）`
    - 正向链接是“数据驱动的” `正向推理`
    - 向后链接是“目标驱动的” `反向推理`
- Expert Systems `专家系统`

Rete算法
=
Rete算法可以分为两部分：规则编译和运行时执行。

- 编译算法描述了如何处理生产存储器中的规则以生成有效的区分网络。

PHREAK算法
=
它是一种渴望的，面向数据的算法。所有工作都是使用插入，更新或删除操作完成的，并急切地为所有规则生成所有部分匹配。

Assert类型
=
- `decision` 决策 具体条件判断
- `Model` 实体模型 OOP
- `process` 过程  整体流程
- `Form` 搞啥子？？
- `Optimization` 优化 暂时不知道搞啥子



实战
===
   - 第一步
     - `Menu ->Project ->Add Project` 高级设置可修改命名空间
     - 点击对应项目 `Add Asset`->选择`Data Object`添加实体对象模型
     - 选择 `DRL file` 添加规则文件
     ```
     package com.myspace.demo;
     
     import com.myspace.demo.model.Goods;
     import com.myspace.demo.model.RuleMsg;
     
     rule "Is of valid goods price"
         salience -50
     when
         Goods( cosePrice < 1 )
         $a : RuleMsg()
     then
         $a.setValid( false );
     end
     
     rule "Is of valid goods category"
     
     when
         Goods( categoryId< 1 )
         $a : RuleMsg()
     then
         $a.setValid( false );
     end
     ``` 
 启动 drool-wb
 =   
 
docker run -p 8080:8080 -p 8001:8001 -v ~/git/drools:/opt/jboss/wildfly/bin/.niogit:Z -d --name drools-wb jboss/drools-workbench-showcase:7.17.0.Final
 
 启动 kie-server
 =
 docker run -p 8180:8080 -d --name kie-server --link drools-wb:kie-wb jboss/kie-server-showcase:7.17.0.Final
 
 
 docker run -p 8080:8080 -p 8001:8001 -d --name drools-wb jboss/drools-workbench-showcase
 
 docker run -p 8180:8080 -d --name kie-server --link drools-wb:kie-wb jboss/kie-server-showcase
 
 Project面板中的Asset说明
 ===
 大类
 ==
 - Process 过程|流程|工序
 - Model 对象实体|枚举
 - Form  暂不清楚
 - Decision 决策|策略|规则设计
 - Other 包啥的
 - Optimization 配置啥的
 
 Decision
 ==
 - Decision Table (Spreadsheet) 决策表（电子表格形式）
 - DMN
 - **DRL file 按drools语法手写决策**
 - DSL definition DSL定义
 - Global Variable(s) 全局变量
 - **Guided Decision Table 向导型决策表**
 - Guided Decision Table Graph 向导型决策表图标
 - Guided Decision Tree 向导型决策树
 - **Guided Rule 向导型规则** 配置面板创建单个规则并返回结果（单一规则）
 - **Guided Rule Template 向导型规则模板** 配置面板创建多策略表格 适用于多种组合条件判断并返回不同的结果 （多个规则）
 - Guided Score Card 向导型积分卡
 - Score Card (Spreadsheet) 积分卡（电子表格形式）
 - Test Scenario 测试方案（表格形式）
 - Test Scenario (Legacy) 测试方案（旧的方式）
 

注：
谷歌机翻：Guided 向导 指导 引导
 
  
Guided Decision Table 向导型决策表
==
- Hit Policy 命中策略
  - None 可以执行多行，并且验证会警告冲突的行。已上载的任何决策表（使用非指导决策表电子表格）将采用此命中策略。
  - Rule Order 同上 只是不会警告
  - Resolved Hit 根据优先级命中
  - First Hit 将应用列表中首先满足的折扣，并忽略其他折扣
  - Unique Hit 决策表中 每行规则唯一命中1次 
- Add Column 添加列
  - Add a Condition 添加条件
  - Add a Condition BRL fragment 添加条件BRL片段
  - Add a Metadata column 添加元数据列
  - Add an Action BRL fragment 添加Action BRL片段
  - Add an Attribute column 添加属性列
  - Delete an existing fact 删除现有事实
  - Execute a Work Item 执行工作项
  - Set the value of a field 设置字段的值
  - Set the value of a field with a Work Item result 使用工作项结果设置字段的值

# DRL规则关键字说明 官方文档19.8
- `package` 声明包
- `import` 导入包
- `function` 自定义方法 较复杂的逻辑处理
- `query` 待完善。。。
- `declare` 对象声明
- `rule end` 声明规则主体
- `when` 后跟 判断主体 相当于 if后紧跟内容 截止到{}
- `then` 相当于if后紧跟的{} 注意 没有else或else if
- `and or` 且 或
- `exists` 存在事实 必须放在第一判断位置(待验证是上下行还是单行内外) 
- `not` 跟exists相反 不纯在的事实
- `forall` 待完善
- `from` 迭代器 比较两个集合ListA,ListB,求集合ListA是否包含集合ListB中的全部元素，java内置的有containsAll好像用不了，折中方法 以String集合为例 not(String(ListA not contains this.toString()) from ListB) 双重否定得ListA一定包含ListB的每个元素
- `entry-point` 待完善
- `collect` 收集 待完善
- `accumulate` 迭代


# Decision Table (Spreadsheet) 决策表（电子表格形式）
决策表的电子表格可以包含多个RuleTable区域，但只包含一个RuleSet区域。  
建议一个包只用1个xlsx表格，因为多个很容易发生规则等定义的冲突。  
RuleSet电子表格的区域定义要全局应用于同一包（不仅是电子表格）中的所有规则的元素，例如规则集名称或通用规则属性。  
RuleTable区域定义实际规则（行）以及在指定规则集中构成该规则表的条件，操作和其他规则属性（列）。

### RuleSet
`RuleSet` 定义当前规则的包名 com.colipu.hello.OrderCart  
`Sequetial` bool true时表格中的rule集合在转换成dsl规则语言时将从上至下添加 salience ,**表格中单独定义的salience同时失效**  
`SequentialMaxPriority` salience的最大值设定 默认为65535  
`SequentialMinPriority` salience的最小值设定 默认为0

举例：

Sequential true   
SequentialMaxPriority 10  
SequentialMinPriority 2  
如果有8个rule 那么 转换时自动添加 salience 10  salience 9 ... salience 3  
如果有20个rule 那么 报错 因为设定的范围不够

**注意**  
使用测试时 Audit log中的打印日志顺序可能不正确。  

为true，执行顺序从上至下   
但是如果在输入对象中加一个List,每个规则记录Add一个标识，可以确认是按顺序执行的。

为false,执行顺序按salience 从大到小执行，相同等级的先定义先执行。

`EscapeQuotes` 省略或true，转义引号。
 
`Import` 导入对象 com.colipu.hello.SaleGroupInfo,java.util.Arrays,java.util.List

`Variables` 全局变量 待完善。。。

`Functions` 自定义方法函数
```
  function String hello(String applicantName) {
      return ""Hello "" + applicantName + ""!"";
  }
```
`Queries` 待完善。。。

`Declare` 待完善。。。

### RuleTable
`NO-LOOP` 不打环

When this option is set to true, the rule cannot be reactivated (looped) if a consequence of the rule re-triggers a previously met condition.

当将此选项设置为true，如果规则的结果重新触发先前满足的条件，则无法重新激活（循环）规则。

`ACTIVATION-GROUP` 设置激活规则组名称
主要作用 在同一个组中的将按照定义时顺序执行，满足一次立即退出。  
**如果不分组，所有满足条件的规则都会被执行！！！！**

`AGENDA-GROUP` 规则分组 例如分组：groupa  
`RULEFLOW-GROUP` 当关联的分组激活时触发当前规则或规则组 例如：groupa  
`AUTO-FOCUS` 自动激活当前组  

情况一
AUTO-FOCUS | AGENDA-GROUP | RULEFLOW-GROUP
----- | ----- | ------
false | groupa | groupb
true | groupb | groupa

如果这种循环情况，只会执行groupb所在分组


情况二
AUTO-FOCUS | AGENDA-GROUP | RULEFLOW-GROUP
----- | ----- | ------
false | groupa | groupb
true | groupb |

groupb自动激活后会触发groupa的激活，两组规则 按照 定义顺序执行。

not(String($s.skuCodes not contains this.toString()) from ("$param".split(":")))

["示例表格"](assets/drool-rule.xlsx)