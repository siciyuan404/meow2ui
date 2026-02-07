# A2UI Step17 Cost Governance & Budget Control Tasks

## TASK-1701 增加成本与预算迁移表

- Linked Requirements: RQ-1701, RQ-1702
- Linked Design: DSN-1701, DSN-1702
- Description:
  - 新增 model_pricing/cost_usage/budgets/budget_events
- DoD:
  - migration 可 up/down
- Verify:
  - `go run ./cmd/cli db:migrate`

Status: completed

## TASK-1702 实现成本计量服务

- Linked Requirements: RQ-1701, RQ-1706
- Linked Design: DSN-1701
- Description:
  - 基于 token 与价格计算 estimated_cost
  - 写入 cost_usage
- DoD:
  - 同一输入计算结果稳定一致
- Verify:
  - `go test ./pkg/cost/...`

Status: completed

## TASK-1703 实现预算策略引擎

- Linked Requirements: RQ-1702, RQ-1704
- Linked Design: DSN-1703
- Description:
  - 评估预算与阈值，输出 allow/degrade/block 决策
- DoD:
  - 可根据策略返回正确动作
- Verify:
  - `go test ./pkg/cost/policy/...`

Status: completed

## TASK-1704 集成 provider 路由成本控制

- Linked Requirements: RQ-1704
- Linked Design: DSN-1704
- Description:
  - 调用前预测成本并执行策略
  - 超预算时降级模型或阻断
- DoD:
  - 高成本请求可被策略控制
- Verify:
  - `go test ./pkg/provider/...`

Status: completed

## TASK-1705 实现成本报表 API

- Linked Requirements: RQ-1703
- Linked Design: DSN-1705
- Description:
  - 提供 summary/usage/budget 查询与配置 API
- DoD:
  - 报表可按时间区间与 scope 查询
- Verify:
  - `go test ./cmd/server/...`

Status: completed

## TASK-1706 实现预算告警与审计

- Linked Requirements: RQ-1705
- Linked Design: DSN-1706
- Description:
  - 阈值命中写 budget_events
  - 接入 alerting 通知
- DoD:
  - 告警事件可查询与追踪
- Verify:
  - `go test ./pkg/observability/alerting/...`

Status: completed

## TASK-1707 文档与回归测试

- Linked Requirements: RQ-1708
- Linked Design: DSN-1705, DSN-1706
- Description:
  - 新增 `docs/cost-governance.md`
  - 增加成本控制回归用例
- DoD:
  - 团队可按文档配置预算并验证策略
- Verify:
  - `go test ./...`

Status: completed
