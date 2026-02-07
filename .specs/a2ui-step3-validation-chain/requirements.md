# A2UI Step3 Validation Chain Requirements

## Overview

本规范聚焦 A2UI 第 3 步：强化校验链。目标是把当前基础校验升级为结构化、可追踪、可修复的校验体系。

## Scope

- In Scope
  - JSON 结构与组件白名单校验
  - props 类型校验
  - 结构化错误码与错误路径输出
  - Agent 在校验失败时返回可读错误对象
- Out of Scope
  - 完整 JSON Schema 引擎接入
  - UI 渲染层联动告警

## Functional Requirements

### RQ-301 组件白名单校验

- WHEN A2UI 输出 schema THEN THE SYSTEM SHALL 校验每个节点 `type` 必须在注册组件白名单内。
- WHEN 发现未注册组件 THEN THE SYSTEM SHALL 输出错误码 `A2UI_COMPONENT_NOT_ALLOWED`。

### RQ-302 Props 类型校验

- WHEN 节点包含 props THEN THE SYSTEM SHALL 按组件 props 规则校验类型。
- WHEN props 类型不匹配 THEN THE SYSTEM SHALL 输出错误码 `A2UI_PROP_TYPE_INVALID`。

### RQ-303 结构完整性校验

- WHEN 校验 schema THEN THE SYSTEM SHALL 校验 `schema.version`、`root.id`、`root.type`、子节点 `id/type`。
- WHEN 结构缺失 THEN THE SYSTEM SHALL 输出错误码 `A2UI_SCHEMA_REQUIRED_FIELD_MISSING`。

### RQ-304 结构化错误输出

- WHEN 校验失败 THEN THE SYSTEM SHALL 返回结构化错误数组（code/path/message）。
- WHEN Agent 修复失败 THEN THE SYSTEM SHALL 将结构化错误透传到 API/CLI 调用方。

## Non-functional Requirements

### RQ-305 可维护性

- WHEN 新增组件 THEN THE SYSTEM SHALL 支持通过注册表配置其 props 规则，无需改动校验主流程。

### RQ-306 可测试性

- WHEN 提交代码 THEN THE SYSTEM SHALL 覆盖正向与反向校验测试，并可通过 `go test ./...`。
