# Issue #9 引导式空状态与 Starter Prompts Tasks

## TASK-0901 实现统一空状态组件并接入四个页面

- Linked Requirements: RQ-0901, RQ-0905
- Linked Design: DSN-0901
- Description:
  - 为 Workspaces/Agent/Debug/Provider 增加统一空状态块
- DoD:
  - 每个页面空状态至少 1 个 CTA
- Verify:
  - `npm --prefix web run test -- --run`

Status: completed

## TASK-0902 实现 Agent Starter Prompts

- Linked Requirements: RQ-0902
- Linked Design: DSN-0902
- Description:
  - 增加 Dashboard/Form/Landing 预设 prompt
  - 点击可填充并发送
- DoD:
  - starter prompt 至少覆盖 3 个场景
- Verify:
  - `npm --prefix web run test -- --run`

Status: completed

## TASK-0903 实现 Debug 摘要卡

- Linked Requirements: RQ-0903
- Linked Design: DSN-0903
- Description:
  - 在 Debug 页面增加成功率/平均耗时/成本摘要
- DoD:
  - 首屏可见关键指标
- Verify:
  - `npm --prefix web run test -- --run`

Status: completed

## TASK-0904 串联首次使用闭环引导

- Linked Requirements: RQ-0904
- Linked Design: DSN-0904
- Description:
  - 按页面动作路径提供下一步提示
- DoD:
  - 新用户可按提示完成首次生成
- Verify:
  - 人工 walkthrough

Status: completed

## TASK-0905 补充测试与文档

- Linked Requirements: RQ-0906
- Linked Design: DSN-0901, DSN-0902, DSN-0903
- Description:
  - 补充关键交互测试
  - 更新 UX 指南文档
- DoD:
  - 测试覆盖关键路径
- Verify:
  - `npm --prefix web run test -- --run`

Status: completed
