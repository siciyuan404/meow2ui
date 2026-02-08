# Issue Priority Spec Queue

## Queue Strategy

- Mode: `strict`
- Ordering Rule:
  1. 先处理 `priority:medium`
  2. 在同优先级内按系统影响面与可落地性排序
  3. `priority:low` 仅做预研，不进入本轮执行

## Queue Items

| Queue ID | Source Issue | Spec | Priority | Status | Depends On |
|---|---|---|---|---|---|
| QUEUE-101 | #5 可视化 Agent 调试工具 | `.specs/issue-005-visual-agent-debugger/tasks.md` | medium | completed | QUEUE-015 |
| QUEUE-102 | #1 Flow 编排引擎 | `.specs/issue-001-flow-orchestration/tasks.md` | medium | completed | QUEUE-101 |
| QUEUE-103 | #3 CONTRIBUTING 与社区建设 | `.specs/issue-003-contributing-community/tasks.md` | medium | completed | - |
| QUEUE-201 | #2 多模态支持 | `.specs/issue-002-multimodal-support/tasks.md` | low | completed | QUEUE-102 |
| QUEUE-202 | #4 公开 Benchmark 对比 | `.specs/issue-004-public-benchmark/tasks.md` | low | completed | QUEUE-101, QUEUE-102 |
| QUEUE-203 | #6 模板市场生态建设 | `.specs/issue-006-template-marketplace-ecosystem/tasks.md` | low | completed | QUEUE-103 |

## Notes

- `QUEUE-101` 依赖既有可观测能力（Step15）作为数据底座。
- `QUEUE-102` 先依赖调试可观测能力，降低编排引擎联调风险。
- `QUEUE-103` 可并行推进，但不阻塞运行时能力交付。

## Spec Readiness

- `QUEUE-101`: `requirements.md` / `design.md` / `tasks.md` 已完成，可进入编码。
- `QUEUE-102`: `requirements.md` / `design.md` / `tasks.md` 已完成，等待 `QUEUE-101` 关键能力落地后执行。
- `QUEUE-103`: `requirements.md` / `design.md` / `tasks.md` 已完成，可与 `QUEUE-101` 并行执行。
- `QUEUE-201`: `requirements.md` / `design.md` / `tasks.md` 已完成（low）。
- `QUEUE-202`: `requirements.md` / `design.md` / `tasks.md` 已完成（low）。
- `QUEUE-203`: `requirements.md` / `design.md` / `tasks.md` 已完成（low）。

## Next Window (Low Priority)

1. `QUEUE-201` 多模态支持
2. `QUEUE-202` 公开 Benchmark 对比
3. `QUEUE-203` 模板市场生态建设
