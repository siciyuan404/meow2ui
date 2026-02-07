#!/usr/bin/env bash
set -euo pipefail

THRESHOLD=${1:-40}
if [ ! -f coverage.out ]; then
  echo "[CI][UNIT] coverage.out not found"
  exit 1
fi

TOTAL=$(go tool cover -func=coverage.out | grep total: | awk '{print $3}' | sed 's/%//')
echo "[CI][UNIT] total coverage=${TOTAL}% threshold=${THRESHOLD}%"

TOTAL_INT=${TOTAL%.*}
if [ "$TOTAL_INT" -lt "$THRESHOLD" ]; then
  echo "[CI][UNIT] coverage below threshold"
  exit 1
fi

echo "[CI][UNIT] coverage ok"
