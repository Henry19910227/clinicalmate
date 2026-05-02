# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**Clinicalmate** 是一個陪診系統（Clinical Companion System），使用 Golang 開發，協助病患預約與管理陪診服務。

## Tech Stack

- **Language**: Go
- **Web Framework**: Gin
- **ORM**: GORM
- **Database**: MySQL (via `gorm.io/driver/mysql`)

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

Standard layered Go project layout:

```
cmd/            # Application entrypoint (main.go)
internal/
  handler/      # Gin HTTP handlers (request/response binding)
  service/      # Business logic layer
  repository/   # GORM database access layer
  model/        # GORM model structs
  middleware/   # Gin middleware (auth, logging, etc.)
  router/       # Route registration
config/         # Configuration loading (env/yaml)
```

**Request flow**: `Router → Middleware → Handler → Service → Repository → DB`

Handlers should only handle HTTP concerns (binding, status codes). Business logic lives in services. Repositories own all DB queries.
