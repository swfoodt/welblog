---
title: "Linux 用户与权限管理"
date: 2025-06-12
slug: "linux-user-permission"
categories: ["操作系统", "Linux", "SRE笔记"]
tags: ["Linux基础", "权限管理", "用户管理", "ACL", "安全"]
draft: false
author: "swfoodt"
description: "系统梳理 Linux 用户、用户组、权限、安全模型等内容，包含文件权限设置、特殊权限、ACL、用户与组管理命令，以及通配符、重定向与管道示例。"
---



### 4. 用户和权限管理

---

#### Linux 安全模型与用户类型

- Linux 安全体系建立在“**用户、用户组、权限**”三者关系基础之上。
- **3A 安全模型**：
  - **Authentication（认证）**：确认用户身份（如 `/etc/passwd` 和 `/etc/shadow` 文件）。
  - **Authorization（授权）**：控制用户对资源的访问权限。
  - **Accounting（审计）**：记录用户行为。
<!--more-->
- 用户分类：
  - **超级用户**（root）：UID=0，拥有系统最高权限。
  - **系统用户**：用于运行系统服务，通常不登录系统。
  - **普通用户**：UID ≥ 1000，用于登录和日常操作。

{{<notice tip>}}
判断是否为超级用户：`id -u` 若输出为 0，则当前为 root 用户。
{{</notice>}}

---

#### 用户与用户组配置文件

| 文件路径        | 作用说明                             |
|----------------|--------------------------------------|
| `/etc/passwd`  | 存储用户基本信息（用户名、UID、GID） |
| `/etc/shadow`  | 存储加密的用户密码                   |
| `/etc/group`   | 存储用户组信息（组名、GID 等）       |
| `/etc/gshadow` | 存储用户组的密码                     |

常见配置工具：
```bash
vipw      # 编辑 /etc/passwd
vigr      # 编辑 /etc/group
pwck      # 检查用户配置文件合法性
grpck     # 检查组配置文件合法性
getent    # 查询系统实体信息，如用户、组、服务等
```

---

#### 用户与用户组管理命令

```bash
useradd username      # 添加用户
usermod -aG group user  # 添加用户到附加组
userdel -r username   # 删除用户及其主目录

groupadd groupname    # 创建用户组
groupmod -n new old   # 修改组名
groupdel groupname    # 删除用户组

id                    # 查看用户UID、GID、组等
whoami                # 当前用户名
su - username         # 切换用户
passwd username       # 修改指定用户密码

chage -l username     # 查看密码有效期策略
gpasswd -A admin group  # 设置组管理员
groups                # 当前用户所属所有组
```

---

#### Linux 权限系统概览

Linux 权限由三种角色与三种权限构成：

三种角色：
- 所有者（owner）
- 所属组（group）
- 其他用户（other）

三种权限：
- 读（r = 4）
- 写（w = 2）
- 执行（x = 1）

例：
-rw-r--r--  表示：
 所有者 可读写（rw-）
 所属组 只读（r--）
 其他用户 只读（r--）

文件权限修改命令：

```bash
chmod 755 file       # 使用数字修改权限
chmod u+x script.sh  # 为所有者添加执行权限
chown user file      # 修改文件所有者
chgrp group file     # 修改文件所属组
```

---

#### 特殊权限

```bash
SUID   # 设置程序以文件所有者身份运行（常用于二进制命令）
SGID   # 设置新建文件继承组；对目录生效时新文件自动继承组
STICKY # 目录下只有文件所有者或 root 可删除（常用于 /tmp）

ls -l 可看到：
s 表示设置了 SUID/SGID
t 表示设置了 STICKY
```

设置特殊权限示例：
```bash
chmod u+s /path/to/file    # 添加 SUID
chmod g+s /path/to/dir     # 添加 SGID
chmod +t /path/to/dir      # 添加 STICKY
```

---

#### 默认权限与 umask

```bash
新建文件默认权限：666 - umask
新建目录默认权限：777 - umask

# 查看/设置当前 umask 值：
umask
umask 022   # 默认新建文件权限为 644，目录为 755
```

---

#### ACL 权限访问控制

```bash
# 设置额外访问权限
setfacl -m u:username:rw file
# 查看 ACL 权限
getfacl file
```

---

#### Linux 文件属性管理

```bash
lsattr    # 查看扩展属性
chattr    # 修改文件属性（如防删除、防修改等）

# 常用示例
chattr +i important.txt   # 添加“不可更改”属性
chattr -i important.txt   # 取消该属性
```

---

#### 通配符与重定向、管道

```bash
# 通配符（用于文件匹配）
*       # 匹配任意长度任意字符
?       # 匹配任意单个字符
[a-z]   # 匹配 a 到 z 之间的任意单字符
[!0-9]  # 匹配非数字字符

# 重定向
> file      # 覆盖输出
>> file     # 追加输出
< file      # 从文件读取输入
2> file     # 错误输出重定向
&> file     # 标准输出+错误输出 一起重定向

# 管道（将前一命令输出传给下一命令）
cat file | grep "pattern"    # 查找包含 pattern 的行
ps aux | grep nginx          # 查找 nginx 进程
ls -l | less                 # 分页显示
```

---

#### 示例综合

```bash
# 创建用户并设置密码
useradd bob
passwd bob

# 将用户 bob 添加到 developers 组
usermod -aG developers bob

# 创建目录并赋予权限
mkdir /data/project
chown bob:developers /data/project
chmod 770 /data/project

# 给 /data/log 设置 STICKY 位，只允许拥有者删除
chmod +t /data/log

# 给 script.sh 添加执行权限
chmod +x script.sh
```

{{<notice info>}}
ACL 和 chattr 适用于细粒度权限控制或防止误操作，建议仅在必要场景使用。
{{</notice>}}
