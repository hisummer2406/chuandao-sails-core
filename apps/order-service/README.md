# 订单域（Order）：订单生命周期管理

## 1. 目录结构

```text
order-service/
├── api/                                    # HTTP API服务
├── rpc/                                   # gRPC服务
├── mq/                                   # 消息队列消费者
│   ├── internal/
│   │   ├── config/
│   │   │   └── config.go                # MQ配置
│   │   ├── logic/
│   │   │   ├── order_create_consumer.go  # 订单创建消息消费
│   │   │   ├── status_update_consumer.go # 状态更新消息消费
│   │   │   ├── callback_consumer.go      # 回调消息消费
│   │   │   └── retry_consumer.go         # 重试消息消费
│   │   └── svc/
│   │       └── servicecontext.go        # MQ服务上下文
│   └── mq.go                            # MQ入口文件
├── job/                                 # 定时任务
│   ├── internal/
│   │   ├── config/
│   │   │   └── config.go               # 任务配置
│   │   ├── logic/
│   │   │   ├── order_timeout_job.go    # 订单超时处理任务
│   │   │   ├── penalty_settlement_job.go # 违约金结算任务
│   │   │   ├── data_archive_job.go     # 数据归档任务
│   │   │   └── status_sync_job.go      # 状态同步任务
│   │   └── svc/
│   │       └── servicecontext.go       # 任务服务上下文
│   └── job.go                          # 任务入口文件
├── model/                              # 数据模型
├── pkg/                                # 业务包
│   ├── engine/                         # 核心引擎
│   │   ├── lifecycle/                  # 订单生命周期引擎。管理订单从创建到完成的整个生命周期：
│   │   │   ├── manager.go             # 生命周期管理器
│   │   │   ├── state/                 # 状态机。定义订单状态转换规则
│   │   │   │   ├── state_machine.go   # 状态机核心
│   │   │   │   ├── transitions.go     # 状态转换规则
│   │   │   │   └── validators.go      # 状态转换验证
│   │   │   ├── processor/             # 状态处理器。每个状态都有对应的处理器
│   │   │   │   ├── created_processor.go    # 已创建状态处理
│   │   │   │   ├── confirmed_processor.go  # 已确认状态处理
│   │   │   │   ├── assigned_processor.go   # 已分配状态处理
│   │   │   │   ├── picked_processor.go     # 已取件状态处理
│   │   │   │   ├── delivered_processor.go  # 已送达状态处理
│   │   │   │   ├── cancelled_processor.go  # 已取消状态处理
│   │   │   │   └── completed_processor.go  # 已完成状态处理
│   │   │   └── event/                 # 事件处理。状态变化时发布事件通知其他服务
│   │   │       ├── event_publisher.go # 事件发布器。事件驱动架构是核心设计模式
│   │   │       └── event_handlers.go  # 事件处理器
│   │   ├── penalty/                   # 违约金引擎。专门处理订单取消和违约情况：
│   │   │   ├── calculator.go         # 违约金计算器。根据不同规则计算违约金
│   │   │   ├── rule/                 # 违约规则
│   │   │   │   ├── interface.go      # 规则接口
│   │   │   │   ├── time_rule.go      # 时间违约规则
│   │   │   │   ├── status_rule.go    # 状态违约规则
│   │   │   │   └── amount_rule.go    # 金额违约规则
│   │   │   ├── strategy/             # 计算策略
│   │   │   │   ├── fixed_strategy.go # 固定金额策略
│   │   │   │   ├── ratio_strategy.go # 比例策略
│   │   │   │   └── tiered_strategy.go # 阶梯策略
│   │   │   └── recorder.go           # 违约记录器。记录违约历史和原因
│   │   ├── track/                    # 轨迹追踪引擎。跟踪订单执行过程：
│   │   │   ├── tracker.go           # 轨迹追踪器
│   │   │   ├── collector/           # 数据收集器。从各个平台收集轨迹数据
│   │   │   │   ├── platform_collector.go # 平台数据收集
│   │   │   │   └── internal_collector.go # 内部数据收集
│   │   │   ├── analyzer/            # 轨迹分析器。分析时间线和异常情况
│   │   │   │   ├── timeline_analyzer.go # 时间线分析
│   │   │   │   └── anomaly_analyzer.go  # 异常分析
│   │   │   └── formatter/           # 数据格式化。统一输出格式
│   │   │       ├── timeline_formatter.go # 时间线格式化
│   │   │       └── summary_formatter.go  # 摘要格式化
│   │   └── validation/              # 订单验证引擎
│   │       ├── validator.go         # 验证器主类
│   │       ├── rule/                # 验证规则
│   │       │   ├── business_rule.go # 业务规则验证
│   │       │   ├── data_rule.go     # 数据规则验证
│   │       │   └── constraint_rule.go # 约束规则验证
│   │       └── sanitizer/           # 数据清洗
│   │           ├── address_sanitizer.go # 地址清洗
│   │           └── phone_sanitizer.go   # 电话清洗
│   ├── client/                      # 外部客户端
│   │   ├── delivery_client.go       # 配送服务客户端
│   │   ├── finance_client.go        # 财务服务客户端
│   │   ├── platform_client.go       # 平台网关客户端
│   │   └── notification_client.go   # 通知服务客户端
│   ├── cache/                       # 缓存组件
│   │   ├── order_cache.go          # 订单缓存
│   │   ├── status_cache.go         # 状态缓存
│   │   └── sequence_cache.go       # 序列号缓存
│   ├── mq/                         # 消息队列组件
│   │   ├── producer.go             # 消息生产者
│   │   ├── event_publisher.go      # 事件发布器
│   │   └── message_formatter.go    # 消息格式化
│   ├── utils/                      # 工具类
│   │   ├── order_utils.go         # 订单工具
│   │   ├── id_generator.go        # ID生成器
│   │   ├── time_utils.go          # 时间工具
│   │   └── phone_utils.go         # 电话号码工具
│   ├── constants/                  # 常量定义
│   │   ├── order_status.go        # 订单状态常量
│   │   ├── order_type.go          # 订单类型常量
│   │   ├── penalty_type.go        # 违约类型常量
│   │   └── event_type.go          # 事件类型常量
│   └── validator/                  # 验证器
│       ├── order_validator.go      # 订单数据验证
│       ├── status_validator.go     # 状态验证
│       └── phone_validator.go      # 手机号验证
├── etc/                           # 配置文件
├── test/                          # 测试
│   ├── api_test.go               # API测试
│   ├── rpc_test.go               # RPC测试
│   ├── integration_test.go        # 集成测试
│   ├── state_machine_test.go      # 状态机测试
│   └── penalty_test.go           # 违约金测试
└── README.md                      # 项目说明文档
```

## 2. 微服务划分

1. order-service（订单服务）

* 订单状态管理
* 订单生命周期追踪
* 订单取消处理
* 违约金计算

### 订单业务引擎模块化

**1. lifecycle/ - 订单生命周期引擎**
这是订单服务的核心，管理订单从创建到完成的整个生命周期：

* state 状态机：定义订单状态转换规则
* processor 处理器：每个状态都有对应的处理器
* event 事件发布：状态变化时发布事件通知其他服务

**2. penalty/ - 违约金引擎**
专门处理订单取消和违约情况：

* 计算器：根据不同规则计算违约金
* 策略模式：支持固定金额、比例、阶梯等多种计算方式
* 记录器：记录违约历史和原因

**3. track/ - 轨迹追踪引擎**
跟踪订单执行过程：

* 数据收集：从各个平台收集轨迹数据
* 分析器：分析时间线和异常情况
* 格式化：统一输出格式

## 3. 限界上下文

### 1. 先MQ异步处理，再创建订单的方案
- 临时订单号机制
```text
立即返回：TMP_20240904_001234567
异步生成：ORD_20240904_001234567
```
- 幂等性保证
   * 基于临时订单号去重
   * 分布式锁防止并发处理
   * 双重检查确保数据一致性
- 监控告警
   * MQ消息堆积监控
   * 订单处理成功率监控
   * 异常订单补偿机制

### 2. 订单状态机设计
```go
// 订单状态枚举
const (
    OrderStatusCreated     = "CREATED"     // 已创建
    OrderStatusConfirmed   = "CONFIRMED"   // 已确认
    OrderStatusAssigned    = "ASSIGNED"    // 已分配配送员
    OrderStatusPicked      = "PICKED"      // 已取件
    OrderStatusDelivered   = "DELIVERED"   // 已送达
    OrderStatusCompleted   = "COMPLETED"   // 已完成
    OrderStatusCancelled   = "CANCELLED"   // 已取消
)

// 状态转换规则
var StateTransitions = map[string][]string{
    OrderStatusCreated:   {OrderStatusConfirmed, OrderStatusCancelled},
    OrderStatusConfirmed: {OrderStatusAssigned, OrderStatusCancelled},
    OrderStatusAssigned:  {OrderStatusPicked, OrderStatusCancelled},
    OrderStatusPicked:    {OrderStatusDelivered},
    OrderStatusDelivered: {OrderStatusCompleted},
}
```

### 3. 关键业务流程

**1. 订单创建流程**
```text
MQ接收订单消息 → 数据验证 → 生成订单号 → 写入数据库 → 发布订单创建事件 → 触发配送调度
```
**2.状态更新流程**
```text
接收状态更新 → 状态机验证 → 更新数据库 → 记录状态日志 → 发布状态变更事件 → 通知相关服务
```
**3. 违约金处理流程**
```text
订单取消 → 违约规则匹配 → 计算违约金 → 记录违约记录 → 通知财务服务 → 发送违约通知
```

### 4. 事件驱动架构



### 5. 订单状态机