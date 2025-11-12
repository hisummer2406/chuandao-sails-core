CREATE TABLE `dispatch_inquiry_detail`
(
    `id`              bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `inquiry_id`      bigint      NOT NULL COMMENT '询价ID',
    `account_id`      int         NOT NULL COMMENT '询价账号ID',
    `delivery_code`   varchar(32) NOT NULL DEFAULT '' COMMENT '平台代码',
    `price`           int         NOT NULL DEFAULT 0 COMMENT '配送费用(分)',
    `distance`        int         NOT NULL DEFAULT 0 COMMENT '配送距离(米)',
    `duration`        int         NOT NULL DEFAULT 0 COMMENT '接口耗时(毫秒)',
    `estimate_time`   int         NOT NULL DEFAULT 0 COMMENT '预计时长(分钟)',
    `quote_status`    tinyint     NOT NULL DEFAULT 0 COMMENT '询价状态：0-失败 1-成功',
    `price_token`     json        NOT NULL COMMENT '配送费令牌',
    `result_response` json        NOT NULL COMMENT '原始响应',
    `created_at`      datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY               `idx_inquiry_id` (`inquiry_id`),
    KEY               `idx_delivery_code` (`delivery_code`),
    KEY               `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='询价明细表';