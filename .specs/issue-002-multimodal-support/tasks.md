# Issue #2 多模态支持 Tasks

## TASK-0201 扩展模型能力标签与路由规则

- Linked Requirements: RQ-0201
- Linked Design: DSN-0201
- Description:
  - 扩展 provider 路由能力匹配逻辑
  - 增加 image/audio 任务类型映射
- DoD:
  - 能按能力标签正确选路
- Verify:
  - `go test ./pkg/provider/...`

Status: completed

## TASK-0202 定义多模态输入协议并接入上下文

- Linked Requirements: RQ-0202
- Linked Design: DSN-0202
- Description:
  - 定义 MultimodalInput 结构
  - 接入 session task context
- DoD:
  - 任务上下文可携带 media 数组
- Verify:
  - `go test ./pkg/session/...`

Status: completed

## TASK-0203 实现 Agent 多模态摘要节点

- Linked Requirements: RQ-0203, RQ-0206
- Linked Design: DSN-0203
- Description:
  - image/audio 输入预处理为结构化摘要
  - 失败时降级文本流程
- DoD:
  - 多模态输入可进入统一生成链路
- Verify:
  - `go test ./pkg/agent/...`

Status: completed

## TASK-0204 增加版本资源关联迁移与存储

- Linked Requirements: RQ-0204
- Linked Design: DSN-0204
- Description:
  - 新增 schema_version_assets 表
  - 增加 repository 查询接口
- DoD:
  - 可按版本查询资源引用
- Verify:
  - `go run ./cmd/cli db:migrate`
  - `go test ./internal/infra/sqlstore/...`

Status: completed

## TASK-0205 实现媒体存储抽象与 local 适配器

- Linked Requirements: RQ-0205
- Linked Design: DSN-0205
- Description:
  - 新增 media storage interface
  - 实现 local backend
- DoD:
  - 接口可独立替换
- Verify:
  - `go test ./pkg/media/...`

Status: completed

## TASK-0206 实现资源安全校验

- Linked Requirements: RQ-0202, RQ-0207
- Linked Design: DSN-0206
- Description:
  - 增加资源引用白名单与内网拦截规则
- DoD:
  - 非法引用可被拒绝
- Verify:
  - `go test ./pkg/security/...`

Status: completed

## TASK-0207 增加测试与文档

- Linked Requirements: RQ-0208
- Linked Design: DSN-0201, DSN-0203, DSN-0204
- Description:
  - 增加多模态回归测试
  - 新增 `docs/multimodal.md`
- DoD:
  - 关键场景可复现并通过
- Verify:
  - `go test ./...`

Status: completed
