# Factory / 依賴注入架構

## Factory 組合順序

```
repository.Factory   ← 持有 singleton repository 實例（在 New() 建立）
  └── store.Factory
        └── service.Factory
              └── controller.Factory
```

## main.go 接線範例

```go
rdb      := database.New(&cfg.Database)      // 開啟 MySQL 連線
repoF    := repoFactory.New(rdb.Connect())   // 建立並快取 repo 實例
storeF   := storeFactory.New(repoF)
serviceF := serviceFactory.New(storeF)
ctrlF    := controllerFactory.New(serviceF)

g := gin.Default()
adminR := adminRouter.New(g.Group("/api/v1"))
adminR.Set(ctrlF)                            // 註冊路由
```

## 各層 Factory 規則

- **Repository Factory**：在 `New()` 時建立所有 repo 實例（singleton），之後每次呼叫 `AdminRepository()` 回傳同一個實例。
- **Store / Service / Controller Factory**：依賴上層 factory，每次呼叫方法時可建立新實例或快取，依需求決定。
- 每個 factory package 固定暴露兩個檔案：
  - `interface.go` — 定義 `Factory` interface，列出所有 domain 的存取方法
  - `factory.go` — private struct 實作 + `New(...)` constructor

## 新增 Domain 時的 Factory 接線

在各層 `interface.go` 加入新方法，在 `factory.go` 實作，最後在 `main.go` 掛載 router。詳細步驟見 [`../ARCHITECTURE.md`](../ARCHITECTURE.md)。
