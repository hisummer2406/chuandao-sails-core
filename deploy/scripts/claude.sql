-- =====================================================
-- 聚合配送策略服务数据表设计 (PolarDB MySQL 8.0) - JSON优化版本
-- =====================================================

-- 1. 配送策略配置主表 (JSON存储)
CREATE TABLE `delivery_strategy_configs`
(
    `id`              BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `config_key`      VARCHAR(64)  NOT NULL COMMENT '配置键名(如: default, beijing, shanghai)',
    `config_name`     VARCHAR(128) NOT NULL COMMENT '配置名称',
    `config_type`     ENUM('default', 'city') NOT NULL DEFAULT 'default' COMMENT '配置类型',
    `city_code`       VARCHAR(16) NULL COMMENT '城市编码(city类型时使用)',
    `city_name`       VARCHAR(64) NULL COMMENT '城市名称',
    `version`         VARCHAR(16)  NOT NULL DEFAULT '1.0.0' COMMENT '配置版本',
    `enabled`         TINYINT(1) NOT NULL DEFAULT 1 COMMENT '是否启用',
    `inherit_default` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '是否继承默认配置',
    `strategy_config` JSON         NOT NULL COMMENT '完整策略配置JSON',
    `config_hash`     VARCHAR(64) GENERATED ALWAYS AS (SHA2(strategy_config, 256)) STORED COMMENT '配置哈希值',
    `created_by`      VARCHAR(64)  NOT NULL DEFAULT 'system' COMMENT '创建人',
    `updated_by`      VARCHAR(64)  NOT NULL DEFAULT 'system' COMMENT '更新人',
    `created_at`      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`      TIMESTAMP NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_config_key` (`config_key`),
    UNIQUE KEY `uk_city_code` (`city_code`),
    KEY               `idx_config_type` (`config_type`),
    KEY               `idx_enabled` (`enabled`),
    KEY               `idx_config_hash` (`config_hash`),
    KEY               `idx_version` (`version`),
    KEY               `idx_created_at` (`created_at`),
    -- JSON字段索引 (MySQL 8.0 多值索引)
    INDEX             `idx_platforms` ((CAST(JSON_KEYS(JSON_EXTRACT(strategy_config, '$.platforms')) AS CHAR(255) ARRAY))),
    INDEX             `idx_strategy_type` ((CAST(JSON_UNQUOTE(JSON_EXTRACT(strategy_config, '$.strategy_type')) AS CHAR(32))))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='配送策略配置表(JSON存储)';

-- 为JSON字段创建虚拟列以提升查询性能
ALTER TABLE `delivery_strategy_configs`
    ADD COLUMN `strategy_type_virtual` VARCHAR(32)
        GENERATED ALWAYS AS (JSON_UNQUOTE(JSON_EXTRACT(strategy_config, '$.strategy_type'))) STORED,
ADD INDEX `idx_strategy_type_virtual` (`strategy_type_virtual`);

ALTER TABLE `delivery_strategy_configs`
    ADD COLUMN `fallback_enabled_virtual` TINYINT(1)
GENERATED ALWAYS AS (JSON_EXTRACT(strategy_config, '$.fallback_enabled')) STORED,
ADD INDEX `idx_fallback_enabled_virtual` (`fallback_enabled_virtual`);

-- 2. 配送订单表
CREATE TABLE `delivery_orders`
(
    `id`                   BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `order_id`             VARCHAR(64) NOT NULL COMMENT '订单ID',
    `business_order_id`    VARCHAR(64) NOT NULL COMMENT '业务订单ID',
    `city_code`            VARCHAR(16) NOT NULL COMMENT '城市编码',
    `order_info`           JSON        NOT NULL COMMENT '订单详细信息',
    `order_status`         ENUM('pending', 'quoted', 'assigned', 'picked_up', 'delivered', 'cancelled', 'failed') NOT NULL DEFAULT 'pending' COMMENT '订单状态',
    `selected_platform_id` VARCHAR(32) NULL COMMENT '选中的平台ID',
    `selected_quote_id`    VARCHAR(64) NULL COMMENT '选中的报价ID',
    `actual_fee`           DECIMAL(10, 2) NULL COMMENT '实际费用',
    `delivery_result`      JSON NULL COMMENT '配送结果信息',
    `created_at`           TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`           TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_order_id` (`order_id`),
    KEY                    `idx_business_order_id` (`business_order_id`),
    KEY                    `idx_city_code` (`city_code`),
    KEY                    `idx_order_status` (`order_status`),
    KEY                    `idx_platform_id` (`selected_platform_id`),
    KEY                    `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
PARTITION BY RANGE (TO_DAYS(created_at)) (
    PARTITION p202501 VALUES LESS THAN (TO_DAYS('2025-02-01')),
    PARTITION p202502 VALUES LESS THAN (TO_DAYS('2025-03-01')),
    PARTITION p202503 VALUES LESS THAN (TO_DAYS('2025-04-01')),
    PARTITION p202504 VALUES LESS THAN (TO_DAYS('2025-05-01')),
    PARTITION p202505 VALUES LESS THAN (TO_DAYS('2025-06-01')),
    PARTITION p202506 VALUES LESS THAN (TO_DAYS('2025-07-01')),
    PARTITION pmax VALUES LESS THAN MAXVALUE
) COMMENT='配送订单表';

-- 为订单JSON字段创建虚拟列
ALTER TABLE `delivery_orders`
    ADD COLUMN `distance_km_virtual` DECIMAL(8, 2)
        GENERATED ALWAYS AS (JSON_UNQUOTE(JSON_EXTRACT(order_info, '$.distance_km'))) STORED,
ADD COLUMN `weight_kg_virtual` DECIMAL(8,2)
GENERATED ALWAYS AS (JSON_UNQUOTE(JSON_EXTRACT(order_info, '$.weight_kg'))) STORED,
ADD COLUMN `order_amount_virtual` DECIMAL(12,2)
GENERATED ALWAYS AS (JSON_UNQUOTE(JSON_EXTRACT(order_info, '$.order_amount'))) STORED,
ADD COLUMN `is_cod_virtual` TINYINT(1)
GENERATED ALWAYS AS (JSON_EXTRACT(order_info, '$.is_cod')) STORED,
ADD INDEX `idx_distance_km` (`distance_km_virtual`),
ADD INDEX `idx_weight_kg` (`weight_kg_virtual`),
ADD INDEX `idx_order_amount` (`order_amount_virtual`),
ADD INDEX `idx_is_cod` (`is_cod_virtual`);

-- 3. 平台报价记录表
CREATE TABLE `delivery_platform_quotes`
(
    `id`             BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `quote_id`       VARCHAR(64) NOT NULL COMMENT '报价ID',
    `order_id`       VARCHAR(64) NOT NULL COMMENT '订单ID',
    `platform_id`    VARCHAR(32) NOT NULL COMMENT '平台ID',
    `quote_data`     JSON        NOT NULL COMMENT '完整报价信息',
    `is_selected`    TINYINT(1) NOT NULL DEFAULT 0 COMMENT '是否被选中',
    `selection_info` JSON NULL COMMENT '选择信息(原因、评分等)',
    `quote_time`     TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '报价时间',
    `expire_time`    TIMESTAMP   NOT NULL COMMENT '过期时间',
    `created_at`     TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_quote_id` (`quote_id`),
    KEY              `idx_order_id` (`order_id`),
    KEY              `idx_platform_id` (`platform_id`),
    KEY              `idx_quote_time` (`quote_time`),
    KEY              `idx_expire_time` (`expire_time`),
    KEY              `idx_is_selected` (`is_selected`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
PARTITION BY RANGE (TO_DAYS(created_at)) (
    PARTITION p202501 VALUES LESS THAN (TO_DAYS('2025-02-01')),
    PARTITION p202502 VALUES LESS THAN (TO_DAYS('2025-03-01')),
    PARTITION p202503 VALUES LESS THAN (TO_DAYS('2025-04-01')),
    PARTITION p202504 VALUES LESS THAN (TO_DAYS('2025-05-01')),
    PARTITION p202505 VALUES LESS THAN (TO_DAYS('2025-06-01')),
    PARTITION p202506 VALUES LESS THAN (TO_DAYS('2025-07-01')),
    PARTITION pmax VALUES LESS THAN MAXVALUE
) COMMENT='平台报价记录表';

-- 为报价JSON字段创建虚拟列
ALTER TABLE `delivery_platform_quotes`
    ADD COLUMN `total_price_virtual` DECIMAL(10, 2)
        GENERATED ALWAYS AS (JSON_UNQUOTE(JSON_EXTRACT(quote_data, '$.total_price'))) STORED,
ADD COLUMN `available_virtual` TINYINT(1)
GENERATED ALWAYS AS (JSON_EXTRACT(quote_data, '$.available')) STORED,
ADD COLUMN `estimated_time_virtual` INT
GENERATED ALWAYS AS (JSON_EXTRACT(quote_data, '$.estimated_delivery_time_minutes')) STORED,
ADD INDEX `idx_total_price` (`total_price_virtual`),
ADD INDEX `idx_available` (`available_virtual`),
ADD INDEX `idx_estimated_time` (`estimated_time_virtual`);

-- 4. 平台状态监控表
CREATE TABLE `delivery_platform_monitor`
(
    `id`           BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `platform_id`  VARCHAR(32) NOT NULL COMMENT '平台ID',
    `city_code`    VARCHAR(16) NOT NULL COMMENT '城市编码',
    `monitor_data` JSON        NOT NULL COMMENT '监控数据',
    `monitor_time` TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '监控时间',
    `created_at`   TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY            `idx_platform_city` (`platform_id`, `city_code`),
    KEY            `idx_monitor_time` (`monitor_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='平台状态监控表';

-- 5. 系统配置表
CREATE TABLE `delivery_system_settings`
(
    `id`             BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `setting_group`  VARCHAR(64)  NOT NULL COMMENT '设置分组',
    `setting_name`   VARCHAR(128) NOT NULL COMMENT '设置名称',
    `setting_config` JSON         NOT NULL COMMENT '设置配置',
    `enabled`        TINYINT(1) NOT NULL DEFAULT 1 COMMENT '是否启用',
    `description`    TEXT NULL COMMENT '设置描述',
    `created_at`     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_group_name` (`setting_group`),
    KEY              `idx_enabled` (`enabled`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='系统配置表';

-- 6. 配置变更日志表
CREATE TABLE `delivery_config_change_logs`
(
    `id`             BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `config_id`      BIGINT UNSIGNED NOT NULL COMMENT '配置ID',
    `config_type`    VARCHAR(64) NOT NULL COMMENT '配置类型',
    `operation_type` ENUM('create', 'update', 'delete', 'enable', 'disable') NOT NULL COMMENT '操作类型',
    `change_data`    JSON        NOT NULL COMMENT '变更数据(包含before/after)',
    `operator`       VARCHAR(64) NOT NULL COMMENT '操作人',
    `operator_ip`    VARCHAR(45) NULL COMMENT '操作IP',
    `change_reason`  VARCHAR(255) NULL COMMENT '变更原因',
    `created_at`     TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY              `idx_config_id` (`config_id`),
    KEY              `idx_config_type` (`config_type`),
    KEY              `idx_operator` (`operator`),
    KEY              `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='配置变更日志表';

-- =====================================================
-- 初始化数据
-- =====================================================

-- 插入默认配置
INSERT INTO `delivery_strategy_configs`
(`config_key`, `config_name`, `config_type`, `strategy_config`, `created_by`)
VALUES ('default', '默认配送策略', 'default', '{
  "enabled": true,
  "strategy_type": "price_priority",
  "fallback_enabled": true,
  "max_retry_times": 3,
  "timeout_seconds": 30,
  "platforms": {
    "uu_delivery": {
      "platform_id": "uu",
      "platform_name": "UU跑腿",
      "enabled": true,
      "priority": 1,
      "weight": 100,
      "pricing": {
        "base_price": 5.0,
        "distance_price_per_km": 2.0,
        "weight_price_per_kg": 1.0,
        "time_multiplier": {
          "peak_hours": 1.5,
          "normal_hours": 1.0,
          "off_peak_hours": 0.8
        }
      },
      "service_config": {
        "max_distance_km": 20,
        "max_weight_kg": 30,
        "estimated_delivery_time_minutes": 45,
        "support_cod": true,
        "support_fragile": true,
        "support_cold_chain": false
      },
      "api_config": {
        "endpoint": "https://api.uu.com/delivery",
        "timeout_ms": 5000,
        "retry_times": 2,
        "auth": {
          "type": "api_key",
          "key_field": "X-API-Key"
        }
      }
    },
    "dada_delivery": {
      "platform_id": "dada",
      "platform_name": "达达快送",
      "enabled": true,
      "priority": 2,
      "weight": 90,
      "pricing": {
        "base_price": 4.5,
        "distance_price_per_km": 1.8,
        "weight_price_per_kg": 0.8,
        "time_multiplier": {
          "peak_hours": 1.4,
          "normal_hours": 1.0,
          "off_peak_hours": 0.9
        }
      },
      "service_config": {
        "max_distance_km": 15,
        "max_weight_kg": 25,
        "estimated_delivery_time_minutes": 40,
        "support_cod": true,
        "support_fragile": false,
        "support_cold_chain": true
      },
      "api_config": {
        "endpoint": "https://api.dada.com/delivery",
        "timeout_ms": 4000,
        "retry_times": 3,
        "auth": {
          "type": "oauth2",
          "token_endpoint": "https://api.dada.com/oauth/token"
        }
      }
    }
  },
  "strategy_rules": {
    "price_priority": {
      "enabled": true,
      "weight": 0.6,
      "consider_service_fee": true
    },
    "time_priority": {
      "enabled": true,
      "weight": 0.3,
      "max_acceptable_delay_minutes": 15
    },
    "reliability_priority": {
      "enabled": true,
      "weight": 0.1,
      "min_success_rate": 0.95
    }
  },
  "business_rules": {
    "peak_hours": [
      "11:00-13:00",
      "17:00-20:00"
    ],
    "blacklist_check": true,
    "distance_limit_km": 25,
    "weight_limit_kg": 40,
    "cod_limit_amount": 5000.0
  }
}', 'system');

-- 插入北京特殊配置
INSERT INTO `delivery_strategy_configs`
(`config_key`, `config_name`, `config_type`, `city_code`, `city_name`, `strategy_config`, `created_by`)
VALUES ('beijing', '北京配送策略', 'city', '110000', '北京市', '{
  "enabled": true,
  "strategy_type": "time_priority",
  "inherit_default": true,
  "platforms": {
    "uu_delivery": {
      "platform_id": "uu",
      "platform_name": "UU跑腿",
      "enabled": true,
      "priority": 2,
      "weight": 95,
      "pricing": {
        "base_price": 6.0,
        "distance_price_per_km": 2.5,
        "weight_price_per_kg": 1.2,
        "time_multiplier": {
          "peak_hours": 1.8,
          "normal_hours": 1.2,
          "off_peak_hours": 1.0
        }
      },
      "service_config": {
        "max_distance_km": 25,
        "estimated_delivery_time_minutes": 40
      }
    },
    "dada_delivery": {
      "platform_id": "dada",
      "platform_name": "达达快送",
      "enabled": true,
      "priority": 1,
      "weight": 100,
      "pricing": {
        "base_price": 5.5,
        "distance_price_per_km": 2.2,
        "weight_price_per_kg": 1.0
      }
    }
  },
  "business_rules": {
    "peak_hours": [
      "11:00-14:00",
      "16:30-20:30"
    ],
    "distance_limit_km": 30,
    "special_zones": [
      {
        "zone_name": "CBD区域",
        "coordinates": [
          [
            116.45,
            39.92
          ],
          [
            116.48,
            39.94
          ]
        ],
        "additional_fee": 3.0,
        "priority_platforms": [
          "dada_delivery",
          "uu_delivery"
        ]
      }
    ]
  }
}', 'system');

-- 插入系统全局设置
INSERT INTO `delivery_system_settings`
    (`setting_group`, `setting_name`, `setting_config`, `description`)
VALUES ('global', '全局系统配置', '{
  "monitoring": {
    "success_rate_threshold": 0.90,
    "response_time_threshold_ms": 10000,
    "alert_channels": [
      "email",
      "sms",
      "webhook"
    ]
  },
  "fallback_strategy": {
    "enabled": true,
    "fallback_platforms": [
      "uu_delivery",
      "dada_delivery"
    ],
    "max_fallback_attempts": 2
  },
  "cache_settings": {
    "quote_cache_ttl_seconds": 300,
    "platform_status_cache_ttl_seconds": 60
  },
  "rate_limiting": {
    "requests_per_minute": 1000,
    "requests_per_platform_per_minute": 300
  }
}', '全局系统配置');

-- =====================================================
-- 创建有用的视图和函数
-- =====================================================

-- 创建获取有效配置的视图
CREATE VIEW `v_effective_delivery_configs` AS
SELECT id,
       config_key,
       config_name,
       config_type,
       city_code,
       city_name,
       version,
       enabled,
       strategy_config,
       CASE
           WHEN config_type = 'city' THEN
               JSON_MERGE_PATCH(
                       COALESCE((SELECT strategy_config
                                 FROM delivery_strategy_configs
                                 WHERE config_key = 'default' AND enabled = 1), JSON_OBJECT()),
                       strategy_config
               )
           ELSE strategy_config
           END AS effective_config,
       created_at,
       updated_at
FROM delivery_strategy_configs
WHERE enabled = 1
  AND deleted_at IS NULL;

-- 创建获取城市配置的存储过程
DELIMITER
//
CREATE PROCEDURE `GetCityDeliveryConfig`(
    IN p_city_code VARCHAR (16)
)
BEGIN
    DECLARE
config_exists INT DEFAULT 0;

    -- 检查城市配置是否存在
SELECT COUNT(*)
INTO config_exists
FROM delivery_strategy_configs
WHERE city_code = p_city_code
  AND enabled = 1
  AND deleted_at IS NULL;

IF
config_exists > 0 THEN
        -- 返回城市特定配置
SELECT *
FROM v_effective_delivery_configs
WHERE city_code = p_city_code;
ELSE
        -- 返回默认配置
SELECT *
FROM v_effective_delivery_configs
WHERE config_key = 'default';
END IF;
END
//
DELIMITER ;

-- 创建JSON配置验证函数
DELIMITER
//
CREATE FUNCTION `ValidateDeliveryConfig`(config_json JSON)
    RETURNS BOOLEAN
    READS SQL DATA
    DETERMINISTIC
BEGIN
    DECLARE
is_valid BOOLEAN DEFAULT FALSE;

    -- 检查必需字段
    IF
JSON_CONTAINS_PATH(config_json, 'one', '$.enabled', '$.strategy_type', '$.platforms') THEN
        SET is_valid = TRUE;
END IF;

RETURN is_valid;
END
//
DELIMITER ;

-- =====================================================
-- JSON Schema 验证约束 (MySQL 8.0.17+)
-- =====================================================
ALTER TABLE `delivery_strategy_configs`
    ADD CONSTRAINT `chk_strategy_config_schema`
        CHECK (JSON_SCHEMA_VALID('{
  "type": "object",
  "required": ["enabled", "strategy_type"],
  "properties": {
    "enabled": {"type": "boolean"},
    "strategy_type": {"type": "string", "enum": ["price_priority", "time_priority", "reliability_priority", "hybrid"]},
    "platforms": {
      "type": "object",
      "patternProperties": {
        "^[a-z_]+$": {
          "type": "object",
          "required": ["platform_id", "enabled"],
          "properties": {
            "platform_id": {"type": "string"},
            "enabled": {"type": "boolean"},
            "priority": {"type": "integer", "minimum": 1},
            "pricing": {
              "type": "object",
              "properties": {
                "base_price": {"type": "number", "minimum": 0},
                "distance_price_per_km": {"type": "number", "minimum": 0}
              }
            }
          }
        }
      }
    }
  }
}', strategy_config));

-- =====================================================
-- 性能优化建议
-- =====================================================
/*
JSON存储方案的优势：
1. 配置结构灵活，易于扩展
2. 减少表关联，查询更简单
3. 配置整体性强，原子操作
4. 减少数据库表数量，维护简单

性能优化要点：
1. 为常用JSON字段创建虚拟列和索引
2. 使用JSON_EXTRACT优化查询
3. 合理使用JSON_MERGE_PATCH进行配置合并
4. 利用MySQL 8.0的多值索引特性
5. 配置变更使用事务保证一致性

PolarDB优化配置：
- 启用并行查询提升JSON处理性能
- 合理配置innodb_buffer_pool_size
- 使用读写分离，读请求分发到只读节点
- 启用查询缓存提升重复查询性能
*/