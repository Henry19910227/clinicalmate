-- Notion doc: https://www.notion.so/order-detail-35512ea11a3280348416cb0b29c2511a

CREATE TABLE IF NOT EXISTS `orders` (
    `id`               BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `out_trade_no`     VARCHAR(64)     NOT NULL COMMENT '唯一訂單編號',
    `trade_state`      VARCHAR(20)     NOT NULL DEFAULT '待支付' COMMENT '訂單狀態（待支付 / 待服務 / 已完成 / 已取消）',
    `order_start_time` BIGINT          NOT NULL COMMENT '下單時間戳（毫秒）',
    `price`            VARCHAR(20)     NOT NULL COMMENT '應付金額（如 "100.00"）',
    `service_name`     VARCHAR(100)    NOT NULL COMMENT '預約服務名稱',
    `start_time`       DATETIME        NOT NULL COMMENT '期望就診時間',
    `tel`              VARCHAR(20)     NOT NULL COMMENT '就診人備用聯絡電話',
    `receive_address`  VARCHAR(500)    NOT NULL COMMENT '接送地址',
    `demand`           TEXT                NULL COMMENT '備註需求（選填）',
    `code_url`         VARCHAR(500)        NULL COMMENT '支付二維碼連結（待支付時使用）',
    `client_name`      VARCHAR(100)    NOT NULL COMMENT '就診人姓名',
    `client_mobile`    VARCHAR(20)     NOT NULL COMMENT '就診人手機號碼',
    `user_id`          BIGINT UNSIGNED NOT NULL COMMENT '下單用戶 ID（FK -> users）',
    `hospital_id`      BIGINT UNSIGNED NOT NULL COMMENT '就診醫院 ID（FK -> hospitals）',
    `companion_id`     BIGINT UNSIGNED     NULL COMMENT '指派陪診師 ID（FK -> companions，派單後填入）',
    `created_at`       DATETIME(3)         NULL,
    `updated_at`       DATETIME(3)         NULL,
    `deleted_at`       DATETIME(3)         NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uni_orders_out_trade_no` (`out_trade_no`),
    INDEX `idx_orders_user_id` (`user_id`),
    INDEX `idx_orders_hospital_id` (`hospital_id`),
    INDEX `idx_orders_companion_id` (`companion_id`),
    INDEX `idx_orders_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
