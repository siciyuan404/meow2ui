# Phase 1: 队列状态一致性修复 Requirements

## Overview

修复 `.specs/task-queue.md` 和 `.specs/issue-priority-queue.md` 中的状态数据不一致问题。快照区声称 44 个任务待完成，但实际 tasks.md 文件显示全部已完成。QUEUE-301 标记为 pending 但其任务已全部完成，且依赖它的 QUEUE-302/303 反而标记为 completed。

关联 Issue: [#11](https://github.com/siciyuan404/meow2ui/issues/11)

## Requirement Index

| ID | Title | Priority | Notes |
|----|-------|----------|-------|
| RQ-001 | 修复 task-queue.md 快照区过期数据 | high | Current Snapshot / Active Queue Tasks / Next Execution Window / Start Command |
| RQ-002 | 修复 issue-priority-queue.md QUEUE-301 状态 | high | pending -> completed |
| RQ-003 | 清理 issue-priority-queue.md 过期导航区 | medium | Next Window / UX Next Queue 已无意义 |
| RQ-004 | 建立队列状态校验机制 | low | 防止再次出现不一致 |

---

## User Stories and Acceptance Criteria

### RQ-001: 修复 task-queue.md 快照区过期数据

As a 项目维护者, I want to 看到准确的队列快照, so that 我能正确判断项目进度。

**Acceptance Criteria (EARS):**

**Success Scenarios**
- WHEN 读取 task-queue.md 的 Current Snapshot 区域 THE SYSTEM SHALL 显示 Completed Queues 为 QUEUE-001 ~ QUEUE-018，Pending 为 0
- WHEN 读取 task-queue.md 的 Next Execution Window 区域 THE SYSTEM SHALL 显示所有队列已完成
- WHEN 读取 task-queue.md 的 Start Command 区域 THE SYSTEM SHALL 指向下一阶段规划文件

**Outcomes to Verify**
- Current Snapshot 中 Pending Queues 为 (none)
- Active Queue Tasks 区域已归档或移除
- Start Command 不再指向已完成的 QUEUE-013

---

### RQ-002: 修复 issue-priority-queue.md QUEUE-301 状态

As a 项目维护者, I want to QUEUE-301 状态与实际一致, so that 依赖链不再断裂。

**Acceptance Criteria (EARS):**

**Success Scenarios**
- WHEN 读取 issue-priority-queue.md THE SYSTEM SHALL 显示 QUEUE-301 状态为 completed
- WHEN 检查依赖链 THE SYSTEM SHALL 显示 QUEUE-301 -> QUEUE-302 -> QUEUE-303 全部为 completed

**Outcomes to Verify**
- QUEUE-301 Status 列为 completed
- 依赖链无断裂

---

### RQ-003: 清理 issue-priority-queue.md 过期导航区

As a 项目维护者, I want to 过期的 Next Window 和 UX Next Queue 被替换为完成摘要, so that 文档不产生误导。

**Acceptance Criteria (EARS):**

**Success Scenarios**
- WHEN 读取 issue-priority-queue.md 底部 THE SYSTEM SHALL 显示完成摘要而非待执行列表

**Outcomes to Verify**
- Next Window (Low Priority) 区域已替换
- UX Next Queue 区域已替换

---

### RQ-004: 建立队列状态校验机制

As a 项目维护者, I want to 有一个校验脚本能检测队列文件与 tasks.md 的一致性, so that 未来不再出现类似问题。

**Acceptance Criteria (EARS):**

**Success Scenarios**
- WHEN 运行校验脚本 THE SYSTEM SHALL 对比队列文件中的状态与各 tasks.md 中的实际完成情况
- WHEN 发现不一致 THE SYSTEM SHALL 输出具体的不一致项和建议修复

**Outcomes to Verify**
- 脚本可执行且输出清晰
- 对当前已修复的文件运行时输出 "一致"

---

## Non-Functional Requirements

### Reliability
- WHEN 修改队列文件 THE SYSTEM SHALL 保持文件格式与既有约定一致

---

## Constraints and Assumptions

### Constraints
- 仅修改 `.specs/` 下的队列元数据文件，不涉及代码变更

### Assumptions
- 各 tasks.md 中的 completed 状态是真实的（已通过逐文件审计确认）

---

## Out of Scope

- 修改任何 tasks.md 文件内容
- 修改代码或测试

---

## Open Questions

| ID | Question | Owner | Needed By |
|----|----------|-------|-----------|
| Q-001 | 校验脚本是否需要集成到 CI | engineering | phase 4 |
