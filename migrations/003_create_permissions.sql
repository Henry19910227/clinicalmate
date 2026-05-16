-- Notion doc: https://www.notion.so/menu-permissions-35512ea11a3280719361f87df077ff4d

CREATE TABLE IF NOT EXISTS `permissions` (
    `id`          BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name`        VARCHAR(100)    NOT NULL,
    `label`       VARCHAR(255)    NOT NULL,
    `path`        VARCHAR(255)    NOT NULL DEFAULT '',
    `component`   VARCHAR(255)    NOT NULL DEFAULT '',
    `icon`        VARCHAR(100)    NOT NULL DEFAULT '',
    `parent_id`   BIGINT UNSIGNED NOT NULL DEFAULT 0,
    `disabled`    TINYINT(1)      NOT NULL DEFAULT 0,
    `created_at`  DATETIME(3)     NULL,
    `updated_at`  DATETIME(3)     NULL,
    `deleted_at`  DATETIME(3)     NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uni_permissions_name` (`name`),
    INDEX `idx_permissions_parent_id` (`parent_id`),
    INDEX `idx_permissions_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
