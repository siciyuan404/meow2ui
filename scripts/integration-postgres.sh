#!/usr/bin/env bash
set -euo pipefail

mkdir -p artifacts
LOG_FILE="artifacts/integration-postgres.log"
exec > >(tee "$LOG_FILE") 2>&1

echo "[CI][INT] start"

export STORE_DRIVER=postgres
export PG_HOST=${PG_HOST:-localhost}
export PG_PORT=${PG_PORT:-5432}
export PG_USER=${PG_USER:-postgres}
export PG_PASSWORD=${PG_PASSWORD:-postgres}
export PG_DATABASE=${PG_DATABASE:-a2ui_platform}
export PG_SSLMODE=${PG_SSLMODE:-disable}

go run ./cmd/cli db:create
go run ./cmd/cli db:migrate
go run ./cmd/cli db:rollback 1
go run ./cmd/cli db:migrate

WS_OUT=$(go run ./cmd/cli workspace:create ci "/tmp/a2ui-ci")
WS_ID=${WS_OUT#workspace=}
SSN_OUT=$(go run ./cmd/cli session:create "$WS_ID" "ci-session")
SSN_ID=$(echo "$SSN_OUT" | awk '{print $1}' | sed 's/session=//')
go run ./cmd/cli agent:run "$SSN_ID" "ci integration"

echo "[CI][INT] ok"
