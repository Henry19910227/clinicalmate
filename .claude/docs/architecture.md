## 專案核心架構

**Request flow**: `Router → Middleware → Controller → Service → Store → Repository → DB`

```
cmd/            # Application entrypoint
internal/
  app/
    core/       # App interface + 實作（interface.go + app.go）
  config/
    app/        # App Config interface + 實作（interface.go + config.go）
    mysql/      # MySQL Config interface + 實作（interface.go + config.go）
  infra/        # RDB interface + MySQL 實作
    mysql/
    redis/
  router/       # Gin RouterGroup 設定與路由註冊
  controller/   # HTTP handler（bind request, call service, write response）
  service/      # 業務邏輯（orchestrate use cases, enforce business rules）
  store/        # 跨 repo transaction + domain↔DB model mapping
  repository/   # 單一資料表 GORM 操作
  model/        # GORM model structs 與 domain types
    config/     # Config struct（含 yaml tags）
    app/        # App model
  middleware/   # Gin middleware
  factory/      # 每層的 DI factory
    repository/
    store/
    service/
    controller/
```

## 各層說明文件

| 層 | 說明文件 |
|---|---|
| App | [`internal/app/ARCHITECTURE.md`](../../internal/app/ARCHITECTURE.md) |
| Config | [`internal/config/ARCHITECTURE.md`](../../internal/config/ARCHITECTURE.md) |
| Controller | [`internal/controller/ARCHITECTURE.md`](../../internal/controller/ARCHITECTURE.md) |
| Factory | [`internal/factory/ARCHITECTURE.md`](../../internal/factory/ARCHITECTURE.md) |
| Infra | [`internal/infra/ARCHITECTURE.md`](../../internal/infra/ARCHITECTURE.md) |
| Model | [`internal/model/ARCHITECTURE.md`](../../internal/model/ARCHITECTURE.md) |
| Repository | [`internal/repository/ARCHITECTURE.md`](../../internal/repository/ARCHITECTURE.md) |
| Router | [`internal/router/ARCHITECTURE.md`](../../internal/router/ARCHITECTURE.md) |
| Service | [`internal/service/ARCHITECTURE.md`](../../internal/service/ARCHITECTURE.md) |
| Store | [`internal/store/ARCHITECTURE.md`](../../internal/store/ARCHITECTURE.md) |