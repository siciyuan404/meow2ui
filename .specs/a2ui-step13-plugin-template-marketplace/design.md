# A2UI Step13 Plugin & Template Marketplace Design

## Design Goals

- 构建“模板 + 插件”双生态基础，支持 A2UI 平台扩展。
- 让插件加载与权限控制可审计、可禁用。
- 保持现有 session/agent/security 架构兼容。

## Architecture

### DSN-1301 数据模型

新增表：

- `templates`
  - `id`, `name`, `version`, `category`, `tags(json)`, `schema_snapshot`, `owner_user_id`, `status`, `created_at`
- `template_versions`
  - `id`, `template_id`, `version`, `schema_snapshot`, `changelog`, `created_at`
- `plugins`
  - `id`, `name`, `version`, `entrypoint`, `capabilities(json)`, `permissions(json)`, `owner_user_id`, `status`, `created_at`
- `project_plugins`
  - `workspace_id`, `plugin_id`, `enabled`, `config(json)`

### DSN-1302 插件权限模型

- 权限枚举：`read`, `write`, `exec`, `network`, `provider_call`
- 插件执行前必须通过 `policy engine` 校验
- 拒绝事件写入安全审计：`SECURITY_PLUGIN_PERMISSION_BLOCKED`

### DSN-1303 运行时插件加载

- `PluginLoader`（pkg/plugins）
  - 输入：workspace/session
  - 输出：可用插件列表（enabled && published）
- Agent 在 plan/emit 前注入 `available_plugins` 到 task context

### DSN-1304 模板应用流程

1. 用户在 Web 市场选择模板
2. 后端读取模板最新版本 schema
3. 创建新 schema version（parent 指向当前）
4. 记录事件 `template_applied`

### DSN-1305 审核流

- 状态机：`draft -> reviewed -> published`，异常可 `blocked`
- 审核字段：`reviewed_by`, `reviewed_at`, `review_note`

### DSN-1306 Web 页面

- `web/src/pages/marketplace/`
  - 模板列表与筛选
  - 插件列表与启停
  - 模板应用按钮（Apply to Session）

### DSN-1307 API 设计

- `GET /api/v1/templates`
- `POST /api/v1/templates`
- `POST /api/v1/templates/{id}/apply`
- `GET /api/v1/plugins`
- `POST /api/v1/plugins`
- `POST /api/v1/plugins/{id}/toggle`

## Requirement Coverage

- RQ-1301 -> DSN-1301, DSN-1304, DSN-1307
- RQ-1302 -> DSN-1301, DSN-1302, DSN-1307
- RQ-1303 -> DSN-1302, DSN-1303
- RQ-1304 -> DSN-1305
- RQ-1305 -> DSN-1306
- RQ-1306 -> DSN-1302
- RQ-1307 -> DSN-1301
- RQ-1308 -> DSN-1304, DSN-1306, DSN-1307
