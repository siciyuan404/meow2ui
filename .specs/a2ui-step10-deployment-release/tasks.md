# A2UI Step10 Deployment & Release Tasks

## TASK-1001 增加后端 Dockerfile（多阶段）

- Linked Requirements: RQ-1001, RQ-1007
- Linked Design: DSN-1001
- Description:
  - 新增后端 Dockerfile
  - 使用多阶段构建与轻量运行镜像
- DoD:
  - 镜像可构建并启动
  - 不包含敏感配置
- Verify:
  - `docker build -t a2ui-backend:local .`

Status: completed

## TASK-1002 增加部署环境模板与 compose

- Linked Requirements: RQ-1002
- Linked Design: DSN-1002
- Description:
  - 增加 dev/staging env 示例
  - 增加 compose 文件用于本地/测试部署
- DoD:
  - 可通过 env + compose 启动完整服务
- Verify:
  - `docker compose -f deploy/compose/docker-compose.dev.yml up`

Status: completed

## TASK-1003 迁移流程脚本化

- Linked Requirements: RQ-1003
- Linked Design: DSN-1003
- Description:
  - 新增部署脚本：先迁移再启动
  - 支持失败即退出
- DoD:
  - 迁移失败时不继续部署
- Verify:
  - 人工模拟迁移失败并确认阻断

Status: completed

## TASK-1004 版本信息注入与查询接口

- Linked Requirements: RQ-1004, RQ-1006
- Linked Design: DSN-1004
- Description:
  - 构建时注入版本元信息
  - 暴露 `/version` 或健康接口附带版本字段
- DoD:
  - 可查询当前运行版本与构建信息
- Verify:
  - `curl /version`

Status: completed

## TASK-1005 回滚手册与发布文档

- Linked Requirements: RQ-1004
- Linked Design: DSN-1005
- Description:
  - 新增 `docs/release-runbook.md`
  - 记录回滚步骤与注意事项
- DoD:
  - 文档可指导一次完整回滚演练
- Verify:
  - 人工演练验证

Status: completed

## TASK-1006 发布后自动验收脚本

- Linked Requirements: RQ-1005, RQ-1008
- Linked Design: DSN-1006
- Description:
  - 新增发布验收脚本（healthz/readyz/smoke）
  - 与 CI/CD 集成
- DoD:
  - 发布后能自动判断成功/失败
- Verify:
  - 在测试环境执行脚本通过

Status: completed
