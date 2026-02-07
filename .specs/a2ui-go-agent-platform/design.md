# A2UI Go Agent Platform Design

## Design Goals

- 对齐 RQ-001..RQ-011，形成可实现、可验证、可扩展的 Go `pkg` 架构。
- 保证 Agent 主链路稳定：Plan -> Emit -> Validate -> Repair -> Apply。
- 保证后续新增 Provider/Tool/Theme/Playground 时不破坏核心流程。

## Architecture Overview

### DSN-001 模块边界

```text
cmd/
  server/
  cli/

pkg/
  agent/          # Agent runtime 状态机与任务循环
  workspace/      # 工作目录与会话目录映射
  session/        # 会话上下文、版本链、摘要
  a2ui/           # plan/emit/validate/repair/apply
  provider/       # provider/model 抽象与路由
  tools/          # 工具注册、执行、权限校验
  theme/          # 主题 token 与会话绑定
  playground/     # 分类/标签/案例沉淀
  guardrail/      # 注入防护、风险动作拦截
  events/         # 事件总线与审计模型
  telemetry/      # 指标、日志、trace
  store/          # repository 接口与实现

internal/
  infra/          # db/cache/config/queue 具体实现
```

### DSN-002 分层职责

- `cmd/*` 仅处理 I/O、参数、进程生命周期。
- `pkg/*` 提供可复用业务能力，面向接口编程。
- `internal/infra` 提供外部依赖适配，避免污染可复用包。

## Core Domain Models

### DSN-003 Workspace/Session/Version

- Workspace
  - `id`, `name`, `root_path`, `created_at`, `updated_at`
- Session
  - `id`, `workspace_id`, `title`, `active_theme_id`, `status`, `created_at`, `updated_at`
- SchemaVersion
  - `id`, `session_id`, `parent_version_id`, `schema_path`, `schema_hash`, `summary`, `created_at`

### DSN-004 Provider/Model

- Provider
  - `id`, `name`, `type`, `base_url`, `auth_ref`, `timeout_ms`, `enabled`
- Model
  - `id`, `provider_id`, `name`, `capabilities(json)`, `context_limit`, `enabled`

### DSN-005 Theme / Playground

- Theme
  - `id`, `name`, `token_set(json)`, `is_builtin`, `created_at`
- PlaygroundItem
  - `id`, `title`, `category_id`, `source_session_id`, `source_version_id`, `theme_id`, `schema_snapshot(json)`, `preview_ref`, `created_at`
- PlaygroundTag / PlaygroundItemTag（多对多）

### DSN-006 Audit/Event

- AgentRun
  - `id`, `session_id`, `request_text`, `status`, `started_at`, `ended_at`
- AgentEvent
  - `id`, `run_id`, `step`, `payload(json)`, `latency_ms`, `token_in`, `token_out`, `created_at`

## A2UI Runtime Design

### DSN-007 主流程状态机

状态：`INIT -> PLAN -> EMIT -> VALIDATE -> REPAIR? -> APPLY -> COMPLETE | FAILED`

- INIT: 装载 session 当前版本与上下文
- PLAN: 生成结构化计划（变更目标、范围、约束）
- EMIT: 输出 JSON 或 Patch
- VALIDATE: schema + component registry + props 约束校验
- REPAIR: 基于校验错误重试（上限 2~3 次）
- APPLY: 落盘与版本创建

### DSN-008 上下文分层

- System Context
  - 组件白名单、schema 规范、安全规则
- Session Context
  - 当前 schema、最近 N 次版本摘要、当前主题
- Task Context
  - 当前用户意图、局部修改范围、禁止修改区域

上下文预算策略：优先保留 System + Task，Session 采用摘要压缩。

### DSN-009 输出模式

- `full-json`：适用于初次生成页面
- `patch`：适用于增量编辑，默认优先

Patch 数据结构：`op`, `target`, `value`, `reason`。

## Provider/Tool Design

### DSN-010 Provider 抽象

接口：
- `Generate(ctx, req) (resp, error)`
- `GenerateStructured(ctx, schema, req)`
- `Health(ctx)`

路由器：
- 输入任务类型（plan/emit/repair）+ 能力标签
- 输出首选模型 + 备选模型（降级）

### DSN-011 Tool 执行模型

- Tool 分级：`read`, `write`, `exec`, `network`
- 每个 AgentRun 绑定 tool policy
- 高风险动作（exec/network/write critical path）进入 guardrail

## Theme & Playground Design

### DSN-012 Theme 应用策略

- Theme 通过 token map 注入渲染层，不改 schema 树结构。
- 每个 SchemaVersion 记录 `theme_snapshot_ref` 便于回放。

### DSN-013 Playground 沉淀策略

- 允许从整页或选中节点保存。
- 保存时写入 `schema_snapshot` + `source_version_id` + `theme_id`。
- 支持分类、标签、全文检索（标题/标签）。

## Storage & Migration

### DSN-014 存储与迁移

- 关系型数据库（PostgreSQL/SQLite 二选一可配置）
- 使用 `pressly/goose` 管理迁移版本
- 表拆分原则：高频读写表与审计表分离

## Security & Reliability

### DSN-015 安全策略

- Prompt 注入检测（关键模式 + 黑白名单）
- Tool 参数校验与路径白名单
- Provider 密钥仅存引用（auth_ref），敏感值走密钥服务或环境变量

### DSN-016 可靠性策略

- 外部模型调用设置超时与重试退避
- 版本写入采用事务（版本元数据 + 文件索引一致性）
- 失败不落盘，保证 session 最后一致版本可恢复

## Observability

### DSN-017 指标与追踪

- 关键指标：请求量、失败率、P95 延迟、repair 次数、token 消耗
- 每次 AgentRun 生成 trace-id，贯穿 provider/tool/store

## Requirement Coverage

- RQ-001 -> DSN-003, DSN-014
- RQ-002 -> DSN-007, DSN-008, DSN-009
- RQ-003 -> DSN-004, DSN-010
- RQ-004 -> DSN-008
- RQ-005 -> DSN-012
- RQ-006 -> DSN-013
- RQ-007 -> DSN-011, DSN-015
- RQ-008 -> DSN-006, DSN-017
- RQ-009 -> DSN-007, DSN-017
- RQ-010 -> DSN-014, DSN-016
- RQ-011 -> DSN-001, DSN-002, DSN-010

## Open Questions

- 是否首版就支持多租户（当前建议否）
- schema 文件落地格式是否需要压缩存储（当前建议否）
- Playground 检索是否首版引入向量索引（当前建议否）
