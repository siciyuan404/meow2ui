# Issue #5 可视化 Agent 调试工具 Tasks

## TASK-0501 定义调试查询 DTO 与聚合接口

- Linked Requirements: RQ-0501, RQ-0502, RQ-0503, RQ-0505
- Linked Design: DSN-0501, DSN-0502
- Description:
  - 建立 `pkg/debugger` 服务接口与 DTO
  - 明确 run/step/tool/context/cost 数据契约
- DoD:
  - DTO 与 service 接口可编译
  - 字段覆盖 requirements 关键视图
- Verify:
  - `go test ./pkg/debugger/...`

Status: completed

## TASK-0502 实现 debug runs 查询 API

- Linked Requirements: RQ-0501, RQ-0506
- Linked Design: DSN-0501, DSN-0503
- Description:
  - 实现 `GET /api/v1/debug/runs`
  - 支持 session/status/time/provider/model 过滤
- DoD:
  - 列表查询返回稳定分页结果
- Verify:
  - `go test ./cmd/server/...`

Status: completed

## TASK-0503 实现 debug run 详情 API

- Linked Requirements: RQ-0502, RQ-0503, RQ-0504, RQ-0506
- Linked Design: DSN-0501, DSN-0502, DSN-0503
- Description:
  - 实现 `GET /api/v1/debug/runs/{id}`
  - 聚合步骤、工具链、上下文窗口占比
- DoD:
  - 可返回单次 run 完整调试详情
- Verify:
  - `go test ./pkg/debugger/...`

Status: completed

## TASK-0504 实现成本聚合 API

- Linked Requirements: RQ-0505
- Linked Design: DSN-0501, DSN-0503
- Description:
  - 实现 `GET /api/v1/debug/runs/{id}/cost`
  - 统计 provider/model 维度 token 与估算成本
- DoD:
  - 成本口径与 observability 指标一致
- Verify:
  - `go test ./pkg/observability/...`

Status: completed

## TASK-0505 实现调试输出脱敏

- Linked Requirements: RQ-0507
- Linked Design: DSN-0504
- Description:
  - 新增 redaction 规则并接入 debug API 输出路径
- DoD:
  - 敏感字段不以明文返回
- Verify:
  - `go test ./pkg/debugger/redaction/...`

Status: completed

## TASK-0506 实现 Web 调试页路由与列表

- Linked Requirements: RQ-0501
- Linked Design: DSN-0505
- Description:
  - 增加 `/debug` 路由与 run 列表页
  - 支持基础筛选与跳转详情
- DoD:
  - 页面可完成列表检索与详情跳转
- Verify:
  - `npm run test -- --run`

Status: completed

## TASK-0507 实现 Web 调试详情视图

- Linked Requirements: RQ-0502, RQ-0503, RQ-0504, RQ-0505
- Linked Design: DSN-0505
- Description:
  - 实现流程时间线、工具调用链、上下文占比、成本面板
- DoD:
  - 失败步骤可高亮，数据可联动切换
- Verify:
  - `npm run test -- --run`

Status: completed

## TASK-0508 回归测试与文档

- Linked Requirements: RQ-0508
- Linked Design: DSN-0501, DSN-0504, DSN-0505
- Description:
  - 增加后端与前端测试
  - 新增 `docs/debugger.md`
- DoD:
  - 核心测试通过，文档可按步骤复现
- Verify:
  - `go test ./...`
  - `npm run test -- --run`

Status: completed
