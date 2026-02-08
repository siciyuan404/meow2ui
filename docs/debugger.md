# Debugger Guide

## Overview

Debugger APIs provide run-level diagnostics for agent executions.

## Endpoints

- `GET /api/v1/debug/runs`
- `GET /api/v1/debug/runs/{runId}`
- `GET /api/v1/debug/runs/{runId}/cost`

## Query Parameters

- `session_id`
- `status`
- `from` (RFC3339)
- `to` (RFC3339)

## Response Highlights

- Run summary with duration
- Step timeline with token and latency
- Tool-call chain (when available)
- Context token split
- Cost estimation by provider/model

## Security

Sensitive fields are redacted in debugger payloads:

- `api_key`
- `token`
- `secret`
- `authorization`
