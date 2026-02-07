# A2UI Step7 Security & Secrets Governance Design

## Design Goals

- 将当前基础 guardrail 升级为策略化安全控制。
- 将密钥访问、工具权限、注入防护纳入统一审计。
- 保持与现有 provider/agent/guardrail 最小侵入集成。

## Architecture

### DSN-701 Secret Provider 抽象

新增 `pkg/security/secrets`：

- `SecretProvider` 接口
  - `Get(ctx, ref string) (string, error)`

- `EnvSecretProvider` 默认实现
  - 从环境变量读取
  - 统一返回脱敏错误

Provider adapter 不再直接 `os.Getenv(auth_ref)`，改为依赖 `SecretProvider`。

### DSN-702 策略引擎

新增 `pkg/security/policy`：

- `PolicyEngine`
  - 输入：`ToolAction`、会话上下文、全局策略
  - 输出：`Decision{allowed,risk,reason,rule_id}`

策略源：
- 静态配置（YAML/JSON）
- 默认内置策略（可覆盖）

### DSN-703 注入防护器

新增 `pkg/security/injection`：

- `Detector`
  - `DetectPrompt(input) DetectionResult`
  - `DetectContext(context) DetectionResult`

结果包含：
- `severity`
- `patterns`
- `sanitized_text`

### DSN-704 安全审计事件模型

新增安全事件类型：

- `SECURITY_SECRET_ACCESS_FAILED`
- `SECURITY_POLICY_BLOCKED`
- `SECURITY_PROMPT_INJECTION_BLOCKED`
- `SECURITY_CONTEXT_SANITIZED`

字段：
- `event_id`, `trace_id`, `run_id`, `event_type`, `severity`, `payload`, `created_at`

### DSN-705 告警钩子

基础版告警：

- 阈值规则（例如 5 分钟内 `SECURITY_POLICY_BLOCKED` >= 10）
- 触发动作：
  - 提升日志级别
  - 可选 webhook 通知

### DSN-706 集成点

- `provider/openai_compatible_adapter.go`
  - 改为通过 `SecretProvider` 取密钥
- `agent/service.go`
  - 执行前调用注入检测
  - 工具动作通过 `PolicyEngine`
- `events/service.go`
  - 增加安全事件写入入口

## Requirement Coverage

- RQ-701 -> DSN-701
- RQ-702 -> DSN-702, DSN-706
- RQ-703 -> DSN-703, DSN-706
- RQ-704 -> DSN-704, DSN-705
- RQ-705 -> DSN-701, DSN-704
- RQ-706 -> DSN-702
- RQ-707 -> DSN-701, DSN-702, DSN-703, DSN-704
