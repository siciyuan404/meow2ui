# A2UI Step16 Data Lifecycle, Backup & Recovery Design

## Design Goals

- 为生产环境提供可执行的数据保护机制。
- 将备份恢复流程标准化、可自动化、可演练。
- 与现有 Postgres、ops、alerting 模块协同。

## Architecture

### DSN-1601 数据分级策略

按表分级：

- Tier-1（核心）：`workspaces`, `sessions`, `schema_versions`, `playground_items`
- Tier-2（审计）：`agent_runs`, `agent_events`, `product_events`
- Tier-3（临时）：缓存/临时中间表

保留策略：
- Tier-1 长期保留
- Tier-2 保留 180 天（可配置）
- Tier-3 保留 7-30 天（可配置）

### DSN-1602 备份模块

新增 `pkg/backup`：

- `BackupService`
  - `RunFullBackup()`
  - `RunIncrementalBackup()`（可选）
  - `ListBackups()`

新增备份元数据表：
- `backup_jobs`
  - `id`, `type`, `status`, `started_at`, `ended_at`, `artifact_uri`, `size_bytes`, `checksum`, `error`

### DSN-1603 恢复模块

新增 `pkg/recovery`：

- `RecoveryService`
  - `Restore(backupID)`
  - `ValidatePostRestore()`

恢复后校验：
- 关键表计数
- `workspace -> session -> agent` 最小链路 smoke

### DSN-1604 导入导出模块

新增 `pkg/data-transfer`：

- `ExportWorkspace(workspaceID)`
- `ImportWorkspace(bundle)`

数据包格式：
- `manifest.json`
- `workspaces.json`
- `sessions.json`
- `schema_versions.json`
- `playground.json`

### DSN-1605 调度与告警

- 使用 cron/定时任务触发备份
- 失败事件接入 alerting（Step15）
- 备份失败阈值告警：连续失败 >= 2 次

### DSN-1606 API 与 CLI

新增接口：
- `POST /api/v1/ops/backup/run`
- `GET /api/v1/ops/backup/jobs`
- `POST /api/v1/ops/recovery/restore`
- `POST /api/v1/data/export`
- `POST /api/v1/data/import`

CLI 增强：
- `db:backup`
- `db:restore <backup-id>`
- `data:export <workspace-id>`
- `data:import <bundle-path>`

## Requirement Coverage

- RQ-1601 -> DSN-1601
- RQ-1602 -> DSN-1602, DSN-1605
- RQ-1603 -> DSN-1603
- RQ-1604 -> DSN-1604, DSN-1606
- RQ-1605 -> DSN-1603, DSN-1605
- RQ-1606 -> DSN-1602, DSN-1603
- RQ-1607 -> DSN-1602
- RQ-1608 -> DSN-1602, DSN-1603, DSN-1604
