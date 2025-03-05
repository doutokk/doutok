
<div align="center">
  <h1>🎬 DouTok Shop（抖音电商）</h1>
  <p>基于 Kitex 和 Hertz 框架构建的高性能、易扩展的云原生商城项目</p>
  <p>
    <img src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white" alt="Go"/>
    <img src="https://img.shields.io/badge/mysql-%2300f.svg?style=for-the-badge&logo=mysql&logoColor=white" alt="MySQL"/>
    <img src="https://img.shields.io/badge/kitex-%23FF6A00.svg?style=for-the-badge&logo=bytedance&logoColor=white" alt="Kitex"/>
    <img src="https://img.shields.io/badge/hertz-%23008ECF.svg?style=for-the-badge&logo=bytedance&logoColor=white" alt="Hertz"/>
    <img src="https://img.shields.io/badge/redis-%23D82C20.svg?style=for-the-badge&logo=redis&logoColor=white" alt="Redis"/>
  </p>
</div>

---

## 🌟 项目简介

DouTok Shop（抖音电商）是一个面向云原生时代的综合性电商平台，旨在提供高性能、易扩展的电商解决方案。项目基于字节跳动的 Kitex 作为 RPC 框架和 Hertz 作为 HTTP 框架，结合成熟的后端技术与现代 DevOps 工具链，构建用户、管理与运维多端统一的生态系统。

**主要功能：**

- **用户端：**
  - 浏览商品、商品搜索与筛选
  - 添加购物车、下单及订单管理
  - 跳转支付宝支付等集成

- **管理端：**
  - 商品的新增、修改与删除
  - 图片直传至阿里云 OSS，实现高效媒体管理

- **运维端：**
  - 自动化的 Pod 扩缩容与节点弹性扩容
  - 统一日志管理、性能监控（Metrics）及链路追踪
  - GitOps 工作流实现持续集成与持续部署

---

## 🔗 项目地址

- **展示地址：** [https://shopnext.kl.suyiiyii.top/](https://shopnext.kl.suyiiyii.top/)
- **CDN 地址：** [https://shopnext.suyiiyii.top/](https://shopnext.suyiiyii.top/)
- **后端仓库：** [https://github.com/doutokk/doutok](https://github.com/doutokk/doutok)
- **前端仓库：** [https://github.com/doutokk/doutok-frontend](https://github.com/doutokk/doutok-frontend)

---

## 👥 项目分工

- **队长：** 杜家楷  
  *主要负责流水线部署、运维、前端开发、OSS 文件直传及实现商品管理、购物车、订单和支付模块。*

- **队员：** 钟宝骏  
  *主要负责压测、搜索、缓存、测试、仓库管理以及用户、鉴权和网关模块的实现。*

---

## 🛠 技术栈

### 后端
- **语言：** Go
- **RPC & HTTP 框架：** Kitex & Hertz
- **鉴权：** JWT + Casbin（基于角色的权限控制）
- **数据库：** MySQL（关系型数据存储）
- **缓存与分布式锁：** Redis
- **搜索引擎：** Elasticsearch（商品搜索性能提升约10倍）
- **对象存储：** 阿里云 OSS

### 前端
- **框架：** React
- **静态资源托管：** Nginx 配合 CDN 加速

### 部署与运维
- **容器化与编排：** Kubernetes
- **日志与监控：** Prometheus、Grafana、Loki、Promtail
- **链路追踪：** Jaeger
- **CI/CD：** GitHub Actions、GitHub Packages、ArgoCD（GitOps 模式）

---

## 🏗 项目架构

DouTok Shop 的架构设计注重高可用与弹性扩展，主要包括以下层级：

1. **客户端与内容分发**
   - 浏览器访问 → 静态资源通过 CDN 加速

2. **接入层**
   - Ingress（Traefik）统一入口，路由至网关或各微服务

3. **业务服务层**
   - **Gateway：** 统一处理用户请求与服务间聚合
   - **微服务：** 拆分为用户、商品、订单、支付等独立模块，通过 RPC/HTTP 通信
   - **Redis：** 提供高频数据缓存及分布式锁支持

4. **数据与存储层**
   - **MySQL：** 核心业务数据存储，确保数据一致性
   - **对象存储：** 阿里云 OSS 存储图片、视频等非结构化数据

5. **可观测性与监控**
   - **日志：** Loki + Promtail 集中管理日志
   - **监控：** Prometheus + Grafana 实时监控关键指标
   - **链路追踪：** Jaeger 分布式调用链追踪

6. **基础设施层**
   - Kubernetes 集群管理，支持自动扩缩容与高可用部署

---

## 🚀 快速开始

### 环境准备

- **Go 版本：** ≥ 1.23
- **MySQL 版本：** ≥ 8.0
- **Redis 版本：** ≥ 6.0
- **Kubernetes 集群环境**（可选，用于生产部署）
- **阿里云 OSS 账户**（用于对象存储）

### 本地部署步骤

1. **克隆仓库**
   ```bash
   git clone https://github.com/doutokk/doutok.git
   cd doutok
   ```

2. **安装依赖**
   ```bash
   go mod download
   或
   go mod tidy
   ```

3. **数据库配置**
   - 根据 `docs/database` 下的说明，创建并初始化 MySQL 数据库及数据表

4. **编辑配置文件**
   - 在各微服务的 `conf` 目录下，使用您的数据库、Redis 和 OSS 凭证更新 YAML 配置文件

5. **启动服务**
   - 使用 Kubernetes 清单文件进行服务启动和部署

---

## 📚 API 文档

API 文档位于项目文档目录中，方便您了解各微服务接口定义与调用方式：
```bash
cd docs/api
```

---

## 🚢 部署与 GitOps 工作流

DouTok Shop 完整采用 GitOps 工作流，通过 GitHub Actions 触发持续集成和镜像构建，再由 ArgoCD 实现 Kubernetes 集群的持续部署。主要步骤包括：

1. **代码提交触发流水线**
   - 自动化执行编译、单元测试、静态检查

2. **镜像构建与推送**
   - Dockerfile 构建镜像，推送至 GitHub Container Registry

3. **配置声明更新**
   - 使用 kustomize/helm 更新 Kubernetes 清单，提交至 Manifest 仓库

4. **持续部署**
   - ArgoCD 实时监控 Manifest 仓库变更，并同步至集群，实现滚动更新与零宕机部署

---

## 🤝 贡献指南

欢迎各路开发者共同参与 DouTok Shop 项目的建设！

1. **Fork 本仓库**
2. **创建新分支**
   ```bash
   git checkout -b feature/YourFeature
   ```
3. **提交代码与更新文档**
4. **提交 Pull Request**

请确保您的代码通过所有测试，并附上相关文档说明。

---

## 📄 许可证

本项目采用 [MIT 许可证](LICENSE)。

---

## 🙋‍♀ 作者 & 团队

- **杜家楷**  
  邮箱：suyiiyii@gmail.com  
  *（流水线部署、运维、前端、OSS 直传、商品管理、购物车、订单、支付模块）*

- **钟宝骏**  
  邮箱：3168078770@qq.com  
  *（压测、搜索、缓存、测试、仓库管理、用户鉴权、网关模块）*


---

## 🔍 项目总结与展望

- **当前问题：** 部分业务场景仍需进一步优化，如高并发下的细粒度权限控制与缓存策略
- **已识别优化点：** 架构性能瓶颈、日志监控精细化、微服务间通信优化
- **未来方向：** 持续演进架构、引入更多自动化测试和监控手段、提升用户体验

---

<div align="center">
  <br/>
  ⭐ 如果您觉得本项目有帮助，请给我们点个 star！  
  📖 更多信息请查阅项目 Wiki 或提交 issue。
</div>