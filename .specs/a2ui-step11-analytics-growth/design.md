# A2UI Step11 Analytics & Growth Design

## Design Goals

- 用统一事件模型支撑产品指标分析。
- 在不引入复杂 BI 系统前提下，提供可决策的数据视图。
- 保持埋点对主链路的低侵入与低开销。

## Architecture

### DSN-1101 事件模型标准化

新增 `product_events` 表：

- `id`
- `event_type`
- `user_id`
- `workspace_id`
- `session_id`
- `run_id`
- `properties (json)`
- `occurred_at`

事件命名建议：

- `workspace_created`
- `session_created`
- `agent_run_started`
- `agent_run_completed`
- `agent_run_failed`
- `playground_saved`

### DSN-1102 埋点 SDK（后端）

新增 `pkg/analytics`：

- `Tracker`
  - `Track(ctx, Event)`

- `Event`
  - 标准字段 + 扩展 properties

通过 service 层调用 tracker，避免在 handler 散落埋点。

### DSN-1103 指标聚合服务

新增 `pkg/analytics/metrics`：

- `GetDailyMetrics(range)`
  - DAU
  - workspace/session 新增
  - agent 成功率
  - agent 平均耗时

- `GetFunnel(range)`
  - 逐步转化率

- `GetRetention(range)`
  - D1/D7 留存

### DSN-1104 查询 API

新增接口：

- `GET /analytics/metrics?start=&end=`
- `GET /analytics/funnel?start=&end=`
- `GET /analytics/retention?start=&end=`

### DSN-1105 数据最小化策略

- prompt 不记录原文，记录长度/分类标签/哈希摘要
- properties 做字段白名单过滤

### DSN-1106 性能策略

- 写入事件同步落库（MVP）
- 后续可切异步队列
- 为 `event_type`, `occurred_at`, `user_id` 建索引

## Requirement Coverage

- RQ-1101 -> DSN-1101, DSN-1102
- RQ-1102 -> DSN-1103, DSN-1104
- RQ-1103 -> DSN-1103, DSN-1104
- RQ-1104 -> DSN-1103, DSN-1104
- RQ-1105 -> DSN-1105
- RQ-1106 -> DSN-1106
- RQ-1107 -> DSN-1101, DSN-1102
- RQ-1108 -> DSN-1102, DSN-1103
