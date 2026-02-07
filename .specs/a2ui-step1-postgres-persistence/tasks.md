# A2UI Step1 PostgreSQL Persistence Tasks

## TASK-101 添加 PostgreSQL 连接与配置模块

- Linked Requirements: RQ-104, RQ-106
- Linked Design: DSN-101, DSN-102
- Description:
  - 新增 `internal/infra/db`，实现 DSN 构建、连接、ping
  - 支持环境变量读取与默认值
- DoD:
  - `STORE_DRIVER=postgres` 时可建立连接
  - 连接失败输出明确错误
- Verify:
  - `go test ./internal/infra/db/...`

Status: completed

## TASK-102 实现数据库创建命令

- Linked Requirements: RQ-101
- Linked Design: DSN-103
- Description:
  - 在 CLI 新增 `db:create`
  - 检测并创建 `a2ui_platform` 数据库
- DoD:
  - 未创建时可成功创建
  - 已存在时幂等返回
- Verify:
  - `go run ./cmd/cli db:create`

Status: completed

## TASK-103 接入 goose 迁移执行入口

- Linked Requirements: RQ-102, RQ-107
- Linked Design: DSN-104
- Description:
  - 新增 `db:migrate` 命令执行 `migrations`
  - 可选 `AUTO_MIGRATE=true` 启动时执行
- DoD:
  - 迁移成功后 `goose status` 为 up
  - 迁移失败阻断启动（自动迁移场景）
- Verify:
  - `go run ./cmd/cli db:migrate`

Status: completed

## TASK-104 实现 sqlstore 仓储（核心表）

- Linked Requirements: RQ-103, RQ-108
- Linked Design: DSN-105, DSN-106
- Description:
  - 实现 Workspace/Session/Version/Provider/Theme/Playground/Event SQL 仓储
  - 对接 `pkg/store` 接口
- DoD:
  - 编译通过
  - 关键 CRUD 可执行
- Verify:
  - `go test ./internal/infra/sqlstore/...`

Status: completed

## TASK-105 改造 bootstrap 支持驱动切换

- Linked Requirements: RQ-104
- Linked Design: DSN-101, DSN-102
- Description:
  - `bootstrap.New` 按 `STORE_DRIVER` 初始化 memory/postgres
  - 保持上层服务初始化代码不变
- DoD:
  - memory 模式可回归跑通
  - postgres 模式可跑通
- Verify:
  - `go test ./internal/infra/bootstrap/...`

Status: completed

## TASK-106 端到端验证与文档

- Linked Requirements: RQ-105, RQ-108
- Linked Design: DSN-107
- Description:
  - 跑通 `db:create -> db:migrate -> server -> API`
  - 输出简明使用文档（命令与环境变量）
- DoD:
  - 能在本地 postgres 完成一条完整链路
  - 文档可复现
- Verify:
  - `go test ./...`
  - 手动 API 验证通过

Status: completed
