# Issue #4 公开 Benchmark 与性能对比 Tasks

## TASK-0401 扩展 benchmark 数据模型与迁移

- Linked Requirements: RQ-0401, RQ-0402, RQ-0406
- Linked Design: DSN-0401
- Description:
  - 新增 benchmark_targets 与环境快照表
  - 增加 repository 访问接口
- DoD:
  - 可持久化 target 与环境信息
- Verify:
  - `go run ./cmd/cli db:migrate`
  - `go test ./internal/infra/sqlstore/...`

Status: completed

## TASK-0402 实现 benchmark orchestrator

- Linked Requirements: RQ-0401, RQ-0402
- Linked Design: DSN-0402
- Description:
  - 编排多 target 执行，隔离失败
- DoD:
  - 单次 run 可输出多对象聚合结果
- Verify:
  - `go test ./pkg/evaluation/...`

Status: completed

## TASK-0403 扩展回归判定引擎

- Linked Requirements: RQ-0403
- Linked Design: DSN-0403
- Description:
  - 支持多维指标阈值判断
- DoD:
  - 可输出回归指标列表与摘要
- Verify:
  - `go test ./pkg/evaluation/...`

Status: completed

## TASK-0404 增加 CI benchmark 定时任务

- Linked Requirements: RQ-0404
- Linked Design: DSN-0404
- Description:
  - 在 CI 增加 schedule/workflow_dispatch benchmark job
- DoD:
  - CI 可执行 benchmark 并上传 artifacts
- Verify:
  - GitHub Actions 手动触发通过

Status: completed

## TASK-0405 实现公开报告生成与发布脚本

- Linked Requirements: RQ-0405
- Linked Design: DSN-0405
- Description:
  - 生成 Markdown/JSON 报告并发布
- DoD:
  - 报告可被公开访问
- Verify:
  - `bash scripts/publish-benchmark-report.sh`

Status: completed

## TASK-0406 接入预算与降采样控制

- Linked Requirements: RQ-0407
- Linked Design: DSN-0406
- Description:
  - 增加 max_cases/max_tokens/max_cost 限制
- DoD:
  - 超预算会降采样并记录
- Verify:
  - `go test ./pkg/evaluation/...`

Status: completed

## TASK-0407 补充测试与文档

- Linked Requirements: RQ-0408
- Linked Design: DSN-0402, DSN-0403, DSN-0405
- Description:
  - 增加 benchmark 回归测试
  - 新增 `docs/benchmark-public.md`
- DoD:
  - 团队可按文档执行并查看公开报告
- Verify:
  - `go test ./...`

Status: completed
