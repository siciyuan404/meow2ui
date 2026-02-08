# Phase 4: 收尾与 V1 正式发布 Requirements

## Overview

在 Phase 1~3 完成队列修复、发布验证、文档与测试补全后，本阶段执行最终收尾：获取 Product sign-off、标记正式版本、生成 changelog、归档发布产物，并规划 v1.1 延后项。

## Requirement Index

| ID | Title | Priority | Notes |
|----|-------|----------|-------|
| RQ-001 | 获取 Product sign-off | high | release-gate-v1.md 最后一项 |
| RQ-002 | 标记正式版本 | high | git tag v1.0.0 |
| RQ-003 | 生成 Changelog | medium | 基于 commit history |
| RQ-004 | 归档发布产物 | medium | 确保所有 release-artifacts 完整 |
| RQ-005 | 规划 v1.1 延后项 | medium | SDK/SSO/SCIM/占位页面等 |
| RQ-006 | 关闭相关 Issue | low | 清理 GitHub issue tracker |

---

## User Stories and Acceptance Criteria

### RQ-001: 获取 Product sign-off

As a 发布负责人, I want to 获取 Product 团队的正式签字, so that 发布决策有完整的三方确认。

**Acceptance Criteria (EARS):**

**Success Scenarios**
- WHEN 所有技术验证通过且 Engineering/QA 已签字 THE SYSTEM SHALL 在 release-gate-v1.md 中记录 Product sign-off 状态为 done
- WHEN Product sign-off 完成 THE SYSTEM SHALL 将 Decision 从 `ready_for_release_candidate` 更新为 `released`

**Outcomes to Verify**
- release-gate-v1.md 中 Product 行显示 `done / {date}`
- Decision 字段为 `released`

---

### RQ-002: 标记正式版本

As a 开发者, I want to 有一个明确的 git tag 标记 v1.0.0, so that 版本可追溯且可复现。

**Acceptance Criteria (EARS):**

**Success Scenarios**
- WHEN 发布决策为 released THE SYSTEM SHALL 在 main 分支创建 annotated tag `v1.0.0`
- WHEN tag 创建后 THE SYSTEM SHALL 推送到 remote

**Error Scenarios**
- WHEN main 分支有未合并的 hotfix THE SYSTEM SHALL 先合并再打 tag

**Outcomes to Verify**
- `git tag -l v1.0.0` 返回 v1.0.0
- GitHub Releases 页面有对应 release

---

### RQ-003: 生成 Changelog

As a 用户, I want to 有一份完整的 v1.0.0 changelog, so that 我了解这个版本包含哪些功能。

**Acceptance Criteria (EARS):**

**Success Scenarios**
- WHEN 发布 v1.0.0 THE SYSTEM SHALL 生成 CHANGELOG.md 或在 GitHub Release 中包含变更摘要
- WHEN changelog 生成 THE SYSTEM SHALL 按功能分类（Features / Fixes / Breaking Changes / Known Issues）

**Outcomes to Verify**
- CHANGELOG.md 或 GitHub Release body 包含完整的功能列表

---

### RQ-004: 归档发布产物

As a 项目维护者, I want to 所有发布产物完整归档, so that 未来审计和回溯有据可查。

**Acceptance Criteria (EARS):**

**Success Scenarios**
- WHEN 发布完成 THE SYSTEM SHALL 确认以下文件存在且内容完整：
  - `docs/release-artifacts/release-gate-v1.md`
  - `docs/release-artifacts/migration-report.md`
  - `docs/release-artifacts/regression-matrix-v1.md`
  - `docs/release-artifacts/release-notes-v1-rc.md`

**Outcomes to Verify**
- 4 个发布产物文件均存在且非空

---

### RQ-005: 规划 v1.1 延后项

As a 产品负责人, I want to 有一份明确的 v1.1 待办清单, so that 延后项不会被遗忘。

**Acceptance Criteria (EARS):**

**Success Scenarios**
- WHEN v1.0.0 发布后 THE SYSTEM SHALL 创建 v1.1 规划文档，列出所有延后项
- WHEN 延后项列出 THE SYSTEM SHALL 包含优先级和初步估算

**Outcomes to Verify**
- 存在 v1.1 规划文档
- 包含：Go/TS SDK、企业 SSO/SCIM、CLI backup/restore 真实实现、占位前端页面功能化、Canvas 预览渲染器

---

### RQ-006: 关闭相关 Issue

As a 项目维护者, I want to 已完成的 Issue 被关闭, so that issue tracker 反映真实状态。

**Acceptance Criteria (EARS):**

**Success Scenarios**
- WHEN v1.0.0 发布 THE SYSTEM SHALL 关闭所有 v1 范围内已完成的 GitHub Issue
- WHEN Issue #11（队列不一致）已修复 THE SYSTEM SHALL 关闭该 Issue 并引用修复 commit

**Outcomes to Verify**
- v1 范围内的 Issue 全部 closed

---

## Non-Functional Requirements

### Reliability
- WHEN 打 tag THE SYSTEM SHALL 使用 annotated tag 而非 lightweight tag

### Security
- WHEN 发布 THE SYSTEM SHALL 不包含 .env 或 credentials 文件

---

## Constraints and Assumptions

### Constraints
- Product sign-off 需要人工确认，不可自动化
- git tag 必须在所有 CI 通过后执行

### Assumptions
- Phase 1~3 已全部完成
- release-checklist-v1.md 30 项已勾选

---

## Out of Scope

- v1.1 功能开发
- 生产环境部署（由运维团队执行）
- 市场推广材料

---

## Open Questions

| ID | Question | Owner | Needed By |
|----|----------|-------|-----------|
| Q-001 | Product sign-off 由谁执行 | product | Phase 4 开始前 |
| Q-002 | 是否需要 GitHub Release 自动化 | engineering | Phase 4 开始前 |
