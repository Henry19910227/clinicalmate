# Infra 層說明

`internal/infra/` 負責管理外部基礎設施的連線（資料庫、快取等），是整個應用中最底層的技術接線層。  
每個子目錄代表一種技術棧，在應用啟動時建立連線並向上層暴露存取物件。

---

## 目錄結構

```
internal/infra/
├── mysql/
│   ├── interface.go  # 定義 Infra interface，暴露 GORM() 等存取方法
│   └── infra.go      # 實作 infra struct 與 New()，負責建立 DB 連線
└── redis/            # 預留，尚未實作
```

---

## 請求流程中的位置

```
config → Infra → Repository → Store → Service → Controller
```

Infra 在啟動時建立連線，`Repository` 層透過 `factory/infra.Factory` 取得連線物件後使用。

---

## 檔案職責

| 檔案 | 職責 |
|------|------|
| `interface.go` | 宣告 `Infra` interface，列出對外暴露的連線存取方法 |
| `infra.go` | 宣告 `infra` struct、實作 `New()`、建立連線並設定連線池 |

---

## 命名規則

### Package

取技術棧名稱，全部小寫。

```
internal/infra/mysql/ → package mysql
internal/infra/redis/ → package redis
```

### Interface 與 struct

| 項目 | 名稱 |
|------|------|
| 對外 interface | `Infra` |
| 內部 struct | `infra`（未匯出） |

### 建構子

`New()` 接收對應的 config interface，連線失敗時直接 `log.Fatalf` 終止啟動。

```go
func New(cfg {tech}Cfg.Config) Infra
```

範例：

```go
func New(cfg mysqlCfg.Config) Infra {
    // 建立連線、設定連線池
    return &infra{db: db}
}
```

### struct 欄位

持有底層連線物件，欄位名取連線物件慣用縮寫。

```go
// MySQL
type infra struct {
    db *gorm.DB
}

// Redis（範例）
type infra struct {
    client *redis.Client
}
```

### import alias

格式：`{tech}Cfg`，camelCase。

```go
import mysqlCfg "clinicalmate/internal/config/mysql"
```

### 對外方法

暴露底層連線物件，方法名取連線套件的慣用名稱，PascalCase。

| 技術棧 | 方法 | 回傳型別 |
|--------|------|----------|
| MySQL | `GORM()` | `*gorm.DB` |
| Redis | `Client()` | `*redis.Client` |

---

## MySQL 連線細節

`mysql.New()` 在建立連線時會一併設定連線池，所有參數來自 `mysqlCfg.Config`：

| Config 方法 | 對應 DSN / 設定 |
|-------------|----------------|
| `Host()` | DSN host |
| `Port()` | DSN port |
| `Database()` | DSN dbname |
| `Username()` | DSN user |
| `Password()` | DSN password |
| `MaxIdleConns()` | `sql.DB.SetMaxIdleConns()` |
| `MaxOpenConns()` | `sql.DB.SetMaxOpenConns()` |

開發環境下以 `db.Debug()` 模式啟動，會將所有 SQL 輸出至 log。

---

## 新增技術棧的步驟

以新增 `redis` 為例：

1. 建立目錄 `internal/infra/redis/`。
2. `interface.go`：宣告 `Infra` interface，暴露 `Client() *redis.Client`。
3. `infra.go`：實作 `infra` struct，`New()` 接收 `redisCfg.Config`，建立連線後回傳。
4. 至 `internal/factory/infra/interface.go` 新增 `RedisInfra() redisInfra.Infra`。
5. 至 `internal/factory/infra/factory.go` 初始化並實作該方法。

---

## 設計原則

- **連線失敗直接終止啟動**（`log.Fatalf`），不做重試或降級，確保應用以健康狀態運行。
- **只負責建立連線與設定連線池**，不執行任何 SQL 或業務操作。
- **連線物件以單例方式共用**，由 factory 層在啟動時建立一次，之後注入所有需要的 repository。
