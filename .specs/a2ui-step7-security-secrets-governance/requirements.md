# A2UI Step7 Security & Secrets Governance Requirements

## Overview

本规范定义 A2UI 平台第 7 步：安全与密钥治理。目标是降低密钥泄露、越权执行、提示注入带来的风险，并建立最小安全基线。

## Scope

- In Scope
  - Provider 密钥治理与加载策略
  - 工具调用权限分级与策略化控制
  - Prompt/Tool 注入防护增强
  - 安全审计事件与告警钩子（基础版）
  - 安全文档与开发规范
- Out of Scope
  - 完整企业级 KMS 集成（Vault/KMS 云服务）
  - 细粒度多租户权限系统

## Functional Requirements

### RQ-701 密钥治理

- WHEN Provider 需要鉴权 THEN THE SYSTEM SHALL 通过 `auth_ref` 从环境变量或密钥适配层读取密钥。
- WHEN 密钥缺失 THEN THE SYSTEM SHALL 返回 `PROVIDER_AUTH_ERROR`，且日志不得包含明文密钥。
- WHEN 输出调试日志 THEN THE SYSTEM SHALL 对敏感字段进行脱敏。

### RQ-702 工具权限控制

- WHEN Agent 执行工具 THEN THE SYSTEM SHALL 按工具等级（read/write/exec/network）执行策略检查。
- WHEN 命中高风险策略 THEN THE SYSTEM SHALL 阻断执行并写入安全事件。

### RQ-703 注入防护增强

- WHEN 用户输入包含注入模式 THEN THE SYSTEM SHALL 阻断或降级执行路径。
- WHEN 上下文中检测到越权指令 THEN THE SYSTEM SHALL 清理高风险片段并记录审计。

### RQ-704 安全审计

- WHEN 发生安全相关决策（允许/阻断） THEN THE SYSTEM SHALL 记录结构化安全事件。
- WHEN 同类风险在短时间内频繁出现 THEN THE SYSTEM SHALL 触发告警钩子（日志级别提升或 webhook）。

## Non-functional Requirements

### RQ-705 合规基础

- WHEN 处理密钥与敏感数据 THEN THE SYSTEM SHALL 满足“最小可见性”原则（need-to-know）。

### RQ-706 可维护性

- WHEN 新增工具或 provider THEN THE SYSTEM SHALL 可通过策略配置扩展而非修改核心逻辑。

### RQ-707 可测试性

- WHEN 提交实现 THEN THE SYSTEM SHALL 覆盖密钥缺失、策略阻断、注入检测、审计写入四类测试。
