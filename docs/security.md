# Security Baseline

## Secrets

- 使用 `auth_ref` 引用环境变量
- 日志中不得输出密钥明文

## Tool Policy

- `read` 低风险
- `write` 中风险
- `exec/network` 默认阻断

## Injection

- 检测常见注入模式
- 高风险输入阻断或净化后再执行
