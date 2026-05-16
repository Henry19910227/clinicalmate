-- Notion doc: https://www.notion.so/menu-list-35512ea11a32802d913efbec075875bf
-- permission_groups 與 permissions 的中間表，記錄每個權限組擁有的菜單權限 ID 集合。
-- 無軟刪除（無 deleted_at），複合主鍵確保唯一性。

CREATE TABLE IF NOT EXISTS `permission_group_permissions` (
    `permission_group_id` BIGINT UNSIGNED NOT NULL COMMENT '權限組 ID',
    `permission_id`       BIGINT UNSIGNED NOT NULL COMMENT '菜單權限 ID',
    PRIMARY KEY (`permission_group_id`, `permission_id`),
    INDEX `idx_pgp_permission_group_id` (`permission_group_id`),
    INDEX `idx_pgp_permission_id` (`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='權限組與菜單權限中間表';
