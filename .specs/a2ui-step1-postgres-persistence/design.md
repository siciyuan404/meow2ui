# A2UI Step1 PostgreSQL Persistence Design

## Design Goals

- 将当前内存仓储平滑替换为 PostgreSQL 持久化。
- 保持 `pkg/store` 接口稳定，最小化上层服务改动。
- 让迁移、启动、回滚流程可脚本化和可重复执行。

## Architecture

### DSN-101 存储驱动抽象

保留 `pkg/store` 接口不变，在 `internal/infra` 增加：

- `sqlstore/`：PostgreSQL 实现
- `db/`：连接初始化与健康检查

`bootstrap.New(ctx)` 根据 `STORE_DRIVER` 选择：
- `memory` -> `memorystore.New()`
- `postgres` -> `sqlstore.New(db)`

### DSN-102 配置项

新增环境变量：

- `STORE_DRIVER`（`memory|postgres`，默认 `memory`）
- `PG_HOST`（默认 `localhost`）
- `PG_PORT`（默认 `5432`）
- `PG_USER`（默认 `postgres`）
- `PG_PASSWORD`（默认 `postgres`）
- `PG_DATABASE`（默认 `a2ui_platform`）
- `PG_SSLMODE`（默认 `disable`）

### DSN-103 数据库创建策略

新增启动前脚本或 CLI 子命令：

- 连接 `postgres` 默认库
- 检查目标库是否存在
- 不存在则创建 `a2ui_platform`

### DSN-104 goose 迁移策略

- 保持现有 `migrations/00001_init.up.sql` 和 `down.sql`
- 增加运行入口：
  - CLI: `db:migrate`
  - 可选启动钩子：`AUTO_MIGRATE=true` 时服务启动自动执行

### DSN-105 SQL Repository 设计

实现以下仓储：

- `WorkspaceRepository`
- `SessionRepository`
- `VersionRepository`
- `ProviderRepository`
- `ThemeRepository`
- `PlaygroundRepository`
- `EventRepository`

实现要点：
- 查询参数化，避免 SQL 注入
- `created_at/updated_at` 使用 `time.Time`
- JSON 字段（capabilities/token_set/payload/schema_json）使用 `[]byte` 与 JSON 编解码

### DSN-106 错误映射

统一错误映射：
- `sql.ErrNoRows` -> `ErrNotFound`
- 约束冲突 -> `ErrConflict`
- 连接错误 -> `ErrUnavailable`

### DSN-107 验证路径

最小 E2E：
1. 创建 DB
2. 运行迁移
3. 启动 server（postgres 模式）
4. 调用 `/workspace/create`、`/session/create`、`/agent/run`
5. 在 DB 中确认记录存在

## Requirement Coverage

- RQ-101 -> DSN-103
- RQ-102 -> DSN-104
- RQ-103 -> DSN-101, DSN-105, DSN-106
- RQ-104 -> DSN-101, DSN-102
- RQ-105 -> DSN-107
- RQ-106 -> DSN-106
- RQ-107 -> DSN-104
- RQ-108 -> DSN-105, DSN-107
