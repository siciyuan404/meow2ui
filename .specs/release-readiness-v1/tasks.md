# Release Readiness V1 Tasks

## TASK-RR-001 建立发布清单与门禁模板

- Linked Requirements: RQ-RR-001, RQ-RR-007
- Linked Design: DSN-RR-001, DSN-RR-005
- Description:
  - 新增 `docs/release-checklist-v1.md`
  - 新增放行报告模板 `docs/release-artifacts/release-gate-v1.md`
- DoD:
  - 清单覆盖 API/迁移/回归/观测/回滚
- Verify:
  - 人工审阅通过

Status: completed

## TASK-RR-002 补齐关键 API 契约样例

- Linked Requirements: RQ-RR-001
- Linked Design: DSN-RR-002
- Description:
  - 新增 `docs/contracts/` 下关键接口请求/响应样例
- DoD:
  - 覆盖 Agent/Flow/Debugger/Marketplace
- Verify:
  - 人工对照接口返回验证

Status: completed

## TASK-RR-003 执行迁移 up/down 演练并产出报告

- Linked Requirements: RQ-RR-002
- Linked Design: DSN-RR-003
- Description:
  - 执行 migration 与 rollback 演练
  - 产出 `migration-report.md`
- DoD:
  - 新增迁移可安全回滚
- Verify:
  - `go run ./cmd/cli db:migrate`
  - rollback 命令执行记录

Status: completed

## TASK-RR-004 执行发布回归矩阵

- Linked Requirements: RQ-RR-003, RQ-RR-005
- Linked Design: DSN-RR-004
- Description:
  - 执行并记录 Go/Web 核心回归
  - 产出 `regression-matrix-v1.md`
- DoD:
  - 所有核心链路状态可追踪
- Verify:
  - `go test ./...`
  - `npm --prefix web run test -- --run`

Status: completed

## TASK-RR-005 观测与日志链路验收

- Linked Requirements: RQ-RR-006
- Linked Design: DSN-RR-004, DSN-RR-005
- Description:
  - 验证 trace_id、关键事件、核心指标输出
- DoD:
  - 问题可通过日志/trace 快速定位
- Verify:
  - 人工演练 debugger + runbook

Status: completed

## TASK-RR-006 形成最终放行结论

- Linked Requirements: RQ-RR-007
- Linked Design: DSN-RR-005
- Description:
  - 汇总风险与回滚策略
  - 完成放行结论文档
- DoD:
  - 形成可审计的 release decision
- Verify:
  - `docs/release-artifacts/release-gate-v1.md` 完整

Status: completed
