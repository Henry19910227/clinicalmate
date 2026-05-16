-- Notion doc: https://www.notion.so/auth-admin-35512ea11a32805fbc78f8c15c13bcd3

CREATE TABLE IF NOT EXISTS `admins` (
    `id`             BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name`           VARCHAR(255)    NOT NULL,
    `mobile`         VARCHAR(20)     NOT NULL,
    `permissions_id` INT             NOT NULL,
    `active`         TINYINT         NOT NULL DEFAULT 1,
    `created_at`     DATETIME(3)     NULL,
    `updated_at`     DATETIME(3)     NULL,
    `deleted_at`     DATETIME(3)     NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uni_admins_mobile` (`mobile`),
    INDEX `idx_admins_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
