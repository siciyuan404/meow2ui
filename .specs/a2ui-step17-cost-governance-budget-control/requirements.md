# A2UI Step17 Cost Governance & Budget Control Requirements

## Overview

本规范定义 A2UI 平台第 17 步：成本治理与预算控制。目标是可观测、可控制、可预警地管理模型调用成本，防止失控增长。

## Scope

- In Scope
  - Provider/Model 成本计量
  - 租户/用户/工作区预算策略
  - 预算阈值告警与限流/阻断策略
  - 成本报表与趋势查询
  - 成本回溯（按 run/session/provider）
- Out of Scope
  - 自动动态采购优化
  - 企业级财务对账系统

## Functional Requirements

### RQ-1701 成本计量

- WHEN 发生模型调用 THEN THE SYSTEM SHALL 记录 token 输入/输出与估算成本。
- WHEN 使用不同 provider/model THEN THE SYSTEM SHALL 按对应单价策略计算成本。

### RQ-1702 预算策略

- WHEN 配置预算 THEN THE SYSTEM SHALL 支持日/月预算（user/workspace/tenant 维度）。
- WHEN 达到预算阈值 THEN THE SYSTEM SHALL 触发告警并执行策略（提醒/降级/阻断）。

### RQ-1703 成本查询

- WHEN 查询成本报表 THEN THE SYSTEM SHALL 返回时间区间内的成本趋势与分布。
- WHEN 查询单次运行成本 THEN THE SYSTEM SHALL 可回溯到 run、session、provider、model。

### RQ-1704 策略执行

- WHEN 成本超预算 THEN THE SYSTEM SHALL 支持：
  - 降级到低成本模型
  - 阻断高成本请求
  - 允许白名单紧急放行

### RQ-1705 审计与解释

- WHEN 成本决策触发 THEN THE SYSTEM SHALL 记录决策原因、规则 ID 与执行动作。

## Non-functional Requirements

### RQ-1706 准确性

- WHEN 成本统计计算 THEN THE SYSTEM SHALL 保证同一数据源重复计算结果一致。

### RQ-1707 可维护性

- WHEN 模型价格变化 THEN THE SYSTEM SHALL 支持配置化更新单价而不改核心逻辑。

### RQ-1708 可测试性

- WHEN 提交实现 THEN THE SYSTEM SHALL 覆盖计量、预算触发、降级/阻断、报表查询测试。
