# A2UI Step12 OpenAPI & SDK Platform Design

## Design Goals

- 将当前内部 API 升级为对外可消费的平台接口。
- 提供低门槛 SDK，减少外部接入成本。
- 保持版本治理与兼容性策略明确。

## Architecture

### DSN-1201 API 版本化

- 路由前缀：`/api/v1/*`
- 现有内部接口逐步迁移到版本化路径
- 非兼容变更通过 `/api/v2` 发布

### DSN-1202 OpenAPI 产物

新增：

- `openapi/openapi.v1.yaml`

覆盖接口：

- auth（如已有）
- workspace
- session/version
- agent/run
- playground
- analytics（若启用）

### DSN-1203 API Key 模型

新增数据模型：

- `api_keys`
  - `id`, `name`, `key_hash`, `status`, `owner_user_id`, `created_at`, `last_used_at`

请求头：
- `X-API-Key: <raw_key>`

### DSN-1204 鉴权中间件扩展

- `AuthMode`：Bearer Token / API Key
- 对开放接口可配置接受 API Key

### DSN-1205 限流策略

MVP：
- 内存滑动窗口限流（按 key 或 user）
- 默认阈值：60 req/min

后续：
- Redis 分布式限流

### DSN-1206 SDK 结构

Go SDK：
- `sdk/go/`
  - client
  - auth
  - workspace
  - session
  - agent
  - playground

TypeScript SDK：
- `sdk/ts/`
  - `client.ts`
  - `resources/*`

### DSN-1207 错误与重试

SDK 统一错误对象：
- `code`
- `message`
- `status`
- `traceId`

对 `RATE_LIMITED`/`5xx` 提供可选重试策略。

### DSN-1208 文档体系

新增文档：

- `docs/openapi.md`
- `docs/sdk-go.md`
- `docs/sdk-ts.md`
- `docs/api-key.md`

## Requirement Coverage

- RQ-1201 -> DSN-1201, DSN-1202
- RQ-1202 -> DSN-1206, DSN-1207
- RQ-1203 -> DSN-1203, DSN-1204
- RQ-1204 -> DSN-1205, DSN-1207
- RQ-1205 -> DSN-1208
- RQ-1206 -> DSN-1201
- RQ-1207 -> DSN-1203
- RQ-1208 -> DSN-1202, DSN-1206, DSN-1205
