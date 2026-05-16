## Permission Domain

**資料表**：`permissions`
**Notion 來源**：https://www.notion.so/menu-permissions-35512ea11a3280719361f87df077ff4d
**最後更新**：2026-05-16

### 欄位

| Go 欄位 | DB 欄位 | 型別 | 約束 | 說明 |
|---------|---------|------|------|------|
| Name | name | VARCHAR(100) | NOT NULL, UNIQUE | 路由名稱，前端 router.push 使用，全域唯一 |
| Label | label | VARCHAR(255) | NOT NULL | 菜單顯示標題（中文），對應 Notion meta.name |
| Path | path | VARCHAR(255) | NOT NULL, DEFAULT '' | 瀏覽器路由路徑，如 /auth/admin |
| Component | component | VARCHAR(255) | NOT NULL, DEFAULT '' | 前端組件映射鍵值，如 auth/admin |
| Icon | icon | VARCHAR(100) | NOT NULL, DEFAULT '' | Element Plus 圖標名稱，如 UserFilled |
| ParentID | parent_id | BIGINT UNSIGNED | NOT NULL, DEFAULT 0 | 父菜單 ID，0 代表頂層節點，支撐樹狀結構 |
| Disabled | disabled | TINYINT(1) | NOT NULL, DEFAULT 0 | 是否停用：0 啟用，1 停用 |

> `gorm.Model` 內含 `id`、`created_at`、`updated_at`、`deleted_at`，不重複列出。

### 索引

| 索引名稱 | 欄位 | 類型 |
|---------|------|------|
| uni_permissions_name | name | UNIQUE |
| idx_permissions_parent_id | parent_id | INDEX（樹狀查詢） |
| idx_permissions_deleted_at | deleted_at | INDEX（soft delete） |

### 業務說明

`permissions` 資料表儲存系統的**動態菜單權限**定義，決定登入後使用者在側邊欄可見的導航選單。

資料結構為**樹狀**（Tree）：頂層節點 `parent_id = 0`，子節點的 `parent_id` 指向父節點 `id`。API `/menu/permissions` 在使用者登入後呼叫，後端依該使用者所屬 Admin 的 `permissions_id` 查詢對應的菜單集合，組裝樹狀 JSON 回傳前端。

前端收到後透過 Vuex `store/menu.js` 的 `dynamicMenu` mutation：
1. 將 `component` 字串映射至實際 `.vue` 檔案路徑
2. 以 `router.addRoute()` 動態注冊路由
3. 由 `AppAside.vue` 遞迴渲染左側導航欄
