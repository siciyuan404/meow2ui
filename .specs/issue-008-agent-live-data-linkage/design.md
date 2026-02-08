# Issue #8 Agent 页面真实数据联动 Design

## Design Goals

- 复用现有后端接口，最小侵入实现 Agent 工作台可用闭环。
- 保持页面结构与既有设计一致，仅替换数据源与交互行为。
- 失败可恢复、状态可追踪。

## Architecture

### DSN-0801 Agent Frontend State Model

- 在 `/agent` 页面建立状态模型：
  - `inputText`
  - `messages[]`
  - `currentRunId`
  - `runStatus` (`idle|running|success|error`)
  - `monitorData`
  - `errorState`

### DSN-0802 API Integration Layer

- 统一封装接口调用：
  - `POST /agent/run`
  - `GET /api/v1/debug/runs/{id}`
  - `GET /api/v1/debug/runs/{id}/cost`
- 失败时统一提取 message 与 trace_id。

### DSN-0803 Message & Status Rendering

- 对话区渲染规则：
  - 用户消息：提交后立即追加
  - 系统消息：运行中显示占位，完成后替换
  - 错误消息：高亮展示，保留重试入口

### DSN-0804 Monitor Mapping

- `tokens` 来自 debug/cost 聚合
- `thinking` 来自 steps/payload 摘要
- `files` 来自 output payload 的文件相关字段
- `output` 展示 schema 片段或结构化摘要

### DSN-0805 Retry & Recovery

- 提供 `rerun` 快捷操作（重用最后一次 prompt）
- 失败不清空历史消息，允许用户继续输入。

## Requirement Coverage

- RQ-0801 -> DSN-0801, DSN-0802
- RQ-0802 -> DSN-0801, DSN-0803, DSN-0805
- RQ-0803 -> DSN-0801, DSN-0804
- RQ-0804 -> DSN-0802, DSN-0804
- RQ-0805 -> DSN-0803
- RQ-0806 -> DSN-0805
- RQ-0807 -> DSN-0801, DSN-0803, DSN-0804
