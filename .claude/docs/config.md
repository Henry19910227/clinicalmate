## Config 管理

- **Model**：`internal/model/config/model.go`，含 `Config`、`AppConfig`、`DatabaseConfig` struct（yaml tags 用 snake_case）
- **實作**：`internal/config/app/interface.go` + `config.go`，遵循與其他層相同的 interface 模式
- **config.yaml**：專案根目錄，欄位與 model struct 保持一致

新增欄位時：先改 `model.go` 加欄位（含 yaml tag），再同步更新 `config.yaml`。
