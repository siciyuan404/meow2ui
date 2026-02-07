#!/usr/bin/env bash
set -euo pipefail

echo "[CI][SMOKE] start"

go test ./...

go run ./cmd/cli db:create
go run ./cmd/cli db:migrate

export STORE_DRIVER=postgres
WS_OUT=$(go run ./cmd/cli workspace:create smoke "/tmp/a2ui-smoke")
WS_ID=${WS_OUT#workspace=}

SSN_OUT=$(go run ./cmd/cli session:create "$WS_ID" "smoke-session")
SSN_ID=$(echo "$SSN_OUT" | awk '{print $1}' | sed 's/session=//')

go run ./cmd/cli agent:run "$SSN_ID" "smoke test ui"

echo "[CI][SMOKE] ok"
