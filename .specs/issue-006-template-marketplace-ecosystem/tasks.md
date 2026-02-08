# Issue #6 模板市场与生态建设 Tasks

## TASK-0601 扩展模板检索索引与查询

- Linked Requirements: RQ-0601
- Linked Design: DSN-0601
- Description:
  - 增加 popularity/download/rating 指标字段
  - 扩展模板检索排序和过滤
- DoD:
  - 可按多条件稳定检索与排序
- Verify:
  - `go test ./pkg/...`

Status: completed

## TASK-0602 实现发布审核工作流扩展

- Linked Requirements: RQ-0602, RQ-0606
- Linked Design: DSN-0602
- Description:
  - 新增 submitted 状态与审核记录表
  - 增加审核 API
- DoD:
  - 审核通过与驳回路径完整可用
- Verify:
  - `go test ./pkg/...`

Status: completed

## TASK-0603 实现评分与评论模块

- Linked Requirements: RQ-0603
- Linked Design: DSN-0603
- Description:
  - 新增 ratings/comments 表与 API
  - 增加违规标记状态流
- DoD:
  - 可提交评分评论并可审核状态变更
- Verify:
  - `go test ./pkg/...`

Status: completed

## TASK-0604 实现 Session 模板打包与版本清单

- Linked Requirements: RQ-0604, RQ-0607
- Linked Design: DSN-0604
- Description:
  - 从 session 打包 schema/theme/dependencies
  - 增加版本 manifest 记录
- DoD:
  - 可回溯版本依赖和变更摘要
- Verify:
  - `go test ./pkg/session/...`

Status: completed

## TASK-0605 实现模板应用兼容校验

- Linked Requirements: RQ-0605, RQ-0606
- Linked Design: DSN-0605
- Description:
  - 应用前执行依赖/兼容校验
  - 失败返回结构化缺失项
- DoD:
  - 非兼容模板不会污染 session 版本链
- Verify:
  - `go test ./pkg/...`

Status: completed

## TASK-0606 实现 Web 市场体验增强

- Linked Requirements: RQ-0601, RQ-0602, RQ-0603
- Linked Design: DSN-0606
- Description:
  - 增加模板详情、评分评论、审核状态展示
- DoD:
  - 市场主流程（发现-查看-应用）完整可用
- Verify:
  - `npm run test -- --run`

Status: completed

## TASK-0607 补充测试与文档

- Linked Requirements: RQ-0608
- Linked Design: DSN-0601, DSN-0602, DSN-0603, DSN-0605
- Description:
  - 增加市场回归测试
  - 新增 `docs/marketplace-ecosystem.md`
- DoD:
  - 团队可按文档执行完整流程
- Verify:
  - `go test ./...`
  - `npm run test -- --run`

Status: completed
