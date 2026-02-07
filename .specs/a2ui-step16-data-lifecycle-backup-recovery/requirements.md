# A2UI Step16 Data Lifecycle, Backup & Recovery Requirements

## Overview

本规范定义 A2UI 平台第 16 步：数据生命周期、备份与恢复能力。目标是保障关键数据可恢复、可归档、可清理，降低数据丢失与存储失控风险。

## Scope

- In Scope
  - 关键数据分级与保留策略
  - PostgreSQL 备份与恢复流程（全量 + 可选增量）
  - 业务级导出与恢复（workspace/session/schema/playground）
  - 数据归档与清理任务
  - 恢复演练与验收标准
- Out of Scope
  - 跨地域多活容灾
  - 云厂商专有备份服务深度集成

## Functional Requirements

### RQ-1601 数据分级与保留

- WHEN 系统存储数据 THEN THE SYSTEM SHALL 按业务价值分级（核心业务数据、审计数据、临时数据）。
- WHEN 达到保留周期 THEN THE SYSTEM SHALL 对可清理数据执行归档或删除。

### RQ-1602 备份任务

- WHEN 到达备份计划时间 THEN THE SYSTEM SHALL 自动执行数据库备份并记录备份元数据。
- WHEN 备份失败 THEN THE SYSTEM SHALL 记录告警并支持重试。

### RQ-1603 恢复流程

- WHEN 发生数据故障 THEN THE SYSTEM SHALL 支持按指定备份点恢复数据库。
- WHEN 恢复完成 THEN THE SYSTEM SHALL 执行一致性校验（关键表行数、主链路可用性）。

### RQ-1604 业务级导入导出

- WHEN 用户请求导出 THEN THE SYSTEM SHALL 支持导出 workspace/session/playground 数据包。
- WHEN 用户请求导入 THEN THE SYSTEM SHALL 支持将数据包导入到新工作区并保留版本关系。

### RQ-1605 演练与审计

- WHEN 执行恢复演练 THEN THE SYSTEM SHALL 记录演练时间、耗时、成功率与问题清单。

## Non-functional Requirements

### RQ-1606 RPO/RTO 目标

- WHEN 制定备份恢复策略 THEN THE SYSTEM SHALL 定义并跟踪 RPO/RTO 目标（MVP 可设 RPO<=24h，RTO<=2h）。

### RQ-1607 安全性

- WHEN 存储备份文件 THEN THE SYSTEM SHALL 支持加密存储与访问控制。

### RQ-1608 可测试性

- WHEN 提交变更 THEN THE SYSTEM SHALL 覆盖备份、恢复、导入导出、清理策略测试。
