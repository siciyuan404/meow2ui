# Issue #5 可视化 Agent 调试工具 Requirements

## Overview

本规范定义可视化 Agent 调试工具的首版能力，目标是把现有可观测数据转化为可操作的调试视图，降低排障和迭代成本。

## Scope

- In Scope
  - Agent 执行流程可视化（Plan -> Emit -> Validate -> Repair -> Apply）
  - 单次运行的输入/输出与耗时展示
  - 工具调用链与错误传播路径展示
  - 上下文窗口占用可视化
  - 基于既有 metrics 的成本统计视图
- Out of Scope
  - 新增第三方 APM 深度集成
  - 自动根因分析与自动修复建议
  - 多租户权限模型重构

## Functional Requirements

### RQ-0501 调试会话列表与详情

- WHEN 用户进入调试页面 THEN THE SYSTEM SHALL 展示最近 Agent Run 列表并支持按 session_id、status、时间范围筛选。
- WHEN 用户选择某次 Agent Run THEN THE SYSTEM SHALL 返回该 run 的完整执行详情与步骤摘要。

### RQ-0502 决策流程可视化

- WHEN 展示单次运行详情 THEN THE SYSTEM SHALL 以时间序列展示 Plan/Emit/Validate/Repair/Apply 各步骤状态与耗时。
- WHEN 某步骤失败 THEN THE SYSTEM SHALL 标注失败步骤、错误码与失败原因。

### RQ-0503 工具调用链追踪

- WHEN run 包含工具调用 THEN THE SYSTEM SHALL 展示调用顺序、关键参数摘要、返回状态。
- WHEN 调用失败 THEN THE SYSTEM SHALL 展示错误传播路径与最终失败节点。

### RQ-0504 上下文窗口可视化

- WHEN 展示 run 上下文 THEN THE SYSTEM SHALL 展示 system/session/task 三层上下文的 token 占比与压缩前后变化。
- WHEN 上下文发生压缩 THEN THE SYSTEM SHALL 展示压缩触发原因与压缩结果摘要。

### RQ-0505 成本分析面板

- WHEN 展示成本数据 THEN THE SYSTEM SHALL 按 provider/model 统计 token 消耗、调用次数、估算成本。
- WHEN 用户切换时间范围 THEN THE SYSTEM SHALL 同步刷新成本聚合结果。

## Non-functional Requirements

### RQ-0506 性能

- WHEN 查询单次 run 调试详情 THEN THE SYSTEM SHALL 在 P95 <= 800ms 内返回首屏关键数据（不含前端渲染时间）。

### RQ-0507 安全与隐私

- WHEN 返回调试详情 THEN THE SYSTEM SHALL 对敏感字段执行脱敏（如密钥、凭证、用户隐私片段）。

### RQ-0508 可测试性

- WHEN 提交实现 THEN THE SYSTEM SHALL 覆盖流程图数据组装、工具链聚合、上下文占比计算、成本聚合的测试。
