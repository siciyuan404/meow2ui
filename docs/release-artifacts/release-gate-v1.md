# Release Gate V1

- Candidate Version: `v1-release-candidate`
- Date: `2026-02-08`
- Owner: `engineering`

## Gate Summary

- API Contract: `pass`
- Migration Validation: `pass`
- Regression Matrix: `pass`
- Observability Validation: `pass`
- Rollback Preparedness: `pass`

Overall: `pass`

## Open Risks

| Risk ID | Level | Description | Mitigation | Owner | Status |
|---|---|---|---|---|---|
| RISK-001 | low | rollback 自动化已接入，需观察稳定性 | 保持 CI 监控并保留日志产物（tracking: #7） | engineering | mitigated |

## Rollback Plan

1. Trigger condition:
   - 发布后关键 API 持续 5xx 或迁移后数据不可读。
2. Technical rollback:
   - 回退到前一稳定镜像并重启服务。
3. Data rollback:
   - 按 `00014 -> 00013 -> 00012` 顺序执行 down SQL。
4. Communication:
   - 在工程频道和 issue 线程同步回滚状态与恢复 ETA。

## Verification Links

- Checklist: `docs/release-checklist-v1.md`
- Migration Report: `docs/release-artifacts/migration-report.md`
- Regression Matrix: `docs/release-artifacts/regression-matrix-v1.md`
- Runbook: `docs/runbook.md`

## Sign-off

- Engineering: `done / 2026-02-08`
- QA: `done / 2026-02-08`
- Product: `pending`

Decision: `ready_for_release_candidate`
