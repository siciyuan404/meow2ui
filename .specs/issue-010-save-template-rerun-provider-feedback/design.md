# Issue #10 一键保存模板 + Rerun + Provider 反馈 Design

## Design Goals

- 把“生成 -> 沉淀 -> 继续迭代”做成单页可达闭环。
- 尽量复用已有 marketplace 与 provider API。
- 保证失败时有明确反馈，不影响继续操作。

## Architecture

### DSN-1001 Agent Action Bar 扩展

- 在 Agent 页结果区域增加动作按钮：
  - Save to Template
  - Rerun Last
- 状态：`idle/loading/success/error`。

### DSN-1002 Save Template Integration

- 复用 `POST /api/v1/marketplace/templates`。
- 发送字段：name/category/tags/schema/theme/sessionId/versionId。
- 保存成功后展示跳转链接到 marketplace。

### DSN-1003 Rerun Strategy

- 复用 `use-agent-runtime` 中最近 prompt 与当前 session。
- rerun 时沿用 monitor 更新路径。

### DSN-1004 Provider Feedback UX

- Provider 页增加：
  - 当前激活 provider 显示
  - 连接测试按钮 loading 与结果 toast/inline message
- 失败展示错误摘要。

## Requirement Coverage

- RQ-1001 -> DSN-1001, DSN-1002
- RQ-1002 -> DSN-1001, DSN-1003
- RQ-1003 -> DSN-1004
- RQ-1004 -> DSN-1002
- RQ-1005 -> DSN-1001, DSN-1004
- RQ-1006 -> DSN-1001, DSN-1002, DSN-1004
