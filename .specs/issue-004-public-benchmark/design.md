# Issue #4 公开 Benchmark 与性能对比 Design

## Design Goals

- 复用 Step14 评估能力，增加“公开对比 + 定时发布”能力。
- 保持执行链可追踪、可复跑、可回归判断。
- 报告产物同时面向机器与人类阅读。

## Architecture

### DSN-0401 Benchmark Suite 扩展

- 在现有 benchmark 数据模型上补充：
  - `benchmark_targets`（对比对象配置）
  - `benchmark_env_snapshots`（执行环境快照）
- 记录 git sha、数据集版本、执行参数。

### DSN-0402 执行编排器

- 新增 `pkg/evaluation/benchmark/orchestrator`：
  - 按 target 顺序执行
  - 失败 target 隔离
  - 聚合 run 结果

### DSN-0403 回归引擎

- 复用 Step14 阈值比较逻辑，扩展多维指标判定。
- 输出结构：`status`, `regressed_metrics`, `summary`。

### DSN-0404 CI 集成

- 在 `.github/workflows/ci.yml` 增加定时/手动 benchmark job：
  - 拉取基准数据
  - 执行 benchmark
  - 上传 artifacts

### DSN-0405 报告发布

- 新增 `scripts/publish-benchmark-report.sh`：
  - 生成 `docs/benchmarks/*.md` 与 `*.json`
  - 发布到 GitHub Pages 分支或目录

### DSN-0406 成本与预算控制

- 执行器支持 `max_cases`, `max_tokens`, `max_cost` 配置。
- 超预算自动降采样并在报告中标注。

## Requirement Coverage

- RQ-0401 -> DSN-0401, DSN-0402
- RQ-0402 -> DSN-0401, DSN-0402
- RQ-0403 -> DSN-0403
- RQ-0404 -> DSN-0404
- RQ-0405 -> DSN-0405
- RQ-0406 -> DSN-0401
- RQ-0407 -> DSN-0406
- RQ-0408 -> DSN-0402, DSN-0403, DSN-0405
