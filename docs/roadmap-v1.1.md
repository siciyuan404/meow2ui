# Roadmap v1.1

## Overview

v1.1 包含从 v1.0.0 延后的工作项。优先级基于用户价值和技术依赖排序。

## Deferred Items

| # | Item | Priority | Estimate | Dependencies | Notes |
|---|------|----------|----------|--------------|-------|
| 1 | Go SDK 客户端 | high | 1 week | OpenAPI spec | 基于 openapi.v1.yaml 自动生成 |
| 2 | TypeScript SDK 客户端 | high | 1 week | OpenAPI spec | 基于 openapi.v1.yaml 自动生成 |
| 3 | Canvas 预览真实渲染器 | high | 2 weeks | - | 将 UISchema 渲染为可交互的 Canvas 预览 |
| 4 | 企业 SSO 实现 | medium | 2 weeks | Auth foundation | 替换 `/api/v1/enterprise/sso/config` 的 501 stub |
| 5 | 企业 SCIM 同步实现 | medium | 1 week | SSO | 替换 `/api/v1/enterprise/scim/sync` 的 501 stub |
| 6 | CLI db:backup 真实实现 | medium | 3 days | PostgreSQL | pg_dump wrapper + 存储管理 |
| 7 | CLI db:restore 真实实现 | medium | 3 days | db:backup | pg_restore wrapper + 验证 |
| 8 | CLI data:export 真实实现 | medium | 3 days | Store layer | Workspace 数据导出为 JSON bundle |
| 9 | CLI data:import 真实实现 | medium | 3 days | data:export | JSON bundle 导入到目标 workspace |
| 10 | Workspaces 页面功能化 | medium | 1 week | Web app | 列表、创建、切换 workspace |
| 11 | Playground 页面功能化 | medium | 1 week | Web app | 浏览、搜索、预览 playground items |
| 12 | Sessions 页面功能化 | low | 3 days | Web app | Session 列表、历史版本浏览 |
| 13 | Editor 页面功能化 | low | 1 week | Canvas renderer | 实时编辑 + 预览 |
| 14 | Settings 页面功能化 | low | 3 days | Web app | Provider 配置、主题管理 |
| 15 | API 路由动态化 | low | 3 days | - | 替换 benchmark-runs/projects/audit 中的硬编码 ID |

## Timeline

- **Sprint 1 (Week 1-2):** Go SDK + TS SDK + Canvas 渲染器启动
- **Sprint 2 (Week 3-4):** SSO/SCIM + CLI backup/restore
- **Sprint 3 (Week 5-6):** 前端页面功能化 + data:export/import
- **Sprint 4 (Week 7-8):** Editor + Settings + API 动态化 + 收尾

## Success Criteria

- SDK 可通过 `npm install` / `go get` 安装并调用所有 v1 API
- SSO/SCIM 端点返回 200 而非 501
- CLI backup/restore 可在 CI 中端到端验证
- 前端页面可完成核心用户流程（创建 workspace -> 创建 session -> agent run -> 查看结果）
- Canvas 渲染器可将 UISchema 渲染为可视化预览
