# A2UI Step12 OpenAPI & SDK Platform Requirements

## Overview

本规范定义 A2UI 平台第 12 步：开放 API 与 SDK 平台化。目标是让外部系统可稳定接入 A2UI 能力，并形成可版本化的开发者体验。

## Scope

- In Scope
  - OpenAPI 规范（核心接口）
  - API 版本化策略（v1）
  - Go/TypeScript SDK（MVP）
  - API Key 鉴权（服务到服务）
  - 限流与基础配额控制
  - 开发者文档与示例
- Out of Scope
  - 完整开发者门户（计费、工单、应用市场）
  - 高级企业 SLA 管理

## Functional Requirements

### RQ-1201 OpenAPI 规范

- WHEN 对外发布 API THEN THE SYSTEM SHALL 提供可机读的 OpenAPI 文档（v1）。
- WHEN 接口发生非兼容变更 THEN THE SYSTEM SHALL 通过新版本路径发布。

### RQ-1202 SDK 提供

- WHEN 开发者接入 THEN THE SYSTEM SHALL 提供 Go 与 TypeScript SDK 基础能力（auth、workspace/session、agent、playground）。
- WHEN SDK 调用失败 THEN THE SYSTEM SHALL 返回标准化错误对象。

### RQ-1203 API Key 鉴权

- WHEN 外部系统调用开放接口 THEN THE SYSTEM SHALL 支持 API Key 鉴权。
- WHEN API Key 无效/禁用 THEN THE SYSTEM SHALL 返回 `UNAUTHORIZED`。

### RQ-1204 限流与配额

- WHEN 客户端请求超过阈值 THEN THE SYSTEM SHALL 返回 `RATE_LIMITED` 并附带重试提示。
- WHEN 查询配额状态 THEN THE SYSTEM SHALL 支持返回剩余额度（MVP 可仅日志观测）。

### RQ-1205 开发者文档

- WHEN 发布 SDK/API THEN THE SYSTEM SHALL 提供快速开始、认证、错误码、示例代码文档。

## Non-functional Requirements

### RQ-1206 稳定性

- WHEN 对外接口运行 THEN THE SYSTEM SHALL 保持向后兼容并记录 deprecation 计划。

### RQ-1207 安全性

- WHEN 处理 API Key THEN THE SYSTEM SHALL 脱敏存储与日志脱敏。

### RQ-1208 可测试性

- WHEN 提交开放平台改动 THEN THE SYSTEM SHALL 覆盖 OpenAPI 校验、SDK 集成、限流测试。
