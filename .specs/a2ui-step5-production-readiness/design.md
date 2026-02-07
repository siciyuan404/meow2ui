# A2UI Step5 Production Readiness Design

## Design Goals

- 建立“可部署、可观测、可定位问题”的基础运行能力。
- 降低线上故障时的排障时间（MTTR）。
- 与现有 server/cli 架构兼容，保持最小侵入。

## Architecture

### DSN-501 配置校验器

新增 `internal/infra/config`：

- `AppConfig`
  - `StoreDriver`
  - `PostgresConfig`
  - `ServerAddr`
  - `AutoMigrate`

- `Validate() error`
  - 校验 `STORE_DRIVER` 枚举
  - postgres 模式强校验 `PG_*`

### DSN-502 健康检查扩展

在 `cmd/server` 增加：

- `/healthz`：应用级（进程存活、基础组件已初始化）
- `/readyz`：依赖级（postgres ping、migrations version 可读）

返回结构：

```json
{
  "ok": true,
  "checks": {
    "app": "up",
    "db": "up"
  }
}
```

### DSN-503 API 错误响应中间件

新增 `pkg/httpx`：

- `APIError`
  - `Code`
  - `Message`
  - `Detail`
  - `TraceID`

- `WriteError(w, err)`
  - 识别 `store.ErrNotFound` -> `NOT_FOUND`
  - 识别校验错误 -> `VALIDATION_FAILED`
  - 未知错误 -> `INTERNAL_ERROR`

### DSN-504 Trace ID

新增 request middleware：

- 优先读取 `X-Trace-Id`
- 否则生成新 trace id
- 写入响应头并贯穿日志/错误响应

### DSN-505 Smoke 检查脚本

新增脚本（可在 CI 执行）：

- `go test ./...`
- `go run ./cmd/cli db:create`
- `go run ./cmd/cli db:migrate`
- 启动 server 后调用 `/healthz` 与 `/readyz`

### DSN-506 文档更新

更新 `docs/`：

- `runbook.md`：启动、迁移、排障
- `api-errors.md`：错误码说明

## Requirement Coverage

- RQ-501 -> DSN-501
- RQ-502 -> DSN-502
- RQ-503 -> DSN-503, DSN-504
- RQ-504 -> DSN-505
- RQ-505 -> DSN-504
- RQ-506 -> DSN-502, DSN-503, DSN-505
