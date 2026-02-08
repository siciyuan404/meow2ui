# Phase 2: V1 发布验证 Requirements

## Overview

对照 `docs/release-checklist-v1.md` 的 30 项检查清单，逐项执行验证并记录结果。确保 API 合约、数据库迁移、回归测试、可观测性、回滚方案全部通过，为正式发布扫清技术障碍。

## Requirement Index

| ID | Title | Priority | Notes |
|----|-------|----------|-------|
| RQ-001 | API 合约验证 | high | 5 项端点对齐检查 |
| RQ-002 | 数据库迁移验证 | high | apply/rollback/re-apply/integrity |
| RQ-003 | 回归矩阵执行 | high | 9 项功能回归 |
| RQ-004 | 可观测性验证 | medium | trace/events/logs/runbook |
| RQ-005 | 回滚方案验证 | high | CI 环境 db:rollback 自动化 |
| RQ-006 | 发布决策记录 | medium | 风险评估与最终决策 |

---

## User Stories and Acceptance Criteria

### RQ-001: API 合约验证

As a 发布负责人, I want to 确认所有 API 端点与文档/合约一致, so that 客户端集成不会因接口变更而中断。

**Acceptance Criteria (EARS):**

**Success Scenarios**
- WHEN 对比 Agent API (`/agent/run`) 的请求/响应与文档 THE SYSTEM SHALL 完全对齐
- WHEN 对比 Flow APIs (`/api/v1/flows*`) THE SYSTEM SHALL 完全对齐
- WHEN 对比 Debugger APIs (`/api/v1/debug/runs*`) THE SYSTEM SHALL 完全对齐
- WHEN 对比 Marketplace APIs (`/api/v1/marketplace/*`) THE SYSTEM SHALL 完全对齐
- WHEN 存在非向后兼容变更 THE SYSTEM SHALL 已识别并有迁移方案

**Outcomes to Verify**
- release-checklist-v1.md 第 1 节 5 项全部勾选

---

### RQ-002: 数据库迁移验证

As a 发布负责人, I want to 确认迁移在干净环境中可正确执行, so that 生产部署不会因迁移失败而中断。

**Acceptance Criteria (EARS):**

**Success Scenarios**
- WHEN 在干净 PostgreSQL 中执行全部迁移 THE SYSTEM SHALL 成功完成无错误
- WHEN 执行 rollback (down) THE SYSTEM SHALL 成功回退
- WHEN rollback 后重新 apply THE SYSTEM SHALL 成功完成
- WHEN 检查核心表数据完整性 THE SYSTEM SHALL 通过

**Error Scenarios**
- WHEN 迁移失败 THE SYSTEM SHALL 输出明确的错误信息和失败的迁移文件

**Outcomes to Verify**
- migration-report.md 已生成
- release-checklist-v1.md 第 2 节 5 项全部勾选

---

### RQ-003: 回归矩阵执行

As a 发布负责人, I want to 确认所有核心功能路径通过回归测试, so that 发布不引入已知回归。

**Acceptance Criteria (EARS):**

**Success Scenarios**
- WHEN 执行 `go test ./...` THE SYSTEM SHALL 全部通过
- WHEN 执行 `npm --prefix web run test -- --run` THE SYSTEM SHALL 全部通过
- WHEN 执行 Agent 基本流程 THE SYSTEM SHALL 成功生成 schema
- WHEN 执行多模态流程 THE SYSTEM SHALL 正确处理图片/音频引用
- WHEN 执行 Flow 编排 create/bind/run THE SYSTEM SHALL 成功完成
- WHEN 执行 Debugger list/detail/cost THE SYSTEM SHALL 返回正确数据
- WHEN 执行 Marketplace create/review/rate/apply THE SYSTEM SHALL 成功完成
- WHEN 执行 Benchmark 报告生成 THE SYSTEM SHALL 输出报告

**Outcomes to Verify**
- regression-matrix-v1.md 已生成
- release-checklist-v1.md 第 3 节 9 项全部勾选

---

### RQ-004: 可观测性验证

As a 运维人员, I want to 确认追踪链路和日志在请求链中正常工作, so that 生产问题可快速定位。

**Acceptance Criteria (EARS):**

**Success Scenarios**
- WHEN 发起 API 请求 THE SYSTEM SHALL 在响应链中传播 Trace ID
- WHEN Agent 运行完成 THE SYSTEM SHALL 在 run timeline 中记录关键事件
- WHEN 发生错误 THE SYSTEM SHALL 输出可操作的日志信息
- WHEN 按照 runbook 操作 THE SYSTEM SHALL 可复现预期结果

**Outcomes to Verify**
- release-checklist-v1.md 第 4 节 4 项全部勾选

---

### RQ-005: 回滚方案验证

As a 发布负责人, I want to 确认回滚方案可执行, so that 发布后出现问题时能快速恢复。

**Acceptance Criteria (EARS):**

**Success Scenarios**
- WHEN 在 CI Postgres 环境执行 `db:rollback` THE SYSTEM SHALL 成功回退指定步数
- WHEN 回滚后检查数据 THE SYSTEM SHALL 数据完整可读

**Outcomes to Verify**
- release-checklist-v1.md 第 5 节 4 项全部勾选
- RISK-001 状态从 mitigated 更新为 resolved

---

### RQ-006: 发布决策记录

As a 项目负责人, I want to 在 release-gate-v1.md 中记录最终发布决策, so that 发布过程有据可查。

**Acceptance Criteria (EARS):**

**Success Scenarios**
- WHEN 所有验证通过 THE SYSTEM SHALL 在 release-gate-v1.md 中记录风险评估和最终决策
- WHEN 存在未解决阻塞项 THE SYSTEM SHALL 明确记录接受理由

**Outcomes to Verify**
- release-checklist-v1.md 第 6 节 3 项全部勾选

---

## Non-Functional Requirements

### Reliability
- WHEN 执行验证 THE SYSTEM SHALL 在隔离环境中运行，不影响开发环境

### Security
- WHEN 验证涉及 API 密钥 THE SYSTEM SHALL 使用测试密钥而非生产密钥

---

## Constraints and Assumptions

### Constraints
- 需要可用的 PostgreSQL 16 实例
- 需要 Docker 环境运行集成测试

### Assumptions
- CI 环境已配置 PostgreSQL 服务
- release-gate-v1.md 中的 Engineering/QA sign-off 有效

---

## Out of Scope

- 性能基准测试（非 v1 阻塞项）
- 负载测试
- 安全渗透测试

---

## Open Questions

| ID | Question | Owner | Needed By |
|----|----------|-------|-----------|
| Q-001 | CI 环境是否已配置 PostgreSQL 16 | engineering | phase 2 开始前 |
