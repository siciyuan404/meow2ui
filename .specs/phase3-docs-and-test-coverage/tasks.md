# Phase 3: 文档与测试补全 Implementation Tasks

## Progress Summary

- Total Tasks: 7
- Completed: 7
- In Progress: 0
- Pending: 0

---

## TASK-P3-001: 提取服务器路由清单

**Status:** completed

**Type:** implementation

**Traceability:**
- Requirements: RQ-001
- Design: DSN-001

**Description:**
从 `cmd/server/main.go` 提取所有已注册的 HTTP 路由，生成端点清单。

**Definition of Done:**
- 输出完整的路由列表（方法 + 路径 + 处理函数）
- 与现有 openapi.v1.yaml 对比，标记缺失端点

**Dependencies:** None

**Verification:**
- Command: `grep -c "HandleFunc\|Handle(" cmd/server/main.go`
- Expected: 匹配到所有路由注册

---

## TASK-P3-002: 补全 OpenAPI 规范

**Status:** completed

**Type:** docs

**Traceability:**
- Requirements: RQ-001
- Design: DSN-002

**Description:**
基于路由清单，将 openapi.v1.yaml 从骨架扩展为完整规范。

**Definition of Done:**
- 所有已注册端点有对应的 path 定义
- 每个端点有 request/response schema
- 通过 OpenAPI 3.0.3 校验

**Dependencies:** TASK-P3-001

**Verification:**
- Command: `npx @redocly/cli lint openapi/openapi.v1.yaml`
- Expected: 无错误

---

## TASK-P3-003: 补全后端核心包测试（agent/a2ui/flow）

**Status:** completed

**Type:** test

**Traceability:**
- Requirements: RQ-002
- Design: DSN-003

**Description:**
为 pkg/agent、pkg/a2ui、pkg/flow 三个核心包补充单元测试，目标覆盖率 50%+。

**Definition of Done:**
- 每个包有 _test.go 文件
- 核心函数有 table-driven tests
- 三个包覆盖率均 >= 50%

**Dependencies:** None

**Parallelizable:** yes（与 TASK-P3-001/002 并行）

**Verification:**
- Command: `go test ./pkg/agent/... ./pkg/a2ui/... ./pkg/flow/... -coverprofile=core.out`
- Expected: 覆盖率 >= 50%

---

## TASK-P3-004: 补全后端数据层与安全层测试

**Status:** completed

**Type:** test

**Traceability:**
- Requirements: RQ-002
- Design: DSN-003

**Description:**
为 pkg/store、internal/infra/sqlstore、pkg/guardrail、pkg/security 补充测试。

**Definition of Done:**
- 数据层和安全层覆盖率 >= 40%

**Dependencies:** None

**Parallelizable:** yes

**Verification:**
- Command: `go test ./pkg/store/... ./pkg/guardrail/... ./pkg/security/... -coverprofile=infra.out`
- Expected: 覆盖率 >= 40%

---

## TASK-P3-005: 补全前端关键路径测试

**Status:** completed

**Type:** test

**Traceability:**
- Requirements: RQ-003
- Design: DSN-004

**Description:**
为 use-agent-runtime.ts、Zustand store、Agent 页面组件补充测试。

**Definition of Done:**
- use-agent-runtime.ts 有对应测试
- store 有状态变更测试
- Agent 页面有渲染测试

**Dependencies:** None

**Parallelizable:** yes

**Verification:**
- Command: `npm --prefix web run test -- --coverage`
- Expected: 关键文件有覆盖

---

## TASK-P3-006: 增强 CLI 占位命令提示

**Status:** completed

**Type:** implementation

**Traceability:**
- Requirements: RQ-004
- Design: DSN-005

**Description:**
为 db:backup、db:restore、data:export、data:import 添加明确的未实现提示。

**Definition of Done:**
- 执行占位命令时输出 "Not yet implemented. Planned for v1.1."
- `--help` 输出中标注 `[NOT IMPLEMENTED]`

**Dependencies:** None

**Verification:**
- Command: `go run ./cmd/cli db:backup`
- Expected: 输出包含 "not yet implemented"

---

## TASK-P3-007: 更新 CI 覆盖率门槛

**Status:** completed

**Type:** implementation

**Traceability:**
- Requirements: RQ-002
- Design: DSN-003

**Description:**
将 CI 覆盖率检查脚本中的门槛从 15% 提升到 40%。

**Definition of Done:**
- scripts/coverage-check.sh 中的阈值更新为 40
- CI 流水线通过

**Dependencies:** TASK-P3-003, TASK-P3-004

**Verification:**
- Command: `grep "threshold\|THRESHOLD\|MIN_COVERAGE" scripts/coverage-check.sh`
- Expected: 值为 40

---

## Execution Notes

- TASK-P3-001 -> TASK-P3-002 串行（依赖路由清单）
- TASK-P3-003、TASK-P3-004、TASK-P3-005、TASK-P3-006 可并行
- TASK-P3-007 依赖 TASK-P3-003 和 TASK-P3-004（确保覆盖率达标后再提升门槛）
