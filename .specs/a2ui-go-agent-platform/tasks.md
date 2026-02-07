# A2UI Go Agent Platform Tasks

## Execution Rules

- 状态流转：`pending -> in_progress -> completed`
- 每个任务必须可独立验证
- 默认串行，标记可并行任务可交叉执行

## Task List

### TASK-001 初始化仓库结构与模块骨架

- Linked Design: DSN-001, DSN-002
- Linked Requirements: RQ-011
- Description:
  - 创建 `cmd/server`, `cmd/cli`, `pkg/*`, `internal/infra` 目录骨架
  - 初始化模块边界与基础 README（每包职责）
- DoD:
  - 目录结构与设计一致
  - 基础构建通过
- Verify:
  - `go list ./...`

Status: completed

### TASK-002 建立数据层与 goose 迁移

- Linked Design: DSN-003, DSN-004, DSN-005, DSN-006, DSN-014
- Linked Requirements: RQ-001, RQ-003, RQ-006, RQ-008, RQ-010
- Description:
  - 引入 `pressly/goose`
  - 创建首批迁移：workspace/session/schema_version/provider/model/theme/playground/agent_run/agent_event
- DoD:
  - 迁移可 up/down
  - 表结构满足主流程最小需求
- Verify:
  - `goose up`
  - `goose status`

Status: completed

### TASK-003 实现 store repository 接口

- Linked Design: DSN-014
- Linked Requirements: RQ-001, RQ-006, RQ-008
- Description:
  - 定义 `pkg/store` 接口（WorkspaceRepo/SessionRepo/VersionRepo/PlaygroundRepo/EventRepo）
  - 在 `internal/infra` 提供 SQL 实现
- DoD:
  - 接口与实现编译通过
  - 关键增删查改具备单元测试
- Verify:
  - `go test ./...`

Status: completed

### TASK-004 实现 provider 与模型路由器

- Linked Design: DSN-010
- Linked Requirements: RQ-003, RQ-011
- Description:
  - 实现 Provider Adapter 接口
  - 实现任务类型到模型能力标签路由
  - 支持主备降级
- DoD:
  - 可配置至少一个 provider + 两个模型（plan/emit）
  - provider 故障时可回退
- Verify:
  - `go test ./pkg/provider/...`

Status: completed

### TASK-005 实现 session 上下文聚合与摘要

- Linked Design: DSN-008
- Linked Requirements: RQ-004
- Description:
  - 构建 system/session/task 三层上下文
  - 实现历史摘要压缩策略
- DoD:
  - 上下文构建可复现
  - 超预算时能稳定压缩
- Verify:
  - `go test ./pkg/session/...`

Status: completed

### TASK-006 实现 A2UI 主状态机

- Linked Design: DSN-007, DSN-009
- Linked Requirements: RQ-002
- Description:
  - 实现 INIT/PLAN/EMIT/VALIDATE/REPAIR/APPLY 流程
  - patch 优先，full-json 兜底
- DoD:
  - 可执行单轮任务并输出新版本
  - 失败路径可观测
- Verify:
  - `go test ./pkg/agent/...`

Status: completed

### TASK-007 实现 schema 校验与自动修复

- Linked Design: DSN-007
- Linked Requirements: RQ-002
- Description:
  - 集成 JSON schema + component/props 校验
  - repair 重试上限与错误聚合输出
- DoD:
  - 非法输出可被拒绝并有明确信息
  - 合法修复可自动通过
- Verify:
  - `go test ./pkg/a2ui/...`

Status: completed

### TASK-008 实现 guardrail 与工具执行策略

- Linked Design: DSN-011, DSN-015
- Linked Requirements: RQ-007
- Description:
  - 工具分级、策略匹配、危险操作拦截
  - prompt 注入基础检测
- DoD:
  - 高风险操作默认阻断或需显式确认
  - 注入样例可被识别
- Verify:
  - `go test ./pkg/guardrail/...`

Status: completed

### TASK-009 实现 theme 模块与会话绑定

- Linked Design: DSN-012
- Linked Requirements: RQ-005
- Description:
  - 实现主题 token 存取与会话绑定
  - 版本保存 theme 快照引用
- DoD:
  - 切换主题不改 schema 结构
  - 版本可回放主题
- Verify:
  - `go test ./pkg/theme/...`

Status: completed

### TASK-010 实现 playground 分类/标签/检索

- Linked Design: DSN-013
- Linked Requirements: RQ-006
- Description:
  - 实现案例保存、分类、标签、检索 API
  - 记录 source session/version/theme
- DoD:
  - 可按分类和标签查询
  - 保存内容可追溯来源
- Verify:
  - `go test ./pkg/playground/...`

Status: completed

### TASK-011 实现事件审计与 telemetry

- Linked Design: DSN-006, DSN-017
- Linked Requirements: RQ-008, RQ-009
- Description:
  - 记录 run/event、时延、token
  - 暴露基础 metrics
- DoD:
  - 每次 AgentRun 有完整链路事件
  - 能查询失败步骤原因
- Verify:
  - `go test ./pkg/events/...`
  - `go test ./pkg/telemetry/...`

Status: completed

### TASK-012 交付最小 API 与 CLI 命令

- Linked Design: DSN-001, DSN-002
- Linked Requirements: RQ-001, RQ-002, RQ-006
- Description:
  - 提供最小接口：创建 workspace/session、提交 agent 请求、版本查询、playground 保存
  - CLI 对应命令实现
- DoD:
  - API 可端到端跑通一次生成链路
  - CLI 可触发同等流程
- Verify:
  - `go test ./...`
  - 手动 E2E 脚本通过

Status: completed

## Milestones

- M1: TASK-001 ~ TASK-004 完成（底座可运行）
- M2: TASK-005 ~ TASK-008 完成（核心 Agent 能闭环）
- M3: TASK-009 ~ TASK-012 完成（主题/案例/可观测与入口齐备）
