# A2UI Step9 Auth & Tenant Foundation Tasks

## TASK-901 增加用户与认证会话表迁移

- Linked Requirements: RQ-901, RQ-904
- Linked Design: DSN-901, DSN-907
- Description:
  - 新增 users/auth_sessions 表
  - 增加核心资源 owner 字段
- DoD:
  - 迁移可 up/down
- Verify:
  - `go run ./cmd/cli db:migrate`

Status: completed

## TASK-902 实现 auth 服务（注册/登录/登出）

- Linked Requirements: RQ-901, RQ-905
- Linked Design: DSN-903, DSN-905, DSN-906
- Description:
  - 实现密码哈希、token 发放与失效
  - 暴露 auth API
- DoD:
  - 可注册登录并拿到 token
- Verify:
  - `go test ./pkg/auth/...`

Status: completed

## TASK-903 接入鉴权中间件到受保护 API

- Linked Requirements: RQ-903
- Linked Design: DSN-903, DSN-906
- Description:
  - 为 workspace/session/agent/playground 接口加鉴权
- DoD:
  - 未登录访问被拒绝
  - 登录后可访问
- Verify:
  - `go test ./cmd/server/...`

Status: completed

## TASK-904 实现 owner 授权校验

- Linked Requirements: RQ-902
- Linked Design: DSN-904
- Description:
  - 增加 owner 资源访问校验
  - 越权返回 FORBIDDEN
- DoD:
  - owner 可访问
  - 非 owner 被拒绝
- Verify:
  - `go test ./pkg/authz/...`

Status: completed

## TASK-905 更新 repository 审计字段写入

- Linked Requirements: RQ-904
- Linked Design: DSN-902
- Description:
  - 写入 created_by/updated_by
  - 读接口返回 owner 信息
- DoD:
  - 新增资源有 owner
- Verify:
  - `go test ./internal/infra/sqlstore/...`

Status: completed

## TASK-906 安全与越权测试覆盖

- Linked Requirements: RQ-907
- Linked Design: DSN-903, DSN-904, DSN-905
- Description:
  - 增加 auth/authz 关键测试
  - 全量回归
- DoD:
  - 关键安全路径测试通过
- Verify:
  - `go test ./...`

Status: completed
