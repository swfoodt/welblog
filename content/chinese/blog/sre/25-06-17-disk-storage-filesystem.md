---
title: "磁盘存储与文件系统管理"
date: 2025-06-17
slug: "disk-storage-filesystem"
categories: ["操作系统", "Linux"]
tags: ["磁盘管理", "分区", "文件系统", "RAID", "LVM", "存储管理"]
draft: false
description: "本节系统讲解 Linux 中的磁盘结构、分区类型、文件系统管理命令，结合 RAID 与 LVM 的概念、工具与应用场景，帮助建立完整的存储系统理解。"
---

### 一、磁盘结构与基础术语

在 Linux 中，所有设备（包括硬盘）都被视为文件，称为“**设备文件**”。

#### 1. 设备文件

- 位于 `/dev` 目录中；
- 命名规则：
  - `/dev/sda`：第一个SCSI/SATA磁盘
  - `/dev/sdb1`：第二块磁盘的第一个分区
<!--more-->

#### 2. 常见术语解释

| 名称 | 含义 |
|------|------|
| 块设备（Block device） | 支持随机访问的存储设备，如硬盘、U盘 |
| 分区（Partition） | 将一个磁盘划分为若干逻辑区域 |
| 文件系统（Filesystem） | 用于组织和管理文件数据的结构 |
| 挂载（Mount） | 将设备或目录接入主文件系统目录树 |
| 挂载点（Mount point） | 用于访问挂载设备的目录 |

---

### 二、分区类型与分区表结构

#### 1. 分区表类型对比：MBR vs GPT

| 项目     | MBR                        | GPT                          |
|----------|----------------------------|-------------------------------|
| 最大容量 | 2TB                        | 理论上无限（目前为 9.4ZB）    |
| 分区数   | 最多 4 个主分区            | 128 个以上                   |
| 兼容性   | 与旧 BIOS 兼容             | 需 UEFI 启动支持              |
| 结构     | 主引导记录 + 分区表        | 主 GUID 表 + 备份分区表       |
| 安全性   | 无冗余                     | 有分区表备份，防破坏         |

#### 2. 管理分区工具

```bash
fdisk /dev/sdX       # 针对 MBR 分区
parted /dev/sdX      # 支持 GPT 和 MBR
lsblk                # 查看块设备结构
blkid                # 查看设备 UUID 与类型
```

---

### 三、文件系统基础与管理

#### 1. 常见文件系统类型

| 类型   | 特点                         |
|--------|------------------------------|
| ext4   | 默认文件系统，稳定可靠       |
| xfs    | 大文件支持好，写入性能佳     |
| btrfs  | 支持快照、自修复、高可用性   |
| vfat   | 跨平台兼容性高               |

#### 2. 文件系统创建与管理命令

```bash
mkfs.ext4 /dev/sda1     # 创建 ext4 文件系统
mkfs.xfs /dev/sdb1      # 创建 xfs
fsck /dev/sda1          # 检查并修复文件系统
tune2fs -l /dev/sda1    # 查看文件系统详细信息
```

---

### 四、挂载与挂载管理

#### 1. 挂载操作

```bash
mount /dev/sda1 /mnt         # 手动挂载
umount /mnt                  # 卸载
```

#### 2. 查看挂载状态

```bash
mount        # 显示所有挂载项
df -h        # 查看磁盘空间与挂载点
findmnt      # 层级结构查看挂载
```

#### 3. 永久挂载

编辑 `/etc/fstab` 文件，添加一行：

```bash
UUID=xxxxx  /mnt/data  ext4  defaults  0  2
```

获取 UUID：

```bash
blkid /dev/sda1
```

#### 4. 交换分区操作

```bash
mkswap /dev/sdX
swapon /dev/sdX
swapoff /dev/sdX
```

#### 5. 可移动设备管理

- 通常由系统自动挂载到 `/media/用户名/`
- 使用 `udisksctl` 或 `mount` 可手动挂载

---

### 五、磁盘配额管理（Quota）

用于限制用户/组使用磁盘资源：

```bash
quota username              # 查看用户配额
edquota username            # 编辑配额
repquota -a                 # 查看所有配额统计
```

启用步骤：
1. 修改挂载选项支持 `usrquota` 或 `grpquota`；
2. 重新挂载分区；
3. 初始化配额数据库 `quotacheck -cum /mnt`;
4. 启用配额 `quotaon /mnt`

---

### 六、RAID 技术概览

RAID（Redundant Array of Independent Disks）：冗余磁盘阵列，用于提高性能与可靠性。

#### 各级别对比

| RAID级别 | 最少磁盘数 | 利用率     | 冗余性   | 性能           | 特点                        |
|----------|------------|------------|----------|----------------|-----------------------------|
| RAID 0   | 2          | 100%       | 无       | 读写最快       | 条带化，无冗余，性能优先     |
| RAID 1   | 2          | 50%        | 高       | 读取快，写入略慢 | 镜像，数据完全复制            |
| RAID 5   | 3          | (n-1)/n    | 可容忍1块 | 读取快，写入中等 | 条带化+奇偶校验               |
| RAID 10  | 4          | 50%        | 高       | 高读写性能      | RAID 1+0，镜像后再条带化      |
| RAID 01  | 4          | 50%        | 中       | 较快           | RAID 0+1，条带化后再镜像      |

- 💡 RAID 10 优于 RAID 01，在故障容忍方面更稳定。RAID 不等于备份，仍建议定期备份数据。

---

### 七、LVM（逻辑卷管理）

LVM（Logical Volume Manager）为 Linux 提供灵活的磁盘管理方式，支持动态扩容、快照等。

#### 基本结构

```text
物理卷（PV） -> 卷组（VG） -> 逻辑卷（LV）
```

#### 1. 创建流程

```bash
# 创建 PV
pvcreate /dev/sdb1

# 创建 VG
vgcreate myvg /dev/sdb1

# 创建 LV
lvcreate -L 10G -n mylv myvg

# 格式化并挂载
mkfs.ext4 /dev/myvg/mylv
mount /dev/myvg/mylv /mnt
```

#### 2. 扩容与缩减

```bash
lvextend -L +5G /dev/myvg/mylv     # 扩容
resize2fs /dev/myvg/mylv           # 同步文件系统

lvreduce -L 5G /dev/myvg/mylv      # 缩减（需先 umount）
```

- ⚠️ 缩减卷存在数据丢失风险，建议备份数据并确保挂载已卸载。

#### 3. 快照卷（snapshot）

```bash
lvcreate -s -L 1G -n snap /dev/myvg/mylv
```

- 快照通常用于备份或短期实验；
- 恢复快照前需卸载原卷。

---

### 八、使用场景建议

| 场景 | 推荐方式 |
|------|-----------|
| 单用户系统，简单挂载 | ext4 + 手动分区 |
| 多用户或动态扩展 | LVM + ext4/xfs |
| 高可用服务器 | RAID 1 / 5 / 10 + LVM |
| 快速备份 | LVM 快照 + `rsync` |
| 移动设备兼容性 | 使用 `vfat` 或 `exfat` 文件系统 |

---

