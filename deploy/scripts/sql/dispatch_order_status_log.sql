CREATE TABLE `dispatch_order_status_log`
(
    `id`                bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `order_no`          varchar(64)  NOT NULL COMMENT '配送单号',
    `delivery_code`     varchar(32)  NOT NULL DEFAULT '' COMMENT '配送平台代码',
    `delivery_order_no` varchar(128) NOT NULL DEFAULT '' COMMENT '平台订单号',
    `old_status`        varchar(32)  NOT NULL DEFAULT '' COMMENT '原状态',
    `new_status`        varchar(32)  NOT NULL COMMENT '新状态',
    `status_desc`       varchar(255) NOT NULL DEFAULT '' COMMENT '状态描述',
    `remark`            varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
    `created_at`        datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY                 `idx_order_no` (`order_no`),
    KEY                 `idx_delivery_order` (`delivery_code`, `delivery_order_no`),
    KEY                 `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='订单状态流转记录表';