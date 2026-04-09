# ShipFast

> A modern microservices platform for rapid product delivery

[![Pullminder Health](https://api.pullminder.com/badge/t_51387964e55ce153472693ccc965a930.svg)](https://pullminder.com)

## About

ShipFast is a microservices-based platform demonstrating modern development practices across multiple languages and frameworks. This repository is monitored by [Pullminder](https://pullminder.com) — the verification layer for AI-assisted pull requests.

## Pullminder in Action

### Risk Analysis

<!-- Screenshot: PR detail view showing risk score, findings, and reviewer brief -->

_Screenshot coming soon_

### Dashboard Overview

<!-- Screenshot: Dashboard analytics with risk trends -->

_Screenshot coming soon_

### GitHub Check Runs

<!-- Screenshot: GitHub Check Run result on a PR -->

_Screenshot coming soon_

### AI Reviewer Brief

<!-- Screenshot: PR comment with AI-generated reviewer brief -->

_Screenshot coming soon_

### Policy Enforcement

<!-- Screenshot: Policy configuration page -->

_Screenshot coming soon_

### Slack Alerts

<!-- Screenshot: Slack notification for high-risk PR -->

_Screenshot coming soon_

## Architecture

ShipFast is composed of six services:

| Service                  | Language         | Purpose                                                  |
| ------------------------ | ---------------- | -------------------------------------------------------- |
| `apps/api`               | Go               | REST API backend — handles business logic, auth, billing |
| `apps/web`               | React/TypeScript | Dashboard frontend — user-facing SPA                     |
| `apps/worker`            | Python           | Data pipeline — ETL jobs, background processing          |
| `apps/marketing`         | PHP              | Marketing site — landing pages, blog                     |
| `services/auth`          | Rust             | Auth microservice — token validation, session management |
| `services/notifications` | Ruby             | Notification service — email, SMS, push notifications    |

Infrastructure is managed with Terraform (AWS), Kubernetes manifests, and Docker, with CI/CD via GitHub Actions.

## Demo Pull Requests

These PRs demonstrate how Pullminder analyzes real-world changes:

| PR                                                | Title                                | Risk     | What Pullminder Catches                                                       |
| ------------------------------------------------- | ------------------------------------ | -------- | ----------------------------------------------------------------------------- |
| [#24](https://github.com/Upmate/shipfast/pull/24) | Add user authentication endpoint     | Critical | Leaked AWS keys, JWT secrets, sensitive `auth/` path, insecure patterns       |
| [#32](https://github.com/Upmate/shipfast/pull/32) | Update React dashboard components    | High     | Unsafe HTML rendering, direct DOM manipulation, console.log, disabled ESLint  |
| [#23](https://github.com/Upmate/shipfast/pull/23) | Add data pipeline ETL script         | High     | Shell injection, unsafe deserialization, SQL injection, hardcoded DB password |
| [#25](https://github.com/Upmate/shipfast/pull/25) | Upgrade dependencies across services | Medium   | Dependency changes across 6 lock files simultaneously                         |
| [#27](https://github.com/Upmate/shipfast/pull/27) | Refactor billing module              | High     | 600+ line diff, 20+ files changed, sensitive `billing/` path                  |
| [#26](https://github.com/Upmate/shipfast/pull/26) | Add notification templates           | Medium   | AI-generated code (Co-authored-by: Cursor, cursor/ branch, .cursorrules)      |
| [#28](https://github.com/Upmate/shipfast/pull/28) | Add user registration with tests     | Medium   | New Go handlers with no corresponding test files                              |
| [#31](https://github.com/Upmate/shipfast/pull/31) | Update CI/CD pipeline and Terraform  | High     | GitHub Actions workflows, Terraform IaC, Kubernetes RBAC                      |
| [#29](https://github.com/Upmate/shipfast/pull/29) | Fix login page styling               | Low      | Clean PR — small CSS-only change, passes all checks                           |
| [#30](https://github.com/Upmate/shipfast/pull/30) | Add Redis caching layer              | Medium   | Hardcoded Redis connection string, new Cargo.lock dependency                  |

## Analyzer Coverage

Every Pullminder analyzer is triggered by at least one demo PR:

| Analyzer                 | What It Detects                                         | Triggered By          |
| ------------------------ | ------------------------------------------------------- | --------------------- |
| Secrets Detection        | AWS keys, API tokens, private keys, connection strings  | PR #24                |
| Insecure Patterns        | SQL injection, shell injection, unsafe HTML             | PR #24, #32, #23, #30 |
| Sensitive Paths          | Changes to auth/, billing/, security/, migrations/      | PR #24, #27           |
| AI Detection             | Co-author trailers, AI branch prefixes, AI config files | PR #26                |
| Test Gap                 | Source files without corresponding test files           | PR #28                |
| Diff Size                | Large PRs (500+ lines added/removed)                    | PR #27                |
| Files Changed            | PRs touching many files (15+)                           | PR #25, #27           |
| Dependency Changes       | Lock file and manifest updates                          | PR #25, #30           |
| Config/Permissions       | CI/CD workflows, IaC, RBAC, CODEOWNERS                  | PR #31                |
| Framework Rules (React)  | React-specific anti-patterns and security issues        | PR #32                |
| Framework Rules (Go)     | Go-specific anti-patterns                               | PR #24                |
| Framework Rules (Python) | Python-specific anti-patterns                           | PR #23                |

## Features Demonstrated

- [x] 9 concurrent risk analyzers
- [x] Framework-specific rule packs (React, Go, Python)
- [x] AI-generated code detection
- [x] Policy enforcement (block/warn)
- [x] AI reviewer briefs (Claude-powered)
- [x] GitHub Check Runs
- [x] PR comments with findings
- [x] Public health badges
- [x] Multi-language support (Go, TypeScript, Python, PHP, Rust, Ruby)
- [x] Dependency analysis across ecosystems

## Getting Started

```bash
git clone https://github.com/Upmate/shipfast.git

# Each service has its own setup — see docs/architecture.md
```

## License

MIT
