# css3

## border

- border-radius 圆角 左上角，右上角，右下角，左下角 可用百分比
  - 四个值: 第一个值为左上角，第二个值为右上角，第三个值为右下角，第四个值为左下角。
  - 三个值: 第一个值为左上角, 第二个值为右上角和左下角，第三个值为右下角
  - 两个值: 第一个值为左上角与右下角，第二个值为右上角与左下角
  - 一个值： 四个圆角值相同
- box-shadow:y轴 x轴 模糊距离 阴影

## transform:rotate(90deg) 旋转度数

## background-image

- linear-gradient 线性渐变
  - (#e66465, #9198e5);Y轴渐变
  - (to right, red , yellow); X轴渐变
  - to bottom right 对角线
  - deg 角度
- radial-gradient 径向渐变
  - circle 表示圆形，ellipse 表示椭圆形
  - red 5%


## 动画

```css3
@keyframes myfirst
{
    0%   {background: red; left:0px; top:0px;}
    25%  {background: yellow; left:200px; top:0px;}
    50%  {background: blue; left:200px; top:200px;}
    75%  {background: green; left:0px; top:200px;}
    100% {background: red; left:0px; top:0px;}
}
```


## box-sizing: border-box;

内边距也包含在width height 中