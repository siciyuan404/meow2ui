# A2UI Step5 Production Readiness Tasks

## TASK-501 实现配置加载与校验模块

- Linked Requirements: RQ-501
- Linked Design: DSN-501
- Description:
  - 增加 `internal/infra/config`，统一加载运行配置
  - 启动前执行 Validate
- DoD:
  - 非法配置会阻断启动并输出明确错误
- Verify:
  - `go test ./internal/infra/config/...`

Status: completed

## TASK-502 增强健康检查接口

- Linked Requirements: RQ-502
- Linked Design: DSN-502
- Description:
  - 保留 `/healthz`
  - 新增 `/readyz`，postgres 模式检查数据库连接
- DoD:
  - 两个接口都可返回结构化健康状态
- Verify:
  - `go test ./cmd/server/...`

Status: completed

## TASK-503 增加统一错误响应与 trace id

- Linked Requirements: RQ-503, RQ-505
- Linked Design: DSN-503, DSN-504
- Description:
  - 新增 `pkg/httpx`
  - API 错误统一返回 `{code,message,detail,traceId}`
  - 接入 trace id middleware
- DoD:
  - 已知错误码可稳定映射
  - 响应头含 `X-Trace-Id`
- Verify:
  - `go test ./pkg/httpx/...`

Status: completed

## TASK-504 增加 smoke 检查脚本

- Linked Requirements: RQ-504, RQ-506
- Linked Design: DSN-505
- Description:
  - 新增脚本串联测试、迁移、健康检查
  - 失败时非零退出
- DoD:
  - 本地可一键执行基础 smoke 流程
- Verify:
  - `bash scripts/smoke.sh` 或等价命令

Status: completed

## TASK-505 补充运维文档

- Linked Requirements: RQ-504, RQ-505
- Linked Design: DSN-506
- Description:
  - 新增 `docs/runbook.md`
  - 新增 `docs/api-errors.md`
- DoD:
  - 文档可覆盖常见启动与排障路径
- Verify:
  - 人工审阅 + 按文档实操通过

Status: completed
