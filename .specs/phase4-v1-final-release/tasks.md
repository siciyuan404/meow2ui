# Phase 4: 收尾与 V1 正式发布 Implementation Tasks

## Progress Summary

- Total Tasks: 7
- Completed: 0
- In Progress: 0
- Pending: 7

---

## TASK-P4-001: 准备发布摘要并提交 Product 审阅

**Status:** pending

**Type:** docs

**Traceability:**
- Requirements: RQ-001
- Design: DSN-001

**Description:**
汇总 v1 功能列表、已知风险、延后项，形成发布摘要文档，提交 Product 负责人审阅。

**Definition of Done:**
- 发布摘要文档已创建
- 包含功能列表、已知风险（RISK-001）、延后项清单
- 已提交审阅

**Dependencies:** Phase 2 和 Phase 3 全部完成

**Verification:**
- Command: `ls docs/release-artifacts/release-summary-v1.md`
- Expected: 文件存在

---

## TASK-P4-002: 获取 Product sign-off 并更新 release-gate

**Status:** pending

**Type:** docs

**Traceability:**
- Requirements: RQ-001
- Design: DSN-001

**Description:**
获取 Product 确认后，更新 release-gate-v1.md 中的 Product sign-off 和 Decision 字段。

**Definition of Done:**
- release-gate-v1.md Product 行：`done / {date}`
- Decision：`released`

**Dependencies:** TASK-P4-001

**Verification:**
- Command: `grep "Product" docs/release-artifacts/release-gate-v1.md`
- Expected: 包含 `done`

---

## TASK-P4-003: 归档发布产物完整性检查

**Status:** pending

**Type:** test

**Traceability:**
- Requirements: RQ-004
- Design: DSN-004

**Description:**
检查 docs/release-artifacts/ 下 4 个产物文件的存在性和内容完整性。

**Definition of Done:**
- release-gate-v1.md 存在且 Decision 为 released
- migration-report.md 存在且非空
- regression-matrix-v1.md 存在且非空
- release-notes-v1-rc.md 存在且非空

**Dependencies:** TASK-P4-002

**Verification:**
- Command: `ls -la docs/release-artifacts/`
- Expected: 4 个文件均存在

---

## TASK-P4-004: 生成 Changelog

**Status:** pending

**Type:** docs

**Traceability:**
- Requirements: RQ-003
- Design: DSN-003

**Description:**
基于 release-notes-v1-rc.md 和 git log 生成结构化 CHANGELOG.md。

**Definition of Done:**
- CHANGELOG.md 包含 v1.0.0 条目
- 按 Features / Infrastructure / Security / Enterprise / Known Limitations 分类
- 内容与实际功能一致

**Dependencies:** TASK-P4-002

**Parallelizable:** yes（与 TASK-P4-003 并行）

**Verification:**
- Command: `grep "v1.0.0" CHANGELOG.md`
- Expected: 匹配到版本标题

---

## TASK-P4-005: 创建 git tag 和 GitHub Release

**Status:** pending

**Type:** implementation

**Traceability:**
- Requirements: RQ-002
- Design: DSN-002

**Description:**
在 main 分支创建 annotated tag v1.0.0，推送到 remote，并通过 gh 创建 GitHub Release。

**Definition of Done:**
- `git tag -a v1.0.0 -m "Release v1.0.0"` 成功
- `git push origin v1.0.0` 成功
- `gh release create v1.0.0` 成功，body 包含 changelog 摘要

**Dependencies:** TASK-P4-002, TASK-P4-004

**Verification:**
- Command: `git tag -l v1.0.0`
- Expected: 输出 v1.0.0

---

## TASK-P4-006: 创建 v1.1 规划文档

**Status:** pending

**Type:** docs

**Traceability:**
- Requirements: RQ-005
- Design: DSN-005

**Description:**
创建 docs/roadmap-v1.1.md，列出所有从 v1 延后的工作项及优先级。

**Definition of Done:**
- 文件包含 Go/TS SDK、SSO/SCIM、CLI 真实实现、占位页面功能化、Canvas 渲染器等延后项
- 每项有优先级和初步估算

**Dependencies:** None

**Parallelizable:** yes（可与其他任务并行）

**Verification:**
- Command: `grep -c "SDK\|SSO\|Canvas" docs/roadmap-v1.1.md`
- Expected: >= 3

---

## TASK-P4-007: 关闭 v1 范围 GitHub Issue

**Status:** pending

**Type:** implementation

**Traceability:**
- Requirements: RQ-006
- Design: DSN-006

**Description:**
关闭所有 v1 范围内已完成的 GitHub Issue，为延后项添加 v1.1 label。

**Definition of Done:**
- Issue #11（队列不一致）已关闭并引用修复说明
- 其他 v1 范围已完成 issue 已关闭
- 延后项 issue 已添加 `v1.1` label

**Dependencies:** TASK-P4-005（发布完成后再关闭）

**Verification:**
- Command: `gh issue list --state open --label "v1"`
- Expected: 无 open issue（或仅剩 v1.1 标记的）

---

## Execution Notes

- TASK-P4-001 -> TASK-P4-002 串行（需要人工审阅）
- TASK-P4-003 和 TASK-P4-004 可并行
- TASK-P4-005 依赖 TASK-P4-002 和 TASK-P4-004
- TASK-P4-006 可随时并行执行
- TASK-P4-007 在 TASK-P4-005 之后执行
