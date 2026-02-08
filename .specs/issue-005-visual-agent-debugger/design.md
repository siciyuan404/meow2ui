# Issue #5 可视化 Agent 调试工具 Design

## Design Goals

- 复用 Step15 可观测能力，避免重复采集。
- 提供 `run -> step -> tool -> context -> cost` 的统一调试视图。
- 以最小侵入方式接入后端与 Web。

## Architecture

### DSN-0501 Debug Query Service

- 新增 `pkg/debugger/service.go` 聚合层：
  - 输入：`run_id` 或过滤条件
  - 输出：调试视图 DTO（run 概览、步骤、工具链、上下文、成本）
- 依赖：`pkg/events`, `pkg/observability/*`, `pkg/session`

### DSN-0502 数据模型与 DTO

- 新增 DTO：
  - `DebugRunSummary`
  - `DebugStepView`
  - `ToolCallView`
  - `ContextWindowView`
  - `CostBreakdownView`
- 不新增核心业务表，优先复用现有 `agent_runs/agent_events` 与观测聚合。

### DSN-0503 API 设计

- 新增 API：
  - `GET /api/v1/debug/runs`
  - `GET /api/v1/debug/runs/{id}`
  - `GET /api/v1/debug/runs/{id}/cost`
- 查询参数：`session_id`, `status`, `from`, `to`, `provider`, `model`

### DSN-0504 敏感信息脱敏

- 增加 `pkg/debugger/redaction`：
  - key 模式：`api_key`, `token`, `secret`, `authorization`
  - value 模式：长凭证串、Bearer token
- 在 API 输出前统一执行脱敏。

### DSN-0505 Web 调试页

- 新增页面：`web/src/pages/debug/`
  - 列表页：Run 列表 + 筛选
  - 详情页：流程时间线、工具调用链、上下文占比、成本面板
- 交互要求：
  - 步骤失败高亮
  - 工具链按时间排序
  - 成本图支持 provider/model 切换

## Requirement Coverage

- RQ-0501 -> DSN-0501, DSN-0503, DSN-0505
- RQ-0502 -> DSN-0501, DSN-0502, DSN-0505
- RQ-0503 -> DSN-0501, DSN-0502, DSN-0505
- RQ-0504 -> DSN-0501, DSN-0502, DSN-0505
- RQ-0505 -> DSN-0501, DSN-0503, DSN-0505
- RQ-0506 -> DSN-0501, DSN-0503
- RQ-0507 -> DSN-0504
- RQ-0508 -> DSN-0501, DSN-0504
