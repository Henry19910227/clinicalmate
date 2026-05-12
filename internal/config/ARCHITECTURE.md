# Config 層說明

`internal/config/` 負責將原始 config model（YAML 解析結果）包裝成帶有方法的 interface，  
讓下游層（infra、app）透過方法取值，而不直接存取 struct 欄位。

---

## 目錄結構

```
internal/config/
├── app/
│   ├── interface.go  # 定義 Config interface（App 設定方法簽名）
│   └── config.go     # 實作 config struct 與 New()
└── mysql/
    ├── interface.go  # 定義 Config interface（MySQL 設定方法簽名）
    └── config.go     # 實作 config struct 與 New()
```

原始資料來源：`internal/model/config/model.go`（純 struct，對應 `config.yaml`）

---

## 資料流

```
config.yaml
    ↓ yaml.Unmarshal
model.Config          (internal/model/config)
    ↓ factory/config.New()
config.Factory        (internal/factory/config)
    ↓ .AppConfig() / .MysqlConfig()
app.Config / mysql.Config   (internal/config/{tech})
    ↓
infra / app 各層使用
```

---

## 檔案職責

| 檔案 | 職責 |
|------|------|
| `interface.go` | 宣告 `Config` interface，列出所有設定項的 getter 方法 |
| `config.go` | 宣告 `config` struct、實作 `New()`、實作各 getter 方法 |

---

## 命名規則

### Package

取技術棧或用途名稱，全部小寫。

```
internal/config/app/   → package app
internal/config/mysql/ → package mysql
```

### Interface 與 struct

| 項目 | 名稱 |
|------|------|
| 對外 interface | `Config` |
| 內部 struct | `config`（未匯出） |

### 建構子

`New()` 接收對應 model 的指標。

```go
func New(cfg *model.{Name}Config) Config
```

範例：

```go
// internal/config/app/config.go
func New(cfg *model.AppConfig) Config

// internal/config/mysql/config.go
func New(cfg *model.MysqlConfig) Config
```

### struct 欄位

固定為 `cfg`，持有 model 指標。

```go
type config struct {
    cfg *model.AppConfig
}
```

### import alias

固定為 `model`，指向 `internal/model/config`。

```go
import model "clinicalmate/internal/model/config"
```

### Getter 方法

每個 YAML 欄位對應一個 getter，方法名為欄位名的 PascalCase，無參數，回傳對應型別。

```go
func (c *config) Name() string { return c.cfg.Name }
func (c *config) Ip() string   { return c.cfg.Ip }
func (c *config) Port() int    { return c.cfg.Port }
```

---

## 現有設定項

### `app.Config`

對應 `config.yaml` 的 `app` 區塊：

| 方法 | YAML key | 型別 |
|------|----------|------|
| `Name()` | `app.name` | `string` |
| `Ip()` | `app.ip` | `string` |
| `Port()` | `app.port` | `int` |

### `mysql.Config`

對應 `config.yaml` 的 `database` 區塊：

| 方法 | YAML key | 型別 |
|------|----------|------|
| `Host()` | `database.host` | `string` |
| `Port()` | `database.port` | `int` |
| `Database()` | `database.database` | `string` |
| `Username()` | `database.username` | `string` |
| `Password()` | `database.password` | `string` |
| `MaxIdleConns()` | `database.max_idle_conns` | `int` |
| `MaxOpenConns()` | `database.max_open_conns` | `int` |

---

## 新增設定區塊的步驟

以新增 `redis` 設定為例：

1. 至 `internal/model/config/model.go` 新增 `RedisConfig` struct 與對應 YAML tag，並加入 `Config` struct。
2. 至 `config.yaml` 新增對應的 `redis` 區塊。
3. 建立目錄 `internal/config/redis/`，新增 `interface.go` 與 `config.go`。
4. 至 `internal/factory/config/interface.go` 新增 `RedisConfig() redisCfg.Config`。
5. 至 `internal/factory/config/factory.go` 實作該方法。

---

## 設計原則

- **model 只做資料容器**（純 struct + YAML tag），不帶任何方法。
- **config 層只做 getter 封裝**，不做轉換或計算。
- **下游層依賴 `Config` interface 而非 model struct**，方便測試時替換假實作。
