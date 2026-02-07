# A2UI Step2 Real Provider Routing Tasks

## TASK-201 扩展 Model 元数据与仓储字段

- Linked Requirements: RQ-202, RQ-204
- Linked Design: DSN-202, DSN-207
- Description:
  - 在 `domain.Model` 增加 `Metadata`
  - 更新 repository 存取逻辑支持 metadata
- DoD:
  - 编译通过
  - metadata 可读写
- Verify:
  - `go test ./pkg/...`

Status: completed

## TASK-202 实现 OpenAI 兼容 Adapter

- Linked Requirements: RQ-201, RQ-204
- Linked Design: DSN-201, DSN-205
- Description:
  - 新增 `OpenAICompatibleAdapter`
  - 实现请求构造、响应解析、超时控制
- DoD:
  - 成功调用返回 text/tokens
  - 错误码映射可用
- Verify:
  - `go test ./pkg/provider/...`

Status: completed

## TASK-203 重构 Router 以支持 role/priority

- Linked Requirements: RQ-202, RQ-203
- Linked Design: DSN-202, DSN-203
- Description:
  - 按 task role 过滤模型
  - 按 priority 排序并返回候选列表
- DoD:
  - plan/emit/repair 可路由不同模型
  - 候选顺序稳定
- Verify:
  - `go test ./pkg/provider/...`

Status: completed

## TASK-204 实现 ProviderService 重试与降级

- Linked Requirements: RQ-203, RQ-206
- Linked Design: DSN-203, DSN-204
- Description:
  - 对单模型引入有限重试
  - 失败后切换下一候选模型
  - 输出结构化 ProviderError
- DoD:
  - 429/5xx/timeout 可重试
  - 所有候选失败时返回清晰错误
- Verify:
  - `go test ./pkg/provider/...`

Status: completed

## TASK-205 增加调用观测信息

- Linked Requirements: RQ-205
- Linked Design: DSN-206
- Description:
  - 在 provider 调用处记录 provider/model/task/latency/token/error
  - 与现有 events/telemetry 对齐
- DoD:
  - 成功与失败都可追踪
- Verify:
  - `go test ./pkg/provider/...`

Status: completed

## TASK-206 集成验证与文档

- Linked Requirements: RQ-208
- Linked Design: DSN-209
- Description:
  - 用 mock server 做最小集成测试
  - 更新运行文档（环境变量、模型配置示例）
- DoD:
  - `go test ./...` 通过
  - 文档可复现
- Verify:
  - `go test ./...`

Status: completed
