# A2UI Step15 Observability & Ops Center Requirements

## Overview

本规范定义 A2UI 平台第 15 步：可观测性与运维中心。目标是构建统一运行视图，支持故障定位、容量评估与稳定性治理。

## Scope

- In Scope
  - 指标（Metrics）标准化
  - 日志（Logs）结构化与检索字段规范
  - 分布式追踪（Tracing）基础接入
  - 运维中心 API（健康、异常、容量）
  - 告警规则模板（错误率/延迟/队列积压）
- Out of Scope
  - 商业化 APM 平台深度集成
  - 全自动根因分析（RCA AI）

## Functional Requirements

### RQ-1501 指标标准化

- WHEN 服务运行 THEN THE SYSTEM SHALL 暴露核心指标：QPS、错误率、P95/P99 延迟、agent 成功率、provider 失败率。
- WHEN 指标采集 THEN THE SYSTEM SHALL 区分全局维度与关键标签（provider/model/task）。

### RQ-1502 结构化日志

- WHEN 记录日志 THEN THE SYSTEM SHALL 采用结构化格式并包含 trace_id、run_id、session_id。
- WHEN 发生错误 THEN THE SYSTEM SHALL 输出错误码、错误类别与调用阶段。

### RQ-1503 追踪链路

- WHEN Agent 请求执行 THEN THE SYSTEM SHALL 形成从 API -> agent -> provider -> store 的链路追踪。
- WHEN 查询 trace THEN THE SYSTEM SHALL 可定位慢步骤与失败节点。

### RQ-1504 运维中心查询

- WHEN 调用 ops API THEN THE SYSTEM SHALL 返回当前健康态、最近错误分布、关键容量指标。
- WHEN 某项依赖异常 THEN THE SYSTEM SHALL 标注 degraded 状态。

### RQ-1505 告警模板

- WHEN 指标触发阈值 THEN THE SYSTEM SHALL 产生告警事件并记录。
- WHEN 告警恢复 THEN THE SYSTEM SHALL 记录恢复时间与持续时长。

## Non-functional Requirements

### RQ-1506 性能开销

- WHEN 开启可观测性 THEN THE SYSTEM SHALL 将额外开销控制在可接受范围（请求路径额外开销 < 5%）。

### RQ-1507 可维护性

- WHEN 新增服务模块 THEN THE SYSTEM SHALL 可通过统一埋点接口扩展观测数据。

### RQ-1508 可测试性

- WHEN 提交实现 THEN THE SYSTEM SHALL 覆盖指标输出、日志字段、追踪链路与告警判定测试。
