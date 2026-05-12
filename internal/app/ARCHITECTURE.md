# App 層說明

`internal/app/` 是應用的頂層協調者，持有 Gin engine 與所有 controller，  
對外只暴露一個 `Run()` 方法，由 `main.go` 呼叫以啟動 HTTP server。

---

## 目錄結構

```
internal/app/
└── core/
    ├── interface.go  # 定義 App interface
    └── app.go        # 實作 app struct 與 New()、Run()
```

---

## 請求流程中的位置

```
main.go → App.Run() → Gin Engine → Router → Controller → ...
```

`main.go` 完成所有 factory 與 router 的初始化後，呼叫 `app.Run()` 交由 Gin 接管。

---

## 啟動流程（`main.go`）

```go
// 1. 建立 factory 鏈
configFac     := configFactory.New(cfg)
infraFac      := infraFactory.New(configFac)
repoF         := repoFactory.New(infraFac)
storeF        := storeFactory.New(repoF)
serviceF      := serviceFactory.New(storeF)
controllerF   := controllerFactory.New(serviceF)

// 2. 建立 Gin engine 並掛載 router
g := gin.Default()
adminR := adminRouter.New(g.Group("/api/v1"))
adminR.Set(controllerF)

// 3. 建立 app factory 並取得 core app
appFac := appFactory.New(g, controllerF)
appFac.CoreApp().Run()
```

---

## 檔案職責

| 檔案 | 職責 |
|------|------|
| `interface.go` | 宣告 `App` interface，目前只有 `Run() error` |
| `app.go` | 宣告 `app` struct、實作 `New()`、實作 `Run()` |

---

## 命名規則

### Package

固定為 `core`。

```
internal/app/core/ → package core
```

### Interface 與 struct

| 項目 | 名稱 |
|------|------|
| 對外 interface | `App` |
| 內部 struct | `app`（未匯出） |

### 建構子

```go
func New(engine *gin.Engine, adminCtrl adminController.Controller) App
```

### struct 欄位

持有 Gin engine 與各 domain 的 controller。

```go
type app struct {
    engine    *gin.Engine
    adminCtrl adminController.Controller
}
```

新增 domain 時，在 struct 加入對應的 controller 欄位。

### import alias

格式：`{domain}Controller`，camelCase。

```go
import adminController "clinicalmate/internal/controller/admin"
```

### `Run()` 方法

直接委派給 `gin.Engine.Run()`，不帶任何參數（host 與 port 已在 engine 建立時設定）。

```go
func (a *app) Run() error {
    return a.engine.Run()
}
```

---

## 新增 Domain 的步驟

以新增 `patient` domain 為例：

1. 在 `app.go` 的 `app` struct 新增欄位 `patientCtrl patientController.Controller`。
2. 更新 `New()` 簽名，加入 `patientCtrl patientController.Controller` 參數。
3. 至 `factory/app/factory.go` 更新 `New()` 呼叫，傳入 `controllerFac.PatientController()`。

---

## 設計原則

- **不負責路由定義**，路由由 `router` 層在 `main.go` 中掛載完成後再傳入 engine。
- **`Run()` 是唯一對外方法**，`main.go` 只需呼叫這一行即可啟動服務。
- **不包含任何業務邏輯**，app 層只做最終的啟動協調。
