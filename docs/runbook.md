# Runbook

## Start (memory)

```bash
go run ./cmd/server
```

## Start (postgres)

```bash
export STORE_DRIVER=postgres
export PG_HOST=localhost
export PG_PORT=5432
export PG_USER=postgres
export PG_PASSWORD=postgres
export PG_DATABASE=a2ui_platform
export PG_SSLMODE=disable

go run ./cmd/cli db:create
go run ./cmd/cli db:migrate
go run ./cmd/server
```

## Health checks

- `GET /healthz`
- `GET /readyz`
- `GET /version`

## Common issues

- `config invalid`: 检查 `STORE_DRIVER` 和 `PG_*` 环境变量
- `migration failed`: 先执行 `db:create` 再执行 `db:migrate`
- `NOT_FOUND`: 检查 workspace/session ID 是否存在
