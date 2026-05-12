# Factory 層說明

`internal/factory/` 是 Clinicalmate 的依賴注入（DI）接線層。  
每個子目錄對應一個業務層，負責建立該層的物件並向上層暴露介面。

---

## 目錄結構

```
internal/factory/
├── config/       # 包裝 config model，提供各層設定物件
├── infra/        # 建立基礎設施（DB 連線等）
├── repository/   # 建立 repository 物件（直接操作 DB）
├── store/        # 建立 store 物件（組合 repository）
├── service/      # 建立 service 物件（業務邏輯）
├── controller/   # 建立 controller 物件（HTTP handler）
└── app/          # 建立 core app（掛載 router）
```

每個子目錄都有兩個檔案：

| 檔案 | 職責 |
|------|------|
| `interface.go` | 定義 `Factory` interface，供上層依賴 |
| `factory.go` | 實作 `factory` struct 與 `New()` 建構子 |

---

## 初始化順序

`cmd/main.go` 按以下順序依序建立各層 factory：

```
config → infra → repository → store → service → controller → app
```

每一層只依賴**下一層的 Factory interface**，不直接持有具體實作。

---

## 各層職責

### `config`

- 輸入：`model.Config`（從 `config.yaml` 解析而來）
- 輸出：`AppConfig()`、`MysqlConfig()`

### `infra`

- 輸入：`config.Factory`
- 輸出：`MysqlInfra()`（持有 GORM `*DB`）

### `repository`

- 輸入：`infra.Factory`
- 從 `infra.MysqlInfra().GORM()` 取得 DB 連線，建立各 domain 的 repository
- 輸出：`AdminRepository()` …

### `store`

- 輸入：`repository.Factory`
- 建立各 domain 的 store（組合多個 repository）
- 輸出：`AdminStore()` …

### `service`

- 輸入：`store.Factory`
- 建立各 domain 的 service（業務邏輯）
- 輸出：`AdminService()` …

### `controller`

- 輸入：`service.Factory`
- 建立各 domain 的 controller（HTTP handler）
- 輸出：`AdminController()` …

### `app`

- 輸入：`*gin.Engine`、`controller.Factory`
- 建立 `core.App`，負責將 router 與 controller 接在一起
- 輸出：`CoreApp()`

---

## 命名規則

### Factory 方法（`interface.go`）

格式：`{Domain}{LayerType}()`，全部 PascalCase。

| 範例 | 說明 |
|------|------|
| `AdminRepository()` | admin domain 的 repository |
| `AdminStore()` | admin domain 的 store |
| `AdminService()` | admin domain 的 service |
| `AdminController()` | admin domain 的 controller |
| `MysqlInfra()` | MySQL 基礎設施 |
| `AppConfig()` | App 設定 |

Domain 名稱取自 `internal/{layer}/{domain}/` 的目錄名，首字母大寫。  
LayerType 取整個層名，全字拼寫（`Repository`、`Store`、`Service`、`Controller`、`Infra`、`Config`）。

---

### `factory` struct 欄位

格式：`{domain}{suffix}`，camelCase，suffix 為層名縮寫。

| 層 | 縮寫 | 範例 |
|----|------|------|
| repository | `Repo` | `adminRepo` |
| store | `Sto` | `adminSto` |
| service | `Svc` | `adminSvc` |
| controller | `Ctrl` | `adminCtrl` |
| infra | `Inf` | `mysqlInf` |
| factory 依賴 | `Fac` | `configFac`、`infraFac`、`repositoryFac` |

---

### import alias

格式：`{domain}{LayerType}`，camelCase。

```go
// 業務層
import adminRepository "clinicalmate/internal/repository/admin"
import adminStore      "clinicalmate/internal/store/admin"
import adminService    "clinicalmate/internal/service/admin"
import adminController "clinicalmate/internal/controller/admin"

// 下層 factory
import configFactory     "clinicalmate/internal/factory/config"
import infraFactory      "clinicalmate/internal/factory/infra"
import repositoryFactory "clinicalmate/internal/factory/repository"
import storeFactory      "clinicalmate/internal/factory/store"
import serviceFactory    "clinicalmate/internal/factory/service"

// 基礎設施
import mysqlInfra "clinicalmate/internal/infra/mysql"
```

---

## 新增 Domain 的步驟

以新增 `patient` domain 為例：

1. **實作業務層**：依序在 `repository/patient`、`store/patient`、`service/patient`、`controller/patient` 下建立實作。
2. **更新各層 factory interface**：在對應的 `interface.go` 新增方法，例如 `PatientRepository()`。
3. **更新各層 factory 實作**：在 `factory.go` 的 `factory` struct 新增欄位，並在 `New()` 中初始化，再實作新方法。
4. **接線順序**：從底層（repository）往上（controller）依序更新，保持單向依賴。

---

## 設計原則

- **每層只依賴下一層的 interface**，不跨層依賴。
- **`New()` 是唯一的初始化入口**，物件在 `main.go` 啟動時一次性建立，之後以單例方式共用。
- **不在 factory 層放業務邏輯**，factory 只負責接線。
