# Model 層說明

`internal/model/` 是純資料容器層，定義應用中所有的 struct。  
不包含任何方法或業務邏輯，只負責資料的結構描述與序列化 tag。

---

## 目錄結構

```
internal/model/
├── config/
│   └── model.go   # config.yaml 對應的 struct（AppConfig、MysqlConfig）
└── {domain}/
    └── model.go   # 業務 domain 的實體 struct（對應資料表）
```

---

## 子目錄分類

| 目錄 | 用途 | Tag |
|------|------|-----|
| `config/` | 對應 `config.yaml` 的設定結構 | `yaml:` |
| `{domain}/` | 對應資料庫資料表的實體結構 | `gorm:` |

---

## 命名規則

### Package

取目錄名，全部小寫。

```
internal/model/config/  → package config
internal/model/admin/   → package admin
```

### Struct 命名

PascalCase，名稱直接描述資料內容。

```go
// config model
type Config struct { ... }
type AppConfig struct { ... }
type MysqlConfig struct { ... }

// domain entity
type Admin struct { ... }
type Patient struct { ... }
```

### 欄位命名

PascalCase，搭配對應的序列化 tag。

**Config model**（對應 YAML）：

```go
type AppConfig struct {
    Name string `yaml:"name"`
    Ip   string `yaml:"ip"`
    Port int    `yaml:"port"`
}
```

**Domain entity**（對應資料表）：

```go
type Admin struct {
    gorm.Model
    Name  string `gorm:"column:name;not null"`
    Email string `gorm:"column:email;uniqueIndex"`
}
```

### 頂層 Config struct

`model/config/model.go` 中的 `Config` 是整份 `config.yaml` 的根節點，  
每個子設定以具名欄位內嵌：

```go
type Config struct {
    App      AppConfig   `yaml:"app"`
    Database MysqlConfig `yaml:"database"`
}
```

新增設定區塊時，在此 struct 加入對應欄位即可。

---

## 現有 Config Model

### `AppConfig`

| 欄位 | YAML key | 型別 |
|------|----------|------|
| `Name` | `app.name` | `string` |
| `Ip` | `app.ip` | `string` |
| `Port` | `app.port` | `int` |

### `MysqlConfig`

| 欄位 | YAML key | 型別 |
|------|----------|------|
| `Driver` | `database.driver` | `string` |
| `Host` | `database.host` | `string` |
| `Port` | `database.port` | `int` |
| `Database` | `database.database` | `string` |
| `Username` | `database.username` | `string` |
| `Password` | `database.password` | `string` |
| `MaxIdleConns` | `database.max_idle_conns` | `int` |
| `MaxOpenConns` | `database.max_open_conns` | `int` |

---

## 新增 Domain Entity 的步驟

以新增 `patient` 為例：

1. 建立 `internal/model/patient/model.go`。
2. 定義 `Patient` struct，嵌入 `gorm.Model`，加上對應的 `gorm:` tag。
3. 在對應的 `repository` 方法中使用此 struct 進行 CRUD。

---

## 設計原則

- **只有 struct 與 tag，不帶任何方法**。
- **不 import 任何內部套件**，避免循環依賴。
- **config model 與 domain entity 分開存放**，兩者的 tag 體系不同（`yaml:` vs `gorm:`）。
