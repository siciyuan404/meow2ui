# A2UI Step10 Deployment & Release Design

## Design Goals

- 建立“可重复部署 + 可回滚发布”的基础能力。
- 与现有 CI、Postgres 迁移、健康检查流程对齐。
- 为后续上云（容器平台）保留扩展空间。

## Architecture

### DSN-1001 Docker 镜像设计

- 多阶段构建：
  1. build 阶段编译 Go 二进制
  2. runtime 阶段使用轻量基础镜像运行
- 仅复制可执行文件与必要配置模板
- 运行时注入环境变量，不把敏感信息 bake 进镜像

### DSN-1002 部署配置分层

建议目录：

```text
deploy/
  env/
    dev.env.example
    staging.env.example
    prod.env.example
  compose/
    docker-compose.dev.yml
    docker-compose.staging.yml
```

### DSN-1003 迁移执行策略

两种可选模式：

- `AUTO_MIGRATE=true` 启动时迁移（适合小规模）
- 独立迁移任务（推荐 staging/prod）
  - `go run ./cmd/cli db:migrate`

生产推荐：先迁移后切换流量。

### DSN-1004 版本元数据

在构建时注入：

- `APP_VERSION`
- `GIT_SHA`
- `BUILD_TIME`

并在 `/healthz` 或 `/version` 返回。

### DSN-1005 回滚方案

应用回滚：
- 切回上一个稳定镜像 tag

数据库回滚：
- 非破坏迁移可使用 goose down
- 破坏性迁移需预先制定“向前修复”策略（优先）

### DSN-1006 发布验证

发布后自动执行：

- `/healthz`
- `/readyz`
- smoke 链路（workspace/session/agent）

失败则标记发布失败并触发回滚步骤。

## Requirement Coverage

- RQ-1001 -> DSN-1001
- RQ-1002 -> DSN-1002
- RQ-1003 -> DSN-1003
- RQ-1004 -> DSN-1004, DSN-1005
- RQ-1005 -> DSN-1006
- RQ-1006 -> DSN-1004
- RQ-1007 -> DSN-1001
- RQ-1008 -> DSN-1001, DSN-1006
