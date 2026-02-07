# A2UI Step14 Evaluation & Quality Benchmark Design

## Design Goals

- 让 A2UI 质量从“主观感觉”变成“可度量指标”。
- 支持版本回归检测，降低模型替换风险。
- 与现有 Agent、analytics、CI 流程无缝衔接。

## Architecture

### DSN-1401 评估模型

新增模块：`pkg/evaluation`

- `Evaluator`
  - 输入：生成结果、校验结果、运行元数据
  - 输出：`EvalScore`

- `EvalScore`
  - `schema_valid`
  - `component_valid_rate`
  - `prop_valid_rate`
  - `repair_count`
  - `success`
  - `failure_type`

### DSN-1402 Benchmark 数据集

新增数据表：

- `benchmark_suites`
  - `id`, `name`, `version`, `created_at`
- `benchmark_cases`
  - `id`, `suite_id`, `prompt`, `expected_constraints(json)`, `tags(json)`

### DSN-1403 执行与报告

新增执行器：`pkg/evaluation/runner`

- 批量运行 case
- 记录每个 case 的 run 结果
- 汇总报告：通过率、平均耗时、失败分布

新增结果表：

- `benchmark_runs`
  - `id`, `suite_id`, `model_id`, `prompt_profile`, `started_at`, `ended_at`
- `benchmark_run_results`
  - `run_id`, `case_id`, `score(json)`, `pass`, `latency_ms`, `tokens`

### DSN-1404 回归判定

- 与 baseline run 对比
- 规则示例：
  - success rate 下降 > 3% 判为回归
  - schema_valid 下降 > 1% 判为回归

### DSN-1405 CI 集成

- 新增脚本 `scripts/benchmark-check.sh`
- 在 CI 可选 job 中执行
- 回归时返回非零退出码

### DSN-1406 API

- `POST /api/v1/evaluation/benchmark-runs`
- `GET /api/v1/evaluation/benchmark-runs/{id}`
- `GET /api/v1/evaluation/benchmark-runs/{id}/results`

## Requirement Coverage

- RQ-1401 -> DSN-1401
- RQ-1402 -> DSN-1402
- RQ-1403 -> DSN-1404, DSN-1405
- RQ-1404 -> DSN-1403, DSN-1406
- RQ-1405 -> DSN-1401
- RQ-1406 -> DSN-1401, DSN-1404
