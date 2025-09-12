#  配送调度域（delivery-dispatch）：配送平台对接与管理（platform）、动态定价与询价策略（pricing）、发单策略与路由规则（strategy）、发单服务（dispatch）

## 1. 目录结构
```text
delivery-dispatch-service/
├── api/                             # API定义目录
├── rpc/                             # RPC服务目录
├── job/                             # 定时任务目录
├── mq/                              # 消息队列消费者目录
├── model/                           # 数据模型目录
├── pkg/                             # 业务包目录
│   ├── engine/                      # 核心引擎 // 一个服务，四个引擎：平台管理、定价、策略、发单
│   │   ├── platform/                # 平台管理引擎
│   │   │   ├── manager.go           # 平台管理器
│   │   │   ├── adapter/             # 平台适配器
│   │   │   │   ├── interface.go     # 适配器接口
│   │   │   │   ├── uu_adapter.go    # UU跑腿适配器
│   │   │   │   ├── flash_adapter.go # 闪送适配器
│   │   │   │   ├── sf_adapter.go    # 顺丰适配器
│   │   │   │   └── fengniao_adapter.go # 蜂鸟适配器
│   │   │   └── capability.go        # 能力管理
│   │   ├── pricing/                 # 定价引擎 
│   │   │   ├── engine.go           # 定价引擎主类
│   │   │   ├── calculator/         # 价格计算器
│   │   │   │   ├── interface.go    # 计算器接口
│   │   │   │   ├── base_calculator.go # 基础计算器
│   │   │   │   ├── dynamic_calculator.go # 动态计算器
│   │   │   │   └── regional_calculator.go # 区域计算器
│   │   │   ├── strategy/           # 定价策略
│   │   │   │   ├── city_strategy.go # 城市策略
│   │   │   │   ├── time_strategy.go # 时间策略
│   │   │   │   └── distance_strategy.go # 距离策略
│   │   │   └── cache/              # 价格缓存
│   │   │       ├── price_cache.go  # 价格缓存管理
│   │   │       └── cache_warmer.go # 缓存预热
│   │   ├── strategy/               # 策略引擎
│   │   │   ├── engine.go          # 策略引擎主类
│   │   │   ├── rule/              # 规则引擎
│   │   │   │   ├── interface.go   # 规则接口
│   │   │   │   ├── priority_rule.go # 优先级规则
│   │   │   │   ├── price_rule.go  # 价格规则
│   │   │   │   ├── time_rule.go   # 时效规则
│   │   │   │   └── capacity_rule.go # 运力规则
│   │   │   ├── router/            # 路由算法
│   │   │   │   ├── sequential_router.go # 顺序路由
│   │   │   │   ├── concurrent_router.go # 并发路由
│   │   │   │   └── intelligent_router.go # 智能路由
│   │   │   └── abtest/            # A/B测试
│   │   │       ├── experiment.go  # 实验管理
│   │   │       └── analyzer.go    # 效果分析
│   │   └── dispatch/              # 发单引擎
│   │       ├── engine.go         # 发单引擎主类
│   │       ├── executor/         # 执行器
│   │       │   ├── interface.go  # 执行器接口
│   │       │   ├── sequential_executor.go # 顺序执行器
│   │       │   ├── concurrent_executor.go # 并发执行器
│   │       │   └── batch_executor.go # 批量执行器
│   │       ├── tracker/          # 状态追踪
│   │       │   ├── status_tracker.go # 状态追踪器
│   │       │   └── callback_handler.go # 回调处理器
│   │       └── retry/            # 重试机制
│   │           ├── retry_manager.go # 重试管理器
│   │           └── backoff.go    # 退避算法
│   ├── client/                   # 外部客户端
│   │   ├── platform_client.go   # 配送平台客户端
│   │   ├── order_client.go      # 订单服务客户端
│   │   └── finance_client.go    # 财务服务客户端
│   ├── cache/                    # 缓存组件
│   │   ├── redis_cache.go       # Redis缓存
│   │   └── local_cache.go       # 本地缓存
│   ├── mq/                      # 消息队列组件
│   │   ├── producer.go          # 消息生产者
│   │   └── consumer.go          # 消息消费者
│   ├── utils/                   # 工具类
│   │   ├── geo_utils.go        # 地理位置工具
│   │   ├── time_utils.go       # 时间工具
│   │   ├── string_utils.go     # 字符串工具
│   │   └── id_generator.go     # ID生成器
│   └── validator/               # 验证器
│       ├── order_validator.go   # 订单验证器
│       └── platform_validator.go # 平台验证器
├── etc/                         # 配置文件目录
│   ├── delivery-api.yaml       # API服务配置
│   ├── delivery-rpc.yaml       # RPC服务配置  
│   ├── delivery-job.yaml       # 定时任务配置
│   └── delivery-mq.yaml        # MQ消费者配置
├── test/                     # 测试目录
│   ├── api_test.go          # API测试
│   ├── rpc_test.go          # RPC测试
│   └── integration_test.go   # 集成测试
└── README.md                 # 项目说明文档
```

## 2. 微服务划分

3. delivery-service（配送服务）

* 配送平台管理
* 配送能力维护
* 配送状态同步
* 配送平台适配器

4. pricing-service（定价服务）

* 动态询价
* 价格策略引擎
* 城市/区域定价
* 实时价格缓存

5. dispatch-service（发单服务）

* 发单策略执行
* 智能路由分配
* 并发发单控制
* 发单重试机制

7. strategy-service（策略服务）

* 发单策略配置
* 规则引擎
* A/B测试支持
* 策略效果分析