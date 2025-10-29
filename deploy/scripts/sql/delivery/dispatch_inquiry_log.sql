CREATE TABLE `dispatch_inquiry_log`
(
    `id`                bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `order_no`          varchar(255) NOT NULL DEFAULT '' COMMENT '订单号',
    `from_address`      json         NOT NULL COMMENT '起点地址',
    `to_address`        json         NOT NULL COMMENT '终点地址',
    `status`            varchar(50)  NOT NULL DEFAULT 'processing' COMMENT '状态：processing/success/failed',
    `goods_type`        varchar(255) NOT NULL DEFAULT '' COMMENT '物品类型',
    `delivery_codes`    varchar(255) NOT NULL DEFAULT '' COMMENT '询价平台(逗号分隔)',
    `success_platforms` varchar(255) NOT NULL DEFAULT '' COMMENT '成功平台(逗号分隔)',
    `total_duration`    int          NOT NULL DEFAULT 0 COMMENT '总耗时(毫秒)',
    `created_at`        datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY                 `idx_order_no` (`order_no`),
    KEY                 `idx_status` (`status`),
    KEY                 `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='询价记录表';