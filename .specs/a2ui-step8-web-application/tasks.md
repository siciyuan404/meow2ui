# A2UI Step8 Web Application Tasks

## TASK-801 搭建 Web 工程骨架与路由

- Linked Requirements: RQ-801, RQ-809
- Linked Design: DSN-801, DSN-802, DSN-803
- Description:
  - 创建 `web/` 工程与基础路由
  - 搭建主布局与导航
- DoD:
  - 可访问主要页面路由
  - 项目可本地启动
- Verify:
  - `npm run dev`（web）

Status: completed

## TASK-802 实现 API Client 与错误映射

- Linked Requirements: RQ-806
- Linked Design: DSN-804, DSN-808
- Description:
  - 新增统一 API client
  - 对接后端错误结构并做前端友好提示
- DoD:
  - API 调用有统一成功/失败处理
- Verify:
  - 单元测试 + 手动调用验证

Status: completed

## TASK-803 Workspace/Session 页面实现

- Linked Requirements: RQ-802
- Linked Design: DSN-802, DSN-804
- Description:
  - 实现 workspace 列表、创建
  - 实现 session 列表、创建与跳转
- DoD:
  - 用户可从 0 创建并进入 session
- Verify:
  - 手工流程验证通过

Status: completed

## TASK-804 Editor+Preview 主流程实现

- Linked Requirements: RQ-803, RQ-808
- Linked Design: DSN-805, DSN-807, DSN-808
- Description:
  - 实现 prompt 提交、agent 生成、版本列表、预览刷新
  - 增加 loading/error 状态
- DoD:
  - `agent/run` 可在页面端到端执行
  - 失败可读、可重试
- Verify:
  - 页面手动 E2E 流程

Status: completed

## TASK-805 Theme 页面实现

- Linked Requirements: RQ-804
- Linked Design: DSN-804, DSN-807
- Description:
  - 展示主题列表与切换
  - 支持创建基础自定义主题
- DoD:
  - 主题切换能影响预览
- Verify:
  - 手工切换验证

Status: completed

## TASK-806 Playground 页面实现

- Linked Requirements: RQ-805
- Linked Design: DSN-806
- Description:
  - 实现保存、分类、标签、检索列表
- DoD:
  - 可保存会话内容到 playground
  - 可按条件检索
- Verify:
  - 页面手工验证

Status: completed

## TASK-807 Web 测试与构建校验

- Linked Requirements: RQ-810
- Linked Design: DSN-809
- Description:
  - 增加关键模块测试
  - 接入 web 构建检查
- DoD:
  - 关键测试通过
  - 构建成功
- Verify:
  - `npm run test`（web）
  - `npm run build`（web）

Status: completed
