---
title: "css盒子模型"
date: 2021-05-31T15:00:00+08:00
draft: false
categories: ["前端基础", "CSS"]
tags: ["HTML", "css", "盒子模型", "前端入门"]
author: "swfoodt"
description: "介绍 CSS 盒子模型，包括内容、内边距、边框和外边距，并解释如何使用 box-sizing 属性实现所见即所得的效果"
slug: "css-box-model"
docmeta:
    id: 前端基础
    path: 前端基础/css
    title: "CSS 盒子模型"
    weight: 5
---
### 盒子模型

- html 标签中**每个标签**都可以看作是一个盒子，盒子模型中包括**内容(content)**、**内边距(padding)**、**边框(border)**、**外边距(margin)**。

#### 内容部分

- 下面是一个简单的盒子模型。  
  ![](https://swfoodt-blog.oss-cn-beijing.aliyuncs.com/img/blog-docs/20221004182736.png)

- 内容部分是指标签中的内容，如下图所示的蓝色区域:  
  ![](https://swfoodt-blog.oss-cn-beijing.aliyuncs.com/img/blog-docs/20221004182901.png)

这个盒子模型的 css 代码如下:

```css title="盒子模型css代码"
div {
  width: 100px;
  height: 100px;
  background: pink;
  padding: 20px;
  margin: 30px;
}
```

从代码中可以看到，我设置了一个宽度为 100px，高度为 100px，背景色为粉色的盒子，同时设置了内边距为 20px，外边距为 30px。

但是值得注意的是这个盒子背景颜色覆盖的粉色区域实际上是 140px，原先我认为我在 css 中设置的盒子宽高就是粉色背景覆盖的区域，我设置`width: 100px; height: 100px;`,那么这个盒子就应该是 100px，但是实际上 100px **只是盒子的内容部分的宽高，而盒子的宽高是包括内边距和边框的**，所以这个盒子的宽高是 140px。

那么如何让我们“所见即所得”呢，可以使用 css 中的`box-sizing: border-box;`属性，这个属性**可以让我们设置的盒子宽高包括内边距和边框**。

#### 内边距

- 内边距 padding 所指的是盒子的内容与边框之间的距离

```css title="padding的复合写法"
div {
  /* 上内边距 上内边距 右内边距 左内边距; 顺时针进行赋值*/
  padding: 10px 20px 30px 40px;
}
```

{{<notice "info">}}

当我们使用 padding 的复合写法时未写够四个值时，编译时会按照顺时针的顺序经行赋值，直到将所有的值赋完。剩下未被赋值的方向会使用对称方向的值进行赋值。

{{</notice>}}

#### 边框

- 边框 border 可以有许多样式。但是常用的一般就以下三种:

  - solid 实线
  - dashed 虚线
  - dotted 点线

```css title="边框的样式"
div {
  /* 边框宽度 边框样式 边框颜色;*/
  /* 实线 */
  border: 1px solid red;
  /* 虚线 */
  border: 1px dashed red;
  /* 点线 */
  border: 1px dotted red;
}
```

#### 外边距

- 外边距 margin 的属性写法与内边距 padding 的属性写法一致，所以不再赘述。

### 外边距折叠现象

#### 合并现象

当两个垂直布局的相邻盒子都设置了**垂直外边距**时，这两个盒子的外边距会发生合并现象，合并后的外边距的大小为两个盒子外边距的最大值。

这个现象带来的影响个人认为不是很大。但是在一些特殊的场景下，这个现象会带来一些不可预知的问题。

#### 塌陷现象

当两个盒子为嵌套关系时，设置子元素的 margin-top 时，会导致**父元素发生塌陷现象**，父元素的**位置将会与子元素一同下移**。

解决方法：设置父元素的`overflow: hidden;`属性。

#### 行内元素的外边距

行内元素的外边距，**垂直方向上无法生效**，水平方向上可以生效。也就是说要控制行内元素的垂直方向上的外边距，需要设置行高，或是将行内元素设置为块级元素。
