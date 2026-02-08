# Issue #3 CONTRIBUTING 与社区建设 Design

## Design Goals

- 让首次贡献者在最短路径内完成环境搭建、代码修改和 PR 提交。
- 让仓库协作流程标准化，减少 reviewer 的沟通成本。
- 通过 ADR 机制沉淀关键决策，降低架构认知门槛。

## Architecture

### DSN-0301 贡献入口文档

- 新增仓库根文档 `CONTRIBUTING.md`，内容包含：
  - 开发环境前置条件（Go/Node 版本、依赖、数据库）
  - 本地启动与测试命令
  - 分支、提交信息、PR 规范
  - 代码风格与评审预期

### DSN-0302 模板与流程标准化

- 在 `.github/` 增补或完善模板：
  - `PULL_REQUEST_TEMPLATE.md`
  - `ISSUE_TEMPLATE/bug_report.md`
  - `ISSUE_TEMPLATE/feature_request.md`
- 模板字段统一要求：背景、方案、风险、验证结果。

### DSN-0303 开发验证基线

- 在 `CONTRIBUTING.md` 明确最小本地验证集：
  - `go test ./...`
  - `npm run test`（如涉及 Web）
  - `npm run lint`（如涉及前端）
- 与现有 `docs/ci.md`、`.github/workflows/ci.yml` 保持一致。

### DSN-0304 ADR 目录规范

- 新增 `docs/adr/` 目录与 `docs/adr/README.md`。
- 新增 ADR 模板：`docs/adr/0000-template.md`。
- 约定命名：`NNNN-short-title.md`，例如 `0001-provider-abstraction.md`。

### DSN-0305 社区入口与行为准则

- 在 `CONTRIBUTING.md` 与 `README`（如已有相关章节）补充：
  - GitHub Discussions 入口
  - 实时沟通渠道占位（如 Discord/Slack）
  - 行为准则（Code of Conduct）链接或后续补齐说明

## Requirement Coverage

- RQ-0301 -> DSN-0301
- RQ-0302 -> DSN-0302
- RQ-0303 -> DSN-0301, DSN-0303
- RQ-0304 -> DSN-0304
- RQ-0305 -> DSN-0305
- RQ-0306 -> DSN-0301, DSN-0303
- RQ-0307 -> DSN-0301, DSN-0302, DSN-0303
