# A2UI Step17 Cost Governance & Budget Control Design

## Design Goals

- 建立统一成本计量与预算控制闭环。
- 将成本策略嵌入 provider 路由与 agent 执行链。
- 提供可追溯报表与审计记录。

## Architecture

### DSN-1701 价格与计量模型

新增表：

- `model_pricing`
  - `id`, `provider_id`, `model_id`, `currency`, `input_per_1k`, `output_per_1k`, `effective_from`
- `cost_usage`
  - `id`, `run_id`, `session_id`, `workspace_id`, `user_id`, `provider_id`, `model_id`, `token_in`, `token_out`, `estimated_cost`, `created_at`

### DSN-1702 预算模型

新增表：

- `budgets`
  - `id`, `scope_type(user/workspace/tenant)`, `scope_id`, `period(day/month)`, `amount`, `currency`, `thresholds(json)`, `action_policy(json)`, `enabled`
- `budget_events`
  - `id`, `budget_id`, `scope_id`, `event_type`, `current_spent`, `threshold`, `action_taken`, `created_at`

### DSN-1703 成本策略引擎

新增模块 `pkg/cost/policy`：

- `Evaluate(scope, predictedCost) -> Decision`
- `Decision`：`allow`, `degrade_model`, `block`, `allow_whitelist`
- 返回规则：`rule_id`, `reason`, `recommended_model`

### DSN-1704 与 Provider/Agent 集成

- 在 provider 调用前预测本次成本
- 调用成本策略引擎决定：允许/降级/阻断
- 调用后写入 `cost_usage`

### DSN-1705 报表与 API

新增接口：

- `GET /api/v1/cost/summary?start=&end=`
- `GET /api/v1/cost/usage?run_id=`
- `GET /api/v1/cost/budgets`
- `POST /api/v1/cost/budgets`

### DSN-1706 告警与审计

- 阈值命中写 `budget_events`
- 接入 observability alerting（Step15）
- 记录 `action_taken` 与 `rule_id`

## Requirement Coverage

- RQ-1701 -> DSN-1701, DSN-1704
- RQ-1702 -> DSN-1702, DSN-1703
- RQ-1703 -> DSN-1705
- RQ-1704 -> DSN-1703, DSN-1704
- RQ-1705 -> DSN-1706
- RQ-1706 -> DSN-1701
- RQ-1707 -> DSN-1701
- RQ-1708 -> DSN-1703, DSN-1705
