# Controller 層說明

`internal/controller/` 負責處理 HTTP 請求，是請求流程中最靠近 Router 的業務層。  
每個子目錄代表一個 domain，只依賴同 domain 的 `service`。

---

## 目錄結構

```
internal/controller/
├── admin/
│   ├── interface.go   # 定義 Controller interface（handler 方法簽名）
│   └── controller.go  # 實作 controller struct 與 New()
└── {domain}/
    ├── interface.go
    └── controller.go
```

---

## 請求流程中的位置

```
Router → Controller → Service → Store → Repository → DB
```

Router 透過 `factory/controller.Factory` 取得 controller，再將 handler 方法註冊到路由上。

---

## 檔案職責

| 檔案 | 職責 |
|------|------|
| `interface.go` | 宣告 `Controller` interface，列出所有對外暴露的 handler 方法 |
| `controller.go` | 宣告 `controller` struct、實作 `New()`、實作各 handler 方法 |

---

## 命名規則

### Package

取 domain 目錄名，全部小寫。

```
internal/controller/admin/   → package admin
internal/controller/patient/ → package patient
```

### Interface 與 struct

| 項目 | 名稱 |
|------|------|
| 對外 interface | `Controller` |
| 內部 struct | `controller`（未匯出） |

### 建構子

```go
func New({domain}Service {alias}.Service) Controller
```

範例：
```go
func New(adminService adminService.Service) Controller {
    return &controller{adminService: adminService}
}
```

### struct 欄位

格式：`{domain}Service`，camelCase。

```go
type controller struct {
    adminService adminService.Service
}
```

### import alias

格式：`{domain}Service`，camelCase。

```go
import adminService "clinicalmate/internal/service/admin"
```

### Handler 方法

格式：動詞 + 資源名，PascalCase，參數固定為 `*gin.Context`。

```go
func (c *controller) GetAdminProfile(ctx *gin.Context) {}
func (c *controller) CreateAdmin(ctx *gin.Context)     {}
func (c *controller) UpdateAdmin(ctx *gin.Context)     {}
func (c *controller) DeleteAdmin(ctx *gin.Context)     {}
```

常用動詞：`Get`（單筆）、`List`（列表）、`Create`、`Update`、`Delete`。

---

## 新增 Domain 的步驟

以新增 `patient` domain 為例：

1. 建立目錄 `internal/controller/patient/`。
2. `interface.go`：宣告 `Controller` interface 與所有 handler 方法。
3. `controller.go`：實作 `controller` struct，`New()` 接收 `patientService patientService.Service`。
4. 至 `internal/factory/controller/interface.go` 新增 `PatientController() patientCtrl.Controller`。
5. 至 `internal/factory/controller/factory.go` 初始化並實作該方法。
