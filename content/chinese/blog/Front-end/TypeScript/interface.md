---
title: "interface 接口"
date: 2022-06-04T15:00:00+08:00
draft: false
categories: ["前端基础", "TypeScript"]
tags: ["TypeScript", "前端入门", "接口"]
author: "swfoodt"
description: "TypeScript 中的接口（Interfaces）是对行为的抽象，具体如何行动需要由类去实现。本文介绍了接口的基本用法和特点。"
slug: "interface"
---

# interface 接口

接口（Interfaces）是 TypeScript 的一个重要概念，它是对行为的抽象，而具体如何行动需要由类（classes）去实现（implement）。
<!--more-->
```ts
//定义一个对象应当长成什么样子
interface Person {
  //属性
  readonly id: number;
  name: string;
  age: number;
  sex?: string; //可选属性
}

let sw: Person = {
  id: 1234,
  name: "sw",
  age: 21,
};
```

上面的代码中，我们定义了一个接口在这个接口中有一个**只读元素 id**，用于表明其只读性质的关键字是 `readonly`，这样的话，我们就不能再给 id 赋值了。

`readonly` 与 `const` 的区别在于，`const` 用于修饰变量，而 `readonly` 用于修饰元素。

`sex?: string;` 这种表达方式表示性别是一个可选元素。

抽象地来说 interface 定义了一个对象应当包含哪些元素。
