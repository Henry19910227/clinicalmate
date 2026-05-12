---
description: 為 Clinicalmate 產生符合規範的 git commit，自動分析變更並選擇正確的 Conventional Commit type
argument-hint: 可選：簡短描述本次 commit 的目的
---

# Git Commit

為 Clinicalmate 建立一個符合 Conventional Commits 規範的 commit。

使用者的補充說明：$ARGUMENTS

---

## 步驟

### 1. 收集變更資訊

同步執行：
- `git status` — 查看工作區狀態與未追蹤檔案
- `git diff --cached` — 查看已 staged 的變更
- `git diff` — 查看未 staged 的變更
- `git log --oneline -10` — 了解近期 commit 風格

### 2. 決定 scope 前先確認架構層

根據變更的路徑對應 Clinicalmate 的層級：

| 路徑 | scope |
|------|-------|
| `cmd/` | `cmd` |
| `config/` 或 `config.yaml` | `config` |
| `internal/app/` | `app` |
| `internal/controller/` | `controller` |
| `internal/service/` | `service` |
| `internal/store/` | `store` |
| `internal/repository/` | `repository` |
| `internal/model/` | `model` |
| `internal/router/` | `router` |
| `internal/factory/` | `factory` |
| `internal/infra/` | `infra` |
| `.claude/` | `claude` |
| `Makefile` / `Dockerfile` | `build` |
| 跨多層 | 省略 scope |

### 3. 選擇 Conventional Commit type

| type | 使用時機 |
|------|---------|
| `feat` | 新功能、新 API、新 domain |
| `fix` | bug 修正 |
| `refactor` | 重構，不改變行為 |
| `chore` | 維護性工作（依賴更新、設定調整、架構文件） |
| `test` | 新增或修改測試 |
| `docs` | 純文件變更（README、ARCHITECTURE.md 等） |
| `perf` | 效能優化 |
| `ci` | CI/CD 設定 |

### 4. 安全檢查

若發現以下情況，**立即停止並告知使用者**，不得繼續 commit：
- `.env` 或含有 secret/token/password 的檔案被 stage
- 大型 binary 檔案（>1 MB）
- `config.yaml` 包含明文密碼或 DSN

### 5. 處理未 staged 的檔案

若存在未 stage 的變更：
1. 列出所有未 staged 的檔案
2. 詢問使用者：「這些檔案要一起加入這次 commit 嗎？」
3. 等待明確回覆後再 `git add`
4. 不使用 `git add -A` 或 `git add .`，只加使用者確認的檔案

### 6. 起草 commit message

格式：
```
<type>(<scope>): <subject>

<body（可選）>
```

規則：
- subject 用**中文**，動詞開頭，不加句號，50 字以內
- 若變更橫跨多個關注點，body 用條列說明 why（不是 what）
- 若使用者在 `$ARGUMENTS` 提供了描述，優先採用其意圖

### 7. 確認後執行

1. 向使用者展示完整 commit message 草稿
2. 詢問：「確認後將執行 commit，是否繼續？」
3. 獲得確認後，用 HEREDOC 執行：
   ```bash
   git commit -m "$(cat <<'EOF'
   <完整 commit message>
   EOF
   )"
   ```
4. 執行後顯示 `git log --oneline -3` 確認結果

---

**注意**：不要 push。Push 需使用者另行明確要求。
