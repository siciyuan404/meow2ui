# A2UI Step9 Auth & Tenant Foundation Design

## Design Goals

- 为 Web 端提供最小可用多用户隔离能力。
- 不破坏现有业务流程，优先增量引入。
- 后续可平滑迁移到更强认证方案。

## Architecture

### DSN-901 用户与会话模型

新增数据模型：

- `users`
  - `id`, `email`, `password_hash`, `status`, `created_at`
- `auth_sessions`
  - `id`, `user_id`, `token_hash`, `expires_at`, `created_at`

### DSN-902 资源归属扩展

为核心表增加归属字段：

- `workspaces.owner_user_id`
- `sessions.owner_user_id`
- `playground_items.owner_user_id`

可选审计字段：
- `created_by`, `updated_by`

### DSN-903 Auth 中间件

新增 `pkg/auth`：

- `TokenService`
  - 颁发 token
  - 校验 token
- `Middleware`
  - 解析 Bearer token
  - 注入 `user_id` 到 context

### DSN-904 Owner 校验器

新增 `pkg/authz`：

- `Authorizer`
  - `CanAccessWorkspace(userID, workspaceID)`
  - `CanAccessSession(userID, sessionID)`
  - `CanAccessPlaygroundItem(userID, itemID)`

### DSN-905 密码安全

- 密码哈希：`bcrypt`（MVP）
- 不存储明文密码

### DSN-906 API 变更

新增接口：

- `POST /auth/register`
- `POST /auth/login`
- `POST /auth/logout`

受保护接口加鉴权：
- `/workspace/*`, `/session/*`, `/agent/*`, `/playground/*`

### DSN-907 迁移与回填

- goose 新迁移增加用户与归属字段
- 对历史数据可用 `system` 用户回填 owner（MVP）

## Requirement Coverage

- RQ-901 -> DSN-901, DSN-903
- RQ-902 -> DSN-902, DSN-904
- RQ-903 -> DSN-903, DSN-906
- RQ-904 -> DSN-902
- RQ-905 -> DSN-905
- RQ-906 -> DSN-903
- RQ-907 -> DSN-903, DSN-904, DSN-905
