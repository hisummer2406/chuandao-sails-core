船到千帆 - 本地生活聚合配送平台架构设计文档

# 📋 项目概述

## 业务背景
船到千帆是一个本地生活聚合配送平台，在订单高峰期接收京东外卖、美团外卖、淘宝闪购等平台的溢出订单，并调度UU跑腿、闪送、顺丰、蜂鸟等配送平台完成配送，从中赚取配送差价。

## 核心指标

- 日订单量: 300万单
- 峰值QPS: 3500（按2-8规律计算）
- 系统可用性: 99.9%
- 平均响应: <100ms

## 技术栈
```
yaml框架: Go-zero
数据库: 阿里云 PolarDB MySQL 8.0 (读写分离)
消息队列: 阿里云 RocketMQ
缓存: Redis Cluster
服务注册: etcd
监控: ARMS （阿里云Prometheus + Grafana）
日志: 阿里云 SLS
链路追踪: 阿里云 ARMS (Jaeger) 
网关: 阿里云云原生网关
容器编排: 阿里云ACK 

---

## 🏗️ 架构设计

### 整体架构图
┌─────────────────────────────────────────────────────────────┐
│           上游平台: 京东/美团/淘宝                           │
└─────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────┐
│              阿里云云原生网关 + Nginx负载均衡                │
└─────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────┐
│                    微服务层 (4核心+3支撑)                    │
│                                                             │
│  ┌─────────────────┐      ┌─────────────────┐             │
│  │ platform-gateway│ ───► │  order-service  │             │
│  │   (协议转换)     │      │   (订单管理)     │             │
│  └─────────────────┘      └─────────────────┘             │
│                                   ↓                         │
│                    ┌──────────────────────────┐            │
│                    │ delivery-dispatch-service│            │
│                    │   (配送调度 - 核心)       │            │
│                    │  ┌──────────────────┐   │            │
│                    │  │ 平台管理引擎      │   │            │
│                    │  │ 智能询价引擎      │   │            │
│                    │  │ 策略评分引擎      │   │            │
│                    │  │ 发单执行引擎      │   │            │
│                    │  └──────────────────┘   │            │
│                    └──────────────────────────┘            │
│                                   ↓                         │
│           ┌────────────────────────────────┐               │
│           │     finance-service            │               │
│           │     (财务结算)                  │               │
│           └────────────────────────────────┘               │
│                                                             │
│  支撑服务: user-service, notification-service, monitor     │
└─────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────┐
│              RocketMQ消息队列 (事件驱动)                     │
└─────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────┐
│          数据层: PolarDB + Redis + ElasticSearch            │
└─────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────┐
│       下游平台: UU跑腿/闪送/顺丰/蜂鸟                         │
└─────────────────────────────────────────────────────────────┘
```

# 🎯 微服务划分
## 核心业务服务（4个）
### 1. platform-gateway-service（平台网关服务）
职责: 上游平台接入与协议转换

```yaml
功能:
  - 接收UU/顺丰/船到推单
  - 协议转换为标准格式
  - 限流熔断保护
  - 发送MQ解耦

API入口:
  - POST /platform/callback/uu/create
  - POST /platform/callback/sf/create
  - POST /platform/callback/chd/create

输出:
  - upstream-push-topic → order-service
```

**目录结构**:
```yaml
app/platform-gateway/
├── api/internal/logic/uu/      # UU平台处理逻辑
├── pkg/platform/
│   ├── adapter/                # 适配器模式
│   │   ├── uu_adapter.go
│   │   ├── sf_adapter.go
│   │   └── factory.go
│   └── events/                 # 标准事件定义
│       └── standard_order_event.go
```

### 2. order-service（订单服务）
职责: 订单生命周期管理

```yaml
功能:
  - 订单CRUD与状态管理
  - 订单状态机流转
  - 违约金计算
  - 订单轨迹追踪

状态流转:
  CREATED → CONFIRMED → ASSIGNED → PICKED → DELIVERED → COMPLETED

核心引擎:
  - lifecycle: 生命周期管理
  - penalty: 违约金计算
  - track: 轨迹追踪
  - validation: 数据验证
```

核心表设计:
```sql
PolarDB MySQL8.0 

-- 订单主表（二级分区）
platform_push_order_master
  一级分区: created_at (按月 Interval)  
  二级分区: city_code (城市 LIST DEFAULT HASH)
```

### 3. delivery-dispatch-service（配送调度服务）🔥
职责: 一站式配送调度核心
```yaml
四大引擎:
  1. platform引擎: 平台管理与适配
  2. pricing引擎: 智能询价与定价
  3. strategy引擎: 策略评分与路由
  4. dispatch引擎: 发单执行与追踪

处理流程:
  接收订单 → 筛选平台 → 并发询价 → 策略评分 → 智能发单 → 状态追踪

对接平台:
  - UU跑腿 (openID发单)
  - 闪送 (门店ID发单)
  - 顺丰 (shopId发单)
  - 蜂鸟 (开发者建店)
```

**目录结构**:
```
apps/
├── pkg/engine/
│   ├── platform/              # 平台管理引擎
│   │   ├── manager.go
│   │   └── adapter/
│   │       ├── uu_adapter.go
│   │       ├── flash_adapter.go
│   │       ├── sf_adapter.go
│   │       └── fengniao_adapter.go
│   ├── pricing/               # 定价引擎
│   │   ├── engine.go
│   │   ├── calculator/        # 价格计算器
│   │   └── cache/             # 价格缓存
│   ├── strategy/              # 策略引擎
│   │   ├── engine.go
│   │   ├── rule/              # 规则引擎
│   │   └── router/            # 路由算法
│   └── dispatch/              # 发单引擎
│       ├── engine.go
│       ├── executor/          # 执行器
│       └── tracker/           # 状态追踪
```

### 4. finance-service（财务服务）
职责: 资金管理与结算

```yaml
功能:
  - 资金流水管理
  - 自动对账
  - 垫付资金管理
  - 结算单生成
  - 违约金收取

核心引擎:
  - account: 账户管理
  - transaction: 交易处理
  - settlement: 结算引擎
  - reconciliation: 对账引擎
```

---

### 支撑服务（3个）

#### 5. user-service（用户服务）
- 平台用户管理
- 权限控制

#### 6. notification-service（通知服务）
- 消息推送
- 短信/邮件/Webhook

#### 7. monitor-service（监控服务）
- 业务指标监控
- 告警管理

---

## 🔄 核心业务流程

### 订单完整流程
```
1. 上游推单
   京东/美团 → platform-gateway → 协议转换

2. 订单创建  
   MQ(OrderPushEvent) → order-service → 创建订单
   
3. 配送调度 (delivery-dispatch核心)
   ├─ platform.Manager: 筛选可用平台
   ├─ pricing.Engine: 并发询价
   ├─ strategy.Engine: 策略评分
   └─ dispatch.Engine: 执行发单

4. 状态同步
   配送平台回调 → 更新订单状态 → 发布事件

5. 财务结算
   订单完成 → finance-service → 生成结算单
```

# 📨 消息队列设计
## Topic规划（4个核心Topic）

```text
1. upstream-push-topic (上游推单)
   Producer: platform-gateway
   Consumer: order-service (order-create-group)
   保留: 72小时

2. order-event-topic (订单事件)
   Producer: order-service
   Consumer: 
     - delivery-dispatch (dispatch-group)
     - finance (finance-group)
     - notification (notify-group)
   
3. delivery-event-topic (配送事件)
   Producer: delivery-dispatch
   Consumer: 
     - order (status-sync-group)
     - finance (settlement-group)

4. system-event-topic (系统事件)
   消息: 配置变更、价格缓存更新
```

# 🗄️ 数据架构设计

## 分库分表策略

```text
分库:
  - platform_db: 平台域
  - order_db: 订单域
  - delivery_db: 配送域
  - finance_db: 财务域

二级分区示例:
  platform_push_order_master
    一级: created_at (月分区) [间隔（Interval）分区](https://help.aliyun.com/zh/polardb/polardb-for-mysql/user-guide/when-to-use-interval-range-partitioning)
    二级: city_code (城市Hash) [LIST DEFAULT HASH分区
](https://help.aliyun.com/zh/polardb/polardb-for-mysql/user-guide/when-to-use-list-default-hash-partitioning)
```    

## 缓存策略
```yaml
三级缓存:
  L1: 本地缓存 (配置、静态数据)
  L2: Redis (热点数据、价格缓存)
  L3: PolarDB (持久化数据)

热点数据:
  - 城市定价规则
  - 平台能力配置
  - 发单策略规则
```

## 分布式事务
- go-zero DTM: [事务消息](https://go-zero.dev/docs/tutorials/mysql/distribute/transaction)
- DTM:  [go-zero 集成](https://dtm.pub/ref/gozero.html)
- 分布式事务: [小米信息部技术团队](https://xiaomi-info.github.io/2020/01/02/distributed-transaction/)


# ⚡ 性能优化方案

## 应对300w/天高并发

```yaml
峰值处理:
  - 峰值QPS: 3500
  - 异步处理: MQ削峰填谷
  - 连接池优化: 数据库/HTTP连接池
  - 限流降级: go-zero内置限流

数据库优化:
  - PolarDB读写分离
  - 索引优化
  - 慢查询监控

缓存预热:
  - 启动时预加载热点数据
  - 定时更新价格缓存
 ```

---

## 📂 项目目录速查
```
chuandao-sails-core/
├── apps/                          # 微服务应用
│   ├── platform-gateway/          # 平台网关 ✅ 70%
│   ├── order-service/             # 订单服务 🚧 40%
│   ├──          # 配送调度 ⚠️ 框架搭建
│   └── finance-service/           # 财务服务 📋 目录完整
│
├── common/                        # 公共组件 ✅
│   ├── rocketmq/                  # MQ客户端封装
│   ├── snowflake/                 # 订单号生成器
│   ├── response/                  # 统一响应
│   └── events/                    # 事件基类
│
├── deploy/                        # 部署配置
│   ├── docker/
│   └── k8s/
│
└── docs/                          # 文档
```

## 阿里云云原生网关
-  代理Nginx：[云原生API网关](https://help.aliyun.com/zh/api-gateway/cloud-native-api-gateway/product-overview/what-is-cloud-native-api-gateway)

- 阿里云SLS：[日志服务](https://help.aliyun.com/zh/sls/developer-reference/go-sdk/)


# 📖 相关文档

- Go-zero官方文档
    - [官方文档](https://go-zero.dev/)
    - [快速上手](https://go-zero.dev/docs/tasks)
    - [指南](https://go-zero.dev/docs/tutorials)
    - [组件](https://go-zero.dev/docs/components/limiter/token)
    - [项目示例](https://github.com/zeromicro/go-zero-examples)

- PolarDB文档
    - [PolarDB文档](https://help.aliyun.com/zh/polardb/polardb-for-mysql/getting-started)
    - [分区表](https://help.aliyun.com/zh/polardb/polardb-for-mysql/user-guide/partitioned-tables)

- RocketMQ文档
    - [文档](https://help.aliyun.com/zh/apsaramq-for-rocketmq/cloud-message-queue-rocketmq-5-x-series)
    - [Golang](https://github.com/apache/rocketmq-clients/tree/master/golang)

- ARMS文档
    - [配置方式](https://www.shiguopeng.cn/posts/2023083117/)

