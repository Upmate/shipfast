# ShipFast Architecture

## Overview

ShipFast follows a microservices architecture with six independent services communicating via REST APIs and message queues.

## Services

### apps/api (Go)

- **Framework:** chi router
- **Database:** PostgreSQL
- **Responsibilities:** User management, authentication, billing, product catalog
- **Port:** 8080

### apps/web (React/TypeScript)

- **Framework:** Vite + React 19 + TanStack Router
- **Responsibilities:** Admin dashboard, user management UI, analytics
- **Port:** 5173

### apps/worker (Python)

- **Framework:** FastAPI + Celery
- **Responsibilities:** Data ingestion, ETL pipelines, report generation
- **Port:** 8081

### apps/marketing (PHP)

- **Framework:** Laravel
- **Responsibilities:** Landing pages, blog, contact forms
- **Port:** 8082

### services/auth (Rust)

- **Framework:** Actix-web
- **Responsibilities:** JWT validation, session management, token refresh
- **Port:** 8090

### services/notifications (Ruby)

- **Framework:** Sinatra
- **Responsibilities:** Email, SMS, push notification dispatch
- **Port:** 8091

## Infrastructure

- **Cloud:** AWS (EC2, RDS, ElastiCache, S3)
- **Orchestration:** Kubernetes (EKS)
- **IaC:** Terraform
- **CI/CD:** GitHub Actions
- **Containers:** Docker with multi-stage builds

## Communication

```
[Web Browser] --> [apps/web] --> [apps/api] --> [PostgreSQL]
                                     |
                              [services/auth] --> [Redis]
                                     |
                              [apps/worker] --> [S3]
                                     |
                          [services/notifications] --> [SES/SNS]
```
