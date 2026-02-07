#!/usr/bin/env bash
set -euo pipefail

echo "[CI][RELEASE] verify start"

curl -fsS "${BASE_URL:-http://localhost:8080}/healthz" >/dev/null
curl -fsS "${BASE_URL:-http://localhost:8080}/readyz" >/dev/null

echo "[CI][RELEASE] verify ok"
