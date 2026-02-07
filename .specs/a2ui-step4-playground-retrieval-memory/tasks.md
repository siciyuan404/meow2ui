# A2UI Step4 Playground Retrieval Memory Tasks

## TASK-401 定义检索数据结构与接口

- Linked Requirements: RQ-401, RQ-406
- Linked Design: DSN-401
- Description:
  - 新增 `Retriever`、`RetrievalQuery`、`RetrievalHit`
  - 放置在 `pkg/playground/retrieval`
- DoD:
  - 接口定义清晰
  - 编译通过
- Verify:
  - `go test ./pkg/playground/...`

Status: completed

## TASK-402 实现 KeywordRetriever

- Linked Requirements: RQ-401, RQ-405
- Linked Design: DSN-402
- Description:
  - 按 title/tags/theme/category 评分排序
  - 支持 limit 与 min score
- DoD:
  - 检索结果排序稳定
  - 支持无命中返回空列表
- Verify:
  - `go test ./pkg/playground/...`

Status: completed

## TASK-403 Agent 集成案例增强注入

- Linked Requirements: RQ-402
- Linked Design: DSN-403
- Description:
  - 在 Agent Run 流程中加入检索与上下文注入
  - plan/emit 均可收到 examples
- DoD:
  - 命中时上下文包含案例摘要
  - 不破坏原有主流程
- Verify:
  - `go test ./pkg/agent/...`

Status: completed

## TASK-404 增加开关与降级逻辑

- Linked Requirements: RQ-403
- Linked Design: DSN-404, DSN-405
- Description:
  - 增加全局与会话级检索开关
  - 检索失败时自动降级
- DoD:
  - 关闭开关时不检索
  - 检索异常不影响生成链路
- Verify:
  - `go test ./pkg/agent/...`

Status: completed

## TASK-405 增加审计事件字段

- Linked Requirements: RQ-404
- Linked Design: DSN-406
- Description:
  - 在 events payload 写入 retrieval 相关字段
  - 成功、跳过、异常都可观测
- DoD:
  - 事件可追踪命中来源
- Verify:
  - `go test ./pkg/events/...`

Status: completed

## TASK-406 扩展 Session 元数据并持久化

- Linked Requirements: RQ-403, RQ-406
- Linked Design: DSN-407
- Description:
  - `domain.Session` 增加 metadata
  - repository 增加 metadata JSON 读写
- DoD:
  - 可配置会话级检索开关
- Verify:
  - `go test ./pkg/session/...`

Status: completed

## TASK-407 补充测试与回归

- Linked Requirements: RQ-407
- Linked Design: DSN-408, DSN-409
- Description:
  - 覆盖命中/无命中/关闭/异常降级
  - 执行全量回归
- DoD:
  - 关键场景测试完备
  - 全量测试通过
- Verify:
  - `go test ./...`

Status: completed
