# CI Guide

## Required Checks

- `unit`
- `integration-postgres`
- `smoke`

## Local Verify

```bash
go test ./... -coverprofile=coverage.out
bash scripts/check-coverage.sh 40
bash scripts/integration-postgres.sh
bash scripts/smoke.sh
```
