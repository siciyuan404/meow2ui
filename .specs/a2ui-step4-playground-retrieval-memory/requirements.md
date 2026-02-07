# A2UI Step4 Playground Retrieval Memory Requirements

## Overview

本规范定义 A2UI 平台第 4 步：将 Playground 案例反哺 Agent，建立检索增强（RAG-like）记忆机制，提升生成一致性与可控性。

## Scope

- In Scope
  - 基于分类、标签、关键词的案例检索
  - 将检索结果注入 Agent task context
  - 检索结果可配置数量与阈值
  - 记录“是否命中案例”与“命中来源”审计信息
- Out of Scope
  - 向量数据库与语义向量检索
  - 自动学习评分与在线反馈训练

## Functional Requirements

### RQ-401 案例检索

- WHEN 用户发起 A2UI 请求 THEN THE SYSTEM SHALL 基于会话主题、分类、标签和关键词检索 Playground 案例。
- WHEN 检索条件为空 THEN THE SYSTEM SHALL 至少使用关键词检索标题与标签。

### RQ-402 上下文注入

- WHEN 命中案例 THEN THE SYSTEM SHALL 将最多 N 条案例摘要注入 task context（默认 N=3）。
- WHEN 案例过长 THEN THE SYSTEM SHALL 使用摘要字段而不是全量 schema。

### RQ-403 可控开关

- WHEN 会话配置关闭案例增强 THEN THE SYSTEM SHALL 跳过检索并走原始 Agent 流程。
- WHEN 命中率过低 THEN THE SYSTEM SHALL 自动降级为无案例模式并继续执行。

### RQ-404 审计与可观测

- WHEN 使用了案例增强 THEN THE SYSTEM SHALL 记录命中条数、案例 ID、耗时。
- WHEN 未使用案例增强 THEN THE SYSTEM SHALL 记录跳过原因（关闭/无命中/异常降级）。

## Non-functional Requirements

### RQ-405 性能

- WHEN 执行案例检索 THEN THE SYSTEM SHALL 将额外耗时控制在 P95 <= 150ms（本地数据库）。

### RQ-406 可维护性

- WHEN 新增检索策略 THEN THE SYSTEM SHALL 通过策略接口扩展，不改动 Agent 主状态机。

### RQ-407 可测试性

- WHEN 提交实现 THEN THE SYSTEM SHALL 覆盖命中、无命中、关闭开关、异常降级四类测试场景。
