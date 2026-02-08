# Release Readiness V1 Requirements

## Overview

本规范定义当前已完成队列（medium + low）进入可发布状态前的统一验收标准，确保功能可上线、可回滚、可观测。

## Scope

- In Scope
  - API 契约一致性验收
  - 数据库迁移与回滚验证
  - 关键链路回归（Agent/Flow/Debugger/Marketplace）
  - 文档可复现性验收
  - 发布前风险清单与放行门禁
- Out of Scope
  - 新功能开发
  - 大规模架构重构

## Functional Requirements

### RQ-RR-001 API 契约验收

- WHEN 执行发布验收 THEN THE SYSTEM SHALL 对新增或变更 API 完成请求/响应契约检查并记录结果。
- WHEN 发现契约破坏性变更 THEN THE SYSTEM SHALL 阻止发布并输出修复清单。

### RQ-RR-002 数据迁移验收

- WHEN 执行数据库迁移验证 THEN THE SYSTEM SHALL 完成 up/down 演练并确认核心数据完整性。
- WHEN 迁移失败或回滚失败 THEN THE SYSTEM SHALL 阻止发布并标记高风险。

### RQ-RR-003 关键链路回归

- WHEN 执行发布回归 THEN THE SYSTEM SHALL 覆盖以下链路：
  - Agent run（含 multimodal）
  - Flow orchestration
  - Debugger 查询与成本接口
  - Marketplace 发布/审核/应用

### RQ-RR-004 文档可复现

- WHEN 新贡献者按文档操作 THEN THE SYSTEM SHALL 在 30 分钟内完成环境启动、核心接口验证和最小流程走通。

## Non-functional Requirements

### RQ-RR-005 质量门禁

- WHEN 进入发布候选 THEN THE SYSTEM SHALL 满足 `go test ./...` 与 `npm --prefix web run test -- --run` 全通过。

### RQ-RR-006 可观测性门禁

- WHEN 发布前验证 THEN THE SYSTEM SHALL 确认核心指标可采集、关键日志可定位、trace_id 可串联。

### RQ-RR-007 风险可控

- WHEN 发布决策评审 THEN THE SYSTEM SHALL 提供风险分级（high/medium/low）与回滚预案。
