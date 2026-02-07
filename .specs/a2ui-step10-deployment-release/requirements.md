# A2UI Step10 Deployment & Release Requirements

## Overview

本规范定义 A2UI 平台第 10 步：部署与发布体系。目标是建立可重复、可回滚、可观测的交付流程，支撑从本地到测试再到生产环境的稳定发布。

## Scope

- In Scope
  - 容器化（后端服务）
  - 环境分层配置（dev/staging/prod）
  - 数据库迁移发布流程
  - 发布版本标记与回滚策略
  - 基础部署文档
- Out of Scope
  - Kubernetes 全自动编排（可后续）
  - 灰度/金丝雀高级流量治理

## Functional Requirements

### RQ-1001 容器化交付

- WHEN 构建发布版本 THEN THE SYSTEM SHALL 生成可运行的后端 Docker 镜像。
- WHEN 容器启动 THEN THE SYSTEM SHALL 支持读取环境变量并正常连接 Postgres。

### RQ-1002 环境分层配置

- WHEN 部署到不同环境 THEN THE SYSTEM SHALL 支持 dev/staging/prod 独立配置。
- WHEN 环境变量缺失 THEN THE SYSTEM SHALL 在启动时明确报错并拒绝启动。

### RQ-1003 发布迁移流程

- WHEN 发布新版本 THEN THE SYSTEM SHALL 在应用启动前或启动阶段执行数据库迁移。
- WHEN 迁移失败 THEN THE SYSTEM SHALL 阻断发布并保留旧版本可回滚。

### RQ-1004 版本与回滚

- WHEN 发布成功 THEN THE SYSTEM SHALL 记录版本号、构建信息与发布时间。
- WHEN 发布失败 THEN THE SYSTEM SHALL 提供可执行的回滚步骤（应用版本 + 数据库策略说明）。

### RQ-1005 发布验收

- WHEN 完成部署 THEN THE SYSTEM SHALL 自动执行健康检查与最小业务验收（smoke）。

## Non-functional Requirements

### RQ-1006 可运维性

- WHEN 线上故障发生 THEN THE SYSTEM SHALL 能基于日志与版本信息快速定位变更来源。

### RQ-1007 安全性

- WHEN 构建镜像 THEN THE SYSTEM SHALL 避免将密钥写入镜像层。

### RQ-1008 可测试性

- WHEN 发布流水线执行 THEN THE SYSTEM SHALL 在 CI/CD 中验证镜像构建与启动。
