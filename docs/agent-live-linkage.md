# Agent Live Data Linkage

## Overview

`/agent` 页面已从静态骨架升级为真实数据联动模式，打通输入、运行、监控、重试闭环。

## Runtime Flow

1. 页面初始化时自动创建 workspace/session。
2. 用户提交 prompt 调用 `POST /agent/run`。
3. 运行完成后拉取：
   - `GET /api/v1/debug/runs/{id}`
   - `GET /api/v1/debug/runs/{id}/cost`
4. 右侧 Monitor 映射真实数据：
   - tokens
   - thinking
   - files
   - output preview

## UX States

- `idle`：可输入并提交
- `running`：按钮禁用，显示发送中
- `success`：消息与 monitor 更新
- `error`：显示错误信息，支持“重试上次”

## Implementation Notes

- 核心状态与行为封装在：`web/src/use-agent-runtime.ts`
- 页面渲染在：`web/src/main.tsx` (`AgentPage`)
- 共享类型在：`web/src/agent-types.ts`
