# A2UI Step3 Validation Chain Tasks

## TASK-301 定义结构化校验错误模型

- Linked Requirements: RQ-303, RQ-304
- Linked Design: DSN-301, DSN-304
- Description:
  - 在 `pkg/a2ui` 增加 `ValidationError{Code,Path,Message}`
  - 调整 `ValidationResult.Errors` 类型
- DoD:
  - 代码编译通过
  - 基础单测可运行
- Verify:
  - `go test ./pkg/a2ui/...`

Status: completed

## TASK-302 实现组件白名单与 props 规则注册表

- Linked Requirements: RQ-301, RQ-302, RQ-305
- Linked Design: DSN-302
- Description:
  - 添加最小组件规则注册表
  - 添加默认白名单组件和 props 类型规则
- DoD:
  - 可通过注册表判定组件是否合法
  - props 规则可查询
- Verify:
  - `go test ./pkg/a2ui/...`

Status: completed

## TASK-303 重构 ValidateSchema 校验流水线

- Linked Requirements: RQ-301, RQ-302, RQ-303, RQ-304
- Linked Design: DSN-303, DSN-304
- Description:
  - 拆分 required/type/props 三类校验
  - 输出稳定错误码和精确路径
- DoD:
  - 不合法 schema 返回结构化错误
  - 合法 schema 返回 `Valid=true`
- Verify:
  - `go test ./pkg/a2ui/...`

Status: completed

## TASK-304 Agent 适配新错误模型并透传

- Linked Requirements: RQ-304
- Linked Design: DSN-305, DSN-306
- Description:
  - 更新 `pkg/agent/service.go` 处理 `[]ValidationError`
  - 失败时记录 `repair_failed` 事件并透传错误细节
- DoD:
  - agent 编译通过
  - 修复失败错误可读
- Verify:
  - `go test ./pkg/agent/...`

Status: completed

## TASK-305 补充/更新测试并跑全量回归

- Linked Requirements: RQ-306
- Linked Design: DSN-307, DSN-308
- Description:
  - 更新 `pkg/a2ui/service_test.go`
  - 增加 props/type/path 反向测试
  - 执行全量测试
- DoD:
  - 关键场景均有测试覆盖
  - `go test ./...` 通过
- Verify:
  - `go test ./...`

Status: completed
