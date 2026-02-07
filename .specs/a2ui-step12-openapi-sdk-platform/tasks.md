# A2UI Step12 OpenAPI & SDK Platform Tasks

## TASK-1201 设计并落地 OpenAPI v1 文档

- Linked Requirements: RQ-1201
- Linked Design: DSN-1202
- Description:
  - 新增 `openapi/openapi.v1.yaml`
  - 覆盖核心资源接口
- DoD:
  - OpenAPI 可通过基础校验
- Verify:
  - 使用 openapi lint 工具校验通过

Status: completed

## TASK-1202 路由版本化改造

- Linked Requirements: RQ-1201, RQ-1206
- Linked Design: DSN-1201
- Description:
  - 将外部可用接口挂载到 `/api/v1`
  - 保持旧路径兼容期（可选）
- DoD:
  - `/api/v1` 路径可访问
- Verify:
  - 接口冒烟测试通过

Status: completed

## TASK-1203 增加 API Key 数据模型与存储

- Linked Requirements: RQ-1203, RQ-1207
- Linked Design: DSN-1203
- Description:
  - 新增 migration 创建 `api_keys`
  - 实现 key hash 存储与查询
- DoD:
  - 可创建/禁用 API Key
- Verify:
  - `go run ./cmd/cli db:migrate`

Status: completed

## TASK-1204 鉴权中间件支持 API Key

- Linked Requirements: RQ-1203
- Linked Design: DSN-1204
- Description:
  - 在中间件中解析 `X-API-Key`
  - 校验通过后注入调用身份
- DoD:
  - 有效 key 可访问
  - 无效 key 返回 UNAUTHORIZED
- Verify:
  - `go test ./cmd/server/...`

Status: completed

## TASK-1205 增加限流能力

- Linked Requirements: RQ-1204
- Linked Design: DSN-1205
- Description:
  - 实现按 key/user 的基础限流
  - 返回 `RATE_LIMITED` 与重试建议
- DoD:
  - 超限请求被正确拦截
- Verify:
  - `go test ./pkg/...`（限流相关）

Status: completed

## TASK-1206 产出 Go SDK MVP

- Linked Requirements: RQ-1202
- Linked Design: DSN-1206, DSN-1207
- Description:
  - 提供 Go SDK client 与核心资源调用
- DoD:
  - 能通过 SDK 调 workspace/session/agent
- Verify:
  - SDK 集成示例可运行

Status: completed

## TASK-1207 产出 TypeScript SDK MVP

- Linked Requirements: RQ-1202
- Linked Design: DSN-1206, DSN-1207
- Description:
  - 提供 TS SDK client 与核心资源调用
- DoD:
  - 能通过 SDK 调 workspace/session/agent
- Verify:
  - TS 示例可运行

Status: completed

## TASK-1208 开发者文档与示例

- Linked Requirements: RQ-1205
- Linked Design: DSN-1208
- Description:
  - 增加 OpenAPI、API Key、Go/TS SDK 文档与 quick start
- DoD:
  - 文档可支持第三方快速接入
- Verify:
  - 按文档完成一次接入演练

Status: completed

## TASK-1209 平台测试与回归

- Linked Requirements: RQ-1208
- Linked Design: DSN-1202, DSN-1205, DSN-1206
- Description:
  - 增加 OpenAPI 校验、SDK 集成、限流回归
- DoD:
  - 全量测试通过
- Verify:
  - `go test ./...`

Status: completed
