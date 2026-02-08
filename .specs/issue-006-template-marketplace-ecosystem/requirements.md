# Issue #6 模板市场与生态建设 Requirements

## Overview

本规范定义模板市场生态化能力，目标是在现有 Step13 基础上补齐创作者体验、治理机制和分发能力，形成可持续增长的模板生态。

## Scope

- In Scope
  - 模板分类、标签、搜索、筛选与预览
  - 模板发布与审核工作流
  - 模板评分与评论基础能力
  - 从 session 一键保存模板与版本管理
  - 模板依赖主题与资源打包规则
- Out of Scope
  - 完整交易与支付结算
  - 高级推荐算法与广告系统

## Functional Requirements

### RQ-0601 市场浏览与发现

- WHEN 用户进入模板市场 THEN THE SYSTEM SHALL 支持分类、标签、关键词、排序筛选。
- WHEN 用户查看模板详情 THEN THE SYSTEM SHALL 提供预览图、版本、评分、作者信息。

### RQ-0602 模板发布与审核

- WHEN 创作者发布模板 THEN THE SYSTEM SHALL 支持草稿保存、提交审核、发布上架。
- WHEN 审核不通过 THEN THE SYSTEM SHALL 返回审核意见并保持草稿可编辑。

### RQ-0603 模板评分与评论

- WHEN 用户使用模板 THEN THE SYSTEM SHALL 允许提交评分与评论并关联版本。
- WHEN 评论含违规内容 THEN THE SYSTEM SHALL 支持标记与审核处理。

### RQ-0604 Session 到模板沉淀

- WHEN 用户在 session 中选择保存模板 THEN THE SYSTEM SHALL 将 schema、主题依赖、元数据打包保存。
- WHEN 生成模板新版本 THEN THE SYSTEM SHALL 记录变更摘要并保留历史版本。

### RQ-0605 一键应用与兼容校验

- WHEN 用户应用模板到当前 session THEN THE SYSTEM SHALL 执行兼容校验并生成新版本。
- WHEN 依赖缺失 THEN THE SYSTEM SHALL 返回缺失项并提供修复建议。

## Non-functional Requirements

### RQ-0606 安全与治理

- WHEN 模板上架 THEN THE SYSTEM SHALL 通过内容与依赖扫描策略，拦截高风险模板。

### RQ-0607 可维护性

- WHEN 模板元数据扩展 THEN THE SYSTEM SHALL 通过版本化 schema 保持向后兼容。

### RQ-0608 可测试性

- WHEN 提交实现 THEN THE SYSTEM SHALL 覆盖检索、审核、评分评论、应用流程与兼容校验测试。
