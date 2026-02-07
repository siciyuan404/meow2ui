# A2UI Step8 Web Application Requirements

## Overview

本规范定义 A2UI 平台第 8 步：建设 Web 端应用（Browser UI），让用户通过网页完成 workspace/session 管理、Agent 生成、主题切换、Playground 管理与检索。

## Scope

- In Scope
  - Web 端基础框架与路由
  - 登录前提下的单用户模式（MVP 不做复杂鉴权）
  - 与现有 Go API 对接
  - 页面：Workspace、Session、Editor+Preview、Playground、Settings
  - 基础错误态/加载态/空态
- Out of Scope
  - 完整多租户权限体系
  - 实时协同编辑（OT/CRDT）
  - 高级可视化拖拽编辑器（先保留 JSON+预览）

## Functional Requirements

### RQ-801 Web 应用基础结构

- WHEN 用户访问 Web 端 THEN THE SYSTEM SHALL 提供清晰导航到 Workspace/Session/Playground/Settings。
- WHEN 页面刷新 THEN THE SYSTEM SHALL 保留当前工作上下文（active workspace/session）。

### RQ-802 Workspace/Session 管理页面

- WHEN 用户创建 workspace THEN THE SYSTEM SHALL 调用后端创建接口并更新列表。
- WHEN 用户在 workspace 下创建 session THEN THE SYSTEM SHALL 初始化 schema 并进入编辑页面。

### RQ-803 A2UI 生成与预览页面

- WHEN 用户输入 prompt 并提交 THEN THE SYSTEM SHALL 调用 `agent/run` 并展示最新 schema。
- WHEN 生成失败 THEN THE SYSTEM SHALL 显示结构化错误信息（错误码 + 可读提示）。
- WHEN 有版本历史 THEN THE SYSTEM SHALL 支持查看历史版本与切换。

### RQ-804 Theme 管理页面

- WHEN 用户切换主题 THEN THE SYSTEM SHALL 触发会话主题更新并刷新预览。
- WHEN 用户创建自定义主题 THEN THE SYSTEM SHALL 保存 theme token 并可复用。

### RQ-805 Playground 页面

- WHEN 用户保存当前会话内容 THEN THE SYSTEM SHALL 支持分类与标签保存。
- WHEN 用户检索 Playground THEN THE SYSTEM SHALL 支持关键词、分类、标签过滤。

### RQ-806 API 状态与错误体验

- WHEN API 请求中 THEN THE SYSTEM SHALL 显示明确 loading 状态。
- WHEN API 失败 THEN THE SYSTEM SHALL 显示可操作错误提示，并支持重试。

## Non-functional Requirements

### RQ-807 兼容性

- WHEN 用户使用现代浏览器（Chrome/Edge/Firefox） THEN THE SYSTEM SHALL 正常运行。

### RQ-808 响应式布局

- WHEN 屏幕宽度变化 THEN THE SYSTEM SHALL 在桌面与中等屏幕下保持可用布局。

### RQ-809 可维护性

- WHEN 新增页面模块 THEN THE SYSTEM SHALL 按功能目录拆分，避免单文件膨胀。

### RQ-810 可测试性

- WHEN 提交前端改动 THEN THE SYSTEM SHALL 提供关键页面与 API flow 的基础测试。
