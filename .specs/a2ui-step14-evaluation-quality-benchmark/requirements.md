# A2UI Step14 Evaluation & Quality Benchmark Requirements

## Overview

本规范定义 A2UI 平台第 14 步：建立模型输出评估与质量基准体系。目标是量化 A2UI 生成质量，支持回归检测与模型/提示词迭代。

## Scope

- In Scope
  - A2UI 输出自动评估指标
  - 基准数据集（golden cases）
  - 回归对比报告
  - 评估结果存储与查询
  - 与 CI 集成的质量门禁（可选阈值）
- Out of Scope
  - 人工标注平台
  - 在线 A/B 实验平台

## Functional Requirements

### RQ-1401 自动评估指标

- WHEN Agent 生成 UI JSON/Patch THEN THE SYSTEM SHALL 计算结构正确率、组件合法率、修复次数、最终成功率。
- WHEN 评估失败 THEN THE SYSTEM SHALL 给出失败分类（schema/props/component/timeout）。

### RQ-1402 基准数据集

- WHEN 团队新增基准样例 THEN THE SYSTEM SHALL 支持保存输入 prompt、预期约束、参考输出摘要。
- WHEN 执行 benchmark THEN THE SYSTEM SHALL 对全部样例批量运行并汇总指标。

### RQ-1403 回归对比

- WHEN 新模型或新提示词上线前 THEN THE SYSTEM SHALL 对比当前基线与候选版本结果差异。
- WHEN 指标下降超过阈值 THEN THE SYSTEM SHALL 标记为回归并阻断发布候选。

### RQ-1404 评估结果查询

- WHEN 查询评估结果 THEN THE SYSTEM SHALL 返回 run 级与 case 级结果（含错误类别、耗时、token、通过/失败）。

## Non-functional Requirements

### RQ-1405 可维护性

- WHEN 新增指标 THEN THE SYSTEM SHALL 支持在不改核心执行器的情况下扩展计算模块。

### RQ-1406 可测试性

- WHEN 提交评估改动 THEN THE SYSTEM SHALL 覆盖指标计算、回归判定、报告生成测试。
