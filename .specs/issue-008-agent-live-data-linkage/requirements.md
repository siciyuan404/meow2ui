# Issue #8 Agent 页面真实数据联动 Requirements

## Overview

本规范定义 `/agent` 页面从静态骨架升级为真实可用工作台的需求，目标是打通输入、运行、监控、错误反馈全链路。

## Scope

- In Scope
  - `/agent` 对话输入接入后端运行接口
  - 发送中/成功/失败状态可视化
  - 右侧 Monitor 绑定真实运行数据（tokens/thinking/files/output）
  - 与 debug run 详情联动
- Out of Scope
  - 多会话并行执行调度
  - 高级可视化图表面板

## Functional Requirements

### RQ-0801 对话输入与运行触发

- WHEN 用户在 `/agent` 输入 prompt 并提交 THEN THE SYSTEM SHALL 调用后端运行接口并创建一次 run。
- WHEN 用户未输入内容提交 THEN THE SYSTEM SHALL 阻止请求并提示必填。

### RQ-0802 运行状态反馈

- WHEN run 正在执行 THEN THE SYSTEM SHALL 显示发送中状态并禁用重复提交。
- WHEN run 执行完成 THEN THE SYSTEM SHALL 在对话区追加成功消息与结果摘要。
- WHEN run 执行失败 THEN THE SYSTEM SHALL 展示错误信息与 trace_id（若可用）。

### RQ-0803 Monitor 数据联动

- WHEN run 完成或失败 THEN THE SYSTEM SHALL 更新右侧 Monitor 的 token、thinking、files、output 数据。
- WHEN 用户切换 run THEN THE SYSTEM SHALL 刷新 Monitor 视图到该 run 的详情数据。

### RQ-0804 Debug 联动

- WHEN run 完成 THEN THE SYSTEM SHALL 支持基于 run_id 拉取 debug 详情并展示关键字段。

## Non-functional Requirements

### RQ-0805 交互性能

- WHEN 用户提交 prompt THEN THE SYSTEM SHALL 在 200ms 内进入 loading 反馈状态。

### RQ-0806 稳定性

- WHEN 任一接口异常 THEN THE SYSTEM SHALL 保持页面可继续操作，不进入不可恢复状态。

### RQ-0807 可测试性

- WHEN 提交实现 THEN THE SYSTEM SHALL 覆盖提交流程、状态切换、错误展示、monitor 更新的前端测试。
