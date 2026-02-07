# A2UI Step6 Quality CI Gates Design

## Design Goals

- 将“本地可运行”升级为“持续可验证”。
- 确保关键功能在 CI 中自动回归。
- 失败日志可直接映射到具体任务/模块。

## Architecture

### DSN-601 GitHub Actions 工作流

新增 `.github/workflows/ci.yml`，包含：

- `unit` job
  - `go mod tidy` 校验
  - `go test ./... -coverprofile=coverage.out`
  - 输出覆盖率摘要

- `integration-postgres` job
  - 启动 postgres service
  - 执行 `go run ./cmd/cli db:create`
  - 执行 `go run ./cmd/cli db:migrate`
  - 执行最小 CLI 链路（workspace/session/agent）

- `smoke` job
  - 执行 `scripts/smoke.sh`

### DSN-602 覆盖率门槛

新增 `scripts/check-coverage.sh`：

- 读取 `coverage.out`
- 计算总覆盖率
- 低于阈值（默认 40）返回非零

### DSN-603 关键链路集成脚本

新增 `scripts/integration-postgres.sh`：

- 设置 `STORE_DRIVER=postgres`
- 跑 `db:create`、`db:migrate`
- 创建 workspace/session
- 执行 `agent:run`

### DSN-604 PR 保护建议

仓库设置（文档化）：

- Required checks:
  - `unit`
  - `integration-postgres`
  - `smoke`

### DSN-605 日志可读性

脚本输出统一前缀：

- `[CI][UNIT]`
- `[CI][INT]`
- `[CI][SMOKE]`

失败时打印最后命令和退出码。

## Requirement Coverage

- RQ-601 -> DSN-601, DSN-603
- RQ-602 -> DSN-602
- RQ-603 -> DSN-601, DSN-603
- RQ-604 -> DSN-604
- RQ-605 -> DSN-601
- RQ-606 -> DSN-605
