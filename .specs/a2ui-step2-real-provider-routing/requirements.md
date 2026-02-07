# A2UI Step2 Real Provider Routing Requirements

## Overview

本规范定义 A2UI 平台第 2 步：接入真实 LLM Provider（优先 OpenAI 兼容接口），并实现按任务类型路由与降级。

## Scope

- In Scope
  - 接入 OpenAI 兼容 provider 适配器
  - 任务路由：`plan` / `emit` / `repair`
  - 主备模型降级策略
  - 基础重试、超时、错误分类
  - Provider 配置持久化与启停开关
- Out of Scope
  - 多模态（图像/语音/视频）调用
  - Provider 成本优化策略引擎

## Functional Requirements

### RQ-201 真实 Provider 接入

- WHEN `provider.type=openai_compatible` THEN THE SYSTEM SHALL 使用 HTTP 调用兼容 Chat Completions API。
- WHEN provider 调用成功 THEN THE SYSTEM SHALL 返回标准化 `GenerateResponse{text,tokens}`。

### RQ-202 任务路由

- WHEN Agent 发起 `plan` 任务 THEN THE SYSTEM SHALL 路由到具备 `text` 能力且标记为 `plan` 首选模型。
- WHEN Agent 发起 `emit` 任务 THEN THE SYSTEM SHALL 路由到具备 `text` 能力且标记为 `emit` 首选模型。
- WHEN Agent 发起 `repair` 任务 THEN THE SYSTEM SHALL 路由到具备 `text` 能力且标记为 `repair` 首选模型。

### RQ-203 降级与重试

- WHEN 首选模型失败 THEN THE SYSTEM SHALL 在同能力模型中按优先级尝试备选模型。
- WHEN 请求出现可重试错误（429/5xx/超时） THEN THE SYSTEM SHALL 执行有限重试（最多 2 次）。
- WHEN 全部失败 THEN THE SYSTEM SHALL 返回结构化失败原因。

### RQ-204 配置与安全

- WHEN Provider 配置保存 THEN THE SYSTEM SHALL 支持 `base_url`、`api_key_ref`、`timeout_ms`、`enabled`。
- WHEN 运行时读取密钥 THEN THE SYSTEM SHALL 从环境变量或安全存储读取真实密钥，不在日志输出明文。

### RQ-205 观测与审计

- WHEN Provider 被调用 THEN THE SYSTEM SHALL 记录 provider/model、耗时、token、错误类型。

## Non-functional Requirements

### RQ-206 可靠性

- WHEN 外部网络波动 THEN THE SYSTEM SHALL 在超时后快速失败并触发降级，不阻塞主线程。

### RQ-207 可维护性

- WHEN 新增 Provider THEN THE SYSTEM SHALL 仅新增 Adapter，不改动 Agent 主流程。

### RQ-208 可测试性

- WHEN 提交实现 THEN THE SYSTEM SHALL 提供 adapter/router 的单测与最小集成测试。
