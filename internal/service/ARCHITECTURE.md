# Service 層說明

`internal/service/` 負責業務邏輯，是請求流程中協調資料存取與業務規則的核心層。  
每個子目錄代表一個 domain，只依賴同 domain 的 `store`。

---

## 目錄結構

```
internal/service/
├── admin/
│   ├── interface.go  # 定義 Service interface（業務方法簽名）
│   └── service.go    # 實作 service struct 與 New()
└── {domain}/
    ├── interface.go
    └── service.go
```

---

## 請求流程中的位置

```
Controller → Service → Store → Repository → DB
```

Controller 透過 `factory/service.Factory` 取得 service，調用業務方法處理請求。

---

## 檔案職責

| 檔案 | 職責 |
|------|------|
| `interface.go` | 宣告 `Service` interface，列出所有對外暴露的業務方法 |
| `service.go` | 宣告 `service` struct、實作 `New()`、實作各業務方法 |

---

## 命名規則

### Package

取 domain 目錄名，全部小寫。

```
internal/service/admin/   → package admin
internal/service/patient/ → package patient
```

### Interface 與 struct

| 項目 | 名稱 |
|------|------|
| 對外 interface | `Service` |
| 內部 struct | `service`（未匯出） |

### 建構子

```go
func New({domain}Store {alias}.Store) Service
```

範例：

```go
func New(adminStore adminStore.Store) Service {
    return &service{adminStore: adminStore}
}
```

### struct 欄位

格式：`{domain}Store`，camelCase。

```go
type service struct {
    adminStore adminStore.Store
}
```

### import alias

格式：`{domain}Store`，camelCase。

```go
import adminStore "clinicalmate/internal/store/admin"
```

### 業務方法

格式：動詞 + 資源名，PascalCase，回傳值帶 `error`。

```go
func (s *service) GetAdminProfile(id int64) (*model.Admin, error) {}
func (s *service) CreateAdmin(req *dto.CreateAdminReq) error      {}
func (s *service) UpdateAdmin(req *dto.UpdateAdminReq) error      {}
func (s *service) DeleteAdmin(id int64) error                     {}
```

常用動詞：`Get`（單筆）、`List`（列表）、`Create`、`Update`、`Delete`。

---

## 新增 Domain 的步驟

以新增 `patient` domain 為例：

1. 建立目錄 `internal/service/patient/`。
2. `interface.go`：宣告 `Service` interface 與所有業務方法。
3. `service.go`：實作 `service` struct，`New()` 接收 `patientStore patientStore.Store`。
4. 至 `internal/factory/service/interface.go` 新增 `PatientService() patientSvc.Service`。
5. 至 `internal/factory/service/factory.go` 初始化並實作該方法。

---

## 設計原則

- **不直接操作 DB**，所有資料存取一律透過 `store`。
- **不持有 HTTP 相關物件**（`*gin.Context` 等），保持與傳輸層解耦。
- **一個 domain 的 service 只依賴同 domain 的 store**，跨 domain 邏輯由上層（controller 或更高層）協調。
