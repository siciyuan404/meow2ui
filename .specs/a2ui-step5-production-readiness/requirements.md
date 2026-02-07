# A2UI Step5 Production Readiness Requirements

## Overview

本规范定义 A2UI 平台第 5 步：生产可用性收口。重点是配置安全、运行健康检查、错误分级、发布前验证。

## Scope

- In Scope
  - 运行配置校验（必需环境变量、默认值策略）
  - 健康检查增强（应用健康 + 数据库健康）
  - 错误分级与统一响应格式
  - 最小发布检查清单（smoke test）
  - 基础运行文档补全
- Out of Scope
  - 完整认证鉴权体系（JWT/RBAC）
  - 多集群部署与服务网格

## Functional Requirements

### RQ-501 配置校验

- WHEN 服务启动 THEN THE SYSTEM SHALL 校验关键配置项并输出明确启动错误。
- WHEN `STORE_DRIVER=postgres` THEN THE SYSTEM SHALL 校验 `PG_*` 配置完整性。

### RQ-502 健康检查

- WHEN 调用 `/healthz` THEN THE SYSTEM SHALL 返回应用状态。
- WHEN 调用 `/readyz` THEN THE SYSTEM SHALL 包含数据库连通性检查结果（postgres 模式下）。

### RQ-503 统一错误响应

- WHEN API 返回错误 THEN THE SYSTEM SHALL 使用统一结构 `{code,message,detail,traceId}`。
- WHEN 发生可预期业务错误 THEN THE SYSTEM SHALL 返回稳定错误码（如 `NOT_FOUND`, `VALIDATION_FAILED`）。

### RQ-504 发布前检查

- WHEN 发布候选版本构建 THEN THE SYSTEM SHALL 执行 smoke 命令集合并输出结果。
- WHEN smoke 检查失败 THEN THE SYSTEM SHALL 阻断发布流程。

## Non-functional Requirements

### RQ-505 可运维性

- WHEN 服务运行 THEN THE SYSTEM SHALL 输出结构化日志，包含 `traceId` 与关键上下文字段。

### RQ-506 可测试性

- WHEN 提交实现 THEN THE SYSTEM SHALL 提供健康检查与错误响应的最小测试覆盖。
