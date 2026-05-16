CREATE TABLE IF NOT EXISTS `companions` (
    `id`         BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name`       VARCHAR(100)    NOT NULL,
    `mobile`     VARCHAR(20)     NOT NULL,
    `avatar`     VARCHAR(500)    NOT NULL,
    `sex`        VARCHAR(2)      NOT NULL,
    `age`        INT             NOT NULL,
    `active`     INT             NOT NULL DEFAULT 1,
    `created_at` DATETIME(3)     NULL,
    `updated_at` DATETIME(3)     NULL,
    `deleted_at` DATETIME(3)     NULL,
    PRIMARY KEY (`id`),
    INDEX `idx_companions_deleted_at` (`deleted_at`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;
