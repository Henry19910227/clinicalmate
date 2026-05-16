## PermissionGroup Domain

**資料表**：`permission_groups`
**Notion 來源**：https://www.notion.so/menu-list-35512ea11a32802d913efbec075875bf
**最後更新**：2026-05-16

### 欄位

| Go 欄位 | DB 欄位 | 型別 | 約束 | 說明 |
|---------|---------|------|------|------|
| Name | name | VARCHAR(100) | NOT NULL | 權限組名稱（如：系統管理員、普通用戶） |
| PermissionName | permission_name | VARCHAR(1000) | NOT NULL, DEFAULT '' | 擁有的菜單名稱彙總，以逗號分隔（供列表展示用，前端以 zhCNtoTW 轉換繁體） |

> `gorm.Model` 內含 `id`、`created_at`、`updated_at`、`deleted_at`，不重複列出。

### 索引

| 索引名稱 | 欄位 | 類型 |
|---------|------|------|
| idx_permission_groups_deleted_at | deleted_at | INDEX（soft delete） |

### 關聯表

`permission_group_permissions`（中間表）：

| 欄位 | 型別 | 說明 |
|------|------|------|
| permission_group_id | BIGINT UNSIGNED | 權限組 ID |
| permission_id       | BIGINT UNSIGNED | 菜單權限 ID |

複合主鍵 (permission_group_id, permission_id)，無軟刪除。

### 業務說明

`permission_groups` 資料表定義系統的**角色（Role）/ 權限組**，每筆記錄代表一個可指派給管理員帳號的權限集合。

API `/menu/list` 回傳的列表結構：
- `id`：權限組唯一識別 ID
- `name`：組名稱，直接顯示於後台列表
- `permissionName`：各菜單名稱彙總字串（以逗號分隔），前端以 `zhCNtoTW` 轉繁體後顯示於表格欄位
- `permission`：該組擁有的 permission ID 陣列（儲存於關聯表，非本表欄位），供前端編輯時以 `el-tree.setCheckedKeys()` 重建勾選狀態

管理員帳號透過外鍵關聯到某一 `permission_group`，決定登入後可存取的菜單範圍。
