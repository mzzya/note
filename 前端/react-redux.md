# Redux 学习

Redux是一个独立的模块。可搭配React、Angular、Ember、jQuery 甚至纯 JavaScript。

[中文官网](https://www.redux.org.cn/)上介绍的一句话`Redux 是 JavaScript 状态容器，提供可预测化的状态管理`。

状态：就是对象，万物皆对象。。。

## 举例

### 案例一

假设我们用水壶作为容器，张三、李四、王五共用一个水壶，该水壶可加热，搅拌等。

张三想喝水，拿着水壶接了一壶自来水，不讲究直接从水壶里倒了一杯凉水就喝了。

李四女朋友来了大姨妈必须喝热水，那么李四需要按下水壶的加热开关，`烧开后`再倒出一杯热水。

李四女朋友喝了一口就不想续杯了，恰好此时王五想喝咖啡，直接往水壶里倒了一些咖啡进去，搅一搅就倒了出来。后边又加了点咖啡伴侣和糖。

#### 分析

水：状态，数据对象

水壶：容器，可以做一些操作，加热泡茶等

接自来水：状态初始化。

张三倒水：可预测倒出来的是杯凉水。

李四倒水：他加热了可预测倒出来的是杯开水。

王五倒水：加了很多东西，可能是纯咖啡，也可能是加过奶或糖的。

水的表象，也就是状态，在每个动作触发后都生成了一个新的表象。

倒水之前都向水壶发起了一个动作，加热，加咖啡，加奶，加糖，这些都是使用者提供的，都是可预测的状态，每个动作都是`action`，但是具体变化是水壶提供的，加热是水壶电热丝发热进行的，加咖啡要搅拌，不然就是一坨，倒出来的可能还是一点味都没有的白开水。`Reducers`就是发热丝，搅拌器。`Store`是水壶。
`Store.getState()`从水壶倒东西。`Store.dispatch()`加东西（简单触发加热，或者加配料）。`subscribe`订阅李四，王五需要等烧开或者泡好后才会倒。

额外：这只是个壶，不能用来煮饭（多功能电饭煲）烧菜（锅）。

### 案例二

依购物车为例，有个全局的购物车入口，入口上写着购物车内存在了几种商品。


## 基础知识介绍

### Array.reduce

reduce() 方法接收一个函数作为累加器，数组中的每个值（从左到右）开始缩减，最终计算为一个值。
reduce() 可以作为一个高阶函数，用于函数的 compose。

array.reduce(function(total `迭代结果`, currentValue `当前数组元素`, currentIndex `当前数组索引`, arr `数组`), initialValue `初始化的迭代结果`)

```js
var numbers = [65, 44, 12, 4];

function getSum(total, num) {
    return total + num;
}
function myFunction(item) {
    document.getElementById("demo").innerHTML = numbers.reduce(getSum);
}
```

combineReducers 组合状态

import { combineReducers, createStore } from 'redux'

创建一个联合状态管理器
let reducer = combineReducers({ visibilityFilter, todos })

创建一个容器管理
let store = createStore(reducer)

依冰箱为例

store 冰箱容器

visibilityFilter 面条

todos 苹果 香蕉 牛奶

reducer 相当于一个目录知道冰箱有那些东西。


`Redux 都不允许程序直接修改数据，而是用一个叫作 “action” 的普通对象来对更改进行描述。`
这句话的意思是：假设购物车内操作商品修改和删除，用户的操作只负责发送命令，实际上页面的数据更新是由命令的订阅者触发的更新。解耦用户的操作和页面更新。