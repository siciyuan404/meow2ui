# A2UI Step13 Plugin & Template Marketplace Requirements

## Overview

本规范定义 A2UI 平台第 13 步：插件与模板市场基础能力。目标是把平台从“内建能力”升级为“可扩展生态”。

## Scope

- In Scope
  - 模板（Template）注册、发布、检索、版本管理
  - 插件（Plugin）注册、启用/禁用、权限声明
  - Agent 运行时按策略加载插件
  - Web 端模板市场基础页面
  - 安全审核状态（draft/reviewed/published）
- Out of Scope
  - 完整商业化计费
  - 第三方支付与结算
  - 高级推荐算法

## Functional Requirements

### RQ-1301 模板市场

- WHEN 用户发布模板 THEN THE SYSTEM SHALL 保存模板元数据（name, version, tags, schema_snapshot, owner）。
- WHEN 用户检索模板 THEN THE SYSTEM SHALL 支持关键词、标签、分类过滤。
- WHEN 模板有新版本 THEN THE SYSTEM SHALL 保留版本历史并支持回滚引用。

### RQ-1302 插件注册与管理

- WHEN 开发者注册插件 THEN THE SYSTEM SHALL 保存插件入口、能力声明、权限清单。
- WHEN 管理员禁用插件 THEN THE SYSTEM SHALL 阻止该插件在 Agent 运行时加载。

### RQ-1303 运行时加载策略

- WHEN Agent 执行任务 THEN THE SYSTEM SHALL 根据会话/项目策略加载可用插件。
- WHEN 插件请求超权限动作 THEN THE SYSTEM SHALL 拒绝执行并写安全审计。

### RQ-1304 审核与发布状态

- WHEN 模板或插件创建 THEN THE SYSTEM SHALL 初始状态为 `draft`。
- WHEN 审核通过 THEN THE SYSTEM SHALL 允许状态变更为 `published`。
- WHEN 存在风险 THEN THE SYSTEM SHALL 可标记为 `blocked`。

### RQ-1305 Web 市场页面

- WHEN 用户访问市场页 THEN THE SYSTEM SHALL 展示模板列表、插件列表与筛选能力。
- WHEN 用户选择模板 THEN THE SYSTEM SHALL 支持一键应用到当前 session。

## Non-functional Requirements

### RQ-1306 安全性

- WHEN 插件声明权限 THEN THE SYSTEM SHALL 强制与安全策略引擎联动校验。

### RQ-1307 可维护性

- WHEN 新增模板/插件字段 THEN THE SYSTEM SHALL 通过版本化 schema 向后兼容。

### RQ-1308 可测试性

- WHEN 提交实现 THEN THE SYSTEM SHALL 覆盖模板检索、插件启停、权限拒绝、应用模板的测试。
