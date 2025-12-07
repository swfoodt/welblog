---
title: "KubeHealer 开发实录 (3)：从面条代码到规则引擎的重构"
date: 2025-12-07
author: "swfoodt"
description: "随着诊断逻辑的增加，单一的函数变得难以维护。本文记录了 KubeHealer 在 Week 2 的一次关键架构重构：引入规则引擎（Rule Engine）模式，通过 Go Interface 实现诊断逻辑的解耦与可扩展性。"
categories: ["Go", "Architecture", "Design Pattern"]
tags: ["Refactoring", "Interface", "RuleEngine"]
---

### 1. 痛点：不断膨胀的 `if-else`

在项目初期，所有的诊断逻辑（OOM、ImagePull、Crash）都堆砌在 `GetContainerStatus` 一个函数中。随着功能的增加，这个函数迅速膨胀，面临以下问题：
* **可读性差**：各种故障的判断逻辑混杂在一起。
* **扩展困难**：每增加一种新故障（比如 Pending），都要修改核心代码，违反了 **开放-关闭原则(Open-Closed Principle)**。

### 2. 架构升级：引入规则引擎

为了解决上述问题，项目引入了规则引擎模式。

#### 2.1 定义规范标准 (Interface)

首先定义了 `Rule` 接口，规范了所有诊断规则的行为：

```go
type Rule interface {
    Name() string
    // Check 方法接收 Pod 上下文，返回诊断结果
    Check(pod *corev1.Pod, container *corev1.Container, status corev1.ContainerStatus) CheckResult
}
```
任何实现了这个接口的结构体，都可以被视为一个诊断规则。

#### 2.2 实现规则引擎 (Engine)
`RuleEngine` 负责维护规则列表，并依次调用它们。采用 **“短路 (Short-circuit)”** 策略：一旦某个规则发现严重问题（Matched=true），引擎立即返回结果，优先暴露最严重的错误，避免信息过载。

### 3. Pending 状态的特殊处理
在实现 `PendingRule` 时，遇到了一个特殊情况：`Pod` 尚未调度，导致 `ContainerStatuses` 列表为空。 为了让规则引擎能覆盖到这种情况，在 Analyzer 中实施了 **“虚拟状态构造 (Synthetic State)”** 策略，手动构造了一个 Dummy 状态对象传入引擎，成功触发了 PendingRule 的检查逻辑，将调度失败诊断纳入了统一的规则框架。

### 4. 成果
重构后，`OOMRule`、`ImagePullRule`、`CrashRule` 和 `PendingRule` 即使逻辑再复杂，也被物理隔离在独立的文件中。这不仅让代码更加整洁，也为即将引入的 **Informer 实时监控** 铺平了道路。