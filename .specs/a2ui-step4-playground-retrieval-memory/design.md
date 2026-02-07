# A2UI Step4 Playground Retrieval Memory Design

## Design Goals

- 将 Playground 从“被动存档”升级为“主动增强上下文”。
- 在不引入向量库的情况下实现稳定且可解释的检索增强。
- 保持 Agent 主流程低侵入，支持未来切换语义检索。

## Architecture

### DSN-401 检索策略接口

新增 `pkg/playground/retrieval`：

- `Retriever` 接口
  - `Retrieve(ctx, query RetrievalQuery) ([]RetrievalHit, error)`

- `RetrievalQuery`
  - `Text string`
  - `ThemeID string`
  - `CategoryID string`
  - `Tags []string`
  - `Limit int`

- `RetrievalHit`
  - `ItemID string`
  - `Title string`
  - `Tags []string`
  - `Summary string`
  - `Score float64`

### DSN-402 默认检索实现

默认 `KeywordRetriever`：

1. 基于 title/tags/category/theme 过滤
2. 简单评分：
   - 标题关键词命中 +3
   - 标签命中 +2
   - 主题匹配 +2
   - 分类匹配 +1
3. 按 score 降序，取 top N

### DSN-403 Agent 上下文注入点

在 `pkg/agent/service.go` 的 `BuildContext` 后注入：

- `task_context.examples` = `[]RetrievalHitSummary`
- 每条仅保留：`title/tags/summary`

并透传给 `plan` 与 `emit` 的 provider 请求。

### DSN-404 配置开关

新增运行时配置：

- `PLAYGROUND_RETRIEVAL_ENABLED`（默认 true）
- `PLAYGROUND_RETRIEVAL_LIMIT`（默认 3）
- `PLAYGROUND_RETRIEVAL_MIN_SCORE`（默认 1）

可在会话级覆盖：
- `session.metadata.retrieval_enabled`

### DSN-405 失败降级

- 检索错误不阻断主流程
- 记录 `retrieval_skipped` 事件并写入原因
- 继续执行原始 Agent 计划/生成链路

### DSN-406 审计字段

在 events payload 中增加：

- `retrieval_used: bool`
- `retrieval_hits: int`
- `retrieval_item_ids: []string`
- `retrieval_latency_ms: int`
- `retrieval_skip_reason: string`

## Data/Interface Changes

### DSN-407 Session 元数据扩展

- `domain.Session` 新增 `Metadata map[string]any`（可复用 provider step 的 metadata 思路）
- repository 增加 metadata 字段持久化（JSON）

## Testing Strategy

### DSN-408 单测

- KeywordRetriever：命中排序、limit、生效阈值
- Agent 注入：有命中/无命中/关闭开关/异常降级

### DSN-409 回归

- `go test ./pkg/playground/...`
- `go test ./pkg/agent/...`
- `go test ./...`

## Requirement Coverage

- RQ-401 -> DSN-401, DSN-402
- RQ-402 -> DSN-403
- RQ-403 -> DSN-404, DSN-405
- RQ-404 -> DSN-406
- RQ-405 -> DSN-402, DSN-406
- RQ-406 -> DSN-401
- RQ-407 -> DSN-408, DSN-409
