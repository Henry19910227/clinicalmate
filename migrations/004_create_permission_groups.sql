-- Notion doc: https://www.notion.so/menu-list-35512ea11a32802d913efbec075875bf

CREATE TABLE IF NOT EXISTS `permission_groups` (
    `id`              BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name`            VARCHAR(100)    NOT NULL,
    `permission_name` VARCHAR(1000)   NOT NULL DEFAULT '',
    `created_at`      DATETIME(3)     NULL,
    `updated_at`      DATETIME(3)     NULL,
    `deleted_at`      DATETIME(3)     NULL,
    PRIMARY KEY (`id`),
    INDEX `idx_permission_groups_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
