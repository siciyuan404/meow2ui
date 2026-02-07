# A2UI Step16 Data Lifecycle, Backup & Recovery Tasks

## TASK-1601 增加备份元数据迁移表

- Linked Requirements: RQ-1602, RQ-1605
- Linked Design: DSN-1602
- Description:
  - 新增 `backup_jobs` 表及索引
- DoD:
  - migration 可 up/down
- Verify:
  - `go run ./cmd/cli db:migrate`

Status: completed

## TASK-1602 实现全量备份服务

- Linked Requirements: RQ-1602, RQ-1606
- Linked Design: DSN-1602
- Description:
  - 实现 PostgreSQL 全量备份执行与元数据记录
- DoD:
  - 可生成备份产物并记录 checksum/size
- Verify:
  - `go test ./pkg/backup/...`

Status: completed

## TASK-1603 实现恢复服务与校验

- Linked Requirements: RQ-1603, RQ-1605
- Linked Design: DSN-1603
- Description:
  - 实现按 backup-id 恢复
  - 恢复后执行一致性与链路校验
- DoD:
  - 恢复后主链路可用
- Verify:
  - `go test ./pkg/recovery/...`

Status: completed

## TASK-1604 实现导出导入数据包

- Linked Requirements: RQ-1604
- Linked Design: DSN-1604
- Description:
  - 实现 workspace 级导出/导入
  - 保留 session/version 关系
- DoD:
  - 导入后数据完整可用
- Verify:
  - `go test ./pkg/data-transfer/...`

Status: completed

## TASK-1605 增加 ops API 与 CLI 命令

- Linked Requirements: RQ-1602, RQ-1603, RQ-1604
- Linked Design: DSN-1606
- Description:
  - 增加 backup/recovery/data-transfer API 与 CLI
- DoD:
  - 可通过 API/CLI 触发备份恢复导入导出
- Verify:
  - `go test ./cmd/server/...`
  - `go run ./cmd/cli --help`

Status: completed

## TASK-1606 接入调度与告警

- Linked Requirements: RQ-1602, RQ-1605
- Linked Design: DSN-1605
- Description:
  - 备份调度任务
  - 失败告警与恢复事件记录
- DoD:
  - 连续失败可触发告警
- Verify:
  - `go test ./pkg/observability/alerting/...`

Status: completed

## TASK-1607 数据保留与清理策略实现

- Linked Requirements: RQ-1601
- Linked Design: DSN-1601
- Description:
  - 按 tier 实现归档/清理策略
- DoD:
  - 可配置保留天数并执行清理
- Verify:
  - `go test ./pkg/...`（retention 相关）

Status: completed

## TASK-1608 演练手册与回归测试

- Linked Requirements: RQ-1605, RQ-1608
- Linked Design: DSN-1603, DSN-1604
- Description:
  - 新增 `docs/backup-recovery.md`
  - 增加备份恢复演练脚本
- DoD:
  - 文档可指导一次完整演练
  - 全量测试通过
- Verify:
  - `go test ./...`

Status: completed
