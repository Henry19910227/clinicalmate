-- Notion doc: https://www.notion.so/admin-order-35412ea11a3280b6a1f3ffebcdb1e4c0

CREATE TABLE IF NOT EXISTS `hospitals` (
    `id`         BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name`       VARCHAR(255)    NOT NULL COMMENT '醫院名稱',
    `address`    VARCHAR(500)    NOT NULL COMMENT '醫院地址',
    `image`      VARCHAR(500)    NOT NULL COMMENT '醫院圖片 URL',
    `phone`      VARCHAR(50)     NOT NULL COMMENT '聯絡電話',
    `created_at` DATETIME(3)     NULL,
    `updated_at` DATETIME(3)     NULL,
    `deleted_at` DATETIME(3)     NULL,
    PRIMARY KEY (`id`),
    INDEX `idx_hospitals_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
