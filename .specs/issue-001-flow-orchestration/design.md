# Issue #1 Flow 编排引擎 Design

## Design Goals

- 在不破坏现有 Agent 主流程的前提下引入可配置编排。
- 兼容当前可观测与审计体系，支持调试视图消费。
- 保持默认路径稳定可回退。

## Architecture

### DSN-0101 Flow Definition Model

- 新增 `pkg/flow/model.go`：
  - `FlowTemplate`（name, version, status）
  - `FlowNode`（id, type, depends_on, timeout_ms）
  - `FlowEdge`（from, to, condition）
  - `FlowPolicy`（parallelism, failure_mode）
- 节点类型首版支持：`plan`, `emit`, `validate`, `repair`, `apply`, `custom`。

### DSN-0102 Flow Validator

- 新增 `pkg/flow/validator`：
  - DAG 校验（禁止循环依赖）
  - 输入输出契约校验
  - 条件表达式静态校验
- 发布前强制校验通过。

### DSN-0103 Orchestrator Runtime

- 新增 `pkg/flow/runtime`：
  - 拓扑调度执行
  - 条件分支选择
  - 有界并行（worker pool）
  - 失败策略：`fail_fast` / `continue_on_error`
- 执行产物输出统一 `FlowRunResult`。

### DSN-0104 Agent 集成

- 在 `pkg/agent` 增加 `FlowExecutor` 抽象：
  - 默认实现映射到固定主流程
  - 编排实现读取会话绑定 flow template
- 未命中配置时自动回退默认执行器。

### DSN-0105 存储与 API

- 新增迁移表：
  - `flow_templates`
  - `flow_template_versions`
  - `session_flow_bindings`
- 新增 API：
  - `GET /api/v1/flows`
  - `POST /api/v1/flows`
  - `POST /api/v1/flows/{id}/bind-session`

### DSN-0106 Observability 集成

- 节点执行事件复用 `agent_events`，新增字段：`flow_node_id`, `flow_template_version`。
- 指标补充：`flow_node_duration_ms`, `flow_branch_selected_total`。

## Requirement Coverage

- RQ-0101 -> DSN-0101, DSN-0105
- RQ-0102 -> DSN-0102, DSN-0103
- RQ-0103 -> DSN-0103
- RQ-0104 -> DSN-0104, DSN-0105
- RQ-0105 -> DSN-0106
- RQ-0106 -> DSN-0103, DSN-0104
- RQ-0107 -> DSN-0103
- RQ-0108 -> DSN-0102, DSN-0103, DSN-0104
