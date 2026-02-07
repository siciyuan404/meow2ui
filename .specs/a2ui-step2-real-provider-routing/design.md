# A2UI Step2 Real Provider Routing Design

## Design Goals

- 在现有 `pkg/provider` 基础上接入真实 OpenAI 兼容 Provider。
- 保持 Adapter 抽象稳定，支持后续扩展 Anthropic/Gemini/本地模型。
- 提供可观测的路由与降级决策链。

## Architecture

### DSN-201 Provider Adapter 扩展

新增 adapter：`OpenAICompatibleAdapter`

- 输入：`domain.Provider` + `domain.Model` + `GenerateRequest`
- 调用：`POST {base_url}/chat/completions`
- 输出：`GenerateResponse{text,tokens}`

响应映射：
- `choices[0].message.content` -> `text`
- `usage.total_tokens` -> `tokens`

### DSN-202 路由元数据

在 `model.capabilities` 之外新增路由标签（可先存在 JSON 扩展字段中）：

- `role:plan|emit|repair`
- `priority:int`（值越小优先级越高）

路由算法：
1. 过滤 `enabled=true` 且能力匹配
2. 过滤角色匹配
3. 按 `priority` 升序
4. 依次尝试直到成功

### DSN-203 重试与降级

重试条件：
- HTTP 429
- HTTP 5xx
- 超时错误

策略：
- 单模型重试 2 次（指数退避）
- 单模型失败后切换下一候选模型

### DSN-204 错误模型

定义 `ProviderError`：
- `Code`（`PROVIDER_TIMEOUT`/`PROVIDER_RATE_LIMIT`/`PROVIDER_UPSTREAM_ERROR`/`PROVIDER_AUTH_ERROR`）
- `ProviderID`
- `ModelID`
- `Retryable`
- `Message`

### DSN-205 安全与密钥

- `provider.auth_ref` 作为密钥引用键（如 `OPENAI_API_KEY`）
- 运行时通过 `os.Getenv(auth_ref)` 读取
- 日志中仅输出 `auth_ref`，不输出 key 值

### DSN-206 可观测性

在 provider service 记录：
- `provider_id`
- `model_id`
- `task_type`
- `latency_ms`
- `token_out`
- `error_code`（如失败）

## Data/Interface Changes

### DSN-207 兼容变更

- `domain.Model` 增加 `Metadata map[string]any`（承载 role/priority）
- `ProviderRepository` 读写支持 metadata 字段

## Testing Strategy

### DSN-208 单测

- Router：角色过滤、优先级排序、降级顺序
- Adapter：成功响应、429/500/401/timeout 分支
- Service：重试与失败回传

### DSN-209 集成验证

- 使用本地 mock server 模拟 OpenAI 兼容接口
- 验证 plan/emit/repair 路由到不同模型

## Requirement Coverage

- RQ-201 -> DSN-201
- RQ-202 -> DSN-202
- RQ-203 -> DSN-203, DSN-204
- RQ-204 -> DSN-205, DSN-207
- RQ-205 -> DSN-206
- RQ-206 -> DSN-203
- RQ-207 -> DSN-201
- RQ-208 -> DSN-208, DSN-209
