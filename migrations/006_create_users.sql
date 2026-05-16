-- Notion doc: N/A
CREATE TABLE IF NOT EXISTS `users` (
    `id`         BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_name`  VARCHAR(100)    NOT NULL COMMENT '暱稱或顯示名稱',
    `avatar`     VARCHAR(500)    NOT NULL COMMENT '大頭照 URL',
    `mobile`     VARCHAR(20)     NOT NULL COMMENT '用戶綁定的手機號',
    `password`   VARCHAR(255)    NOT NULL COMMENT '登入密碼（雜湊儲存）',
    `created_at` DATETIME(3)     NULL,
    `updated_at` DATETIME(3)     NULL,
    `deleted_at` DATETIME(3)     NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uni_users_mobile` (`mobile`),
    INDEX `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
