# Phase 3: 文档与测试补全 Requirements

## Overview

v1 RC 阶段暴露出两个短板：OpenAPI 规范仅为骨架（26 行），测试覆盖率门槛仅 15%。本阶段补全 API 文档并将测试覆盖率提升到可接受水平，为 SDK 开发和后续迭代打下基础。

## Requirement Index

| ID | Title | Priority | Notes |
|----|-------|----------|-------|
| RQ-001 | 补全 OpenAPI 规范 | high | 覆盖所有已注册端点 |
| RQ-002 | 提升后端测试覆盖率 | high | 目标 40%+ |
| RQ-003 | 提升前端测试覆盖率 | medium | 关键路径覆盖 |
| RQ-004 | 补全 CLI 命令文档 | medium | 包含占位命令说明 |

---

## User Stories and Acceptance Criteria

### RQ-001: 补全 OpenAPI 规范

As a API 消费者, I want to 有完整的 OpenAPI 规范, so that 我能基于规范生成客户端代码和文档。

**Acceptance Criteria (EARS):**

**Success Scenarios**
- WHEN 读取 openapi/openapi.v1.yaml THE SYSTEM SHALL 包含所有已注册的 API 端点
- WHEN 每个端点定义 THE SYSTEM SHALL 包含请求/响应 schema、错误码、认证要求
- WHEN 使用 OpenAPI 校验工具 THE SYSTEM SHALL 通过 3.0.3 规范校验

**Error Scenarios**
- WHEN 端点缺少 schema 定义 THE SYSTEM SHALL 至少有 description 和基本类型

**Outcomes to Verify**
- openapi.v1.yaml 端点数量与服务器路由注册数量一致
- 通过 `npx @redocly/cli lint openapi/openapi.v1.yaml` 校验

---

### RQ-002: 提升后端测试覆盖率

As a 开发者, I want to 后端测试覆盖率达到 40%+, so that 核心逻辑有回归保护。

**Acceptance Criteria (EARS):**

**Success Scenarios**
- WHEN 执行 `go test ./... -coverprofile=coverage.out` THE SYSTEM SHALL 总覆盖率 >= 40%
- WHEN 检查关键包覆盖率 THE SYSTEM SHALL pkg/agent、pkg/a2ui、pkg/flow 覆盖率 >= 50%

**Outcomes to Verify**
- CI 覆盖率门槛从 15% 提升到 40%
- 关键包有充分的单元测试

---

### RQ-003: 提升前端测试覆盖率

As a 开发者, I want to 前端关键路径有测试覆盖, so that UI 交互不会因重构而回归。

**Acceptance Criteria (EARS):**

**Success Scenarios**
- WHEN 执行 `npm --prefix web run test -- --coverage` THE SYSTEM SHALL 关键组件有测试
- WHEN 检查 use-agent-runtime.ts THE SYSTEM SHALL 有对应的测试文件

**Outcomes to Verify**
- Agent 页面核心 hook 有测试
- 状态管理 store 有测试

---

### RQ-004: 补全 CLI 命令文档

As a 运维人员, I want to CLI 帮助文档完整, so that 我知道哪些命令是占位实现。

**Acceptance Criteria (EARS):**

**Success Scenarios**
- WHEN 执行 `a2ui-cli --help` THE SYSTEM SHALL 列出所有命令及其状态（可用/占位）
- WHEN 执行占位命令 THE SYSTEM SHALL 输出明确的"未实现"提示而非静默失败

**Outcomes to Verify**
- db:backup、db:restore、data:export、data:import 有明确的占位提示

---

## Non-Functional Requirements

### Performance
- WHEN 执行全量测试 THE SYSTEM SHALL 在 5 分钟内完成

---

## Constraints and Assumptions

### Constraints
- OpenAPI 规范必须与当前代码一致，不能超前定义未实现的端点
- 测试不应依赖外部 LLM 服务（使用 mock provider）

### Assumptions
- 现有 mock provider 足以支撑测试场景

---

## Out of Scope

- Go/TS SDK 开发（延后到 v1.1）
- 性能基准测试文档
- 用户面向的使用手册

---

## Open Questions

| ID | Question | Owner | Needed By |
|----|----------|-------|-----------|
| Q-001 | 是否需要引入 OpenAPI 代码生成工具 | engineering | phase 3 开始前 |
| Q-002 | 覆盖率目标 40% 是否合理 | engineering | phase 3 开始前 |
