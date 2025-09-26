SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
-- ----------------------------
-- 删除原表
-- ----------------------------
DROP TABLE IF EXISTS `platform_push_order_master`;

-- ----------------------------
-- 创建带二级分区的表结构
-- ----------------------------
CREATE TABLE `platform_push_order_master`
(

    `id`                bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `order_no`          varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  NOT NULL COMMENT '系统内部订单号',
    `upstream_order_id` varchar(128) COLLATE utf8mb4_unicode_ci                       NOT NULL COMMENT '上游平台订单号',
    `platform_code`     varchar(16) COLLATE utf8mb4_unicode_ci                        NOT NULL COMMENT '平台标识：UU/SF/CHD',
    `order_source`      varchar(16) COLLATE utf8mb4_unicode_ci                                 DEFAULT '' COMMENT '订单来源：MT/SG/EL/JDDJ等',
    `order_num`         varchar(64) COLLATE utf8mb4_unicode_ci                                 DEFAULT '' COMMENT '订单流水号',
    `city_name`         varchar(32) COLLATE utf8mb4_unicode_ci                                 DEFAULT '' COMMENT '城市名称',
    `county_name`       varchar(32) COLLATE utf8mb4_unicode_ci                                 DEFAULT '' COMMENT '县级地名称',
    `ad_code`           varchar(16) COLLATE utf8mb4_unicode_ci                                 DEFAULT '' COMMENT '城市编码（高德规范）',
    `send_type`         tinyint(1) NOT NULL DEFAULT '0' COMMENT '订单小类：0帮我送 1帮我买 2帮我取',
    `delivery_type`     varchar(8) COLLATE utf8mb4_unicode_ci                                  DEFAULT '' COMMENT '配送类型：1团送 2专送',
    `is_reverse_order`  tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否帮我取订单：1是 0否',
    `push_type`         tinyint(1) NOT NULL DEFAULT '0' COMMENT '推送类型：0正常 1测试',
    `from_address`      json                                                                   DEFAULT NULL COMMENT '发货地址信息',
    `to_address`        json                                                                   DEFAULT NULL COMMENT '收货地址信息',
    `order_time`        datetime                                                               DEFAULT NULL COMMENT '订单时间',
    `is_subscribe`      tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否预约单：1是 0否',
    `subscribe_type`    tinyint(1) NOT NULL DEFAULT '0' COMMENT '预约类型：0实时 1预约取件 2预约送达',
    `subscribe_time`    bigint                                                                 DEFAULT '0' COMMENT '预约时间戳',
    `goods_info`        json                                                                   DEFAULT NULL COMMENT '商品详情信息',
    `price_info`        json                                                                   DEFAULT NULL COMMENT '价格信息',
    `delivery_options`  json                                                                   DEFAULT NULL COMMENT '配送选项',
    `note`              text COLLATE utf8mb4_unicode_ci COMMENT '备注信息',
    `disable_delivery`  varchar(128) COLLATE utf8mb4_unicode_ci                                DEFAULT '' COMMENT '禁用配送方',
    `status`            tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态：1待处理 2已转换 3处理失败',
    `error_msg`         text COLLATE utf8mb4_unicode_ci COMMENT '错误信息',
    `retry_count`       tinyint                                                       NOT NULL DEFAULT '0' COMMENT '重试次数',
    `city_code`         varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '城市编码(List Default Hash 二级分区键)',
    `created_at`        DATETIME                                                    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间(Interval Range 分区键)',
    `updated_at`        DATETIME                                                     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',

    -- 主键必须包含所有分区键
    PRIMARY KEY (`id`, `created_at`, `city_code`),

    -- 唯一索引必须包含分区键
    UNIQUE KEY `uk_upstream_order_platform` (`upstream_order_id`, `city_code`, `created_at`) USING BTREE,

    -- 其他索引
    KEY                 `idx_order_no` (`order_no`),
    KEY                 `idx_created_at` (`created_at`),
    KEY                 `idx_city_code` (`city_code`)

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='平台推单主表'

-- 一级分区：INTERVAL RANGE 按月自动分区
PARTITION BY RANGE COLUMNS(created_at)
INTERVAL(MONTH, 1)

-- 二级分区：LIST DEFAULT HASH 按城市分组
SUBPARTITION BY LIST COLUMNS(city_code)
(
  PARTITION p_0 VALUES LESS THAN ('2025-09-01')
  (
    -- 一线城市独享子分区（高频访问）
    SUBPARTITION p_sub_beijing VALUES IN ('010') COMMENT '北京市',
    SUBPARTITION p_sub_shanghai VALUES IN ('021') COMMENT '上海市',
    SUBPARTITION p_sub_guangzhou VALUES IN ('020') COMMENT '广州市',
    SUBPARTITION p_sub_shenzhen VALUES IN ('0755') COMMENT '深圳市',

    -- 新一线城市分组子分区（按你的分组规则）
    SUBPARTITION p_sub_group1 VALUES IN ('028','023','0532','0551') COMMENT '成都/重庆/青岛/合肥',
    SUBPARTITION p_sub_group2 VALUES IN ('0571','027','0574','0769') COMMENT '杭州/武汉/宁波/东莞',
    SUBPARTITION p_sub_group3 VALUES IN ('029','0512','0757','0731') COMMENT '西安/苏州/佛山/长沙',
    SUBPARTITION p_sub_group4 VALUES IN ('0371','025','022') COMMENT '郑州/南京/天津',

    -- 其他城市DEFAULT HASH分区（6个子分区，系统自动分配）
    SUBPARTITION p_sub_others DEFAULT PARTITIONS 6 COMMENT '其他城市HASH分区'
  )
);
