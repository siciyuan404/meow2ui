# A2UI Step18 Enterprise Readiness & Compliance Baseline Requirements

## Overview

本规范定义 A2UI 平台第 18 步：企业级就绪与合规基线。目标是满足组织级接入要求，包括权限治理、审计导出、合规控制与运营可管理性。

## Scope

- In Scope
  - 组织（Org）与项目空间（Project）基础模型
  - RBAC 角色权限（owner/admin/member/viewer）
  - 审计日志导出与保留策略
  - 合规基线控制（数据最小化、密钥轮换策略、访问审批记录）
  - 企业集成预留（SCIM/SSO 占位接口）
- Out of Scope
  - 完整 SOC2/ISO 认证流程
  - 全功能 IAM 平台替代

## Functional Requirements

### RQ-1801 组织与项目模型

- WHEN 企业创建组织 THEN THE SYSTEM SHALL 支持组织级资源隔离。
- WHEN 组织下创建项目 THEN THE SYSTEM SHALL 将 workspace/session 绑定到项目空间。

### RQ-1802 RBAC 权限

- WHEN 用户访问资源 THEN THE SYSTEM SHALL 按角色权限进行授权判定。
- WHEN 权限不足 THEN THE SYSTEM SHALL 返回 `FORBIDDEN` 并记录审计事件。

### RQ-1803 审计导出

- WHEN 管理员请求审计导出 THEN THE SYSTEM SHALL 支持按时间区间导出审计日志（JSON/CSV）。
- WHEN 导出完成 THEN THE SYSTEM SHALL 记录导出任务元数据与操作者。

### RQ-1804 合规控制

- WHEN 处理敏感配置 THEN THE SYSTEM SHALL 提供密钥轮换记录与策略检查。
- WHEN 保留周期到期 THEN THE SYSTEM SHALL 支持自动归档或删除审计数据。

### RQ-1805 企业接入预留

- WHEN 企业需要身份系统对接 THEN THE SYSTEM SHALL 提供 SSO/SCIM 兼容接口占位和扩展点。

## Non-functional Requirements

### RQ-1806 可追溯性

- WHEN 发生权限变更或高风险操作 THEN THE SYSTEM SHALL 保证可追溯到用户、角色、时间与来源。

### RQ-1807 可维护性

- WHEN 新增角色或权限点 THEN THE SYSTEM SHALL 通过配置扩展，不破坏现有授权逻辑。

### RQ-1808 可测试性

- WHEN 提交变更 THEN THE SYSTEM SHALL 覆盖 RBAC 判定、审计导出、保留策略测试。
