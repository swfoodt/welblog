---
title: "css小tips"
date: 2021-05-31T15:00:00+08:00
draft: false
categories: ["前端基础", "CSS"]
tags: ["HTML", "css", "CSS特性", "前端入门"]
author: "swfoodt"
description: "介绍 CSS 的一些小技巧和常用属性"
slug: "css-tips"
docmeta:
    id: 前端基础
    path: 前端基础/CSS
    title: "CSS 小tips"
    weight: 6
---
### 属性

##### `box-sizing: border-box;`：这个属性**可以让我们设置的盒子宽高包括内边距和边框**。

##### `overflow: hidden;`：这个属性可以让我们设置的盒子超出部分隐藏。**可以解决嵌套块级元素时发生的塌陷现象**。
<!--more-->
### 实操遇到的问题

#### 脱离文档流的方式

`float：left/right`、`position：absolute/fixed`会脱离文档流。
但是`position：relative`不会脱离文档流。

#### css 画三角形

- 将盒子宽高设为 0，颜色设置突出一边的 border 颜色，就能得到一个三角形：

```css
div {
  width: 0;
  height: 0;
  border: 20px solid transparent;
  border-top-color: red;
}
```
