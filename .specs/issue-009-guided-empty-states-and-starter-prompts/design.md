# Issue #9 引导式空状态与 Starter Prompts Design

## Design Goals

- 降低首次进入各页面时的空白感。
- 在 Agent 页提供“可直接用”的起步 prompt。
- 让 Debug 页首屏先看见核心健康度指标。

## Architecture

### DSN-0901 Empty State Component Pattern

- 定义统一空状态块：
  - 标题
  - 描述
  - 主 CTA 按钮
  - 次 CTA（可选）
- 应用页面：Workspaces、Agent、Debug、Provider。

### DSN-0902 Starter Prompt Actions

- 在 Agent 页头部增加 starter prompt 列表：
  - Dashboard
  - Form
  - Landing
- 点击动作：
  - 填充输入框
  - 可立即触发发送

### DSN-0903 Debug Summary Cards

- 在 Debug 页列表上方展示摘要卡：
  - Success Rate
  - Avg Latency
  - Total Cost
- 数据来源：现有 run list + detail/cost 聚合。

### DSN-0904 Guided First-Run Path

- 在 Workspaces 空状态提供“创建工作区”动作。
- 在 Provider 空状态提供“新增 Provider”动作。
- Agent 空状态文案引导使用 starter prompt。

## Requirement Coverage

- RQ-0901 -> DSN-0901, DSN-0904
- RQ-0902 -> DSN-0902
- RQ-0903 -> DSN-0903
- RQ-0904 -> DSN-0904
- RQ-0905 -> DSN-0901
- RQ-0906 -> DSN-0901, DSN-0902, DSN-0903
