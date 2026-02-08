# Contributing Guide

## Scope

This repository contains a Go backend and a React web app. Please keep changes focused, tested, and easy to review.

## Prerequisites

- Go 1.24+
- Node.js 20+
- npm 10+
- Optional: PostgreSQL 16 for integration checks

## Setup

```bash
go mod tidy
npm --prefix web install
```

## Run Locally

Backend (memory mode):

```bash
go run ./cmd/server
```

Web app:

```bash
npm --prefix web run dev
```

## Branch and Commits

- Branch naming: `feature/<topic>`, `fix/<topic>`, `docs/<topic>`
- Commit style: Conventional Commits (e.g. `feat: add debug run API`)

## Pull Request Rules

- Keep PR small and scoped.
- Include background, change summary, risk, and verification.
- Link related issue with `Fixes #<id>` when applicable.

## Required Verification

Backend:

```bash
go test ./... -coverprofile=coverage.out
bash scripts/check-coverage.sh 40
```

Web (if changed):

```bash
npm --prefix web run test -- --run
```

Integration (optional locally, required in CI):

```bash
bash scripts/integration-postgres.sh
bash scripts/smoke.sh
```

## Code Style

- Keep functions focused and deterministic.
- Prefer explicit errors over silent fallback.
- Avoid introducing broad refactors in feature PRs.

## ADR

- Record important architecture decisions in `docs/adr/`.
- Start from `docs/adr/0000-template.md`.

## Community

- Questions and ideas: GitHub Discussions (enable in repository settings if not enabled yet).
- Bug reports and feature requests: GitHub Issues.
- Please follow the repository Code of Conduct when available.
