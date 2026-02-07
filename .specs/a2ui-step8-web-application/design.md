# A2UI Step8 Web Application Design

## Design Goals

- 提供可用的 Web 端产品入口，与 Go 后端能力对齐。
- 保持前后端契约清晰，避免 UI 与后端耦合。
- 使用模块化前端结构，便于后续扩展为多用户/协同场景。

## Frontend Architecture

### DSN-801 技术栈建议

- React + TypeScript + Vite
- 状态管理：Zustand（与当前项目习惯一致）
- 请求层：Fetch 封装或 TanStack Query（二选一，MVP 可先 fetch 封装）
- UI：shadcn/ui + Tailwind

### DSN-802 页面路由

- `/workspaces`
- `/workspaces/:workspaceId/sessions`
- `/sessions/:sessionId/editor`
- `/playground`
- `/settings`

### DSN-803 前端模块拆分

```text
web/
  src/
    app/
    pages/
      workspaces/
      sessions/
      editor/
      playground/
      settings/
    components/
      layout/
      editor/
      preview/
      common/
    services/
      api/
      mappers/
    stores/
    hooks/
    types/
```

### DSN-804 API 契约适配层

新增 `services/api/client.ts`：

- 统一 baseURL
- 统一错误转换（映射 `{code,message,detail,traceId}`）
- 统一超时与重试策略（轻量）

### DSN-805 Editor 页面结构

- 左：会话与版本侧栏
- 中：Prompt + JSON Editor
- 右：Preview 渲染

交互流：
1. 加载 session 当前版本
2. 输入 prompt 调用 `agent/run`
3. 成功后更新版本列表与预览

### DSN-806 Playground 页面结构

- 顶部：搜索框 + 分类 + 标签过滤
- 列表：案例卡片（标题、标签、主题、时间）
- 操作：保存当前会话内容到 playground

### DSN-807 全局状态

核心状态：
- `activeWorkspaceId`
- `activeSessionId`
- `sessionVersions`
- `currentSchema`
- `themes`
- `playgroundQuery`

本地持久化：
- 最近活动 workspace/session
- UI 偏好（面板宽度、主题）

### DSN-808 错误与空态规范

- 网络错误：显示“连接失败 + 重试”
- 业务错误：显示错误码与人类可读消息
- 空态：为 workspace/session/playground 提供引导 CTA

### DSN-809 测试策略

- 单元测试：store、api mapper、关键 hooks
- 组件测试：editor 提交流程、playground 检索流程
- E2E（可选后续）：workspace->session->agent->save playground

## Requirement Coverage

- RQ-801 -> DSN-802, DSN-803, DSN-807
- RQ-802 -> DSN-802, DSN-804
- RQ-803 -> DSN-805, DSN-808
- RQ-804 -> DSN-804, DSN-807
- RQ-805 -> DSN-806
- RQ-806 -> DSN-804, DSN-808
- RQ-807 -> DSN-801
- RQ-808 -> DSN-805, DSN-806
- RQ-809 -> DSN-803
- RQ-810 -> DSN-809
