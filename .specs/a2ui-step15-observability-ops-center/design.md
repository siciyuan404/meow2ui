# A2UI Step15 Observability & Ops Center Design

## Design Goals

- 为平台建立统一可观测能力，支撑生产运维。
- 与当前 telemetry/events 体系兼容，避免重复建设。
- 提供可落地的 ops API 与告警规则。

## Architecture

### DSN-1501 Metrics 模块

新增 `pkg/observability/metrics`：

- 计数器：
  - `http_requests_total`
  - `http_errors_total`
  - `agent_runs_total`
  - `provider_calls_total`
  - `provider_failures_total`
- 直方图：
  - `http_request_duration_ms`
  - `agent_run_duration_ms`
  - `provider_call_duration_ms`

标签建议：`route`, `status`, `provider`, `model`, `task_type`。

### DSN-1502 Logging 模块

新增 `pkg/observability/logging`：

- 统一结构：`timestamp`, `level`, `message`, `trace_id`, `run_id`, `session_id`, `code`, `component`
- 提供 logger 包装器，替代散落 `log.Printf`

### DSN-1503 Tracing 模块

新增 `pkg/observability/tracing`：

- `StartSpan(ctx, name)`
- `EndSpan(span, err)`
- 关键 span：
  - `api.request`
  - `agent.plan`
  - `agent.emit`
  - `provider.generate`
  - `store.query`

MVP 可本地 memory/exporter，后续对接 OTLP。

### DSN-1504 Ops API

新增接口：

- `GET /api/v1/ops/health`
- `GET /api/v1/ops/errors`
- `GET /api/v1/ops/capacity`
- `GET /api/v1/ops/alerts`

返回：当前状态、过去 5/15/60 分钟统计。

### DSN-1505 Alert Engine

新增 `pkg/observability/alerting`：

- 规则示例：
  - 错误率 > 5% 持续 5 分钟
  - P95 延迟 > 3000ms 持续 10 分钟
  - provider 失败率 > 20%
- 告警状态：`firing`, `resolved`

### DSN-1506 与现有模块集成

- `pkg/telemetry` 作为兼容层，逐步迁移到 observability 模块
- `pkg/events` 继续负责业务审计事件
- `cmd/server` 接入中间件：metrics + trace + structured logging

## Requirement Coverage

- RQ-1501 -> DSN-1501
- RQ-1502 -> DSN-1502
- RQ-1503 -> DSN-1503
- RQ-1504 -> DSN-1504
- RQ-1505 -> DSN-1505
- RQ-1506 -> DSN-1501, DSN-1503
- RQ-1507 -> DSN-1506
- RQ-1508 -> DSN-1501, DSN-1502, DSN-1503, DSN-1505
