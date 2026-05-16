## User Domain

**資料表**：`users`
**Notion 來源**：N/A（Notion 無對應資料表設計頁，欄位以使用者提供的 JSON 為準）
**最後更新**：2026-05-16

### 欄位

| Go 欄位 | DB 欄位 | 型別 | 約束 | 說明 |
|---------|---------|------|------|------|
| UserName | user_name | VARCHAR(100) | NOT NULL | 暱稱或顯示名稱 |
| Avatar | avatar | VARCHAR(500) | NOT NULL | 大頭照 URL |
| Mobile | mobile | VARCHAR(20) | NOT NULL, UNIQUE | 用戶綁定的手機號 |
| Password | password | VARCHAR(255) | NOT NULL | 登入密碼（雜湊儲存） |

> `gorm.Model` 內含 `id`、`created_at`、`updated_at`、`deleted_at`，不重複列出。

### 索引

| 索引名稱 | 欄位 | 類型 | 說明 |
|---------|------|------|------|
| uni_users_mobile | mobile | UNIQUE | 手機號全系統唯一 |
| idx_users_deleted_at | deleted_at | INDEX | GORM 軟刪除標準索引 |

### 業務說明

- User 為前台用戶實體，對應使用 Clinicalmate 服務的病患或患者家屬。
- `mobile` 作為用戶身份識別的唯一鍵，不可重複。
- `avatar` 儲存頭像圖片的完整 URL，由外部 CDN 提供。
- `user_name` 為顯示用的暱稱，允許後續由使用者自行修改。
