# Issue #10 一键保存模板 + Rerun + Provider 反馈 Tasks

## TASK-1001 扩展 Agent 动作栏

- Linked Requirements: RQ-1001, RQ-1002, RQ-1005
- Linked Design: DSN-1001
- Description:
  - 在 Agent 页面增加 Save/Rerun 动作按钮与状态
- DoD:
  - 用户可直接触发保存与 rerun
- Verify:
  - `npm --prefix web run test -- --run`

Status: completed

## TASK-1002 接入保存模板 API

- Linked Requirements: RQ-1001, RQ-1004
- Linked Design: DSN-1002
- Description:
  - 调用 marketplace create template API
  - 保存后给出跳转与反馈
- DoD:
  - 成功保存后可在 marketplace 查询到模板
- Verify:
  - `go test ./...`
  - `npm --prefix web run test -- --run`

Status: completed

## TASK-1003 强化 rerun 行为

- Linked Requirements: RQ-1002
- Linked Design: DSN-1003
- Description:
  - rerun 复用最近 prompt 与上下文
  - 结果联动 monitor
- DoD:
  - rerun 可稳定执行并更新视图
- Verify:
  - `npm --prefix web run test -- --run`

Status: completed

## TASK-1004 Provider 连接测试反馈

- Linked Requirements: RQ-1003, RQ-1005
- Linked Design: DSN-1004
- Description:
  - Provider 页增加连接测试状态反馈
- DoD:
  - 可明确区分连接成功与失败
- Verify:
  - `npm --prefix web run test -- --run`

Status: completed

## TASK-1005 测试与文档

- Linked Requirements: RQ-1006
- Linked Design: DSN-1001, DSN-1002, DSN-1004
- Description:
  - 补充关键交互测试
  - 更新 UX/Marketplace 联动文档
- DoD:
  - 关键场景有自动化覆盖
- Verify:
  - `go test ./...`
  - `npm --prefix web run test -- --run`

Status: completed
