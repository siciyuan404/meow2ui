# A2UI Step18 Enterprise Readiness & Compliance Baseline Design

## Design Goals

- 为企业场景提供最小可行治理能力。
- 在现有 auth、security、ops 基础上叠加组织级控制。
- 保持向后兼容，避免破坏当前单组织使用路径。

## Architecture

### DSN-1801 组织与项目数据模型

新增表：

- `orgs`
  - `id`, `name`, `status`, `created_at`
- `projects`
  - `id`, `org_id`, `name`, `status`, `created_at`
- `project_members`
  - `id`, `project_id`, `user_id`, `role`, `created_at`

资源绑定扩展：
- `workspaces.project_id`
- `sessions.project_id`

### DSN-1802 RBAC 引擎

新增模块 `pkg/rbac`：

- `Policy`
  - 资源类型（workspace/session/playground/settings）
  - 动作（read/write/admin/export）
  - 角色允许矩阵
- `Authorize(user, project, action, resource)`

### DSN-1803 审计导出服务

新增模块 `pkg/audit/export`：

- `CreateExportJob(start, end, format)`
- `RunExport(jobID)`
- `GetExportResult(jobID)`

新增表：
- `audit_export_jobs`
  - `id`, `org_id`, `requested_by`, `format`, `range_start`, `range_end`, `status`, `artifact_uri`, `created_at`, `finished_at`

### DSN-1804 合规策略服务

新增模块 `pkg/compliance`：

- `CheckKeyRotationPolicy()`
- `CheckAuditRetentionPolicy()`
- `RunRetentionCleanup()`

### DSN-1805 企业接入扩展点

新增接口占位：

- `POST /api/v1/enterprise/sso/config`
- `POST /api/v1/enterprise/scim/sync`

当前返回 `not_implemented`，但保留契约与鉴权。

### DSN-1806 API 设计

- `POST /api/v1/orgs`
- `POST /api/v1/projects`
- `POST /api/v1/projects/{id}/members`
- `POST /api/v1/audit/exports`
- `GET /api/v1/audit/exports/{id}`

## Requirement Coverage

- RQ-1801 -> DSN-1801
- RQ-1802 -> DSN-1802
- RQ-1803 -> DSN-1803
- RQ-1804 -> DSN-1804
- RQ-1805 -> DSN-1805
- RQ-1806 -> DSN-1802, DSN-1803
- RQ-1807 -> DSN-1802
- RQ-1808 -> DSN-1802, DSN-1803, DSN-1804
