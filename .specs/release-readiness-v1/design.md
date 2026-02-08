# Release Readiness V1 Design

## Design Goals

- 用统一发布门禁把“已开发完成”提升为“可发布”。
- 提供结构化验收产物，支持审计与复盘。
- 尽量复用现有测试与脚本，降低额外成本。

## Architecture

### DSN-RR-001 发布验收清单引擎

- 新增发布清单文档：`docs/release-checklist-v1.md`
- 结构：
  - API Contract
  - DB Migration
  - Regression
  - Observability
  - Rollback

### DSN-RR-002 契约验证层

- 新增契约样例目录：`docs/contracts/`
- 对关键接口定义请求/响应示例，并在验收时逐项比对。

### DSN-RR-003 迁移验证流程

- 使用现有 `cmd/cli db:migrate` + rollback 演练。
- 输出迁移结果记录：`docs/release-artifacts/migration-report.md`

### DSN-RR-004 回归执行矩阵

- 新增回归矩阵文档：`docs/release-artifacts/regression-matrix-v1.md`
- 绑定对应命令与预期结果。

### DSN-RR-005 放行门禁与风险报告

- 新增放行报告：`docs/release-artifacts/release-gate-v1.md`
- 字段：
  - Gate status (pass/fail)
  - Open risks
  - Rollback steps
  - Sign-off

## Requirement Coverage

- RQ-RR-001 -> DSN-RR-001, DSN-RR-002
- RQ-RR-002 -> DSN-RR-003
- RQ-RR-003 -> DSN-RR-004
- RQ-RR-004 -> DSN-RR-001, DSN-RR-004
- RQ-RR-005 -> DSN-RR-004, DSN-RR-005
- RQ-RR-006 -> DSN-RR-004, DSN-RR-005
- RQ-RR-007 -> DSN-RR-005
