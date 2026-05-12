# Store 層說明

`internal/store/` 負責組合 repository，將資料存取操作封裝為業務可直接調用的方法。  
每個子目錄代表一個 domain，只依賴同 domain 的 `repository`。

---

## 目錄結構

```
internal/store/
├── admin/
│   ├── interface.go  # 定義 Store interface（資料存取方法簽名）
│   └── store.go      # 實作 store struct 與 New()
└── {domain}/
    ├── interface.go
    └── store.go
```

---

## 請求流程中的位置

```
Service → Store → Repository → DB
```

Service 透過 `factory/store.Factory` 取得 store，調用資料存取方法。

---

## 檔案職責

| 檔案 | 職責 |
|------|------|
| `interface.go` | 宣告 `Store` interface，列出所有對外暴露的資料存取方法 |
| `store.go` | 宣告 `store` struct、實作 `New()`、實作各資料存取方法 |

---

## 命名規則

### Package

取 domain 目錄名，全部小寫。

```
internal/store/admin/   → package admin
internal/store/patient/ → package patient
```

### Interface 與 struct

| 項目 | 名稱 |
|------|------|
| 對外 interface | `Store` |
| 內部 struct | `store`（未匯出） |

### 建構子

`New()` 直接接收所需的 repository 物件，**不接收整個 `repository.Factory`**。

```go
func New({domain}Repo {alias}.Repository) Store
```

範例：

```go
func New(adminRepo adminRepo.Repository) Store {
    return &store{adminRepo: adminRepo}
}
```

factory 層負責從 `repositoryFac.AdminRepository()` 取出後再傳入：

```go
// factory/store/factory.go
adminSto := adminStore.New(repositoryFac.AdminRepository())
```

### struct 欄位

格式：`{domain}Repo`，camelCase。

```go
type store struct {
    adminRepo adminRepo.Repository
}
```

### import alias

格式：`{domain}Repo`，camelCase。

```go
import adminRepo "clinicalmate/internal/repository/admin"
```

### 資料存取方法

格式：動詞 + 資源名，PascalCase，回傳值帶 `error`。

```go
func (s *store) GetByID(id int64) (*model.Admin, error)          {}
func (s *store) List(filter *dto.AdminFilter) ([]*model.Admin, error) {}
func (s *store) Create(admin *model.Admin) error                 {}
func (s *store) Update(admin *model.Admin) error                 {}
func (s *store) Delete(id int64) error                           {}
```

常用動詞：`Get`（單筆）、`List`（列表）、`Create`、`Update`、`Delete`。

---

## 新增 Domain 的步驟

以新增 `patient` domain 為例：

1. 建立目錄 `internal/store/patient/`。
2. `interface.go`：宣告 `Store` interface 與所有資料存取方法。
3. `store.go`：實作 `store` struct，`New()` 接收 `patientRepo patientRepo.Repository`。
4. 至 `internal/factory/store/interface.go` 新增 `PatientStore() patientSto.Store`。
5. 至 `internal/factory/store/factory.go` 初始化並實作該方法。

---

## 設計原則

- **不直接持有 `*gorm.DB`**，所有 DB 操作一律委派給 `repository`。
- **不包含業務邏輯**，store 只做資料的組合與轉發。
- **一個 domain 的 store 只依賴同 domain 的 repository**。
