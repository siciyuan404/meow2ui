# A2UI Step6 Quality CI Gates Requirements

## Overview

本规范定义 A2UI 平台第 6 步：建立质量门禁与 CI 流程，确保每次变更都能自动验证关键能力并阻断回归。

## Scope

- In Scope
  - CI 工作流（lint/test/smoke）
  - 单元测试覆盖率基线与报告
  - 关键路径集成测试（postgres + migration + agent run）
  - PR 质量门禁（必过检查）
- Out of Scope
  - 性能压测平台
  - 安全扫描平台（SAST/DAST 全套）

## Functional Requirements

### RQ-601 CI 基础流水线

- WHEN 提交 PR THEN THE SYSTEM SHALL 自动执行 `go test ./...`。
- WHEN 使用 postgres 模式校验 THEN THE SYSTEM SHALL 自动执行 `db:create` 与 `db:migrate`。

### RQ-602 覆盖率基线

- WHEN 执行测试 THEN THE SYSTEM SHALL 生成覆盖率报告。
- WHEN 覆盖率低于阈值 THEN THE SYSTEM SHALL 标记检查失败（默认阈值 40%，后续可提高）。

### RQ-603 关键路径集成验证

- WHEN CI 运行 integration job THEN THE SYSTEM SHALL 验证 `workspace -> session -> agent` 流程。
- WHEN 任一关键步骤失败 THEN THE SYSTEM SHALL 阻断合并。

### RQ-604 PR 门禁

- WHEN PR 准备合并 THEN THE SYSTEM SHALL 要求 CI 全绿且无阻塞任务。

## Non-functional Requirements

### RQ-605 可维护性

- WHEN 新增 spec 任务 THEN THE SYSTEM SHALL 可将对应验证命令纳入 CI，避免手工验证遗漏。

### RQ-606 可追溯性

- WHEN CI 失败 THEN THE SYSTEM SHALL 输出可直接定位的日志与失败步骤。
