# Release Notes V1 RC

## Highlights

- 完成中优先级与低优先级队列全部交付。
- 新增 Flow 编排、Debugger、多模态输入、公开 Benchmark、模板市场生态能力。
- 发布门禁完成，当前状态 `ready_for_release_candidate`。

## Key Changes

### Runtime & APIs

- Flow API: `/api/v1/flows`, `/api/v1/flows/bind-session`
- Debugger API: `/api/v1/debug/runs*`
- Marketplace API: `/api/v1/marketplace/*`
- Agent run 支持 `media` 输入参数（image/audio 引用）

### Data & Migration

- 新增迁移：
  - `00012_flow_orchestration.sql`
  - `00013_multimodal_assets.sql`
  - `00014_benchmark_public_targets.sql`
- CLI 新增回滚命令：`db:rollback [steps]`

### Documentation

- Release checklist, contracts, migration report, regression matrix, release gate
- `docs/multimodal.md`
- `docs/benchmark-public.md`
- `docs/marketplace-ecosystem.md`

## Validation

- `go test ./...` pass
- `npm --prefix web run test -- --run` pass
- benchmark 报告脚本执行通过并产出文档

## Known Risk

- 需在 CI Postgres 环境补一次 `db:rollback` 自动化验证（风险等级：low）
