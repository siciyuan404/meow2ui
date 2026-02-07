# A2UI Step13 Plugin & Template Marketplace Tasks

## TASK-1301 增加模板与插件迁移表

- Linked Requirements: RQ-1301, RQ-1302, RQ-1304
- Linked Design: DSN-1301, DSN-1305
- Description:
  - 新增 templates/template_versions/plugins/project_plugins 表
  - 增加状态与审核字段
- DoD:
  - migration 可 up/down
- Verify:
  - `go run ./cmd/cli db:migrate`

Status: completed

## TASK-1302 实现模板服务与检索 API

- Linked Requirements: RQ-1301
- Linked Design: DSN-1301, DSN-1307
- Description:
  - 实现模板 CRUD 与筛选查询
  - 支持模板版本记录
- DoD:
  - 可创建并检索模板
- Verify:
  - `go test ./pkg/...`（template 相关）

Status: completed

## TASK-1303 实现模板应用到 Session

- Linked Requirements: RQ-1301, RQ-1305
- Linked Design: DSN-1304
- Description:
  - 实现 apply API
  - 应用模板后生成新 schema version
- DoD:
  - 版本链正确追加
- Verify:
  - `go test ./pkg/session/...`

Status: completed

## TASK-1304 实现插件注册与启停

- Linked Requirements: RQ-1302
- Linked Design: DSN-1301, DSN-1307
- Description:
  - 实现插件注册 API
  - 实现 workspace 级启停配置
- DoD:
  - 可启用/禁用插件
- Verify:
  - `go test ./pkg/...`（plugin 相关）

Status: completed

## TASK-1305 接入插件权限校验

- Linked Requirements: RQ-1303, RQ-1306
- Linked Design: DSN-1302, DSN-1303
- Description:
  - 运行时加载插件并通过 policy engine 校验
  - 拒绝超权限并写安全事件
- DoD:
  - 超权限调用被阻断
- Verify:
  - `go test ./pkg/security/...`

Status: completed

## TASK-1306 实现审核状态流

- Linked Requirements: RQ-1304
- Linked Design: DSN-1305
- Description:
  - 实现 draft/reviewed/published/blocked 状态流 API
- DoD:
  - 非法状态迁移被拒绝
- Verify:
  - `go test ./pkg/...`（review flow 相关）

Status: completed

## TASK-1307 实现 Web 市场页

- Linked Requirements: RQ-1305
- Linked Design: DSN-1306
- Description:
  - 模板列表、插件列表、筛选与模板应用入口
- DoD:
  - 页面可完成模板应用主流程
- Verify:
  - `npm run dev` + 手工验证

Status: completed

## TASK-1308 回归与文档

- Linked Requirements: RQ-1308
- Linked Design: DSN-1307
- Description:
  - 增加测试并更新 marketplace 文档
- DoD:
  - 全量测试通过
- Verify:
  - `go test ./...`

Status: completed
