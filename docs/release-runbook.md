# Release Runbook

## Build

```bash
docker build -t a2ui-backend:local .
```

## Migrate

```bash
go run ./cmd/cli db:create
go run ./cmd/cli db:migrate
```

## Verify

- `GET /healthz`
- `GET /readyz`
- smoke 脚本通过

## Rollback

1. 回滚到上一个稳定镜像 tag
2. 如有破坏性迁移，优先执行向前修复策略
