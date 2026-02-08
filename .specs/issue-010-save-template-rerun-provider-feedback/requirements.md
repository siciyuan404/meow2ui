# Issue #10 一键保存模板 + Rerun + Provider 反馈 Requirements

## Overview

本规范定义 Agent 到 Marketplace 的复用闭环，目标是让用户可将生成结果一键沉淀为模板，并基于上次结果继续迭代，同时提升 Provider Pool 配置反馈可见性。

## Scope

- In Scope
  - Agent 结果一键保存为模板
  - 基于最近结果 rerun
  - Provider Pool 快速切换与连接测试反馈
- Out of Scope
  - 模板商业交易流程
  - 多用户协同编辑

## Functional Requirements

### RQ-1001 一键保存模板

- WHEN Agent 运行成功 THEN THE SYSTEM SHALL 提供“一键保存为模板”入口。
- WHEN 用户点击保存 THEN THE SYSTEM SHALL 创建 marketplace template 并返回成功反馈。

### RQ-1002 基于上次结果继续生成

- WHEN 用户触发 rerun THEN THE SYSTEM SHALL 复用最近一次 prompt 与上下文继续生成。
- WHEN 最近结果不存在 THEN THE SYSTEM SHALL 给出可读提示并引导先生成一次。

### RQ-1003 Provider 切换与测试反馈

- WHEN 用户切换 provider THEN THE SYSTEM SHALL 明确展示当前激活 provider 状态。
- WHEN 用户执行连接测试 THEN THE SYSTEM SHALL 展示成功/失败反馈与错误信息。

### RQ-1004 闭环可见性

- WHEN 保存模板成功 THEN THE SYSTEM SHALL 提供可跳转 marketplace 的入口。

## Non-functional Requirements

### RQ-1005 交互性能

- WHEN 用户点击保存模板或连接测试 THEN THE SYSTEM SHALL 在 200ms 内反馈 loading 状态。

### RQ-1006 可测试性

- WHEN 提交实现 THEN THE SYSTEM SHALL 覆盖保存模板、rerun、provider 测试反馈的关键测试。
