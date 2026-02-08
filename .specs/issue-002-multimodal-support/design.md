# Issue #2 多模态支持 Design

## Design Goals

- 在现有文本 Agent 架构上最小侵入扩展多模态输入。
- 明确模型能力与资源协议，保证路由可解释。
- 先实现离线引用式输入，后续再扩展实时流。

## Architecture

### DSN-0201 能力标签与路由扩展

- 扩展 `domain.Model.Capabilities` 的约定值：`text`, `image`, `audio`, `video`。
- 在 `pkg/provider` 增加任务能力匹配逻辑：
  - `plan_image` -> `image` + `text`
  - `plan_audio` -> `audio` + `text`
  - `emit` -> `text`

### DSN-0202 多模态输入协议

- 新增 `MultimodalInput` 结构：
  - `type`: `image` | `audio`
  - `ref`: URL 或对象存储键
  - `metadata`: mime、duration、size
- 注入 `session.BuildContext` 的 `TaskInput.media`。

### DSN-0203 Agent 执行器扩展

- 在 Agent Run 中增加多模态分支：
  - image -> 调用视觉摘要器节点 -> 输出结构化描述
  - audio -> 调用转写摘要节点 -> 输出文本摘要
- 摘要产物统一转为文本约束输入现有 `Emit/Validate/Apply`。

### DSN-0204 版本与资源索引

- 新增表：`schema_version_assets`
  - `version_id`, `asset_type`, `asset_ref`, `metadata_json`, `created_at`
- 回放版本时联表返回资源引用。

### DSN-0205 资源存储接口

- 新增 `pkg/media/storage` 抽象：
  - `Put`, `Get`, `Delete`, `SignURL`
- 提供 `local` 实现，预留 `s3`/`oss` 适配器。

### DSN-0206 安全策略

- 在 `pkg/security/policy` 增加资源引用白名单校验。
- 禁止内网地址和未授权 bucket/key 访问。

## Requirement Coverage

- RQ-0201 -> DSN-0201
- RQ-0202 -> DSN-0202, DSN-0206
- RQ-0203 -> DSN-0203
- RQ-0204 -> DSN-0204
- RQ-0205 -> DSN-0205
- RQ-0206 -> DSN-0203
- RQ-0207 -> DSN-0206
- RQ-0208 -> DSN-0201, DSN-0203, DSN-0204
