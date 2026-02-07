# A2UI Step11 Analytics & Growth Tasks

## TASK-1101 增加 product_events 数据表与索引

- Linked Requirements: RQ-1101, RQ-1106
- Linked Design: DSN-1101, DSN-1106
- Description:
  - 新增 goose migration 创建 `product_events`
  - 增加关键查询索引
- DoD:
  - migration 可 up/down
- Verify:
  - `go run ./cmd/cli db:migrate`

Status: completed

## TASK-1102 实现 analytics tracker

- Linked Requirements: RQ-1101, RQ-1107
- Linked Design: DSN-1102
- Description:
  - 新增 `pkg/analytics` 与 Track API
  - 在核心 service 接入基础埋点
- DoD:
  - 关键事件可入库
- Verify:
  - `go test ./pkg/analytics/...`

Status: completed

## TASK-1103 实现日指标聚合查询

- Linked Requirements: RQ-1102
- Linked Design: DSN-1103
- Description:
  - 实现 DAU、session/workspace、agent 成功率、平均耗时统计
- DoD:
  - 指标查询结果正确
- Verify:
  - `go test ./pkg/analytics/metrics/...`

Status: completed

## TASK-1104 实现漏斗与留存计算

- Linked Requirements: RQ-1103, RQ-1104
- Linked Design: DSN-1103
- Description:
  - 实现最小漏斗与 D1/D7 留存查询
- DoD:
  - 漏斗步骤与留存数值可输出
- Verify:
  - `go test ./pkg/analytics/metrics/...`

Status: completed

## TASK-1105 暴露 analytics API

- Linked Requirements: RQ-1102, RQ-1103, RQ-1104
- Linked Design: DSN-1104
- Description:
  - 新增 `/analytics/metrics` `/analytics/funnel` `/analytics/retention`
- DoD:
  - API 返回结构化统计结果
- Verify:
  - `go test ./cmd/server/...`

Status: completed

## TASK-1106 加入数据最小化过滤

- Linked Requirements: RQ-1105
- Linked Design: DSN-1105
- Description:
  - 埋点前过滤敏感字段
  - prompt 仅保留摘要信息
- DoD:
  - 敏感字段不落库
- Verify:
  - `go test ./pkg/analytics/...`

Status: completed

## TASK-1107 增加分析文档与指标口径说明

- Linked Requirements: RQ-1102, RQ-1103, RQ-1104
- Linked Design: DSN-1103, DSN-1104
- Description:
  - 新增 `docs/analytics-metrics.md`
  - 说明字段定义、口径、已知限制
- DoD:
  - 团队可按文档解读指标
- Verify:
  - 人工审阅通过

Status: completed

## TASK-1108 全量回归测试

- Linked Requirements: RQ-1108
- Linked Design: DSN-1102, DSN-1103
- Description:
  - 运行全量测试并验证 analytics 链路
- DoD:
  - `go test ./...` 通过
- Verify:
  - `go test ./...`

Status: completed
