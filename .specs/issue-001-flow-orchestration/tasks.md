# Issue #1 Flow 编排引擎 Tasks

## TASK-0101 定义 Flow 模型与校验规则

- Linked Requirements: RQ-0101, RQ-0102
- Linked Design: DSN-0101, DSN-0102
- Description:
  - 实现 flow template/node/edge/policy 数据结构
  - 实现 DAG 与条件表达式校验
- DoD:
  - 非法流程无法通过校验
- Verify:
  - `go test ./pkg/flow/validator/...`

Status: completed

## TASK-0102 增加 Flow 相关迁移与存储实现

- Linked Requirements: RQ-0101, RQ-0104
- Linked Design: DSN-0105
- Description:
  - 新增 flow templates/version/session binding 表
  - 补充 repository 接口与 SQL 实现
- DoD:
  - migration 可 up/down，CRUD 可用
- Verify:
  - `go run ./cmd/cli db:migrate`
  - `go test ./internal/infra/sqlstore/...`

Status: completed

## TASK-0103 实现编排运行时（含分支与并行）

- Linked Requirements: RQ-0102, RQ-0103, RQ-0107
- Linked Design: DSN-0103
- Description:
  - 实现拓扑调度、条件分支、有界并行执行
  - 支持 fail_fast/continue_on_error
- DoD:
  - 可执行包含分支与并行的 flow
- Verify:
  - `go test ./pkg/flow/runtime/...`

Status: completed

## TASK-0104 集成 Agent FlowExecutor 与默认回退

- Linked Requirements: RQ-0104, RQ-0106
- Linked Design: DSN-0104
- Description:
  - 在 agent 层接入 flow executor
  - 未配置 flow 时回退默认流程
- DoD:
  - 现有主链路回归通过
- Verify:
  - `go test ./pkg/agent/...`

Status: completed

## TASK-0105 提供 Flow 管理 API

- Linked Requirements: RQ-0101, RQ-0104
- Linked Design: DSN-0105
- Description:
  - 实现 flow 列表、创建、绑定 session API
- DoD:
  - API 支持创建并绑定流程版本
- Verify:
  - `go test ./cmd/server/...`

Status: completed

## TASK-0106 接入观测与审计字段

- Linked Requirements: RQ-0105
- Linked Design: DSN-0106
- Description:
  - 写入节点级事件与新增 flow 指标
- DoD:
  - 可查询节点轨迹和分支选择统计
- Verify:
  - `go test ./pkg/events/...`
  - `go test ./pkg/observability/...`

Status: completed

## TASK-0107 补充回归测试与文档

- Linked Requirements: RQ-0108
- Linked Design: DSN-0102, DSN-0103, DSN-0104
- Description:
  - 增加流程回归测试集
  - 新增 `docs/flow-orchestration.md`
- DoD:
  - 关键场景测试可稳定复现
- Verify:
  - `go test ./...`

Status: completed
