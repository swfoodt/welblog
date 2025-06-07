---
title: "虚拟机以及linux系统的安装使用"
meta_title: "虚拟机以及linux系统的安装使用"
description: "虚拟机的安装以及Linux系统的安装和使用"
date: 2025-06-06
categories: ["SRE笔记", "Linux基础"]
author: "swfoodt"
tags: ["虚拟机", "Linux", "SRE", "云计算", "系统安装"]
draft: false
slug: "sre-vmware"
---

### 虚拟机以及Linux系统的安装使用

#### 一、准备工作

整个流程分为以下几个步骤：

1. 安装虚拟机软件（ VMware Workstation ）
2. 下载 Linux 系统镜像（建议使用 Ubuntu 以及 Rocky linux）
<!--more-->
3. 创建并配置虚拟机（分配 CPU、内存、磁盘等资源）
4. 安装 Linux 系统（完成基本配置）
5. 安装并配置 Xshell 工具（用于远程连接终端）
6. 使用 SSH 连接并开始使用 Linux 系统

---



#### 二、安装虚拟机软件

- VMware Workstation Pro 目前对于个人用户免费
- [首先打开官网](https://www.vmware.com/)
![官网](https://swfoodt-blog.oss-cn-beijing.aliyuncs.com/img/blog-docs/20250607174707.png)
![下载](https://swfoodt-blog.oss-cn-beijing.aliyuncs.com/img/blog-docs/20250607174756.png)
**（需要注册登录）**
- 选择适合自己操作系统的版本下载
- 安装完成后，打开 VMware Workstation Pro

---

#### 三、下载 Linux 镜像文件

分别下载：
- Ubuntu 24.04 LTS
- Rocky Linux 9.5

- [推荐阿里云镜像站点](https://developer.aliyun.com/mirror/)

---

#### 四、配置虚拟机并安装系统

- 打开 VMware Workstation Pro，点击“设置” - “虚拟网络编辑器”，设置

![](https://swfoodt-blog.oss-cn-beijing.aliyuncs.com/img/blog-docs/20250607175554.png)
![](https://swfoodt-blog.oss-cn-beijing.aliyuncs.com/img/blog-docs/20250607180217.png)


##### VMware 虚拟网络类型详解

在使用 VMware Workstation 创建 Linux 虚拟机时，理解不同类型的虚拟网络是配置系统联网能力的关键。以下是常见网络类型的介绍及使用建议：

---

###### 🔗 桥接模式（Bridged，`VMnet0`）

- 虚拟机与主机处于同一物理局域网中。
- 虚拟机可直接访问外网，获得与主机类似的局域网 IP（由路由器分配）。
- 适用于需要虚拟机充当服务器、接收外部连接的场景。

---

###### 🌐 NAT 模式（Network Address Translation，`VMnet8`）

- 虚拟机通过主机共享网络连接访问外网。
- 虚拟机处于一个由 VMware 管理的私有子网中，主机充当 NAT 路由器。
- 虚拟机可访问外网，但外网无法直接访问虚拟机。

---

###### 🛡️ Host-only 模式（仅主机模式，`VMnet1`）

- 虚拟机仅能与主机通信，无法访问外网或其他局域网设备。
- 创建一个独立的内部网络，用于本地通信。

---

###### ⚙️ 自定义网络（`VMnet2` ~ `VMnet9`）

- 用户手动添加与配置的虚拟网络，可设定为桥接、NAT 或 Host-only。
- 可用于构建复杂的拓扑结构，比如多台虚拟机模拟子网环境、内网通信、分布式服务部署等。

---

###### 📊 网络类型对比

| 网络类型     | 是否能访问外网 | 是否能被主机访问 | 是否能被外部访问 | 常见用途                   |
|--------------|----------------|------------------|------------------|----------------------------|
| Bridged (桥接) | ✅ 是           | ✅ 是             | ✅ 是             | 服务测试、局域网模拟       |
| NAT          | ✅ 是           | ✅ 是             | ❌ 否             | 开发测试、安全访问外网     |
| Host-only    | ❌ 否           | ✅ 是             | ❌ 否             | 本地通信、离线隔离环境     |
| 自定义网络    | 可配置         | 可配置           | 可配置           | 多网络结构、教学与模拟实验 |


##### 虚拟机创建步骤

1. 打开 **VMware Workstation**，点击左上角 **“文件”** → 选择 **“新建虚拟机”**
2. 选择安装类型为 **“自定义（高级）(C)”**，点击 **“下一步”**
3. 在 **“硬件兼容性”** 页面使用默认值（通常为 Workstation 17.x），点击 **“下一步”**
4. 选择 **“稍后安装操作系统(S)”**，点击 **“下一步”**
5. 在 **“客户机操作系统”** 中选择：
   - 系统类型：**Linux**
   - 版本：**Rocky Linux 64位**
6. 为虚拟机命名，例如：**Rocky9.5-1**，路径可保持默认或自定义，点击 **“下一步”**
7. 处理器配置：
   - 处理器数量：**2**
   - 每个处理器的核心数：默认
8. 分配内存大小：**2048 MB**，点击 **“下一步”**
9. 网络类型选择：**使用网络地址转换 (NAT)**，点击 **“下一步”**
10. SCSI 控制器类型选择：**LSI Logic (推荐)**，点击 **“下一步”**
11. 虚拟磁盘类型选择：**SCSI（推荐）**，点击 **“下一步”**
12. 选择磁盘方式：
    - 选项：**创建新虚拟磁盘**
13. 设置磁盘容量：
    - 最大磁盘大小：**200 GB**
    - 勾选：**将虚拟磁盘存储为单个文件**
    - 不勾选：**立即分配所有磁盘空间**
14. 保持默认磁盘文件名（如：**Rocky9.5-1.vmdk**），点击 **“下一步”**
15. 检查配置后点击 **“完成”**
16. 在虚拟机设置中，点击 **“添加”** → 选择 **“CD/DVD (SATA)”** → 点击 **“下一步”**
17. 选择 **“使用 ISO 映像文件”**，点击 **“浏览”** 选择下载的 Rocky Linux 9.5 镜像文件
![](https://swfoodt-blog.oss-cn-beijing.aliyuncs.com/img/blog-docs/20250607181333.png)


---

#### 五、安装 Xshell

- Xshell 也是有个人免费版的（无标签数量限制），下载地址：[Xshell 官网](https://www.xshell.com/zh/free-for-home-school/)

![](https://swfoodt-blog.oss-cn-beijing.aliyuncs.com/img/blog-docs/20250607180633.png)

---

#### 六、Xshell连接 Linux 系统

1. 打开 **Xshell**
2. 点击左上角 **“新建”** 会话v
3. 配置连接信息：
   - **名称**：Rocky9 / Ubuntu24 等（任意）
   - **主机（Host）**：填写VMware中获得的虚拟机 IP 地址，例如 `192.168.111.135`
   - **协议**：选择 **SSH**
   - **端口号**：默认是 `22`，如自定义请手动修改
4. 点击 **“连接”**





