---
author: swfoodt
categories:
    - 前端基础
    - HTML
date: "2021-05-31T15:00:00+08:00"
description: ' HTML 的基础结构样例'
docmeta:
    id: 前端基础
    path: 前端基础/html
    title: "HTML 部分样例"
    weight: 2
draft: false
slug: html-example
tags:
    - HTML
    - 标签语义
    - 网页结构
title: HTML 部分样例
---



### 说了些定义之后， html 具体可以做到什么呢？
在本记录中，通过一个简单的 HTML 页面示例来展示 HTML 的基本用法和常见标签。  
创建一个简单的网页，包含标题、段落、链接、图片、表格、列表等常见元素。

---

#### 以下是各种标签，以及其渲染结果

---

{{< tabs >}}
{{< tab "运行结果">}}
<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8" />
  <meta http-equiv="X-UA-Compatible" content="IE=edge" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>HTML 基础示例</title>
</head>

<body>

  <h1>这是一个一级标题（h1）</h1>
  <h6>h1~h6 表示从大到小的六级标题，下面使用 <code>hr</code> 标签插入一条分隔线：</h6>
  <hr />

  <p>
    使用 <code>p</code> 标签创建了一个段落。我们还可以使用 <code>br</code> 标签实现换行效果：<br />
    这是换行后的内容。
  </p>

  <p>
    使用 <code>b</code> 标签可以<strong>加粗</strong>文字，<code>i</code> 标签则用于<span style="font-style: italic;">斜体显示</span>。
  </p>

  <a href="https://swfoodt.netlify.app">点击这里访问我的博客（使用 a 标签创建超链接）</a>

  <hr />

  <p>下面通过 <code>img</code> 标签插入一张游戏背景图：</p>
  <img src="https://swfoodt-blog.oss-cn-beijing.aliyuncs.com/img/bg/indexbackground10.png"
    alt="背景图" style="width: 200px; height: 100px;" />

  <hr />

  <p>使用 <code>table</code>、<code>tr</code>、<code>td</code> 标签创建一个简单表格，列出我喜欢和不喜欢的食物：</p>
  <table border="2">
    <tr>
      <td>喜欢的</td>
      <td>不喜欢的</td>
    </tr>
    <tr>
      <td>烤鸭</td>
      <td>白菜</td>
    </tr>
  </table>

  <hr />

  <p>使用 <code>ul</code> 和 <code>li</code> 标签创建一个无序列表：</p>
  <ul>
    <li>白菜</li>
    <li>菠菜</li>
    <li>芹菜</li>
  </ul>

  <p>使用 <code>ol</code> 和 <code>li</code> 标签创建一个有序列表：</p>
  <ol>
    <li>烤鸭</li>
    <li>烤鸡</li>
    <li>烤鱼</li>
    <dd>— 这些我都喜欢</dd>
    <dd>— 使用 <code>dd</code> 标签添加额外说明</dd>
  </ol>

  <hr />

  <div style="width: 500px; height: 85px; background-color: aqua;">
    <p>
      这是一个使用 <code>div</code> 标签创建的容器，常用于网页布局与区域划分。
      <br />
      还可以使用 <code>span</code> 标签来对文字进行局部样式调整：
      <span style="color: blue;">比如这段文字设置了蓝色</span>。
    </p>
  </div>

  <div>
    <p>接下来使用 <code>form</code> 标签创建一个简单的表单，让用户填写信息：</p>
    <form>
      用户名：<input type="text" name="username" /><br />
      密码  ：<input type="password" name="password" /><br />

<p>请选择性别：</p>
      <input type="radio" name="sex" value="male" /> 男
      <input type="radio" name="sex" value="female" /> 女

<p>选择你喜欢的水果：</p>
      <input type="checkbox" name="fruit" value="apple" /> 苹果
      <input type="checkbox" name="fruit" value="banana" /> 香蕉
      <input type="checkbox" name="fruit" value="pear" /> 梨

<p>点击按钮提交表单：</p>
      <input type="submit" value="提交" />
    </form>
  </div>

</body>

</html>


{{< /tab>}}

{{< tab "源代码">}}
```html
<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8" />
  <meta http-equiv="X-UA-Compatible" content="IE=edge" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>HTML 基础示例</title>
</head>

<body>

  <h1>这是一个一级标题（h1）</h1>
  <h6>h1~h6 表示从大到小的六级标题，下面使用 <code>hr</code> 标签插入一条分隔线：</h6>
  <hr />

  <p>
    使用 <code>p</code> 标签创建了一个段落。我们还可以使用 <code>br</code> 标签实现换行效果：<br />
    这是换行后的内容。
  </p>

  <p>
    使用 <code>b</code> 标签可以<strong>加粗</strong>文字，<code>i</code> 标签则用于<span style="font-style: italic;">斜体显示</span>。
  </p>

  <a href="https://swfoodt.netlify.app">点击这里访问我的博客（使用 a 标签创建超链接）</a>

  <hr />

  <p>下面通过 <code>img</code> 标签插入一张游戏背景图：</p>
  <img src="https://swfoodt-blog.oss-cn-beijing.aliyuncs.com/img/bg/indexbackground10.png"
    alt="背景图" style="width: 200px; height: 100px;" />

  <hr />

  <p>使用 <code>table</code>、<code>tr</code>、<code>td</code> 标签创建一个简单表格，列出我喜欢和不喜欢的食物：</p>
  <table border="2">
    <tr>
      <td>喜欢的</td>
      <td>不喜欢的</td>
    </tr>
    <tr>
      <td>烤鸭</td>
      <td>白菜</td>
    </tr>
  </table>

  <hr />

  <p>使用 <code>ul</code> 和 <code>li</code> 标签创建一个无序列表：</p>
  <ul>
    <li>白菜</li>
    <li>菠菜</li>
    <li>芹菜</li>
  </ul>

  <p>使用 <code>ol</code> 和 <code>li</code> 标签创建一个有序列表：</p>
  <ol>
    <li>烤鸭</li>
    <li>烤鸡</li>
    <li>烤鱼</li>
    <dd>— 这些我都喜欢</dd>
    <dd>— 使用 <code>dd</code> 标签添加额外说明</dd>
  </ol>

  <hr />

  <div style="width: 500px; height: 85px; background-color: aqua;">
    <p>
      这是一个使用 <code>div</code> 标签创建的容器，常用于网页布局与区域划分。
      <br />
      还可以使用 <code>span</code> 标签来对文字进行局部样式调整：
      <span style="color: blue;">比如这段文字设置了蓝色</span>。
    </p>
  </div>

  <div>
    <p>接下来使用 <code>form</code> 标签创建一个简单的表单，让用户填写信息：</p>
    <form>
      用户名：<input type="text" name="username" /><br />
      密码：<input type="password" name="password" /><br />

      <p>请选择性别：</p>
      <input type="radio" name="sex" value="male" /> 男
      <input type="radio" name="sex" value="female" /> 女

      <p>选择你喜欢的水果：</p>
      <input type="checkbox" name="fruit" value="apple" /> 苹果
      <input type="checkbox" name="fruit" value="banana" /> 香蕉
      <input type="checkbox" name="fruit" value="pear" /> 梨

      <p>点击按钮提交表单：</p>
      <input type="submit" value="提交" />
    </form>
  </div>

</body>

</html>

```
{{< /tab>}}

{{< /tabs >}}

---

{{<notice "warning">}}
由于本段html使用hugo在md中直接渲染，所以渲染结果会受到本博客样式文件影响，实际运行结果有所不同。此处仅供参考。
{{</notice>}}
<!-- </TabItem>
<TabItem value="result" label="运行后结果">

<iframe height="1400px" width="100%" scrolling="no" title="swfoodt-blog-example" src="https://codepen.io/swfoodt/embed/VwxrboQ?default-tab=result&theme-id=dark" frameborder="no">
</iframe>

</TabItem>
</Tabs> -->

---

在上述的代码中我们使用了包括`<h1>`，`<h6>`，`<hr>`，`<p>`，`<b>`，`<i>`，`<a>`，`<img>`，`<table>`，`<tr>`，`<td>`，`<ul>`，`<li>`，`<ol>`，`<dd>`，`<div>`，`<span>`，`<input>`等标签，这些标签都是 html 中的基本标签，我们可以使用这些标签来创建我们想要的网页。

不过可以在运行后的结果中看到，这个简单网页呈现出的效果**并不美观**，因为在上述代码中，**除了限制部分元素的宽高以外**，我们没有使用其他的 css 样式来美化网页，有关 css 方面的代码应用可以参考本记录的[css 部分](#)。

还有一个很有用的的标签没有展示在上面的样例代码中，`iframe` 标签，**iframe 标签**可以用来在网页中嵌入其他网页，在上方的代码结果演示就是使用了 iframe 标签来嵌入了一个`codepen`的网页，若未能显示成功可能是网络问题，可以自己动手运行一下。

---

{{<notice "w3school">}}

日常使用可以参考`w3school`提供的[速查手册](https://www.w3school.com.cn/html/html_quick.asp)。

{{</notice>}}
