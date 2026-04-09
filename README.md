# ShipFast

> A modern microservices platform for rapid product delivery

<!-- Pullminder Health Badge — will be added after enabling -->
<!-- [![Pullminder Health](BADGE_URL)](https://pullminder.com) -->

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

| PR  | Title                                | Risk     | What Pullminder Catches                                       |
| --- | ------------------------------------ | -------- | ------------------------------------------------------------- |
| #1  | Add user authentication endpoint     | Critical | Leaked secrets, sensitive path changes, insecure patterns     |
| #2  | Update React dashboard components    | High     | Unsafe HTML rendering, React anti-patterns                    |
| #3  | Add data pipeline ETL script         | High     | Shell injection, unsafe deserialization, Python anti-patterns |
| #4  | Upgrade dependencies across services | Medium   | Dependency changes across 6 lock files                        |
| #5  | Refactor billing module              | High     | Large diff (600+ lines), 20+ files, sensitive billing path    |
| #6  | Add notification templates           | Medium   | AI-generated code detection                                   |
| #7  | Add user registration                | Medium   | Missing test coverage for new handlers                        |
| #8  | Update CI/CD & Terraform             | High     | CI/CD workflow changes, infrastructure modifications          |
| #9  | Fix login page styling               | Low      | Clean PR — passes all checks                                  |
| #10 | Add Redis caching layer              | Medium   | Hardcoded connection strings, dependency changes              |

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
# Clone the repo
git clone https://github.com/Upmate/shipfast.git

# Each service has its own setup — see docs/architecture.md
```

## License

MIT
