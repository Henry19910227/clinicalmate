# CLAUDE.md

Clinicalmate 的 AI agent 著陸頁。回答「是什麼、怎麼跑、怎麼驗證」。

## 是什麼

**Clinicalmate** 是陪診系統，使用 Golang 開發，協助病患預約與管理陪診服務。

- **語言**：Go 1.25
- **Web 框架**：Gin v1.10
- **ORM**：GORM v1.31（`gorm.io/driver/mysql`）
- **資料庫**：MySQL

## 怎麼跑

```bash
# 啟動應用（需先有 config.yaml）
go run cmd/main.go

# 指定 config 檔
go run cmd/main.go -config config.yaml

# 建置
go build -o bin/clinicalmate cmd/main.go
```

## 怎麼驗證

```bash
# 執行所有測試
go test ./...

# 執行特定 package
go test ./internal/service/...

# 執行單一測試
go test ./internal/service/... -run TestFunctionName

# Lint
golangci-lint run
```

## 架構一覽

**請求流程**：`Router → Middleware → Controller → Service → Store → Repository → DB`

```
cmd/            # 應用進入點（main.go）
config.yaml     # 應用設定（app + database）
config/app/     # Config struct 與載入邏輯
internal/       # 核心業務邏輯
```

- App 頂層協調 → [`internal/app/ARCHITECTURE.md`](internal/app/ARCHITECTURE.md)
- Config 載入 → [`internal/config/ARCHITECTURE.md`](internal/config/ARCHITECTURE.md)
- Controller 層 → [`internal/controller/ARCHITECTURE.md`](internal/controller/ARCHITECTURE.md)
- Factory / DI 接線 → [`internal/factory/ARCHITECTURE.md`](internal/factory/ARCHITECTURE.md)
- Infra 基礎設施 → [`internal/infra/ARCHITECTURE.md`](internal/infra/ARCHITECTURE.md)
- Model 資料結構 → [`internal/model/ARCHITECTURE.md`](internal/model/ARCHITECTURE.md)
- Repository 層 → [`internal/repository/ARCHITECTURE.md`](internal/repository/ARCHITECTURE.md)
- Router 路由 → [`internal/router/ARCHITECTURE.md`](internal/router/ARCHITECTURE.md)
- Service 層 → [`internal/service/ARCHITECTURE.md`](internal/service/ARCHITECTURE.md)
- Store 層 → [`internal/store/ARCHITECTURE.md`](internal/store/ARCHITECTURE.md)
- DB 硬約束 → [`internal/database/CONSTRAINTS.md`](internal/database/CONSTRAINTS.md)

## 當前進度

見 [`PROGRESS.md`](PROGRESS.md)
