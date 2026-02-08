# Flow Orchestration Guide

## Overview

Flow orchestration enables configurable execution paths for agent runs.

## Endpoints

- `GET /api/v1/flows`
- `POST /api/v1/flows`
- `POST /api/v1/flows/bind-session`

## Flow Definition

Key fields:

- `name`
- `policy.parallelism`
- `policy.failure_mode` (`fail_fast` or `continue_on_error`)
- `nodes[]` (`id`, `type`, `depends_on`)
- `edges[]` (`from`, `to`, `condition`)

Supported node types:

- `plan`
- `emit`
- `validate`
- `repair`
- `apply`
- `custom`

Supported conditions:

- `always`
- `validate.ok`
- `!validate.ok`
- `success`
- `failed`

## Session Binding

Bind flow template/version to a session using:

- `sessionId`
- `templateId`
- `version`

When no binding exists, runtime falls back to default flow.
