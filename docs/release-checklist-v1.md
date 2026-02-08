# Release Checklist V1

## 1) API Contract

- [x] Agent API (`/agent/run`) request/response aligned with docs/contracts
- [x] Flow APIs (`/api/v1/flows*`) request/response aligned
- [x] Debugger APIs (`/api/v1/debug/runs*`) request/response aligned
- [x] Marketplace APIs (`/api/v1/marketplace/*`) request/response aligned
- [x] Non-backward-compatible changes identified and mitigated

## 2) Database Migration

- [x] New migrations applied successfully in clean DB
- [x] Rollback path tested (down)
- [x] Re-apply after rollback works
- [x] Core tables data integrity verified
- [x] Migration report written to `docs/release-artifacts/migration-report.md`

## 3) Regression Matrix

- [x] `go test ./...` pass
- [x] `npm --prefix web run test -- --run` pass
- [x] Agent run basic flow pass
- [x] Multimodal flow (image/audio ref) pass
- [x] Flow orchestration create/bind/run pass
- [x] Debugger list/detail/cost pass
- [x] Marketplace create/review/rate/apply pass
- [x] Benchmark report generation pass
- [x] Matrix result written to `docs/release-artifacts/regression-matrix-v1.md`

## 4) Observability

- [x] Trace ID appears in request chain
- [x] Key events present in run timeline
- [x] Error path has actionable logs
- [x] Runbook instructions are reproducible

## 5) Rollback

- [x] Rollback trigger condition defined
- [x] Data rollback strategy documented
- [x] Service rollback steps documented
- [x] Owner and communication plan documented

## 6) Release Decision

- [x] Open risks assessed (high/medium/low)
- [x] Blockers resolved or explicitly accepted
- [x] Final decision recorded in `docs/release-artifacts/release-gate-v1.md`
