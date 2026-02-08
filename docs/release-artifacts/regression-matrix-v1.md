# Regression Matrix V1

- Date: 2026-02-08

## Automated Commands

| Item | Command | Result |
|---|---|---|
| Go tests | `go test ./...` | pass |
| Web tests | `npm --prefix web run test -- --run` | pass |
| Benchmark report script | `bash scripts/publish-benchmark-report.sh benchmark-run-local` | pass |

## Core Workflow Checks

| Workflow | Validation Target | Result |
|---|---|---|
| Agent run | `/agent/run` basic request | pass |
| Multimodal run | `/agent/run` with `media[]` | pass |
| Flow orchestration | `/api/v1/flows*` create/list/bind | pass |
| Debugger | `/api/v1/debug/runs*` list/detail/cost | pass |
| Marketplace | `/api/v1/marketplace/*` create/review/rate/apply | pass |

## Notes

- Regression baseline currently validated via tests and contract examples.
- Recommend adding smoke HTTP checks in CI for all new API paths.
