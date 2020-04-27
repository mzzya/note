# 前端

## px em rem

任意浏览器的默认字体高都是16px。所有未经调整的浏览器都符合: 1em=16px。

- px px像素（Pixel）。相对长度单位。像素px是相对于显示器屏幕分辨率而言的。
- em 相对于当前对象内文本的字体尺寸。相对单位
  - 一般默认1em=16px
- rem 相对的只是HTML根元素。相对单位

## text-overflow

用于文字超长截断省略等操作

- text-overflow
  - clip 溢出时不显示省略标记(...) 限定高度要搭配 overflow:hidden
  - ellipsis 溢出时显示省略标记(...) 配合 overflow:hidden；white-space:nowrap
    - 有浏览器兼容性问题
- text-indent 文字首行缩进


- white-space
  - normal 默认。空白会被浏览器忽略。
  - pre 空白会被浏览器保留。其行为方式类似 HTML 中的 pre 标签。
  - nowrap 文本不会换行，文本会在在同一行上继续，直到遇到 br 标签为止。
  - pre-wrap 保留空白符序列，但是正常地进行换行。
  - pre-line 合并空白符序列，但是保留换行符。
  - inherit 规定应该从父元素继承 white-space 属性的值。
    - ie兼容问题
- ovverflow
  - visible 默认值。内容不会被修剪，会呈现在元素框之外。
  - hidden 内容会被修剪，并且其余内容是不可见的。
  - scroll 内容会被修剪，但是浏览器会显示滚动条以便查看其余的内容。
  - auto 如果内容被修剪，则浏览器会显示滚动条以便查看其余的内容。
  - inherit 规定应该从父元素继承 overflow 属性的值。

- word-spacing 字体之间的间距

- word-break 字体换行方式（针对英文单词） break-all 直接切断

## flex

- flex-direction 决定布局方向 横竖正反向
  - row 默认值 从左向右
  - row-reverse 从右向左
  - column 从上到下
  - column-reserve 从下到上

- flex-wrap  超长换行规则
  - nowrap 默认值 不换行 挤一挤
  - wrap 换行到下一行
  - wrap-reverse 换行到上方

- flex-flow: flex-direction flex-wrap; 合并两者


- justify-content X轴 对齐方式
  - flex-start 起点对齐 row 从左开始 row-reverse 从右开始
  - flex-end 终点对齐
  - center 1 元素元素 1
  - space-between 0 1 1 0
  - space-around 0.5 1 1 0.5
  - space-evenly 1 1 1 1

- align-items Y轴 对齐方式
  - flex-start 起点对齐 column 从上开始 column-reverse 从下开始
  - center 居中
  - baseline  容器第一行文字 对齐
  - stretch 无高度时顶格

- align-content 多根轴线 flex-wrap 多行时生效
  - flex-start 起点对齐
  - flex-end 终点对齐
  - center 中心轴对齐
  - space-between 边轴贴边 子轴间距相等
  - space-around 轴间距 按 0.5 1 1 1 0.5

- align-self 轴内 容器的对齐方式
  - 摆脱父级flex限制 参数有 flex-start  flex-end 等

- flex: flex-grow、flex-shrink flex-basis; 简写 类似border
  - flex-grow 扩展量
  - flex-shrink 收缩量 所有容器超过长度各自收缩比例。
  - auto 与 1 1 auto 相同。
  - none 与 0 0 auto 相同。

- flex-basis 容器初始长度 支持百分比。

## filter 滤镜

最近的网站变黑实现方法，不同浏览器要带兼容性前缀。

- blur(px) 高斯模糊。实现模糊度（马赛克），px代表多少个像素融合到一起。
- brightness(%)	亮度 黑0<1亮<更亮
- contrast(%)	对比度 黑0<1亮<更亮
- drop-shadow 阴影
- grayscale 灰度
- opacity 透明度


## visibility hidden 隐藏元素但是保持坑位

visible 默认 可见

## position

- static 默认值
  - top right bottom left 不可用

- absolute 绝对定位 相对于离元素最近的设置了绝对或相对定位的父元素决定的
  - top right bottom left 可用
  - 一旦元素设置了就会脱离文档流 后边的兄弟元素就会占领他的位置，重叠

- relative 相对自身正常位置定位
  - top right bottom left 可用

- fixed 相对于浏览器定位
  - top right bottom left 可用

## z-index

一旦使用了 position: absolute,relative,fixed就可能出现样式重叠问题

z-index值越高 越优先显示

## 选择器

- div p
  - div下的所有p元素
- div>p
  - div下的所有一级p元素
- div+p
  - div后边的第一个p元素
- div~p
  - div后边的所有p元素

## [伪类](https://www.runoob.com/css/css-pseudo-classes.html)

## 内联样式（行内）> 页级样式（页面style）> 外联样式（外部文件）
