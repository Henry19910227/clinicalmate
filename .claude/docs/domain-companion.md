## Companion Domain

**資料表**：`companions`
**Notion 來源**：`/companion/list`（id: `35412ea1-1a32-808c-bf37-d613142c4c83`）
**最後更新**：2026-05-14

### 欄位

| Go 欄位 | DB 欄位 | 型別 | 約束 | 說明 |
|---------|---------|------|------|------|
| Name    | name    | VARCHAR(100) | NOT NULL | 陪診師姓名 |
| Mobile  | mobile  | VARCHAR(20)  | NOT NULL | 手機號碼 |
| Avatar  | avatar  | VARCHAR(500) | NOT NULL | 頭像 URL |
| Sex     | sex     | VARCHAR(2)   | NOT NULL | 性別（"1" 男 / "2" 女） |
| Age     | age     | INT          | NOT NULL | 年齡 |
| Active  | active  | INT          | NOT NULL, DEFAULT 1 | 狀態（1: 啟用 / 0: 停用） |

> `gorm.Model` 內含 `id`、`created_at`、`updated_at`、`deleted_at`，不重複列出。

### 索引

| 索引名稱 | 欄位 | 類型 |
|---------|------|------|
| idx_companions_deleted_at | deleted_at | INDEX（soft delete） |

### 業務說明

Companion（陪診師）是提供陪診服務的人員，病患可透過系統預約陪診師。`active` 欄位控制陪診師是否對外開放預約，預設啟用。
