# Phase 3: 文档与测试补全 Design

## Overview

本阶段聚焦两个方向：(1) 将 OpenAPI 规范从骨架补全为完整的 API 文档；(2) 将测试覆盖率从 15% 提升到 40%+。两个方向可并行推进。

## Requirement Coverage (RQ -> DSN)

| Requirement ID | Covered By Design IDs | Notes |
|----------------|-----------------------|-------|
| RQ-001 | DSN-001, DSN-002 | OpenAPI 规范补全 |
| RQ-002 | DSN-003 | 后端测试补全 |
| RQ-003 | DSN-004 | 前端测试补全 |
| RQ-004 | DSN-005 | CLI 文档补全 |

---

## Design Index

| Design ID | Element | Type | Notes |
|-----------|---------|------|-------|
| DSN-001 | OpenAPI 端点提取 | workflow | 从代码提取路由 |
| DSN-002 | OpenAPI 规范编写 | data | openapi.v1.yaml |
| DSN-003 | 后端测试补全策略 | component | 按包优先级补测试 |
| DSN-004 | 前端测试补全策略 | component | 关键 hook/store 测试 |
| DSN-005 | CLI 帮助文档增强 | component | 占位命令标注 |

---

## DSN-001: OpenAPI 端点提取

**Type:** workflow

**Purpose:** 从服务器代码中提取所有已注册的 HTTP 路由，作为 OpenAPI 规范的输入。

**Covers Requirements:** RQ-001

**Responsibilities:**
- 扫描 `cmd/server/main.go` 中的 `mux.HandleFunc` / `http.Handle` 调用
- 提取 HTTP 方法、路径、处理函数名
- 与现有 openapi.v1.yaml 对比，标记缺失端点

---

## DSN-002: OpenAPI 规范编写

**Type:** data

**Purpose:** 将 openapi.v1.yaml 从 26 行骨架扩展为完整规范。

**Covers Requirements:** RQ-001

**Responsibilities:**
- 为每个端点定义 path、method、summary、description
- 定义 request body schema（引用 Go struct）
- 定义 response schema（成功 + 错误）
- 定义认证方式（API Key header）
- 定义通用错误模型（ErrorResponse）
- 按功能分组：Agent、Flow、Debug、Marketplace、Enterprise、Ops

---

## DSN-003: 后端测试补全策略

**Type:** component

**Purpose:** 按优先级为关键包补充单元测试。

**Covers Requirements:** RQ-002

**Responsibilities:**
- 优先级 1（核心逻辑）：`pkg/agent`、`pkg/a2ui`、`pkg/flow` — 目标 50%+
- 优先级 2（数据层）：`pkg/store`、`internal/infra/sqlstore` — 目标 40%+
- 优先级 3（安全层）：`pkg/guardrail`、`pkg/security` — 目标 40%+
- 优先级 4（其他）：`pkg/session`、`pkg/provider`、`pkg/debugger` — 目标 30%+

**Key Decisions:**
- 使用 table-driven tests 风格
- 使用 mock provider 避免外部依赖
- 使用 testify/assert 简化断言

---

## DSN-004: 前端测试补全策略

**Type:** component

**Purpose:** 为前端关键路径补充测试。

**Covers Requirements:** RQ-003

**Responsibilities:**
- `use-agent-runtime.ts` — 测试 submit/rerun/save 流程
- Zustand store — 测试状态变更
- Agent 页面组件 — 测试渲染和交互

**Key Decisions:**
- 使用 Vitest + React Testing Library
- Mock fetch/API 调用

---

## DSN-005: CLI 帮助文档增强

**Type:** component

**Purpose:** 为占位命令添加明确的状态标注。

**Covers Requirements:** RQ-004

**Responsibilities:**
- 在 `cmd/cli/main.go` 中为 db:backup、db:restore、data:export、data:import 添加 `[NOT IMPLEMENTED]` 标签
- 执行时输出 "This command is not yet implemented. Planned for v1.1."

---

## Operational Considerations

### Migration / Backward Compatibility
- OpenAPI 规范变更不影响运行时行为
- 测试补全不修改生产代码

### Observability
- CI 覆盖率报告作为 artifact 保存
