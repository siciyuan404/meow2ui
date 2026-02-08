# Phase 1: 队列状态一致性修复 Design

## Overview

本阶段聚焦于修复两个队列元数据文件中的状态不一致，并建立轻量级校验机制防止复发。

## Requirement Coverage (RQ -> DSN)

| Requirement ID | Covered By Design IDs | Notes |
|----------------|-----------------------|-------|
| RQ-001 | DSN-001 | task-queue.md 四处修复 |
| RQ-002 | DSN-002 | issue-priority-queue.md 状态修复 |
| RQ-003 | DSN-003 | issue-priority-queue.md 导航区清理 |
| RQ-004 | DSN-004 | 校验脚本设计 |

---

## Design Index

| Design ID | Element | Type | Notes |
|-----------|---------|------|-------|
| DSN-001 | task-queue.md 快照修复 | data | 4 处文本替换 |
| DSN-002 | issue-priority-queue.md 状态修复 | data | 1 处状态更新 |
| DSN-003 | issue-priority-queue.md 导航清理 | data | 2 个区域替换 |
| DSN-004 | 队列一致性校验脚本 | component | scripts/check-queue-consistency.sh |

---

## DSN-001: task-queue.md 快照修复

**Type:** data

**Purpose:** 将 Current Snapshot、Next Execution Window、Active Queue Tasks、Start Command 四个区域更新为与实际一致的完成状态。

**Covers Requirements:** RQ-001

**Responsibilities:**
- Current Snapshot: Completed Queues 改为 QUEUE-001 ~ QUEUE-018，Pending 归零
- Next Execution Window: 替换为"所有队列已完成"
- Active Queue Tasks: 替换为归档说明
- Start Command: 指向下一阶段规划文件

---

## DSN-002: issue-priority-queue.md 状态修复

**Type:** data

**Purpose:** 将 QUEUE-301 的 Status 从 pending 更新为 completed。

**Covers Requirements:** RQ-002

**Responsibilities:**
- 修改 Queue Items 表中 QUEUE-301 行的 Status 列

---

## DSN-003: issue-priority-queue.md 导航清理

**Type:** data

**Purpose:** 将已过期的 Next Window 和 UX Next Queue 区域替换为完成摘要。

**Covers Requirements:** RQ-003

**Responsibilities:**
- 删除 Next Window (Low Priority) 区域
- 删除 UX Next Queue 区域
- 添加 Completion Summary 区域

---

## DSN-004: 队列一致性校验脚本

**Type:** component

**Purpose:** 提供一个 shell 脚本，自动对比队列文件中声明的状态与各 tasks.md 中的实际完成情况。

**Covers Requirements:** RQ-004

**Responsibilities:**
- 解析 task-queue.md 中每个 QUEUE 的 Status 和对应 tasks file 路径
- 读取每个 tasks.md，统计 completed/pending/in_progress 数量
- 如果队列标记为 completed 但 tasks.md 中有非 completed 任务，报告不一致
- 如果队列标记为 pending 但 tasks.md 中全部 completed，报告不一致
- 对 issue-priority-queue.md 执行同样的检查

**Risks and Mitigations:**
- Risk: 脚本依赖特定的 markdown 格式
- Mitigation: 使用简单的 grep 模式匹配，容错性优先

---

## Operational Considerations

### Migration / Backward Compatibility
- 纯文档修改，无代码影响

### Rollout / Rollback
- git revert 即可回滚
