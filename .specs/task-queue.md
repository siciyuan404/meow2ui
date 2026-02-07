# A2UI Spec Task Queue

## Queue Strategy

- Mode: `strict`（默认严格串行，当前项阻塞则暂停队列）
- Ordering Rule:
  1. 先做基础设施与持久化
  2. 再做输出质量与错误模型
  3. 再做真实 provider 路由与降级
  4. 最后做 playground 检索增强
  5. 进入产品化能力（生产、质量、安全、Web、多用户、发布、分析、开放平台）

## Current Snapshot

- Completed Queues: `QUEUE-001` ~ `QUEUE-012`
- Pending Queues: `QUEUE-013` ~ `QUEUE-018`
- Pending Task Files: `6`
- Pending Task Items: `44`

## Queue Items

| Queue ID | Spec | Tasks File | Priority | Status | Mode | Depends On |
|---|---|---|---|---|---|---|
| QUEUE-001 | Step1 PostgreSQL Persistence | `.specs/a2ui-step1-postgres-persistence/tasks.md` | high | completed | strict | - |
| QUEUE-002 | Step3 Validation Chain | `.specs/a2ui-step3-validation-chain/tasks.md` | high | completed | strict | QUEUE-001 |
| QUEUE-003 | Step2 Real Provider Routing | `.specs/a2ui-step2-real-provider-routing/tasks.md` | high | completed | strict | QUEUE-001, QUEUE-002 |
| QUEUE-004 | Step4 Playground Retrieval Memory | `.specs/a2ui-step4-playground-retrieval-memory/tasks.md` | medium | completed | strict | QUEUE-001, QUEUE-002 |
| QUEUE-005 | Step5 Production Readiness | `.specs/a2ui-step5-production-readiness/tasks.md` | high | completed | strict | QUEUE-001, QUEUE-002, QUEUE-003 |
| QUEUE-006 | Step6 Quality CI Gates | `.specs/a2ui-step6-quality-ci-gates/tasks.md` | high | completed | strict | QUEUE-005 |
| QUEUE-007 | Step7 Security & Secrets Governance | `.specs/a2ui-step7-security-secrets-governance/tasks.md` | high | completed | strict | QUEUE-005, QUEUE-006 |
| QUEUE-008 | Step8 Web Application | `.specs/a2ui-step8-web-application/tasks.md` | high | completed | strict | QUEUE-005, QUEUE-006 |
| QUEUE-009 | Step9 Auth & Tenant Foundation | `.specs/a2ui-step9-auth-tenant-foundation/tasks.md` | high | completed | strict | QUEUE-005, QUEUE-007, QUEUE-008 |
| QUEUE-010 | Step10 Deployment & Release | `.specs/a2ui-step10-deployment-release/tasks.md` | high | completed | strict | QUEUE-006, QUEUE-007, QUEUE-009 |
| QUEUE-011 | Step11 Analytics & Growth | `.specs/a2ui-step11-analytics-growth/tasks.md` | medium | completed | strict | QUEUE-008, QUEUE-009, QUEUE-010 |
| QUEUE-012 | Step12 OpenAPI & SDK Platform | `.specs/a2ui-step12-openapi-sdk-platform/tasks.md` | medium | completed | strict | QUEUE-009, QUEUE-010, QUEUE-011 |
| QUEUE-013 | Step13 Plugin & Template Marketplace | `.specs/a2ui-step13-plugin-template-marketplace/tasks.md` | medium | completed | strict | QUEUE-008, QUEUE-009, QUEUE-012 |
| QUEUE-014 | Step14 Evaluation & Quality Benchmark | `.specs/a2ui-step14-evaluation-quality-benchmark/tasks.md` | medium | completed | strict | QUEUE-011, QUEUE-012, QUEUE-013 |
| QUEUE-015 | Step15 Observability & Ops Center | `.specs/a2ui-step15-observability-ops-center/tasks.md` | medium | completed | strict | QUEUE-010, QUEUE-011, QUEUE-014 |
| QUEUE-016 | Step16 Data Lifecycle, Backup & Recovery | `.specs/a2ui-step16-data-lifecycle-backup-recovery/tasks.md` | medium | completed | strict | QUEUE-010, QUEUE-015 |
| QUEUE-017 | Step17 Cost Governance & Budget Control | `.specs/a2ui-step17-cost-governance-budget-control/tasks.md` | medium | completed | strict | QUEUE-012, QUEUE-015, QUEUE-016 |
| QUEUE-018 | Step18 Enterprise Readiness & Compliance Baseline | `.specs/a2ui-step18-enterprise-readiness-compliance/tasks.md` | medium | completed | strict | QUEUE-009, QUEUE-015, QUEUE-017 |

## Next Execution Window

按照 strict 模式与依赖关系，建议执行顺序为：

1. `QUEUE-013` Plugin & Template Marketplace
2. `QUEUE-014` Evaluation & Quality Benchmark
3. `QUEUE-015` Observability & Ops Center
4. `QUEUE-016` Data Lifecycle, Backup & Recovery
5. `QUEUE-017` Cost Governance & Budget Control
6. `QUEUE-018` Enterprise Readiness & Compliance Baseline

## Active Queue Tasks

### QUEUE-013 -> `.specs/a2ui-step13-plugin-template-marketplace/tasks.md`

- TASK-1301
- TASK-1302
- TASK-1303
- TASK-1304
- TASK-1305
- TASK-1306
- TASK-1307
- TASK-1308

Progress: 1/8 completed (`TASK-1301`)

### QUEUE-014 -> `.specs/a2ui-step14-evaluation-quality-benchmark/tasks.md`

- TASK-1401
- TASK-1402
- TASK-1403
- TASK-1404
- TASK-1405
- TASK-1406
- TASK-1407

### QUEUE-015 -> `.specs/a2ui-step15-observability-ops-center/tasks.md`

- TASK-1501
- TASK-1502
- TASK-1503
- TASK-1504
- TASK-1505
- TASK-1506
- TASK-1507

### QUEUE-016 -> `.specs/a2ui-step16-data-lifecycle-backup-recovery/tasks.md`

- TASK-1601
- TASK-1602
- TASK-1603
- TASK-1604
- TASK-1605
- TASK-1606
- TASK-1607
- TASK-1608

### QUEUE-017 -> `.specs/a2ui-step17-cost-governance-budget-control/tasks.md`

- TASK-1701
- TASK-1702
- TASK-1703
- TASK-1704
- TASK-1705
- TASK-1706
- TASK-1707

### QUEUE-018 -> `.specs/a2ui-step18-enterprise-readiness-compliance/tasks.md`

- TASK-1801
- TASK-1802
- TASK-1803
- TASK-1804
- TASK-1805
- TASK-1806
- TASK-1807

## Task Assignment

### QUEUE-001 -> `.specs/a2ui-step1-postgres-persistence/tasks.md`

- TASK-101
- TASK-102
- TASK-103
- TASK-104
- TASK-105
- TASK-106

### QUEUE-002 -> `.specs/a2ui-step3-validation-chain/tasks.md`

- TASK-301
- TASK-302
- TASK-303
- TASK-304
- TASK-305

### QUEUE-003 -> `.specs/a2ui-step2-real-provider-routing/tasks.md`

- TASK-201
- TASK-202
- TASK-203
- TASK-204
- TASK-205
- TASK-206

### QUEUE-004 -> `.specs/a2ui-step4-playground-retrieval-memory/tasks.md`

- TASK-401
- TASK-402
- TASK-403
- TASK-404
- TASK-405
- TASK-406
- TASK-407

### QUEUE-005 -> `.specs/a2ui-step5-production-readiness/tasks.md`

- TASK-501
- TASK-502
- TASK-503
- TASK-504
- TASK-505

### QUEUE-006 -> `.specs/a2ui-step6-quality-ci-gates/tasks.md`

- TASK-601
- TASK-602
- TASK-603
- TASK-604
- TASK-605

### QUEUE-007 -> `.specs/a2ui-step7-security-secrets-governance/tasks.md`

- TASK-701
- TASK-702
- TASK-703
- TASK-704
- TASK-705
- TASK-706
- TASK-707

### QUEUE-008 -> `.specs/a2ui-step8-web-application/tasks.md`

- TASK-801
- TASK-802
- TASK-803
- TASK-804
- TASK-805
- TASK-806
- TASK-807

### QUEUE-009 -> `.specs/a2ui-step9-auth-tenant-foundation/tasks.md`

- TASK-901
- TASK-902
- TASK-903
- TASK-904
- TASK-905
- TASK-906

### QUEUE-010 -> `.specs/a2ui-step10-deployment-release/tasks.md`

- TASK-1001
- TASK-1002
- TASK-1003
- TASK-1004
- TASK-1005
- TASK-1006

### QUEUE-011 -> `.specs/a2ui-step11-analytics-growth/tasks.md`

- TASK-1101
- TASK-1102
- TASK-1103
- TASK-1104
- TASK-1105
- TASK-1106
- TASK-1107
- TASK-1108

### QUEUE-012 -> `.specs/a2ui-step12-openapi-sdk-platform/tasks.md`

- TASK-1201
- TASK-1202
- TASK-1203
- TASK-1204
- TASK-1205
- TASK-1206
- TASK-1207
- TASK-1208
- TASK-1209

## Why This Order

- `QUEUE-001` 先完成，确保所有后续模块在真实持久化条件下开发与验证。
- `QUEUE-002` 第二，先收敛 schema 校验与错误码，减少后续 provider 与检索联调成本。
- `QUEUE-003` 第三，在稳定校验链上接入真实模型路由，定位问题更直接。
- `QUEUE-004` 最后，属于增强能力，依赖前面稳定的数据层和 agent 错误模型。
- `QUEUE-005` ~ `QUEUE-007` 构成生产基线（可运行、可验证、可防护）。
- `QUEUE-008` ~ `QUEUE-009` 构成产品可用基线（Web + 多用户隔离）。
- `QUEUE-010` ~ `QUEUE-012` 构成交付与平台化基线（发布、分析、开放能力）。

## Batch Plan

| Batch ID | Queues | Goal | Exit Criteria |
|---|---|---|---|
| BATCH-1 | QUEUE-005, QUEUE-006, QUEUE-007 | 生产可用基线 | 健康检查/错误统一/CI门禁/安全治理完成 |
| BATCH-2 | QUEUE-008, QUEUE-009 | Web 多用户可用 | Web 主流程可用，认证与 owner 隔离生效 |
| BATCH-3 | QUEUE-010 | 部署发布闭环 | 容器化 + 迁移 + 回滚手册 + 发布验收可执行 |
| BATCH-4 | QUEUE-011, QUEUE-012 | 数据驱动与平台开放 | 指标可查询，OpenAPI + SDK 可对外接入 |

## Execution Rules Per Queue Item

- 执行当前队列项时，按对应 `tasks.md` 的 `TASK-*` 顺序推进。
- 每个任务状态流转：`pending -> in_progress -> completed`。
- 每完成一个任务，立即更新对应 `tasks.md`。
- 每个队列项完成后，更新本文件的 `Status` 为 `completed`。

## Start Command (Recommended)

- 当前从 `QUEUE-013` 开始执行：
  - `TASK-1301 -> TASK-1302 -> TASK-1303 -> TASK-1304 -> TASK-1305 -> TASK-1306 -> TASK-1307 -> TASK-1308`
