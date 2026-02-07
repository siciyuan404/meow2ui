# A2UI Go Agent Platform Requirements

## Overview

本规范定义一个以 Go 后端为核心的 A2UI（语言生成 UI JSON）平台。系统支持工作目录与多会话管理、可插拔 AI Provider/Model、Agent 上下文编排、主题与案例沉淀（Playground）。

## Scope

- In Scope
  - Go 后端 `pkg` 模块化架构
  - Workspace/Session/Version 管理
  - Agent 生成链路（Plan -> Emit -> Validate -> Repair -> Apply）
  - Provider/Model 注册与路由（文本优先，预留多模态）
  - Theme 管理与会话绑定
  - Playground 分类/标签/保存
  - 审计日志、基础可观测性、权限护栏
- Out of Scope (MVP)
  - 语音输入实时流处理
  - 视频生成与编辑
  - 多租户计费与复杂 RBAC 控制台

## Personas

- 产品工程师：用自然语言快速生成和迭代 UI 原型
- 设计工程师：切换主题并沉淀可复用案例
- 平台开发者：通过 Go `pkg` 引入 Agent 能力到其他项目

## Functional Requirements

### RQ-001 Workspace 与 Session 管理

- WHEN 用户创建工作目录 THEN THE SYSTEM SHALL 创建 workspace 记录并绑定根路径。
- WHEN 用户在 workspace 下创建会话 THEN THE SYSTEM SHALL 为会话分配唯一 ID 并初始化 schema 版本。
- WHEN 会话执行任何写入操作 THEN THE SYSTEM SHALL 生成新版本快照并保留可回退链。
- WHEN 用户请求查看会话历史 THEN THE SYSTEM SHALL 返回按时间排序的版本列表及摘要。

### RQ-002 A2UI Agent 生成闭环

- WHEN 用户提交自然语言任务 THEN THE SYSTEM SHALL 先生成结构化执行计划（Plan）。
- WHEN Plan 通过基础校验 THEN THE SYSTEM SHALL 生成 UI JSON 或 Patch（Emit）。
- WHEN Emit 结果不符合 Schema 或组件约束 THEN THE SYSTEM SHALL 触发自动修复流程（Repair）。
- WHEN 修复成功 THEN THE SYSTEM SHALL 应用变更并输出最新 schema 版本。
- WHEN 连续修复达到上限仍失败 THEN THE SYSTEM SHALL 返回可读错误并拒绝写入。

### RQ-003 Provider 与 Model 管理

- WHEN 管理员注册 Provider THEN THE SYSTEM SHALL 保存 provider 基础配置（name/baseURL/auth/timeout）。
- WHEN 管理员注册模型 THEN THE SYSTEM SHALL 记录模型能力标签（text/image/audio/video/tool_call）。
- WHEN Agent 触发任务 THEN THE SYSTEM SHALL 按任务类型路由到匹配模型。
- WHEN Provider 不可用 THEN THE SYSTEM SHALL 执行降级路由或返回明确错误原因。

### RQ-004 Agent 上下文管理

- WHEN 进入新一轮 Agent 任务 THEN THE SYSTEM SHALL 构建三层上下文（system/session/task）。
- WHEN 历史上下文超过预算 THEN THE SYSTEM SHALL 对历史进行摘要压缩并保留关键决策。
- WHEN 用户指定“仅修改某区域” THEN THE SYSTEM SHALL 在 task 上下文中强制局部变更约束。

### RQ-005 Theme 管理

- WHEN 用户切换主题 THEN THE SYSTEM SHALL 仅切换 token 映射而不修改结构 schema。
- WHEN 会话保存版本 THEN THE SYSTEM SHALL 记录该版本对应的 theme 快照。

### RQ-006 Playground 管理

- WHEN 用户将会话内容保存到 Playground THEN THE SYSTEM SHALL 支持按分类和标签入库。
- WHEN 保存 Playground 项时 THEN THE SYSTEM SHALL 记录来源会话版本与主题信息。
- WHEN 用户检索 Playground THEN THE SYSTEM SHALL 支持分类过滤、标签过滤、关键词检索。

### RQ-007 安全与护栏

- WHEN Agent 请求执行高风险工具动作 THEN THE SYSTEM SHALL 触发策略校验并执行阻断或确认。
- WHEN 输入包含潜在注入指令 THEN THE SYSTEM SHALL 启用 prompt/tool 注入防护策略。

### RQ-008 可观测与审计

- WHEN Agent 执行任务 THEN THE SYSTEM SHALL 记录步骤事件、耗时、token 与错误类型。
- WHEN 用户追溯某版本生成过程 THEN THE SYSTEM SHALL 提供可审计事件链路。

## Non-functional Requirements

### RQ-009 性能

- WHEN 单次 A2UI 请求处理 THEN THE SYSTEM SHALL 在 P95 <= 8s（不含外部模型长尾）内完成主流程。
- WHEN 同时有 20 个会话活跃 THEN THE SYSTEM SHALL 保持核心 API 成功率 >= 99%。

### RQ-010 可靠性

- WHEN 外部模型调用失败 THEN THE SYSTEM SHALL 提供可重试机制并保证版本数据不损坏。
- WHEN 服务重启 THEN THE SYSTEM SHALL 保证 workspace/session/version 数据完整恢复。

### RQ-011 可扩展性

- WHEN 新增模型供应商 THEN THE SYSTEM SHALL 在不修改 Agent 核心流程的前提下完成接入。
- WHEN 其他项目引入能力包 THEN THE SYSTEM SHALL 通过 `pkg` 暴露稳定接口。

## Acceptance Checklist

- 每个功能需求都有 EARS 验收标准
- 明确 MVP 边界与非目标
- 定义性能、可靠性、扩展性底线
