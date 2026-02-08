# Phase 2: V1 发布验证 Implementation Tasks

## Progress Summary

- Total Tasks: 8
- Completed: 8
- In Progress: 0
- Pending: 0

---

## TASK-P2-001: API 合约对比 — Agent & Flow

**Status:** completed

**Type:** test

**Traceability:**
- Requirements: RQ-001
- Design: DSN-001

**Description:**
对比 `/agent/run` 和 `/api/v1/flows*` 端点的实际路由注册与文档/OpenAPI 规范。

**Definition of Done:**
- Agent API 请求/响应与文档对齐
- Flow APIs 请求/响应与文档对齐
- 不一致项已记录

**Dependencies:** None

**Verification:**
- Command: `grep -c "agent/run\|/api/v1/flows" cmd/server/main.go`
- Expected: 匹配到对应路由注册

---

## TASK-P2-002: API 合约对比 — Debugger & Marketplace

**Status:** completed

**Type:** test

**Traceability:**
- Requirements: RQ-001
- Design: DSN-001

**Description:**
对比 `/api/v1/debug/runs*` 和 `/api/v1/marketplace/*` 端点。

**Definition of Done:**
- Debugger APIs 对齐
- Marketplace APIs 对齐
- 非向后兼容变更已识别

**Dependencies:** None

**Parallelizable:** yes（与 TASK-P2-001 并行）

**Verification:**
- Command: `grep -c "debug/runs\|marketplace" cmd/server/main.go`
- Expected: 匹配到对应路由注册

---

## TASK-P2-003: 数据库迁移全生命周期验证

**Status:** completed

**Type:** test

**Traceability:**
- Requirements: RQ-002
- Design: DSN-002

**Description:**
在干净 PostgreSQL 中执行 apply -> rollback -> re-apply，验证数据完整性，生成 migration-report.md。

**Definition of Done:**
- 14 个迁移全部 apply 成功
- rollback 到 0 成功
- re-apply 成功
- 核心表结构完整
- `docs/release-artifacts/migration-report.md` 已更新

**Dependencies:** None

**Verification:**
- Command: `docker compose -f deploy/docker-compose.dev.yml up -d postgres && go run ./cmd/cli db:migrate`
- Expected: 迁移成功完成

---

## TASK-P2-004: 回归矩阵 — Go 测试 + npm 测试

**Status:** completed

**Type:** test

**Traceability:**
- Requirements: RQ-003
- Design: DSN-003

**Description:**
执行后端 Go 测试和前端 npm 测试。

**Definition of Done:**
- `go test ./...` 全部通过
- `npm --prefix web run test -- --run` 全部通过

**Dependencies:** None

**Verification:**
- Command: `go test ./... && npm --prefix web run test -- --run`
- Expected: 全部 PASS

---

## TASK-P2-005: 回归矩阵 — 功能流程验证

**Status:** completed

**Type:** test

**Traceability:**
- Requirements: RQ-003
- Design: DSN-003

**Description:**
验证 Agent 基本流程、多模态流程、Flow 编排、Debugger、Marketplace 五个核心功能路径。

**Definition of Done:**
- Agent run 基本流程通过
- 多模态流程（image/audio ref）通过
- Flow create/bind/run 通过
- Debugger list/detail/cost 通过
- Marketplace create/review/rate/apply 通过

**Dependencies:** TASK-P2-003（需要数据库就绪）

**Verification:**
- Command: `bash scripts/smoke-test.sh`
- Expected: 所有冒烟测试通过

---

## TASK-P2-006: 可观测性验证

**Status:** completed

**Type:** test

**Traceability:**
- Requirements: RQ-004
- Design: DSN-004

**Description:**
验证 Trace ID 传播、关键事件记录、错误日志、runbook 可复现性。

**Definition of Done:**
- Trace ID 在请求链中传播
- Agent run timeline 有关键事件
- 错误路径有可操作日志
- Runbook 步骤可复现

**Dependencies:** TASK-P2-005

**Verification:**
- Command: 手动验证 + 日志检查
- Expected: 4 项可观测性检查全部通过

---

## TASK-P2-007: CI 环境回滚自动化验证

**Status:** completed

**Type:** test

**Traceability:**
- Requirements: RQ-005
- Design: DSN-005

**Description:**
在 CI Postgres 环境执行 `db:rollback`，验证 00014 -> 00013 -> 00012 顺序回退，关闭 RISK-001。

**Definition of Done:**
- db:rollback 3 成功执行
- 回退后数据库状态正确
- RISK-001 状态更新为 resolved

**Dependencies:** TASK-P2-003

**Verification:**
- Command: `go run ./cmd/cli db:rollback 3`
- Expected: 3 个迁移成功回退

---

## TASK-P2-008: 更新发布检查清单与决策文档

**Status:** completed

**Type:** docs

**Traceability:**
- Requirements: RQ-006
- Design: DSN-006

**Description:**
根据前 7 个任务的验证结果，逐项勾选 release-checklist-v1.md，更新 release-gate-v1.md 的 Decision 字段。

**Definition of Done:**
- release-checklist-v1.md 30 项全部勾选（或标注例外）
- release-gate-v1.md Decision 更新
- regression-matrix-v1.md 已生成

**Dependencies:** TASK-P2-001 ~ TASK-P2-007

**Verification:**
- Command: `grep -c "\[x\]" docs/release-checklist-v1.md`
- Expected: 30

---

## Execution Notes

- TASK-P2-001 和 TASK-P2-002 可并行执行
- TASK-P2-003 和 TASK-P2-004 可并行执行
- TASK-P2-005 依赖 TASK-P2-003
- TASK-P2-008 是最后一个任务，汇总所有结果
