# Changelog

## v1.0.0 (2026-02-08)

### Features

- **Agent Pipeline**: AI-powered UI schema generation with guardrail protection, retrieval-augmented context, schema validation, and auto-repair loop
- **Flow Orchestration**: DAG-based execution engine supporting parallel nodes, conditional edges, fail-fast and continue-on-error policies
- **Multimodal Input**: Image and audio media reference support in agent runs with allowlist-based security validation
- **Visual Debugger**: Run timeline with step-level latency, token usage tracking, and cost breakdown APIs
- **Template Marketplace**: Save templates from sessions, search/filter by category and tags, review workflow (submit/approve/reject), star ratings with comment flagging, one-click apply to session
- **Public Benchmark**: Evaluation framework with orchestrator, runner, and automated markdown/JSON report generation
- **Provider Pool**: Multi-provider routing with capability matching, role-based model selection, 3-attempt retry with backoff, and multimodal-to-text degradation fallback
- **Workspace & Session Management**: Create workspaces, sessions with theme support, schema versioning with parent chain
- **Playground Retrieval**: Theme-aware retrieval-augmented generation for agent context enrichment

### Infrastructure

- **PostgreSQL Persistence**: 14 goose-managed migrations with full rollback support
- **CI/CD Pipeline**: 4-job GitHub Actions — unit tests, Postgres integration, smoke tests, benchmark report
- **Docker**: Multi-stage Dockerfile with dev/staging/prod docker-compose configurations
- **Observability**: Trace ID propagation via `X-Trace-Id` header, 14 event types in agent run timeline, structured error responses with trace context
- **Health & Readiness**: `/healthz`, `/readyz` (with DB connectivity check), `/version` endpoints

### Security

- **Guardrail Service**: Prompt injection detection (4 patterns), tool action policy enforcement (read/write/exec/network)
- **Injection Detection**: 5-pattern scanner (ignore_previous, reveal_system_prompt, run_shell, delete_all_files, prompt_leak)
- **Media Security**: URL allowlist validation, SSRF protection (local address blocking), S3 URI support
- **Secrets Governance**: Auth ref indirection for provider credentials

### Enterprise (Baseline)

- **RBAC**: Organization and project creation, member management
- **Audit**: Export creation and status tracking
- **Cost Governance**: Budget management, usage tracking, cost summary
- **SSO/SCIM**: Endpoint stubs returning 501 Not Implemented (planned for v1.1)

### Documentation

- OpenAPI 3.0.3 specification covering all 27 API endpoints
- Release checklist (30 items), migration report, regression matrix, release gate
- CONTRIBUTING.md with community guidelines
- Release runbook with build, migrate, verify, and rollback instructions
- Architecture Decision Records (ADRs)

### Known Limitations

- Enterprise SSO (`/api/v1/enterprise/sso/config`) returns 501 — planned for v1.1
- Enterprise SCIM (`/api/v1/enterprise/scim/sync`) returns 501 — planned for v1.1
- CLI `db:backup`, `db:restore`, `data:export`, `data:import` are stubs — planned for v1.1
- Frontend pages (Workspaces, Playground, Sessions, Editor, Settings) are placeholder UI — planned for v1.1
- Some API routes use hardcoded IDs (benchmark-runs, projects, audit exports) — will be fully dynamic in v1.1
- Go and TypeScript SDK clients not yet available — planned for v1.1
