# A2UI Step14 Evaluation & Quality Benchmark Tasks

## TASK-1401 增加 benchmark 数据表迁移

- Linked Requirements: RQ-1402, RQ-1404
- Linked Design: DSN-1402, DSN-1403
- Description:
  - 新增 benchmark_suites / benchmark_cases / benchmark_runs / benchmark_run_results
- DoD:
  - migration 可 up/down
- Verify:
  - `go run ./cmd/cli db:migrate`

Status: completed

## TASK-1402 实现评估器核心指标

- Linked Requirements: RQ-1401
- Linked Design: DSN-1401
- Description:
  - 实现 schema/component/props/repair 指标计算
- DoD:
  - 可输出标准 EvalScore
- Verify:
  - `go test ./pkg/evaluation/...`

Status: completed

## TASK-1403 实现 benchmark runner

- Linked Requirements: RQ-1402, RQ-1404
- Linked Design: DSN-1403
- Description:
  - 批量执行 benchmark case 并写入结果
- DoD:
  - 可生成完整 run 报告
- Verify:
  - `go test ./pkg/evaluation/runner/...`

Status: completed

## TASK-1404 实现回归判定规则

- Linked Requirements: RQ-1403
- Linked Design: DSN-1404
- Description:
  - 基于 baseline run 做阈值比较
  - 输出 pass/fail 与回归原因
- DoD:
  - 指标下降可被正确识别
- Verify:
  - `go test ./pkg/evaluation/...`

Status: completed

## TASK-1405 增加评估 API

- Linked Requirements: RQ-1404
- Linked Design: DSN-1406
- Description:
  - 实现 benchmark run 创建与结果查询 API
- DoD:
  - API 可查询 run 与 case 级结果
- Verify:
  - `go test ./cmd/server/...`

Status: completed

## TASK-1406 接入 CI 质量门禁

- Linked Requirements: RQ-1403
- Linked Design: DSN-1405
- Description:
  - 新增 benchmark-check 脚本与 CI job
- DoD:
  - 回归时 CI 失败
- Verify:
  - 在 CI 中执行并验证退出码

Status: completed

## TASK-1407 补充评估文档与案例

- Linked Requirements: RQ-1402, RQ-1404
- Linked Design: DSN-1402, DSN-1403
- Description:
  - 新增 `docs/evaluation.md`
  - 提供 baseline 构建与对比示例
- DoD:
  - 团队可按文档运行 benchmark
- Verify:
  - 人工按文档执行通过

Status: completed
