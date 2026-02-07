# A2UI Step3 Validation Chain Design

## Design Goals

- 实现可扩展校验链，覆盖 RQ-301..RQ-306。
- 让校验结果具备稳定错误码和精确路径。
- 与当前 `pkg/a2ui`、`pkg/agent` 最小侵入集成。

## Architecture

### DSN-301 校验结果模型升级

在 `pkg/a2ui` 定义结构化错误：

- `ValidationError`
  - `Code string`
  - `Path string`
  - `Message string`

在 `ValidationResult` 增加：

- `Errors []ValidationError`

并保留 `Valid bool` 作为快速判断。

### DSN-302 组件注册表（校验用）

新增内置注册表（最小版本）：

- `ComponentRule`
  - `Name string`
  - `PropTypes map[string]string`（`string|number|boolean|object|array`）

- `Registry`
  - `AllowedTypes map[string]ComponentRule`
  - `IsAllowed(type) bool`
  - `Rule(type) (ComponentRule, bool)`

初始白名单包括：
- `Container`
- `Text`
- `Button`
- `Card`

### DSN-303 校验器拆分

将 `ValidateSchema` 内逻辑拆成三段：

1. `validateRequiredFields`（RQ-303）
2. `validateComponentType`（RQ-301）
3. `validatePropsType`（RQ-302）

每段统一返回 `[]ValidationError`。

### DSN-304 错误码规范

- `A2UI_SCHEMA_REQUIRED_FIELD_MISSING`
- `A2UI_COMPONENT_NOT_ALLOWED`
- `A2UI_PROP_TYPE_INVALID`

路径格式：
- `schema.version`
- `root.type`
- `root.children[1].props.title`

### DSN-305 Agent 错误透传

在 `pkg/agent/service.go` 中：

- 当 `validated.Valid == false` 且修复后仍失败：
  - 返回 `error` 文本中包含结构化错误 JSON
  - 事件 `repair_failed` 记录 `validation_errors`

## Data/Interface Changes

### DSN-306 接口兼容策略

- `a2ui.ValidationResult.Errors` 类型从 `[]string` 变更为 `[]ValidationError`
- 受影响代码：
  - `pkg/agent/service.go`（格式化修复提示与最终错误）
  - `pkg/a2ui/service_test.go`

## Testing Strategy

### DSN-307 单测矩阵

- 正向：合法 schema（白名单组件 + 合法 props）
- 反向：
  - 缺少 version/root.id/root.type
  - 非白名单组件
  - props 类型错误
  - 嵌套子节点错误路径

### DSN-308 回归验证

- `go test ./pkg/a2ui/...`
- `go test ./pkg/agent/...`
- `go test ./...`

## Requirement Coverage

- RQ-301 -> DSN-302, DSN-303, DSN-304
- RQ-302 -> DSN-302, DSN-303, DSN-304
- RQ-303 -> DSN-301, DSN-303, DSN-304
- RQ-304 -> DSN-301, DSN-304, DSN-305
- RQ-305 -> DSN-302, DSN-306
- RQ-306 -> DSN-307, DSN-308
