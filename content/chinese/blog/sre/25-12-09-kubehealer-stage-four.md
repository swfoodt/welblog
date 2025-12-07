---
title: "KubeHealer 开发实录 (4)：CLI 工具的多模态报告设计"
date: 2025-12-07
author: "swfoodt"
description: "一个优秀的运维工具不仅要有强大的诊断逻辑，还要有友好的输出形式。本文记录了 KubeHealer 如何通过架构解耦，实现终端表格、Markdown、JSON 和 HTML 四种格式的输出，并重点探讨了 Go 模板技术在生成可视化报表中的应用。"
categories: ["Go", "CLI", "Frontend"]
tags: ["TableWriter", "GoTemplate", "Bootstrap", "JSON"]
---

### 1. 架构解耦：Analyzer vs Reporter

在 Week 2 结束时，我们的诊断逻辑和打印逻辑是耦合在一起的。为了支持多种输出格式，Week 3 的第一步就是实施 **关注点分离 (Separation of Concerns) 的设计思想**。

我们定义了标准的 `DiagnosisResult` 结构体，使得 `Analyzer` 只负责生产纯净的数据，而具体的展示工作交由 `cmd` 层根据用户参数（`--output`）分发给不同的渲染函数。

### 2. 终端体验优化：TableWriter

为了让 CLI 输出更具可读性，我们引入了 `tablewriter` 库。相比于简单的 `fmt.Printf`，它支持：
* **自动换行**：避免长错误信息撑爆终端。
* **ASCII 边框**：清晰分隔不同区域。
* **对齐控制**：让数据列整齐划一。

### 3. 可视化进阶：HTML 与 Go Templates

为了生成让人喜欢看的报告，利用 Go 标准库的 **html/template** 实现了 HTML 生成器。

**技术亮点**：
1.  **单文件分发**：将 HTML 模板作为字符串常量嵌入 Go 代码，避免了发布时需要携带外部 `.html` 文件的问题。
2.  **前端集成**：引入了 **Bootstrap 5** CSS 框架，通过卡片（Card）和徽章（Badge）组件，快速构建了现代化的 UI。
3.  **时间轴可视化**：利用 CSS 绘制了 Events Timeline，将枯燥的日志列表转化为直观的时间线，帮助用户快速理解故障发生的时间顺序。

### 4. 机器交互：JSON 与 Tags

为了让 KubeHealer 能集成到 CI/CD 流水线或被其他平台调用，支持了 JSON 输出。
在 Go 中实现这一点非常简单，只需在结构体字段后添加 **结构体标签(Struct Tags)**，例如 `` `json:"pod_name"` ``，然后调用 `json.MarshalIndent` 即可。

### 5. 成果总结

至此，KubeHealer 不再是一个简单的脚本，而是一个具备完整“输入-分析-输出-归档”能力的成熟工具。