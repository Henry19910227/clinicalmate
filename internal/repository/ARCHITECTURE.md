# Repository 層說明

`internal/repository/` 是最靠近資料庫的一層，直接持有 `*gorm.DB` 並執行 SQL 操作。  
每個子目錄代表一個 domain，對應資料庫中的一張（或一組相關）資料表。

---

## 目錄結構

```
internal/repository/
├── admin/
│   ├── interface.go    # 定義 Repository interface（DB 操作方法簽名）
│   └── repository.go   # 實作 repository struct 與 New()
└── {domain}/
    ├── interface.go
    └── repository.go
```

---

## 請求流程中的位置

```
Store → Repository → DB
```

Store 持有 repository，透過其方法對資料庫進行讀寫。

---

## 檔案職責

| 檔案 | 職責 |
|------|------|
| `interface.go` | 宣告 `Repository` interface，列出所有對外暴露的 DB 操作方法 |
| `repository.go` | 宣告 `repository` struct、實作 `New()`、實作各 DB 操作方法 |

---

## 命名規則

### Package

取 domain 目錄名，全部小寫。

```
internal/repository/admin/   → package admin
internal/repository/patient/ → package patient
```

### Interface 與 struct

| 項目 | 名稱 |
|------|------|
| 對外 interface | `Repository` |
| 內部 struct | `repository`（未匯出） |

### 建構子

```go
func New(db *gorm.DB) Repository
```

範例：

```go
func New(db *gorm.DB) Repository {
    return &repository{db: db}
}
```

factory 層從 `infraFac.MysqlInfra().GORM()` 取得 `*gorm.DB` 後傳入：

```go
// factory/repository/factory.go
adminR := adminRepository.New(infraFac.MysqlInfra().GORM())
```

### struct 欄位

固定為 `db *gorm.DB`。

```go
type repository struct {
    db *gorm.DB
}
```

### import

直接使用套件路徑，不需要 alias。

```go
import "gorm.io/gorm"
```

### DB 操作方法

格式：動詞 + 資源名，PascalCase，回傳值帶 `error`。

```go
func (r *repository) FindByID(id int64) (*model.Admin, error)           {}
func (r *repository) FindAll(filter *dto.AdminFilter) ([]*model.Admin, error) {}
func (r *repository) Insert(admin *model.Admin) error                   {}
func (r *repository) Update(admin *model.Admin) error                   {}
func (r *repository) Delete(id int64) error                             {}
```

常用動詞：`Find`（查詢單筆）、`FindAll`（查詢列表）、`Insert`、`Update`、`Delete`。  
與 store 層的動詞（`Get`、`List`）有所區分，凸顯此層直接對應 DB 操作。

---

## 新增 Domain 的步驟

以新增 `patient` domain 為例：

1. 建立目錄 `internal/repository/patient/`。
2. `interface.go`：宣告 `Repository` interface 與所有 DB 操作方法。
3. `repository.go`：實作 `repository` struct，`New()` 接收 `db *gorm.DB`。
4. 至 `internal/factory/repository/interface.go` 新增 `PatientRepository() patientRepo.Repository`。
5. 至 `internal/factory/repository/factory.go` 初始化並實作該方法。

---

## 設計原則

- **唯一允許直接使用 `*gorm.DB` 的層**，其他層均不得持有 DB 連線。
- **不包含業務邏輯**，只做純粹的 CRUD 操作。
- **一個 domain 的 repository 只操作對應的資料表**，不跨表 join 應拆為多個 repository 方法後由 store 組合。
