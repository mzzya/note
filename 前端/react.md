# React

## 核心概念

- element 元素
  - 比如我们常见的标签`<div></div>`就是一个基础元素
  - `<CustomDiv/>`这种自己封装的也叫元素
- component 组件
  - 由元素组成，功能模块的封装。
  - 函数组件 `function(){return 元素}` 无状态展示首选（参考阿里ant design输出的dav最佳实践）
  - Class组件 `class Welcome extends React.Component{render()}`
    - class 类
      - constructor 构造函数
      - func() 方法 render()就是一个方法 也是特殊函数
      - staic func() 静态方法 特殊静态函数
    - extends 继承
      - constructor 内 super()调用父类构造函数 java c#

### 受控组件和不受控组件

受控组件 表单元素 设置了value后必须onChange定义事件处理，否则页面没变化。

不受控组件 表单元素 可以设置defaultValue 正常输入

[差异参考](https://goshakkk.name/controlled-vs-uncontrolled-inputs-react/)

## 符号

- {} 放表达式
- () jsx语句如果是多行的需要用此包住

## props 属性

不能修改，只读，public

```jsx
// UserInfo是自定于组件 user是组件属性 组件内通过 props.user获取 如果是class组件则使用this.props.user
<UserInfo user={props.author} />

// 这样会把所有父属性传递个子组件，但是看情况不要滥用
<UserInfo {...props} />
```

## state 状态

state类似小程序data，private

```jsx
constructor(){
    this.state={name:'哈哈哈'}
}
onChange(){
    this.setState({name:'呵呵呵'})
}
```

## 生命周期

`render`方法必须实现

- 挂载 加载 初始化
  - constructor()
    - 给state中的字段赋值props的字段无意义
  - static getDerivedStateFromProps()
  - render()
  - componentDidMount() 组件挂载后（插入 DOM 树中）立即调用
    - 获取列表数据后 通过setState传入渲染

- 更新 变化
  - static getDerivedStateFromProps()
    - render之前调用
    - 让组件在 props 变化时更新 state
  - shouldComponentUpdate()
    - 返回false 会阻止render
    - 性能优化点 不建议触发判断是否要更新 而是采用PureComponent(只适合静态数据展示，不适合数据变更)
  - render()
  - getSnapshotBeforeUpdate()
    - 最近一次渲染输出（提交到 DOM 节点）之前调用。它使得组件能在发生更改之前从 DOM 中捕获一些信息（例如，滚动位置）。此生命周期的任何返回值将作为参数传递给 componentDidUpdate()。
  - componentDidUpdate() 更新后立即调用 第一次渲染不执行

- 卸载 删除 析构
  - componentWillUnmount()

只列举了部分事件，有些事件不建议使用或未来版本移除，参考[官网说明](https://zh-hans.reactjs.org/docs/state-and-lifecycle.htmlhttps://zh-hans.reactjs.org/docs/state-and-lifecycle.html)

## Context 上下文

顾名思义 组件多层嵌套，最内层依赖最外层的，传统方法要一层层传递。context包裹住可直接透传下去，但是不建议使用。
解决多层嵌套的方法，可把内层组件作为内层组件父级参数属性参数向下层传递。

## Refs

```js
React.createRef();
```

个人感觉像是局部的状态寄存器，局部行为控制，无需返回给调用方的。`函数组件`不可用，`class组件`可用。


## React.Fragment和<>

组件组合器，假设有两个不同的列表元素，需要组合到1个父级组件中，一般做法用`<div>2个组件</div>`包住会形成1个父级div，但是我们不希望，此时可以用`<></>`包住，这样不会渲染。类似小程序block。

## HOC

## Portals

将子节点渲染到存在于父组件以外的 DOM 节点的优秀的方案。

## 坑

- `<UserInfo {...this.state} />`这样传入到UserInfo组件中，父组件的状态state就变成了子组件的属性props
  - 子组件无法直接修改属性
  - 若通过state接收属性值再修改，只能改子组件的，不会改父组组件的显示。
  - 若子组件直接prop展示，父组件改state，子组件会变。
  - `所以`子组件要想在修改后触发父组件展示变更，应将子组件onChange事件提升为属性，父组件中定义
  Change事件传递个子组件.
  - {...this.state}谨慎使用...，这样会把不必要的也传下去。
- form表单的input select textarea等如果设置value={state.field}，那么需要在onChange里处理赋值问题setState({field:e.targe.value})，否则value不会变化。
- 依购物车为例，静态展示的数据区域（名称价格图片等）和有操作的加减框应拆分为不同的子组件，
子组件应采用函数式无状态组件，购物车最外层组价进行组合，并为子组件提供事件处理的函数属性。

## 总结

- 多组合，少继承
- 万物皆对象，元素 组件 事件 全部都可以当成组件属性向下传递。控制反转。
- 组件的粒度拆分要把握好，可复用。
- 状态提升，子组件最好是无状态函数组件。父状态组件用Class组件。
- 给子组件传递事件属性时不要直接传递匿名函数，应将函数绑定到父组件属性上再传递。好处是，父组件触发子组件更新时不需要再次为匿名函数分配内存空间。

## 参考资料

新手问题 https://dvajs.com/knowledgemap/#javascript-%E8%AF%AD%E8%A8%80

## 实战

```shell
# 方案一 创建简单应用
npx create-react-app react-app
cd react-app
# 方案二 创建生产级别应用
npx @umijs/create-umi-app
npm i @umijs/preset-react -D
```