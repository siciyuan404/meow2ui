# Issue #6 模板市场与生态建设 Design

## Design Goals

- 在 Step13 基础上扩展“发现-发布-治理-复用”闭环。
- 优先保证模板可复用与安全上架，再扩展创作者激励。
- 与现有模板/插件数据模型保持兼容演进。

## Architecture

### DSN-0601 市场索引与检索

- 扩展 `templates` 索引字段：
  - `popularity_score`, `downloads`, `rating_avg`, `rating_count`
- 新增检索查询层支持分类、标签、关键词、排序。

### DSN-0602 发布审核工作流

- 扩展状态机：`draft -> submitted -> reviewed -> published`，异常 `blocked`。
- 审核记录表：`template_reviews`
  - `template_id`, `reviewer_id`, `decision`, `note`, `created_at`

### DSN-0603 评分评论系统

- 新增表：
  - `template_ratings`（user_id, template_id, version, score）
  - `template_comments`（user_id, template_id, content, status）
- 支持评论状态：`visible`, `flagged`, `hidden`。

### DSN-0604 打包与版本策略

- 新增模板打包器：
  - schema snapshot
  - theme token snapshot
  - 依赖资源清单
- `template_versions` 增加 `manifest_json` 与 `changelog`。

### DSN-0605 应用与兼容校验

- 应用模板前执行检查：
  - 组件兼容
  - 主题依赖
  - 资源可达性
- 失败返回结构化缺失列表。

### DSN-0606 Web 市场体验增强

- 扩展 `web` 市场页：
  - 排序/筛选
  - 详情页评分评论
  - 发布与审核状态展示

## Requirement Coverage

- RQ-0601 -> DSN-0601, DSN-0606
- RQ-0602 -> DSN-0602
- RQ-0603 -> DSN-0603
- RQ-0604 -> DSN-0604
- RQ-0605 -> DSN-0605
- RQ-0606 -> DSN-0602, DSN-0605
- RQ-0607 -> DSN-0604
- RQ-0608 -> DSN-0601, DSN-0602, DSN-0603, DSN-0605
