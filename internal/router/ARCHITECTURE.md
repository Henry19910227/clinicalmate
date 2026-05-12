# Router 層說明

`internal/router/` 負責定義路由規則，將 HTTP 路徑與 controller handler 綁定。  
每個子目錄代表一個 domain，持有 `*gin.RouterGroup` 並透過 `controller.Factory` 取得 handler。

---

## 目錄結構

```
internal/router/
├── admin/
│   ├── interface.go  # 定義 Router interface
│   └── router.go     # 實作 router struct 與 New()、Set()
└── {domain}/
    ├── interface.go
    └── router.go
```

---

## 請求流程中的位置

```
HTTP Request → Gin Engine → RouterGroup → Router.Set() → Controller Handler
```

Router 在啟動時由 `main.go` 建立，`Set()` 被呼叫後路由即生效。

---

## 啟動流程（`main.go`）

```go
g := gin.Default()
adminR := adminRouter.New(g.Group("/api/v1"))
adminR.Set(controllerF)
_ = g.Run(...)
```

路由前綴（`/api/v1`）由 `main.go` 在呼叫 `New()` 時決定，router 本身不硬編碼前綴。

---

## 檔案職責

| 檔案 | 職責 |
|------|------|
| `interface.go` | 宣告 `Router` interface，固定方法為 `Set(factory controller.Factory)` |
| `router.go` | 宣告 `router` struct、實作 `New()`、在 `Set()` 中註冊所有路由 |

---

## 命名規則

### Package

取 domain 目錄名，全部小寫。

```
internal/router/admin/   → package admin
internal/router/patient/ → package patient
```

### Interface 與 struct

| 項目 | 名稱 |
|------|------|
| 對外 interface | `Router` |
| 內部 struct | `router`（未匯出） |

### 建構子

```go
func New(group *gin.RouterGroup) Router
```

### struct 欄位

固定為 `group *gin.RouterGroup`。

```go
type router struct {
    group *gin.RouterGroup
}
```

### import alias

```go
import "clinicalmate/internal/factory/controller"
```

controller factory 不需要 alias，直接使用套件名 `controller`。

### `Set()` 方法

固定簽名，接收 `controller.Factory`，在方法內完成所有路由註冊。

```go
func (r *router) Set(factory controller.Factory) {
    r.group.GET("/admin",         factory.AdminController().GetAdmin)
    r.group.POST("/admin",        factory.AdminController().CreateAdmin)
    r.group.PUT("/admin/:id",     factory.AdminController().UpdateAdmin)
    r.group.DELETE("/admin/:id",  factory.AdminController().DeleteAdmin)
}
```

### 路由路徑

- 全部小寫，單字以 `-` 分隔（kebab-case）。
- 資源名用複數：`/patients`、`/appointments`。
- 路徑參數用 `:{name}`：`/patients/:id`。

---

## 新增 Domain 的步驟

以新增 `patient` domain 為例：

1. 建立目錄 `internal/router/patient/`。
2. `interface.go`：宣告 `Router` interface，方法固定為 `Set(factory controller.Factory)`。
3. `router.go`：實作 `router` struct，`New()` 接收 `*gin.RouterGroup`，`Set()` 中完成路由註冊。
4. 至 `main.go` 新增：
   ```go
   patientR := patientRouter.New(g.Group("/api/v1"))
   patientR.Set(controllerF)
   ```

---

## 設計原則

- **只做路由綁定**，不含任何業務邏輯或資料處理。
- **路由前綴由 `main.go` 決定**，router 本身只管相對路徑。
- **Router 不由 factory 管理**，直接在 `main.go` 建立與呼叫，因為路由註冊是一次性的啟動行為。
