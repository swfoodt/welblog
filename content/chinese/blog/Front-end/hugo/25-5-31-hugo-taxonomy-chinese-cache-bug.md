---
title: "Hugo 分类页面中文路径未更新的问题排查记录"
date: 2025-06-01
author: swfoodt
categories: ["技术笔记"]
tags: ["Hugo", "静态博客", "缓存", "模板渲染", "中文路径"]
description: 本文详细记录了 Hugo 分类页面在使用中文路径时渲染模板未生效的问题，以及最终如何通过强制清除缓存解决该 bug。
slug: hugo-taxonomy-chinese-cache-bug
---

### 前言

在使用 Hugo 搭建博客并启用分类页面功能时，遇到了一个非常诡异的问题：**中文分类路径的页面渲染结果和英文分类页面不一致**，并且中文页面似乎始终渲染的是「另外一个模板」。
<!--more-->
即便已经更新了 `layouts/_default/taxonomy.html` 文件。本文将对这个问题的复现过程、调试路径、根因剖析以及最终解决方案进行完整记录。

---

### 问题复现

在 Hugo 项目中，启用了分类功能（categories），并配置了 `layouts/_default/taxonomy.html` 来渲染每个分类页面。此时发现：

- 英文分类路径如：`http://localhost:1313/categories/internet/` 页面渲染正常 ✅
- 中文分类路径如：`http://localhost:1313/categories/linux基础/` 页面渲染旧内容，模板修改无效 ❌

即使把 `taxonomy.html` 文件全部清空，只写一句：

```html
<h1>hello</h1>
```

结果仍然是：

- 英文分类页正确输出 `hello`
- 中文分类页仍然输出旧版模板内容（甚至包括已经删除的 `<div>` 结构）

---

### 初步排查尝试

尝试了以下操作：

- ✅ 检查 `.Params`, `.RelPermalink`, `.Kind`, `.Type` 等上下文是否一致
- ✅ 加入调试语句 `<pre>{{ printf "%#v" . }}</pre>` 观察渲染上下文
- ✅ 使用 `<p>使用模板：taxonomy.html</p>` 来确认是否真正命中当前模板
- ❌ 中文路径页面始终未发生变化，甚至调试输出都无法生效

---

### 误判的方向

一度怀疑是：

1. 中文路径的 taxonomy 页面并未使用 `layouts/_default/taxonomy.html`
2. 是否存在 `layouts/categories/` 或 `layouts/_default/categories.terms.html` 优先级覆盖
3. Hugo 内部是否对中文 slug 做了路径拆分或路径映射

但最终这些方向都被排除了。**问题其实出在 Hugo 的缓存机制**。

---

### 根因分析：Hugo 缓存机制 + Fast Render

在 Hugo 启动时，如果使用了开发模式：

```bash
npm run dev  # 实际等效于：hugo server --disableFastRender
```

即便关闭了 fastRender，**Hugo 仍然会在以下位置缓存构建中间结果**：

- `.hugo_cache/`：构建缓存
- `.hugo_build.lock`：构建状态锁
- `resources/`：资源文件缓存
- `node_modules/.vite/`：前端构建缓存（若使用 Vite）

**这些缓存在路径较复杂、字符集为中文时**，极容易出现模板 hash 未刷新，**导致中文路径页面长时间使用旧版模板结果**。

---

### 解决方案：彻底清除所有缓存

以下是一键清理脚本，推荐在根目录执行：

```bash
# 清除 Hugo 所有缓存文件夹
rm -rf .hugo_cache/ .hugo_build.lock resources/ public/

# 如果你使用的是 hugoplate 模板（含前端构建）
rm -rf node_modules/.vite .vite dist/

# 可选：重新安装依赖
rm -rf node_modules/ package-lock.json
npm install
```

清理完成后重新运行：

```bash
npm run dev
```

然后强制刷新浏览器：

```text
Ctrl + Shift + R 或 Command + Shift + R
```

---

### 效果验证

在 `taxonomy.html` 文件中加入：

```html
<!-- layouts/_default/taxonomy.html -->
<p>使用模板：taxonomy.html</p>
```

最终效果：

| 分类页面路径 | 渲染结果          |
|--------------|-------------------|
| `/categories/internet/` | ✅ 正确显示调试信息 |
| `/categories/linux基础/` | ✅ 成功更新并显示侧边栏 |
| `/categories/css/` | ✅ 正常响应模板逻辑 |

问题彻底解决 ✅

---

### 总结

- Hugo 的渲染机制依赖 `.hugo_cache` 中间产物，尤其是在中文路径存在时极易触发 hash 异常
- 不管你是否使用 `--disableFastRender`，**本地开发建议定期清理 `.hugo_cache/`**
- 如果页面渲染异常，但你“确信”模板改了无效，一定要考虑是缓存问题
- 中文路径会更容易触发路径 hash 异常、页面 fallback 等问题

---

{{< notice "tip" >}}
如果 Hugo 页面更新失败，调试输出无效，**优先执行一次清缓存再继续排查**。
{{< /notice >}}

