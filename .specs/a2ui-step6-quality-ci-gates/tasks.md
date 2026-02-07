# A2UI Step6 Quality CI Gates Tasks

## TASK-601 建立 CI 主工作流

- Linked Requirements: RQ-601, RQ-603
- Linked Design: DSN-601
- Description:
  - 新增 `.github/workflows/ci.yml`
  - 包含 unit / integration-postgres / smoke 三个 job
- DoD:
  - PR 触发后自动运行所有 job
- Verify:
  - 在分支触发一次 CI 并查看结果

Status: completed

## TASK-602 增加覆盖率检查脚本

- Linked Requirements: RQ-602
- Linked Design: DSN-602
- Description:
  - 新增 `scripts/check-coverage.sh`
  - 在 unit job 中接入阈值校验
- DoD:
  - 覆盖率低于阈值时 job 失败
- Verify:
  - 本地执行脚本并观察退出码

Status: completed

## TASK-603 增加 postgres 集成脚本

- Linked Requirements: RQ-601, RQ-603
- Linked Design: DSN-603
- Description:
  - 新增 `scripts/integration-postgres.sh`
  - 跑通 db/create+migrate+agent 链路
- DoD:
  - 脚本可在 CI 环境稳定执行
- Verify:
  - 本地或 CI 执行脚本通过

Status: completed

## TASK-604 文档化 PR 门禁配置

- Linked Requirements: RQ-604
- Linked Design: DSN-604
- Description:
  - 更新贡献文档，说明 required checks 设置方式
- DoD:
  - 文档可指导仓库管理员完成门禁配置
- Verify:
  - 人工审阅通过

Status: completed

## TASK-605 统一 CI 日志输出规范

- Linked Requirements: RQ-606
- Linked Design: DSN-605
- Description:
  - 调整脚本输出前缀与失败信息格式
- DoD:
  - CI 日志可快速定位失败步骤
- Verify:
  - 查看一次失败/成功日志样例

Status: completed
