---
title: "从零实现个人博客的Docker容器化部署"
date: 2025-11-01
draft: false
description: "记录将Hugo博客从传统部署迁移到Docker容器化的完整过程，包括遇到的问题和解决方案"
categories: ["DevOps", "Docker"]
tags: ["Docker", "Hugo", "Nginx", "CI/CD", "GitHub Actions"]
---

## 项目背景

我的个人博客最初使用传统的Git Hook方式部署在云服务器上：代码推送后，服务器通过post-receive钩子自动构建并部署。虽然这种方式稳定可靠，但处于学习目的准备尝试将博客容器化，试一下Docker带来的环境一致性和可移植性。

本文记录了整个容器化过程，包括技术选型、遇到的问题以及最终的解决方案。
<!--more-->
## 技术栈

- **静态网站生成**: Hugo + hugoplate模板
- **样式框架**: Tailwind CSS 4.x
- **容器化**: Docker + Docker Compose
- **CI/CD**: GitHub Actions
- **Web服务器**: Nginx (Alpine)
- **开发环境**: Windows 11 + Docker Desktop

## 容器化方案选择

### 两种构建方案对比

在实施过程中，有两种Docker构建方案：

#### 方案1：完整自动化构建
```dockerfile
FROM node:18-alpine AS builder
# 安装Hugo、Go等构建工具
RUN apk add --no-cache hugo go git
# 安装依赖并构建
COPY . .
RUN npm ci && npm run build

FROM nginx:alpine
COPY --from=builder /src/public /usr/share/nginx/html
```

**优点**：
- 完全自包含，不依赖本地环境
- 适合团队协作和CI/CD
- 任何人都能直接构建

**缺点**：
- 构建时间长（需要下载依赖）
- 需要处理Hugo版本兼容性
- Windows环境下可能遇到网络问题

#### 方案2：本地构建 + Docker打包（最终选择）
```dockerfile
FROM nginx:alpine
COPY public /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf
```

**优点**：
- 镜像构建超快（1分钟内）
- 最终镜像极小（~25MB）
- 适合快速迭代
- 构建和运行环境分离

**缺点**：
- 需要先在本地执行 `npm run build`
- 依赖本地开发环境

### 最终选择

**方案2**，理由如下：

1. **开发效率优先**：本地Hugo构建只需几秒，Docker打包1分钟，总共不到2分钟
2. **环境分离原则**：构建和运行环境分离是最佳实践
3. **CI/CD灵活性**：在GitHub Actions中实现完整自动化，保持本地开发的简洁
4. **镜像优化**：最终镜像只包含必需的文件，更安全更小

## 实施步骤

### 1. 创建Dockerfile
```dockerfile
FROM nginx:alpine

# 复制静态网站文件
COPY public /usr/share/nginx/html

# 复制自定义 nginx 配置
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
```

### 2. 配置Nginx

创建优化的nginx配置：
```nginx
server {
    listen 80;
    server_name localhost;
    
    root /usr/share/nginx/html;
    index index.html index.htm;
    
    # 禁用可能导致端口丢失的重定向
    absolute_redirect off;
    port_in_redirect off;
    server_name_in_redirect off;
    
    # 路由配置
    location / {
        try_files $uri $uri/index.html $uri.html /index.html;
    }
    
    # 静态资源缓存
    location ~* \.(jpg|jpeg|png|gif|ico|css|js|svg|woff|woff2|ttf|eot|webp)$ {
        expires 1y;
        add_header Cache-Control "public, immutable";
    }
    
    # Gzip 压缩
    gzip on;
    gzip_vary on;
    gzip_min_length 1024;
    gzip_comp_level 6;
    gzip_types text/plain text/css text/xml text/javascript 
               application/x-javascript application/xml+rss 
               application/javascript application/json 
               image/svg+xml;
}
```

### 3. Docker Compose配置
为简化本地开发环境，配置了 Docker Compose：
```yaml
version: '3.8'

services:
  blog:
    build:
      context: .
      dockerfile: Dockerfile
    container_name: welblog-nginx
    ports:
      - "8080:80"
    restart: unless-stopped
    healthcheck:
      test: ["CMD", "wget", "--quiet", "--tries=1", "--spider", "http://localhost/"]
      interval: 30s
      timeout: 3s
      retries: 3
      start_period: 10s
```

### 4. GitHub Actions CI/CD
#### GitHub Actions 工作流设计
构建了完整的自动化流水线，实现代码推送后的自动构建、测试、验证、上传hub、直传服务器、部署（仅展示流程）：
```yaml
name: Docker Build, Push and Deploy

on:
  push:
    branches: [ main, master ]
  workflow_dispatch:

env:
  DOCKER_USERNAME: ${{ secrets.DOCKER_USERNAME }}
  DOCKER_IMAGE: welblog
  ENABLE_DEPLOY: true

jobs:
  build-push-deploy:
    runs-on: ubuntu-latest
    
    steps:
    # ===== 构建阶段 =====
    - name: 📥 Checkout code
      uses: actions/checkout@v4
    
    - name: 🔧 Setup Hugo
      uses: peaceiris/actions-hugo@v3
      with:
        hugo-version: 'latest'
        extended: true
    
    - name: 🔧 Setup Node.js
      uses: actions/setup-node@v4
      with:
        node-version: '18'
    
    - name: 📦 Install dependencies
      run: npm ci
    
    - name: 🔨 Build Hugo site
      run: |
        ...
    
    - name: 🐳 Build Docker image
      run: |
        ...
    
    # ===== 测试阶段 =====
    - name: 🧪 Test Docker container
      run: |
        ...
    
    # ===== 推送到Docker Hub =====
    - name: 📤 Login to Docker Hub
      if: github.event_name == 'push'
      uses: docker/login-action@v3
      with:
        username: ${{ secrets.DOCKER_USERNAME }}
        password: ${{ secrets.DOCKER_PASSWORD }}
    
    - name: 📤 Push to Docker Hub
      if: github.event_name == 'push'
      run: |
        ...
    
    # ===== 导出镜像（修复后）=====
    - name: 📦 Export image for deployment
      if: env.ENABLE_DEPLOY == 'true' && github.event_name == 'push'
      run: |
        ...
    
    # ===== 传输到服务器 =====
    - name: 📤 Transfer image to server
      if: env.ENABLE_DEPLOY == 'true' && github.event_name == 'push'
      uses: appleboy/scp-action@v0.1.7
      with:
        host: ${{ secrets.SERVER_HOST }}
        port: ${{ secrets.SERVER_PORT }}
        username: ${{ secrets.SERVER_USER }}
        key: ${{ secrets.SSH_PRIVATE_KEY }}
        source: "${{ env.IMAGE_FILENAME }}"
        target: "/tmp/"
    
    # ===== 部署 =====
    - name: 🚀 Deploy on server
      if: env.ENABLE_DEPLOY == 'true' && github.event_name == 'push'
      uses: appleboy/ssh-action@v1.0.3
      with:
        host: ${{ secrets.SERVER_HOST }}
        port: ${{ secrets.SERVER_PORT }}
        username: ${{ secrets.SERVER_USER }}
        key: ${{ secrets.SSH_PRIVATE_KEY }}
        script: |
          ...
    
    # ===== 报告 =====
    - name: 📊 Generate summary
      if: always()
      run: |
        ...
```

#### 版本管理策略
采用 Git commit SHA 的前7位作为镜像版本标签：
```bash
# 示例
commit: ab22979e0559bf46b6afd07b553eae027ec36c9e
镜像标签: welblog:ab22979

优势：
- 精确追踪每个版本
- 便于问题定位
- 支持快速回滚
```

## 遇到的小问题

### 问题1：Hugo版本兼容性

**问题**：Alpine Linux的Hugo版本（0.139）低于项目要求（0.141），导致构建失败。

**错误信息**：
```yaml
Error: permalink attribute not recognised
```

**解决方案**：
改用本地构建方案，避免在Docker内部构建Hugo，绕过了版本问题。

### 问题2：Windows文件锁

**问题**：在Windows环境下，`npm run build`时遇到文件被占用错误。

**错误信息**：
```yaml
The requested operation cannot be performed on a file with a user-mapped section open
```

**解决方案**：
```powershell
# 清理占用进程
Get-Process hugo -ErrorAction SilentlyContinue | Stop-Process -Force

# 删除缓存
Remove-Item -Recurse -Force public, resources
```

### 问题3：端口号丢失问题

**问题描述**：
容器运行在8080端口，首页能正常访问，但点击导航链接后端口号丢失：
- 期望：`http://localhost:8080/blog`
- 实际：`http://localhost/blog`（404错误）

**问题分析**：

通过浏览器开发者工具（F12）发现：
```yaml
请求: GET http://localhost:8080/blog
响应: 301 Moved Permanently
Location: http://localhost/blog/  ← 端口号丢失！
```

问题根源：
1. HTML中的链接是 `href="/blog"`（相对路径，正确的）
2. Nginx识别到 `/blog` 是目录，自动重定向到 `/blog/`（添加斜杠）
3. 重定向时生成的Location header丢失了端口号

**尝试的方案**：

1. ❌ **修改Hugo配置 `relativeURLs = true`**
   - 会影响RSS、sitemap和SEO
   - 不适合生产环境

2. ❌ **只添加 `port_in_redirect off`**
   - 配置没有完全生效

3. ✅ **最终解决方案**
```nginx
# 禁用所有重定向相关配置
absolute_redirect off;
port_in_redirect off;
server_name_in_redirect off;

# 改进try_files规则，避免触发目录重定向
location / {
    # 原来：try_files $uri $uri/ $uri.html /index.html;
    # $uri/ 会触发301重定向
    
    # 现在：try_files $uri $uri/index.html $uri.html /index.html;
    # 直接查找index.html，不触发重定向
    try_files $uri $uri/index.html $uri.html /index.html;
}
```

**关键点**：
- `$uri/` 会让nginx触发目录处理逻辑，产生301重定向
- `$uri/index.html` 直接查找文件，不触发重定向
- 配合三个redirect off指令，彻底解决问题

### 问题4：Docker镜像拉取失败

**问题**：尝试使用 `klakegg/hugo:0.141.0-ext-alpine` 时遇到403错误。

**解决方案**：
改用本地构建方案，不再依赖特定的Hugo镜像，问题自然解决。

### 问题5：云服务器端docker hub网络问题

**问题**：云服务器拉取Docker Hub镜像时网络不稳定，导致部署失败。

**评估解决方案**：
1. **（不适用）使用国内镜像加速**：
   - 镜像源不会同步个人镜像，只能解决官方镜像的拉取问题，所以无法解决个人镜像拉取失败的问题。

2. **（不适用）手动拉取镜像**：
   - 背离了自动化部署的初衷。

3. **（尚未尝试）使用云厂商容器镜像服务（ACR/CCR）**
   - 将镜像推送到云厂商的容器镜像服务，利用其稳定的网络环境进行拉取。
   - 需要额外配置CI/CD流水线，将镜像同时推送到Docker Hub和云厂商镜像服务。

4. **（最终选择）GitHub Actions构建后直接传输到服务器**
   - 通过SSH将构建好的镜像直接传输到服务器，避免拉取失败问题。
   - 需要在GitHub Actions中增加传输步骤。


**技术方案完整性**：

实现了以下目标：

1. ✅ 本地开发环境：`docker-compose up` 一键启动
2. ✅ CI/CD 流程：GitHub Actions 自动构建、测试、上传镜像、部署
3. ✅ 镜像仓库：Docker Hub 公开仓库，版本管理完整
4. ✅ 在有网络条件的环境下，可以直接hub拉取部署，无网络环境时可通过传输镜像方式部署。


## 技术亮点

### 1. 双轨部署策略

保留了原有的传统部署方式，同时实现了容器化：

**生产环境**（传统方式）：
- Git Hook + Hugo + Nginx
- 推送后1分钟左右自动更新
- 得益于hugo构建速度极快

**开发/测试环境**（容器化）：
- Docker + GitHub Actions
- 环境一致性保证
- 可随时切换到容器部署

这种策略的优势：
- 生产环境稳定性优先
- 学习新技术风险可控
- 理解不同方案的适用场景

### 2. 镜像优化

最终镜像只有25MB：
```yaml
nginx:alpine     ~7MB
+ 网站文件      ~18MB
= 总计          ~25MB
```

优化措施：
- 基于alpine基础镜像
- 只包含必需的运行时文件
- 不包含构建工具

### 3. CI/CD自动化

完整的自动化流程：
```yaml
git push → GitHub Actions触发
  ↓
安装Hugo和npm依赖
  ↓
构建网站（npm run build）
  ↓
构建Docker镜像（commit SHA版本号）
  ↓
启动容器健康检查
  ↓
登录Docker Hub并推送镜像
  ↓
无网络环境下，传输镜像到服务器
  ↓
通过SSH将镜像传输到服务器
  ↓
部署到云服务器
  ↓
生成构建报告
```

每次推送2-3分钟完成验证，确保代码质量。

## 使用方式

### 本地开发
```bash
# 开发模式
npm run dev

# 构建网站
npm run build

# Docker方式运行
docker-compose up -d

# 访问
open http://localhost:8080
```

### 部署到生产
```bash
# 推送代码到github
git push github main
# GitHub Actions会自动构建并部署
# 并且通过vercel托管
# 推送代码到云服务器仓库
git push origin main
# 服务器端Git Hook会自动构建并部署
```

## 性能对比

| 指标 | 传统部署 | Docker容器化 |
|------|----------|-------------|
| 部署时间 | ~2分钟 | ~3分钟（含构建） |
| 镜像大小 | N/A | 25MB |
| 环境一致性 | 依赖服务器 | ✅ 完全一致 |
| 可移植性 | ❌ 需要配置 | ✅ 一键部署 |
| 回滚能力 | 手动Git | ✅ 版本化镜像 |
| 资源占用 | ~50MB | ~70MB |

## 参考资源

- [Hugo官方文档](https://gohugo.io/documentation/)
- [Docker官方文档](https://docs.docker.com/)
- [Nginx配置指南](https://nginx.org/en/docs/)
- [GitHub Actions文档](https://docs.github.com/en/actions)

## 📋 文件清单

完成后有这些文件：
```yaml
welblog/
├── .github/
│   └── workflows/
│       └── docker-build.yml      ✅ 新增
├── content/
│   └── blog/
│       └── docker-containerization.md  ✅ 新增（上面的博客）
├── Dockerfile                     ✅ 新增
├── docker-compose.yml            ✅ 新增
├── nginx.conf                    ✅ 新增
├── .dockerignore                 ✅ 新增
└── ... (其他原有文件)