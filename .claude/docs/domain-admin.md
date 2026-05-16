## Admin Domain

**資料表**：`admins`
**Notion 來源**：https://www.notion.so/auth-admin-35512ea11a32805fbc78f8c15c13bcd3
**最後更新**：2026-05-14

### 欄位

| Go 欄位 | DB 欄位 | 型別 | 約束 | 說明 |
|---------|---------|------|------|------|
| Name | name | VARCHAR(255) | NOT NULL | 管理員暱稱 / 名稱 |
| Mobile | mobile | VARCHAR(20) | NOT NULL, UNIQUE | 手機號碼，不可重複 |
| PermissionsID | permissions_id | INT | NOT NULL | 所屬權限組 ID（如 1 代表超級管理員） |
| Active | active | TINYINT | NOT NULL, DEFAULT 1 | 狀態：1 正常，0 失效 |

> `gorm.Model` 內含 `id`、`created_at`、`updated_at`、`deleted_at`，不重複列出。

### 索引

| 索引名稱 | 欄位 | 類型 |
|---------|------|------|
| uni_admins_mobile | mobile | UNIQUE |
| idx_admins_deleted_at | deleted_at | INDEX（soft delete） |

### 業務說明

此資料表儲存系統管理員帳號資訊，對應 API 端點 `GET /auth/admin`（账号管理列表）。

- `permissions_id` 關聯權限組別，前端透過 `menuSelectList` 接口取得清單後進行 ID 轉換（如 1 -> 超級管理員）。
- `active` 用於前端表格狀態標籤（el-tag）：值為 1 顯示綠色「正常」，否則顯示紅色「失效」。
- `mobile` 為登入識別號，全系統唯一。
- API 支援分頁查詢（pageNum、pageSize），需 x-token header 驗證。
