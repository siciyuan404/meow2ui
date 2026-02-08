# Issue #4 公开 Benchmark 与性能对比 Requirements

## Overview

本规范定义公开 benchmark 体系与对比发布流程，目标是可持续量化 A2UI 在延迟、吞吐、成功率、成本等维度的表现，并对回归进行预警。

## Scope

- In Scope
  - 基准场景定义与数据集版本化
  - 基准执行器与结果存储
  - 基线对比与回归判定
  - GitHub Actions 定时执行
  - 报告导出与公开展示（GitHub Pages）
- Out of Scope
  - 商业基准服务平台
  - 自动化对手项目部署编排

## Functional Requirements

### RQ-0401 Benchmark 维度

- WHEN 执行 benchmark THEN THE SYSTEM SHALL 输出 P50/P95/P99 延迟、吞吐、成功率、token 成本、内存占用。
- WHEN 指标采集完成 THEN THE SYSTEM SHALL 同时保存 run 级和 case 级明细。

### RQ-0402 对比对象管理

- WHEN 配置对比对象 THEN THE SYSTEM SHALL 支持记录内部基线与外部对比实现标识。
- WHEN 某对比对象不可用 THEN THE SYSTEM SHALL 标记该对象状态并不中断其他对象执行。

### RQ-0403 回归判定

- WHEN 新 run 完成 THEN THE SYSTEM SHALL 与基线 run 做阈值比较并输出 pass/fail。
- WHEN 指标下降超过阈值 THEN THE SYSTEM SHALL 标记 regression 并返回原因摘要。

### RQ-0404 自动化执行

- WHEN 到达计划时间或触发手动执行 THEN THE SYSTEM SHALL 在 CI 中执行 benchmark 并归档报告。
- WHEN benchmark 失败 THEN THE SYSTEM SHALL 通过 CI 状态暴露失败并保留日志。

### RQ-0405 报告发布

- WHEN benchmark 执行完成 THEN THE SYSTEM SHALL 生成可阅读报告（Markdown/JSON）并发布到公开页面。
- WHEN 查询历史结果 THEN THE SYSTEM SHALL 支持按日期、分支、版本过滤。

## Non-functional Requirements

### RQ-0406 可重复性

- WHEN 相同版本和数据集重复执行 THEN THE SYSTEM SHALL 保持结果波动在可接受范围内并记录环境信息。

### RQ-0407 成本控制

- WHEN 执行公开 benchmark THEN THE SYSTEM SHALL 支持预算上限与样本降采样策略。

### RQ-0408 可测试性

- WHEN 提交 benchmark 逻辑变更 THEN THE SYSTEM SHALL 覆盖指标计算、回归判定、报告生成测试。
