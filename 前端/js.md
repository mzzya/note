# javascrit

## ==和===

```javascript
1=="1" true
1==="1" false
null==undefined true
null===undefined false

// 特殊场景
function Person(uname, age) {
    this.uname = uname;
    this.age = age;
    this.say = function() {
        console.log('我叫' + this.uname + '，今年' + this.age + '岁。');
    }
}
var zhangsan = new Person('张三', 18);
var lisi = new Person('李四', 20);
zhangsan.say == lisi.say false
zhangsan.say === lisi.say false

```

null 定义了变量，但是值为空
undefined 未定义变量

## BOM

Brower Object Model 浏览器对象模型

- Window 对象表示浏览器中打开的窗口
- Navigator 浏览器信息
- Screen 客户端显示屏幕的信息
- History 访问过的 URL  window.history
- Location 当前 URL 的信息 window.location
  - location.href 常见跳转方法
  - location.replace 不会产生history的跳转发放 假设有个A页面跳转到B页面进行回话初始话，然后replace到C页面，那么浏览器后退时可以不经过B页面后退到A页面

## DOM

Document Object Model 文档对象模型


关键在于 object model 可以这么理解`面向对象操作文档`

document.getElementById("id").innerHTML="";

document.getElementById("id").style.color="red";

document.getElementById("id").onclick=function(){};

document.getElementById("id").appendChild(document.createElement("p"));
