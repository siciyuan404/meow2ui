# A2UI Step15 Observability & Ops Center Tasks

## TASK-1501 实现 metrics 统一埋点模块

- Linked Requirements: RQ-1501, RQ-1506
- Linked Design: DSN-1501
- Description:
  - 新增 metrics 包并接入 HTTP/Agent/Provider 关键路径
- DoD:
  - 可输出核心计数器与时延指标
- Verify:
  - `go test ./pkg/observability/metrics/...`

Status: completed

## TASK-1502 实现结构化日志模块

- Linked Requirements: RQ-1502
- Linked Design: DSN-1502
- Description:
  - 新增 logging 包并替换 server 关键日志输出
- DoD:
  - 错误日志包含 trace_id/code/component
- Verify:
  - `go test ./pkg/observability/logging/...`

Status: completed

## TASK-1503 接入追踪链路

- Linked Requirements: RQ-1503
- Linked Design: DSN-1503
- Description:
  - 在 API、Agent、Provider、Store 注入 span
- DoD:
  - 可在单次请求中看到完整 span 链路
- Verify:
  - `go test ./pkg/observability/tracing/...`

Status: completed

## TASK-1504 实现 ops 查询 API

- Linked Requirements: RQ-1504
- Linked Design: DSN-1504
- Description:
  - 增加 ops health/errors/capacity/alerts API
- DoD:
  - 可返回最近窗口的运行状态摘要
- Verify:
  - `go test ./cmd/server/...`

Status: completed

## TASK-1505 实现告警引擎与状态机

- Linked Requirements: RQ-1505
- Linked Design: DSN-1505
- Description:
  - 实现阈值规则判定与 firing/resolved 状态转换
- DoD:
  - 告警触发与恢复可追踪
- Verify:
  - `go test ./pkg/observability/alerting/...`

Status: completed

## TASK-1506 集成与迁移 telemetry

- Linked Requirements: RQ-1507
- Linked Design: DSN-1506
- Description:
  - 将现有 telemetry 关键调用对接新观测模块
- DoD:
  - 旧接口不破坏，新指标可用
- Verify:
  - `go test ./...`

Status: completed

## TASK-1507 文档与运维手册更新

- Linked Requirements: RQ-1504, RQ-1505
- Linked Design: DSN-1504, DSN-1505
- Description:
  - 新增 `docs/observability.md`
  - 更新告警规则与排障流程
- DoD:
  - 运维人员可按文档定位常见问题
- Verify:
  - 人工按文档演练通过

Status: completed
