# A2UI Step1 PostgreSQL Persistence Requirements

## Overview

本规范定义 A2UI 平台第 1 步：从内存存储切换到 PostgreSQL 持久化，并通过 goose 迁移管理数据库结构。

## Environment Baseline

- Host: `localhost`
- Port: `5432`
- User: `postgres`
- Password: `postgres`
- Database: 尚未创建（本步骤需创建）

## Scope

- In Scope
  - 创建目标数据库（默认 `a2ui_platform`）
  - 接入 PostgreSQL 驱动与连接配置
  - 使用 `pressly/goose` 执行迁移
  - 实现 `sqlstore` 替代 `memorystore`
  - `bootstrap` 支持通过配置切换 store 类型
- Out of Scope
  - 多租户数据库隔离
  - 高可用主从与连接池高级调优

## Functional Requirements

### RQ-101 数据库初始化

- WHEN 服务首次部署 THEN THE SYSTEM SHALL 自动或脚本化创建业务数据库 `a2ui_platform`。
- WHEN 数据库已存在 THEN THE SYSTEM SHALL 跳过创建并继续后续流程。

### RQ-102 goose 迁移执行

- WHEN 服务启动前执行迁移 THEN THE SYSTEM SHALL 按版本顺序执行 `migrations/*.up.sql`。
- WHEN 迁移失败 THEN THE SYSTEM SHALL 返回明确错误并阻止服务以不一致状态运行。

### RQ-103 SQL Repository 替换

- WHEN 使用 PostgreSQL 模式 THEN THE SYSTEM SHALL 使用 `sqlstore` 提供 Workspace/Session/Version/Provider/Theme/Playground/Event 仓储实现。
- WHEN Repository 操作失败 THEN THE SYSTEM SHALL 返回可追踪错误并不中断进程外的状态一致性。

### RQ-104 Bootstrap 配置化

- WHEN `STORE_DRIVER=postgres` THEN THE SYSTEM SHALL 从环境变量构建 DSN 并初始化 `sqlstore`。
- WHEN `STORE_DRIVER=memory` THEN THE SYSTEM SHALL 保持当前内存实现用于本地快速开发。

### RQ-105 基础运行验证

- WHEN 完成迁移并启动服务 THEN THE SYSTEM SHALL 支持创建 workspace/session 并生成 version 到数据库。

## Non-functional Requirements

### RQ-106 可靠性

- WHEN 数据库短暂不可用 THEN THE SYSTEM SHALL 在连接阶段快速失败并输出可定位日志。

### RQ-107 可维护性

- WHEN 新增表或字段 THEN THE SYSTEM SHALL 仅通过 goose migration 变更，不允许手工改生产库结构。

### RQ-108 可测试性

- WHEN 提交实现 THEN THE SYSTEM SHALL 保证 `go test ./...` 通过，且至少包含 SQL repository 的关键路径测试。
