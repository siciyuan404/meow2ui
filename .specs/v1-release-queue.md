# V1 Release Spec Queue

## Queue Strategy

- Mode: `strict`
- Ordering Rule: 按阶段顺序执行，Phase 1 -> Phase 2 -> Phase 3 -> Phase 4

## Current Snapshot

- Completed Phases: `Phase 1`, `Phase 2`, `Phase 3`
- In Progress Phases: (none)
- Pending Phases: `Phase 4`
- Total Pending Tasks: `7`

## Queue Items

| Queue ID | Spec | Tasks File | Priority | Status | Mode | Depends On |
|---|---|---|---|---|---|---|
| PHASE-1 | 队列状态一致性修复 | `.specs/phase1-queue-consistency-fix/tasks.md` | high | completed | strict | - |
| PHASE-2 | V1 发布验证 | `.specs/phase2-v1-release-verification/tasks.md` | high | completed | strict | PHASE-1 |
| PHASE-3 | 文档与测试补全 | `.specs/phase3-docs-and-test-coverage/tasks.md` | medium | completed | strict | PHASE-1 |
| PHASE-4 | 收尾与 V1 正式发布 | `.specs/phase4-v1-final-release/tasks.md` | high | pending | strict | PHASE-2, PHASE-3 |

## Task Summary

| Phase | Total | Completed | Pending | Key Deliverables |
|---|---|---|---|---|
| PHASE-1 | 4 | 4 | 0 | 队列文件修复 + 校验脚本 |
| PHASE-2 | 8 | 8 | 0 | release-checklist 30 项勾选 + migration-report + regression-matrix |
| PHASE-3 | 7 | 7 | 0 | OpenAPI 规范 + 测试覆盖率 40%+ + CLI 文档 |
| PHASE-4 | 7 | 0 | 7 | Product sign-off + git tag v1.0.0 + CHANGELOG + v1.1 roadmap |
| **Total** | **26** | **19** | **7** | |

## Active Queue Tasks

### PHASE-1 -> `.specs/phase1-queue-consistency-fix/tasks.md` (completed)

- ~~TASK-P1-001~~ completed
- ~~TASK-P1-002~~ completed
- ~~TASK-P1-003~~ completed
- ~~TASK-P1-004~~ completed

Progress: 4/4 completed

### PHASE-2 -> `.specs/phase2-v1-release-verification/tasks.md` (completed)

- ~~TASK-P2-001~~ completed — API 合约对比 Agent & Flow（27 路由与 OpenAPI 完全对齐）
- ~~TASK-P2-002~~ completed — API 合约对比 Debugger & Marketplace（无不兼容变更）
- ~~TASK-P2-003~~ completed — 数据库迁移验证（smoke test 通过，14 迁移全部 apply）
- ~~TASK-P2-004~~ completed — 回归矩阵 Go + npm 测试全部 PASS
- ~~TASK-P2-005~~ completed — 功能流程验证（smoke test 通过 agent run 基本流程）
- ~~TASK-P2-006~~ completed — 可观测性验证（TraceID 传播 + 14 事件 Emit + 错误日志）
- ~~TASK-P2-007~~ completed — CI 回滚自动化（integration-postgres.sh 含 rollback 验证）
- ~~TASK-P2-008~~ completed — release-checklist 30 项全部勾选

Progress: 8/8 completed

### PHASE-3 -> `.specs/phase3-docs-and-test-coverage/tasks.md` (completed)

- ~~TASK-P3-001~~ completed — 提取 27 条路由清单
- ~~TASK-P3-002~~ completed — OpenAPI 规范已覆盖全部端点（无缺失）
- ~~TASK-P3-003~~ completed — pkg/agent 新增 service_test.go（7 测试）
- ~~TASK-P3-004~~ completed — pkg/store 新增 interfaces_test.go（4 测试）
- ~~TASK-P3-005~~ completed — web/src 新增 use-agent-runtime.test.ts（8 测试）
- ~~TASK-P3-006~~ completed — CLI 占位命令输出 "not yet implemented. Planned for v1.1."
- ~~TASK-P3-007~~ completed — CI 覆盖率门槛从 15% 提升到 40%

Progress: 7/7 completed

## Dependency Graph

```
PHASE-1 (队列修复)
  |
  +---> PHASE-2 (发布验证)  ----+
  |                              |
  +---> PHASE-3 (文档与测试) ---+---> PHASE-4 (正式发布)
```

- PHASE-2 和 PHASE-3 已在 PHASE-1 完成后并行执行完毕
- PHASE-4 依赖已满足，可以开始执行

## Execution Rules

- 执行当前队列项时，按对应 `tasks.md` 的 `TASK-*` 顺序推进。
- 每个任务状态流转：`pending -> in_progress -> completed`。
- 每完成一个任务，立即更新对应 `tasks.md`。
- 每个队列项完成后，更新本文件的 `Status` 为 `completed` 并更新 Current Snapshot。

## Start Command (Recommended)

- PHASE-1、PHASE-2、PHASE-3 已全部完成。
- PHASE-4 依赖已满足，可从 `TASK-P4-001` 开始执行。

## Relation to Previous Queues

- `task-queue.md`：QUEUE-001 ~ QUEUE-018 全部 completed（基础功能开发阶段）
- `issue-priority-queue.md`：QUEUE-101 ~ QUEUE-303 全部 completed（Issue 驱动功能阶段）
- 本文件：PHASE-1 ~ PHASE-3 completed，PHASE-4 pending（v1 发布收尾阶段）
