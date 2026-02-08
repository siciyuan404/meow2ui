# Issue #2 多模态支持 Requirements

## Overview

本规范定义 A2UI 多模态能力的预研到落地路径，目标是在不破坏现有文本主链路的前提下，引入图像/音频输入输出扩展点。

## Scope

- In Scope
  - Provider 与 Model 能力标签扩展（text/image/audio/video）
  - Agent 输入层支持图像/音频引用元数据
  - 会话版本支持多媒体资源引用
  - 多媒体资源存储抽象（local/s3/oss）
  - 首版多模态推理链路（image->ui 描述）
- Out of Scope
  - 实时音视频流处理
  - 端到端多模态生成编辑器
  - 复杂版权和内容审核平台

## Functional Requirements

### RQ-0201 模型能力扩展

- WHEN 管理员注册模型 THEN THE SYSTEM SHALL 支持 `text/image/audio/video` 能力标签与组合能力声明。
- WHEN 路由任务到模型 THEN THE SYSTEM SHALL 基于任务所需能力选择可用模型并在不满足时返回明确错误。

### RQ-0202 多模态输入协议

- WHEN 用户提交图像或音频输入 THEN THE SYSTEM SHALL 接收资源引用（URL/对象存储键）并写入 task context。
- WHEN 输入资源不可访问 THEN THE SYSTEM SHALL 拒绝执行并返回可读错误码。

### RQ-0203 Agent 多模态执行路径

- WHEN 任务包含图像输入 THEN THE SYSTEM SHALL 在 Plan 阶段生成结构化视觉理解摘要并用于后续 Emit。
- WHEN 任务包含音频输入 THEN THE SYSTEM SHALL 支持音频转文本摘要后进入现有文本流程。

### RQ-0204 版本与资源关联

- WHEN 生成新的 schema version THEN THE SYSTEM SHALL 保存关联的多媒体资源引用与处理元数据。
- WHEN 回放历史版本 THEN THE SYSTEM SHALL 返回版本对应的多媒体引用信息。

### RQ-0205 存储抽象

- WHEN 保存多媒体资源 THEN THE SYSTEM SHALL 通过统一存储接口写入 local/s3/oss 任一后端。
- WHEN 切换存储实现 THEN THE SYSTEM SHALL 不修改 Agent 核心流程代码。

## Non-functional Requirements

### RQ-0206 可靠性

- WHEN 多模态模型调用失败 THEN THE SYSTEM SHALL 支持重试或降级回文本流程并保持版本一致性。

### RQ-0207 安全性

- WHEN 处理外部资源引用 THEN THE SYSTEM SHALL 校验来源白名单与访问权限，防止 SSRF 与越权读取。

### RQ-0208 可测试性

- WHEN 提交实现 THEN THE SYSTEM SHALL 覆盖能力路由、资源校验、版本关联、降级流程测试。
