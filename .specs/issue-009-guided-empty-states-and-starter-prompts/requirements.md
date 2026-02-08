# Issue #9 引导式空状态与 Starter Prompts Requirements

## Overview

本规范定义 Workspaces/Agent/Debug/Provider 四个页面的引导式空状态与 Agent starter prompts，目标是提升首轮成功率并降低新用户学习成本。

## Scope

- In Scope
  - 四个页面空状态引导（含明确 CTA）
  - Agent starter prompts（Dashboard/Form/Landing）
  - Debug 页关键指标摘要区
- Out of Scope
  - 个性化推荐算法
  - 复杂新手教程系统

## Functional Requirements

### RQ-0901 空状态引导

- WHEN 页面无可展示数据 THEN THE SYSTEM SHALL 展示清晰的空状态文案与至少 1 个可执行 CTA。

### RQ-0902 Agent Starter Prompts

- WHEN 用户进入 Agent 页面 THEN THE SYSTEM SHALL 提供可点击 starter prompt。
- WHEN 用户点击 starter prompt THEN THE SYSTEM SHALL 自动填充输入框并可一键发送。

### RQ-0903 Debug 指标摘要

- WHEN 用户进入 Debug 页面 THEN THE SYSTEM SHALL 在首屏展示成功率、平均耗时、成本摘要。

### RQ-0904 引导闭环

- WHEN 新用户按引导流程操作 THEN THE SYSTEM SHALL 可在无外部文档情况下完成首次生成闭环。

## Non-functional Requirements

### RQ-0905 一致性

- WHEN 渲染引导 UI THEN THE SYSTEM SHALL 保持统一字体、颜色、按钮样式。

### RQ-0906 可测试性

- WHEN 提交实现 THEN THE SYSTEM SHALL 覆盖空状态与 starter prompt 关键路径测试。
