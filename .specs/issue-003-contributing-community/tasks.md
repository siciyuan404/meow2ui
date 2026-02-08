# Issue #3 CONTRIBUTING 与社区建设 Tasks

## TASK-0301 起草 CONTRIBUTING 文档骨架

- Linked Requirements: RQ-0301, RQ-0303, RQ-0307
- Linked Design: DSN-0301, DSN-0303
- Description:
  - 新增 `CONTRIBUTING.md`，覆盖环境、流程、规范、验证
- DoD:
  - 新贡献者可按文档完成本地运行与一次最小提交
- Verify:
  - 人工按文档演练通过

Status: completed

## TASK-0302 补齐 PR/Issue 模板

- Linked Requirements: RQ-0302, RQ-0307
- Linked Design: DSN-0302
- Description:
  - 新增或更新 PR 与 Issue 模板，统一必填字段
- DoD:
  - 模板字段覆盖背景、方案、风险、验证
- Verify:
  - GitHub 新建 PR/Issue 页面可见模板并可用

Status: completed

## TASK-0303 对齐本地验证命令与 CI 文档

- Linked Requirements: RQ-0303, RQ-0306
- Linked Design: DSN-0303
- Description:
  - 统一文档中的构建、测试、lint 命令口径
  - 对齐 `docs/ci.md` 与 `CONTRIBUTING.md`
- DoD:
  - 文档命令不冲突，能在仓库内执行
- Verify:
  - `go test ./...`
  - `npm run lint`

Status: completed

## TASK-0304 建立 ADR 目录与模板

- Linked Requirements: RQ-0304
- Linked Design: DSN-0304
- Description:
  - 新增 `docs/adr/README.md` 与 `docs/adr/0000-template.md`
  - 约定 ADR 编号、状态、评审流程
- DoD:
  - 团队可基于模板新增 ADR
- Verify:
  - 人工使用模板创建示例 ADR 通过

Status: completed

## TASK-0305 补充社区入口与行为准则说明

- Linked Requirements: RQ-0305
- Linked Design: DSN-0305
- Description:
  - 在贡献文档加入 Discussions/聊天渠道/行为准则入口
- DoD:
  - 新用户可在 1 分钟内找到反馈与交流入口
- Verify:
  - 人工检查链接可达

Status: completed

## TASK-0306 回归检查与发布说明

- Linked Requirements: RQ-0306, RQ-0307
- Linked Design: DSN-0301, DSN-0302, DSN-0303
- Description:
  - 执行文档一致性检查并更新变更说明
- DoD:
  - 文档、模板、流程三者一致
- Verify:
  - 人工 walkthrough 完整通过

Status: completed
