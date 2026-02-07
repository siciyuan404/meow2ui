# A2UI Step9 Auth & Tenant Foundation Requirements

## Overview

本规范定义 A2UI 平台第 9 步：建立认证与租户基础能力，为 Web 端多用户使用做准备。

## Scope

- In Scope
  - 用户登录与会话令牌（MVP 可先本地账号）
  - 基础租户（workspace 归属 user）
  - API 鉴权中间件
  - 资源访问控制（仅 owner 可读写）
  - 审计字段（created_by / updated_by）
- Out of Scope
  - 企业级 SSO（OIDC/SAML）
  - 复杂 RBAC（仅 owner 模型）

## Functional Requirements

### RQ-901 用户认证

- WHEN 用户提交合法凭证 THEN THE SYSTEM SHALL 返回访问令牌并建立会话。
- WHEN 令牌过期或无效 THEN THE SYSTEM SHALL 拒绝访问受保护接口。

### RQ-902 资源归属

- WHEN 用户创建 workspace/session/playground THEN THE SYSTEM SHALL 记录 owner_user_id。
- WHEN 非 owner 请求访问资源 THEN THE SYSTEM SHALL 返回 `FORBIDDEN`。

### RQ-903 API 鉴权

- WHEN 调用受保护 API THEN THE SYSTEM SHALL 从 `Authorization: Bearer` 校验身份。
- WHEN 鉴权失败 THEN THE SYSTEM SHALL 返回统一错误结构。

### RQ-904 审计一致性

- WHEN 任意资源被创建/更新 THEN THE SYSTEM SHALL 写入 `created_by/updated_by`。

## Non-functional Requirements

### RQ-905 安全基础

- WHEN 存储密码 THEN THE SYSTEM SHALL 使用安全哈希（bcrypt/argon2）。

### RQ-906 可维护性

- WHEN 后续引入 OAuth/SSO THEN THE SYSTEM SHALL 通过 AuthProvider 抽象扩展。

### RQ-907 可测试性

- WHEN 提交实现 THEN THE SYSTEM SHALL 覆盖认证成功/失败、越权访问、owner 校验测试。
