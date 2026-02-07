# A2UI Step7 Security & Secrets Governance Tasks

## TASK-701 抽象密钥提供器并替换 provider 读取方式

- Linked Requirements: RQ-701, RQ-705
- Linked Design: DSN-701, DSN-706
- Description:
  - 新增 `SecretProvider` 与 `EnvSecretProvider`
  - OpenAI 兼容 adapter 改为依赖 SecretProvider
- DoD:
  - 无明文密钥日志
  - 缺失密钥时返回标准错误
- Verify:
  - `go test ./pkg/provider/...`

Status: completed

## TASK-702 实现策略引擎并接入工具调用

- Linked Requirements: RQ-702, RQ-706
- Linked Design: DSN-702, DSN-706
- Description:
  - 新增 `PolicyEngine` 与默认策略
  - 工具执行链接入决策结果
- DoD:
  - 高风险动作可被阻断
  - 输出规则 ID 与原因
- Verify:
  - `go test ./pkg/security/policy/...`

Status: completed

## TASK-703 增强注入检测与上下文净化

- Linked Requirements: RQ-703
- Linked Design: DSN-703, DSN-706
- Description:
  - 增加 prompt/context 检测器
  - 高风险场景阻断，中风险净化后继续
- DoD:
  - 关键注入样例可识别
  - 净化结果可追踪
- Verify:
  - `go test ./pkg/security/injection/...`

Status: completed

## TASK-704 增加安全事件审计能力

- Linked Requirements: RQ-704
- Linked Design: DSN-704
- Description:
  - 在事件系统中新增安全事件写入接口
  - 统一安全事件结构
- DoD:
  - 安全决策都可落审计
- Verify:
  - `go test ./pkg/events/...`

Status: completed

## TASK-705 增加基础告警钩子

- Linked Requirements: RQ-704
- Linked Design: DSN-705
- Description:
  - 实现阈值告警触发器
  - 默认日志升级，可选 webhook
- DoD:
  - 可配置阈值与窗口
  - 可观测触发记录
- Verify:
  - `go test ./pkg/security/...`

Status: completed

## TASK-706 安全文档与开发规范补充

- Linked Requirements: RQ-705, RQ-706
- Linked Design: DSN-701, DSN-702, DSN-703
- Description:
  - 新增 `docs/security.md`
  - 新增密钥与日志脱敏规范
- DoD:
  - 文档可指导新开发遵循安全基线
- Verify:
  - 人工审阅 + 按文档抽检代码

Status: completed

## TASK-707 安全测试覆盖

- Linked Requirements: RQ-707
- Linked Design: DSN-701, DSN-702, DSN-703, DSN-704
- Description:
  - 覆盖密钥缺失、策略阻断、注入检测、审计写入
- DoD:
  - 测试覆盖关键风险路径
  - 全量测试通过
- Verify:
  - `go test ./...`

Status: completed
