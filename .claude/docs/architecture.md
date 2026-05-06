## 專案核心架構

**Request flow**: `Router → Middleware → Controller → Service → Store → Repository → DB`

```
cmd/            # Application entrypoint
config/
  app/          # Config interface + 實作（interface.go + config.go）
internal/
  model/
    config/     # Config、AppConfig、DatabaseConfig struct（含 yaml tags）
internal/
  database/     # RDB interface + MySQL 實作
  router/       # Gin RouterGroup 設定與路由註冊
  controller/   # HTTP handler（bind request, call service, write response）
  service/      # 業務邏輯（orchestrate use cases, enforce business rules）
  store/        # 跨 repo transaction + domain↔DB model mapping
  repository/   # 單一資料表 GORM 操作
  model/        # GORM model structs 與 domain types
  middleware/   # Gin middleware
  factory/      # 每層的 DI factory
    repository/
    store/
    service/
    controller/
```