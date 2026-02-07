# PostgreSQL Setup

## Prerequisites

- PostgreSQL running on `localhost:5432`
- Credentials:
  - user: `postgres`
  - password: `postgres`

## Environment Variables

```bash
export STORE_DRIVER=postgres
export PG_HOST=localhost
export PG_PORT=5432
export PG_USER=postgres
export PG_PASSWORD=postgres
export PG_DATABASE=a2ui_platform
export PG_SSLMODE=disable
```

## Initialize Database

```bash
go run ./cmd/cli db:create
go run ./cmd/cli db:migrate
```

## Run CLI Flow

```bash
STORE_DRIVER=postgres go run ./cmd/cli workspace:create demo "/tmp/a2ui"
STORE_DRIVER=postgres go run ./cmd/cli session:create <workspace-id> "demo-session"
STORE_DRIVER=postgres go run ./cmd/cli agent:run <session-id> "生成一个仪表盘标题"
```

## Run Server

```bash
STORE_DRIVER=postgres AUTO_MIGRATE=true go run ./cmd/server
```
