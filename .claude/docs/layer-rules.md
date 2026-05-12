## 層職責邊界（強制規則）

| 層 | 職責 | 絕對不可 |
|---|---|---|
| **Router** | 在 RouterGroup 上註冊路由，呼叫 `Set(controllerFactory)` | 包含業務邏輯 |
| **Controller** | Bind request、呼叫 service、寫 response | 直接查 DB、包含業務邏輯 |
| **Service** | 協調 use case、執行業務規則 | 知道 HTTP 細節或 GORM model |
| **Store** | 包裹多 repo 操作在 transaction 中、對應 domain↔DB model | 包含業務規則 |
| **Repository** | 執行單一資料表 GORM 查詢 | 跨資料表 JOIN 或管理 transaction |
| **Model** | GORM struct，只對應 DB schema | 包含業務邏輯 |

## 命名規範

每個 domain（如 `patient`）在每一層都有子 package：
- `internal/controller/patient/`、`internal/service/patient/` 等
- 每個子 package 必須包含：
  - `interface.go` — 定義公開 interface（`Controller`、`Service`、`Store`、`Repository`）
  - `<layer>.go` — 私有 struct 實作 interface + `New(...)` constructor

## 新增 Domain 標準流程

新增 domain（如 `patient`）時，必須依序完成：

1. `internal/model/` — 新增 GORM model struct
2. `internal/repository/<domain>/` — `interface.go` + `repository.go`
3. `internal/store/<domain>/` — `interface.go` + `store.go`
4. `internal/service/<domain>/` — `interface.go` + `service.go`
5. `internal/controller/<domain>/` — `interface.go` + `controller.go`
6. `internal/router/<domain>/` — `interface.go` + `router.go`
7. `internal/factory/repository/` — 註冊 `<Domain>Repository()`
8. `internal/factory/store/` — 註冊 `<Domain>Store()`
9. `internal/factory/service/` — 註冊 `<Domain>Service()`
10. `internal/factory/controller/` — 註冊 `<Domain>Controller()`
11. `cmd/main.go` — 實例化並掛載 router

## Factory 組合模式

```
repository.Factory   ← 持有 singleton repo 實例（在 New() 中建立）
  └── store.Factory
        └── service.Factory
              └── controller.Factory
```

每個 factory 暴露如 `AdminRepository()`、`AdminStore()` 等方法，回傳對應 interface。
