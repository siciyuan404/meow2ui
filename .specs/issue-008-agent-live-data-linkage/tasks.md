# Issue #8 Agent 页面真实数据联动 Tasks

## TASK-0801 建立 Agent 页面状态模型与类型

- Linked Requirements: RQ-0801, RQ-0802, RQ-0803
- Linked Design: DSN-0801
- Description:
  - 定义消息、运行状态、monitor 数据类型
  - 初始化状态与默认值
- DoD:
  - 状态可驱动页面核心区域渲染
- Verify:
  - `npm --prefix web run test -- --run`

Status: completed

## TASK-0802 接入 /agent/run 提交链路

- Linked Requirements: RQ-0801, RQ-0802, RQ-0805
- Linked Design: DSN-0802, DSN-0803
- Description:
  - 输入提交调用 `/agent/run`
  - 实现 loading/success/error 状态切换
- DoD:
  - 用户提交后可看到实时状态反馈
- Verify:
  - `npm --prefix web run test -- --run`

Status: completed

## TASK-0803 接入 debug 详情与 cost 联动

- Linked Requirements: RQ-0803, RQ-0804
- Linked Design: DSN-0802, DSN-0804
- Description:
  - run 完成后拉取 debug detail + cost
  - 将数据映射到 monitor 面板
- DoD:
  - tokens/thinking/files/output 非占位
- Verify:
  - `npm --prefix web run test -- --run`

Status: completed

## TASK-0804 增加错误展示与重试

- Linked Requirements: RQ-0802, RQ-0806
- Linked Design: DSN-0803, DSN-0805
- Description:
  - 失败消息展示与 trace_id 显示
  - 增加 rerun/重试按钮
- DoD:
  - 失败后可继续操作并重试
- Verify:
  - `npm --prefix web run test -- --run`

Status: completed

## TASK-0805 补充测试与文档

- Linked Requirements: RQ-0807
- Linked Design: DSN-0801, DSN-0803, DSN-0804
- Description:
  - 增加 Agent 页面数据联动测试
  - 更新相关文档（Agent 页面交互说明）
- DoD:
  - 关键交互路径有自动化覆盖
- Verify:
  - `npm --prefix web run test -- --run`

Status: completed
