# A2UI Step11 Analytics & Growth Requirements

## Overview

本规范定义 A2UI 平台第 11 步：产品指标与增长分析体系。目标是让团队能量化评估 Web 端与 Agent 能力的实际价值，并指导后续迭代。

## Scope

- In Scope
  - 关键业务事件埋点（workspace/session/agent/playground）
  - 指标定义与统计口径
  - 漏斗与留存基础分析
  - 指标查询 API（MVP）
  - 分析文档与看板字段定义
- Out of Scope
  - 全量 BI 平台建设
  - 高级归因模型与实验平台

## Functional Requirements

### RQ-1101 关键事件埋点

- WHEN 用户创建 workspace/session THEN THE SYSTEM SHALL 记录创建事件。
- WHEN 用户触发 `agent/run` THEN THE SYSTEM SHALL 记录请求、成功/失败、耗时、重试情况。
- WHEN 用户保存 Playground 项 THEN THE SYSTEM SHALL 记录保存事件与来源信息。

### RQ-1102 指标统计

- WHEN 请求指标 API THEN THE SYSTEM SHALL 返回日维度核心指标（DAU、会话创建数、agent 成功率、平均耗时）。
- WHEN 查询时间区间 THEN THE SYSTEM SHALL 支持按起止时间过滤。

### RQ-1103 漏斗分析

- WHEN 查询漏斗 THEN THE SYSTEM SHALL 返回最小漏斗：
  - `访问 Web` -> `创建 Workspace` -> `创建 Session` -> `运行 Agent` -> `保存 Playground`
- WHEN 漏斗步骤中断 THEN THE SYSTEM SHALL 可识别主要流失节点。

### RQ-1104 留存分析（基础）

- WHEN 查询留存 THEN THE SYSTEM SHALL 返回 D1/D7 基础留存数据（按用户或会话）。

### RQ-1105 隐私与安全

- WHEN 记录事件 THEN THE SYSTEM SHALL 避免存储敏感明文（如 API key、密码、完整私密 prompt）。

## Non-functional Requirements

### RQ-1106 性能

- WHEN 写入事件 THEN THE SYSTEM SHALL 对主流程开销控制在可接受范围（单次额外开销 < 20ms）。

### RQ-1107 可维护性

- WHEN 新增功能模块 THEN THE SYSTEM SHALL 可通过统一事件模型扩展埋点。

### RQ-1108 可测试性

- WHEN 提交实现 THEN THE SYSTEM SHALL 覆盖事件写入、指标查询、漏斗计算基础测试。
