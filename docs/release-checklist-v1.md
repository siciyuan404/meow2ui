# Release Checklist V1

## 1) API Contract

- [ ] Agent API (`/agent/run`) request/response aligned with docs/contracts
- [ ] Flow APIs (`/api/v1/flows*`) request/response aligned
- [ ] Debugger APIs (`/api/v1/debug/runs*`) request/response aligned
- [ ] Marketplace APIs (`/api/v1/marketplace/*`) request/response aligned
- [ ] Non-backward-compatible changes identified and mitigated

## 2) Database Migration

- [ ] New migrations applied successfully in clean DB
- [ ] Rollback path tested (down)
- [ ] Re-apply after rollback works
- [ ] Core tables data integrity verified
- [ ] Migration report written to `docs/release-artifacts/migration-report.md`

## 3) Regression Matrix

- [ ] `go test ./...` pass
- [ ] `npm --prefix web run test -- --run` pass
- [ ] Agent run basic flow pass
- [ ] Multimodal flow (image/audio ref) pass
- [ ] Flow orchestration create/bind/run pass
- [ ] Debugger list/detail/cost pass
- [ ] Marketplace create/review/rate/apply pass
- [ ] Benchmark report generation pass
- [ ] Matrix result written to `docs/release-artifacts/regression-matrix-v1.md`

## 4) Observability

- [ ] Trace ID appears in request chain
- [ ] Key events present in run timeline
- [ ] Error path has actionable logs
- [ ] Runbook instructions are reproducible

## 5) Rollback

- [ ] Rollback trigger condition defined
- [ ] Data rollback strategy documented
- [ ] Service rollback steps documented
- [ ] Owner and communication plan documented

## 6) Release Decision

- [ ] Open risks assessed (high/medium/low)
- [ ] Blockers resolved or explicitly accepted
- [ ] Final decision recorded in `docs/release-artifacts/release-gate-v1.md`
