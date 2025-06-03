---
title: "ienum 枚举"
date: 2022-06-08T15:00:00+08:00
draft: false
categories: ["前端基础", "TypeScript"]
tags: ["TypeScript", "前端入门", "枚举"]
author: "swfoodt"
description: "TypeScript 中的枚举（Enums）是一种特殊的类型，用于定义一组命名的常量。本文介绍了枚举的基本用法和特点。"
slug: "enum"
docmeta:
    id: TypeScript
    path: TypeScript/TS语法
    title: "枚举"
    weight: 5
---
# 枚举

## 为什么要使用枚举

对于枚举，我印象较深的是在编写 unity 的 C#脚本中用到过，当我们使用常量（const）时，有时会遇到一些场景需要限定这个常量的范围，这个时候就可以使用枚举。

枚举的值有两种类型：常量值（constant member）和计算值（computed member）。

## 代码实例

```ts
const enum Direction {
  Up = "UP",
  Down = "Down",
  Left = "Left",
  Right = "Right",
}

console.log(Direction.Down); // down

const value = "Right";
if (value === Direction.Right) {
  console.log("go right...");
}
```

{{<notice "info">}}
使用常量枚举时，枚举成员会被内联进使用它们的地方，可以提升性能。
{{</notice>}}
