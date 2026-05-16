## Hospital Domain

**資料表**：`hospitals`
**Notion 來源**：https://www.notion.so/admin-order-35412ea11a3280b6a1f3ffebcdb1e4c0
**最後更新**：2026-05-17

### 欄位

| Go 欄位 | DB 欄位 | 型別 | 約束 | 說明 |
|---------|---------|------|------|------|
| Name | name | VARCHAR(255) | NOT NULL | 醫院名稱 |
| Address | address | VARCHAR(500) | NOT NULL | 醫院地址 |
| Image | image | VARCHAR(500) | NOT NULL | 醫院圖片 URL |
| Phone | phone | VARCHAR(50) | NOT NULL | 聯絡電話 |

> `gorm.Model` 內含 `id`、`created_at`、`updated_at`、`deleted_at`，不重複列出。

### 索引

| 索引名稱 | 欄位 | 類型 |
|---------|------|------|
| idx_hospitals_deleted_at | deleted_at | INDEX（soft delete） |

### 業務說明

此資料表儲存醫院基本資訊，對應 API 端點 `GET /admin/order`（訂單列表）與 `GET /order/detail`（訂單詳情）回傳的 `hospital_id` 與 `hospital_name` 欄位。

- 訂單資料中以 `hospital_id` 作為外鍵關聯此表，`hospital_name` 為冗餘顯示名稱。
- 欄位來源依據 Notion `/admin/order` 回應範例（`hospital_id: 5`、`hospital_name: "武汉中心医院"`）。
