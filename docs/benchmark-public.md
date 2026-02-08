# Public Benchmark Guide

## Overview

Public benchmark runs are executed in CI and exported as artifacts and docs reports.

## CI Job

- Workflow job: `benchmark-public`
- Triggers: push, pull_request, workflow_dispatch, schedule

## Artifacts

- `docs/benchmarks/<run-id>.md`
- `docs/benchmarks/<run-id>.json`

## Local Run

```bash
go test ./pkg/evaluation/...
bash scripts/publish-benchmark-report.sh benchmark-run-local
```

## Budget Controls

Orchestrator supports:

- `max_cases`
- `max_tokens`
- `max_cost`

When budget is exceeded, execution samples are reduced and report summary contains budget notes.
