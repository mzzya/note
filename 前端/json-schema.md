# json-schema

json-schema 是一种json格式数据格式规范。便于开发人员理解使用。仅列举常见格式，更多请至[官网](http://json-schema.org/)

## 数据类型

### string 字符串

- length 长度限制
  - minLength 最小长度
  - maxLength 最大长度
- pattern `^(\\([0-9]{3}\\))?[0-9]{3}-[0-9]{4}$`
- format 仅列举常见格式
  - date-time `ISO8601` `2018-11-13T20:20:39+00:00`
  - date `2018-11-13`
  - email
  - hostname
  - ipv4
  - ipv6
  - uri
  - regex 说明存的是一个正则表达式

### number integer 数字类型

- minimum 最小值
- exclusiveMinimum `true` x > minimum `false` x ≥ minimum
- maximum 最大值
- exclusiveMaximum true `true` x < maximum `false` x ≤ maximum

### object

- properties KV结构
  - key 可以认为是子的schema名称
  - value 子的schema
- additionalProperties `bool类型`是否可以扩展属性字段
- required 必须字段 数组类型 `properties keys`
- propertyNames 属性名称限制 `不常用`
  - pattern 格式限制 正则
- minProperties 最小属性数量 `不常用`
- maxProperties 最大属性数量 `不常用`
- dependencies KV结构 依赖关系 key（字段名）填写了 value（字段名）就必填 `不常用`
- patternProperties KV结构 `不常用`
  - key 字段名正则
  - value schema


### array

items {}或[] 字段schema
contains 部分语言中数组（集合，列表，切片）内可包含不同类型数据 必须要包含的类型 `不常用`
minItems 最大数量
maxItems 最小数量
uniqueItems 元素是否唯一

### boolean


### 通用字段

- enum 枚举
- title 简称
- description 详细描述
- default 默认值
- examples 示例
- const 常量
- contentMediaType 媒体类型
  - text/html
  - image/png
- contentEncoding 编码
  - binary
  - base64
- $schema `http://example/schema#`
- definitions `父级通用字段` {address:{}}
- $ref 引用 `#/definitions/address`

## 大道至简

很多规范 anyOf allOf not if then else 并没有列举的原因是：

会增加开发人员负担！

会增加开发人员负担！

会增加开发人员负担！

一个人单挑的项目不需要，多人的项目只需要简单的规则就好了，不是所有人都愿意学习复杂规则！
