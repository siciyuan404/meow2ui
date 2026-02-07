# Backup & Recovery

Commands:
- `go run ./cmd/cli db:backup`
- `go run ./cmd/cli db:restore <backup-id>`
- `go run ./cmd/cli data:export <workspace-id>`
- `go run ./cmd/cli data:import <bundle-path>`

MVP targets:
- RPO <= 24h
- RTO <= 2h
