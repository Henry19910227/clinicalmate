# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**Clinicalmate** 是一個陪診系統（Clinical Companion System），使用 Golang 開發，協助病患預約與管理陪診服務。

## Tech Stack

- **Language**: Go 1.25
- **Web Framework**: Gin v1.10
- **ORM**: GORM v1.31 (via `gorm.io/driver/mysql`)
- **Database**: MySQL

## Common Commands

```bash
# Run the application
go run cmd/main.go

# Build
go build -o bin/clinicalmate cmd/main.go

# Run all tests
go test ./...

# Run tests for a specific package
go test ./internal/service/...

# Run a single test
go test ./internal/service/... -run TestFunctionName

# Lint (requires golangci-lint)
golangci-lint run
```

## Architecture

```
cmd/            # Application entrypoint (main.go)
config/         # Configuration loading (env vars: PORT, DB_DSN)
internal/
  database/     # DB connection (RDB interface + MySQL implementation)
  router/       # Gin RouterGroup setup and route registration per domain
  controller/   # Gin HTTP handlers (request binding, status codes, response serialization)
  service/      # Business logic layer
  store/        # Cross-repo transaction coordination and domain↔DB model mapping
  repository/   # GORM database access (single-table operations only)
  model/        # GORM model structs and domain types (currently empty)
  middleware/   # Gin middleware (auth, logging, etc.) (currently empty)
  factory/      # Dependency injection factories for each layer
    repository/ # Constructs and caches repository instances (singleton per factory)
    store/      # Constructs store instances (depends on repository factory)
    service/    # Constructs service instances (depends on store factory)
    controller/ # Constructs controller instances (depends on service factory)
```

**Request flow**: `Router → Middleware → Controller → Service → Store → Repository → DB`

### Layer Responsibilities

| Layer | Responsibility | Must NOT |
|---|---|---|
| **Router** | Register routes on a `*gin.RouterGroup`, call `Set(controllerFactory)` | Contain business logic |
| **Controller** | Bind request, call service, write response | Query DB directly, contain business logic |
| **Service** | Orchestrate use cases, enforce business rules | Know about HTTP or GORM models |
| **Store** | Wrap multi-repo operations in a transaction, map between domain and DB models | Contain business rules |
| **Repository** | Execute single-table GORM queries | Join across tables or manage transactions |

### Database Layer (`internal/database/`)

- `interface.go` — defines `RDB` interface with `Connect() *gorm.DB`
- `msdb.go` — MySQL implementation; `New() RDB` opens the connection via GORM, returns `*ms`

> **Note**: DB credentials are currently hardcoded in `msdb.go`. Migrate to `config.Load()` before production.

## Factory Pattern (Dependency Injection)

All layers are wired via factory interfaces under `internal/factory/`. Factories are composed in `main.go`:

```
repository.Factory   ← holds singleton repository instances (created in New())
  └── store.Factory
        └── service.Factory
              └── controller.Factory
```

Each factory exposes methods like `AdminRepository()`, `AdminStore()`, `AdminService()`, `AdminController()` that return the corresponding interface. Concrete implementations are private structs inside each factory package.

**Wiring order** (current `cmd/main.go`):

```go
rdb      := database.New()                   // opens MySQL connection
repoF    := repoFactory.New(rdb.Connect())   // creates & caches repo instances
storeF   := storeFactory.New(repoF)
serviceF := serviceFactory.New(storeF)
ctrlF    := controllerFactory.New(serviceF)

g := gin.Default()
adminR := adminRouter.New(g.Group("/api/v1"))
adminR.Set(ctrlF)                            // registers routes
_ = g.Run(":8080")
```

## Naming Conventions

- Each domain (e.g., `admin`) has its own sub-package inside every layer:
  `internal/controller/admin/`, `internal/service/admin/`, etc.
- Each sub-package exposes:
  - `interface.go` — defines the public interface (`Controller`, `Service`, `Store`, `Repository`)
  - `<layer>.go` — private struct implementing the interface + `New(...)` constructor
- Router sub-packages follow the same pattern: `interface.go` + `router.go`
  - `interface.go` — defines `Router` interface with `Set(factory controller.Factory)`
  - `router.go` — registers routes on the `*gin.RouterGroup`
- Factory sub-packages follow the same pattern: `interface.go` + `factory.go`

## Current Domains

| Domain | Status |
|---|---|
| `admin` | Scaffold only — interfaces are empty, routes are placeholder |

## Adding a New Domain

To add a new domain (e.g., `patient`):

1. `internal/model/` — add GORM model struct
2. `internal/repository/patient/` — `interface.go` + `repository.go`
3. `internal/store/patient/` — `interface.go` + `store.go`
4. `internal/service/patient/` — `interface.go` + `service.go`
5. `internal/controller/patient/` — `interface.go` + `controller.go`
6. `internal/router/patient/` — `interface.go` + `router.go`
7. Register `PatientRepository()` in `internal/factory/repository/`
8. Register `PatientStore()` in `internal/factory/store/`
9. Register `PatientService()` in `internal/factory/service/`
10. Register `PatientController()` in `internal/factory/controller/`
11. Instantiate and mount the router in `cmd/main.go`

## Configuration

`config/config.go` loads from environment variables:

| Env Var | Default | Description |
|---|---|---|
| `PORT` | `8080` | HTTP listen port |
| `DB_DSN` | _(empty)_ | MySQL DSN (`user:pass@tcp(host:port)/dbname`) |

> **Note**: `config.Load()` is defined but not yet consumed by `main.go` or `database.New()`.
