# A2UI Step18 Enterprise Readiness & Compliance Baseline Tasks

## TASK-1801 增加组织与项目迁移表

- Linked Requirements: RQ-1801
- Linked Design: DSN-1801
- Description:
  - 新增 orgs/projects/project_members
  - 扩展 workspace/session project_id
- DoD:
  - migration 可 up/down
- Verify:
  - `go run ./cmd/cli db:migrate`

Status: completed

## TASK-1802 实现 RBAC 授权引擎

- Linked Requirements: RQ-1802, RQ-1807
- Linked Design: DSN-1802
- Description:
  - 实现角色矩阵与授权判定
  - 接入 API 鉴权流程
- DoD:
  - 无权限请求返回 FORBIDDEN
- Verify:
  - `go test ./pkg/rbac/...`

Status: completed

## TASK-1803 实现审计导出任务

- Linked Requirements: RQ-1803
- Linked Design: DSN-1803
- Description:
  - 实现 audit 导出 job 创建、执行、查询
  - 支持 JSON/CSV 导出
- DoD:
  - 可按时间区间导出审计日志
- Verify:
  - `go test ./pkg/audit/export/...`

Status: completed

## TASK-1804 实现合规策略检查

- Linked Requirements: RQ-1804
- Linked Design: DSN-1804
- Description:
  - 实现密钥轮换与审计保留策略检查
  - 支持清理任务执行
- DoD:
  - 策略检查结果可查询
- Verify:
  - `go test ./pkg/compliance/...`

Status: completed

## TASK-1805 增加企业接入接口占位

- Linked Requirements: RQ-1805
- Linked Design: DSN-1805
- Description:
  - 增加 SSO/SCIM 配置接口占位
  - 接入鉴权与审计
- DoD:
  - 接口契约稳定，返回 not_implemented
- Verify:
  - `go test ./cmd/server/...`

Status: completed

## TASK-1806 集成 API 与文档

- Linked Requirements: RQ-1801, RQ-1803, RQ-1805
- Linked Design: DSN-1806
- Description:
  - 新增 org/project/audit-export API
  - 新增 `docs/enterprise-readiness.md`
- DoD:
  - 文档可指导企业管理员完成基础配置
- Verify:
  - 人工按文档演练

Status: completed

## TASK-1807 回归测试与队列验证

- Linked Requirements: RQ-1808
- Linked Design: DSN-1802, DSN-1803, DSN-1804
- Description:
  - 增加 RBAC/审计导出/合规策略回归测试
- DoD:
  - 全量测试通过
- Verify:
  - `go test ./...`

Status: completed
