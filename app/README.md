# 配送调度域（delivery-dispatch）：配送平台对接与管理（delivery）、动态定价与询价策略（pricing）、发单策略与路由规则（strategy）、发单服务（dispatch）

## 📋 项目简介
核心功能： 聚合16个配送平台，提供统一的配送调度API服务，支持询价、发单、状态追踪、订单管理等完整配送生命周期管理。
性能要求： 日均300万单，峰值3500 QPS

## 🏗️ 技术架构

### 服务部署架构
```
┌─────────────────────────────────────────────────┐
│ 外部系统（ERP/WMS/TMS/订单系统）                 │
└────────────────┬────────────────────────────────┘
                 │ HTTP/RPC
                 ↓
┌─────────────────────────────────────────────────┐
│ delivery-dispatch-service（配送调度服务）        │
│ ├── delivery-api (对外HTTP API)                │
│ ├── delivery-rpc (内部RPC)                     │
│ ├── callback-api (平台回调接收)                │
│ └── delivery-mq (MQ消费者)                     │
└────────────────┬────────────────────────────────┘
                 │ 并发调用
                 ↓
┌─────────────────────────────────────────────────┐
│ 16个配送平台 API                                 │
│ UU/达达/蜂鸟/闪送/顺丰/美团/京东/...             │
└─────────────────────────────────────────────────┘
```

---

### 📂 项目结构

```
apps/
├── api/                          # 对外HTTP服务
│   ├── desc/delivery.api        # API定义
│   └── internal/
│       ├── handler/             # 请求处理
│       └── logic/               # 业务逻辑
│
├── rpc/                         # 内部RPC服务
│   ├── pb/delivery.proto
│   └── internal/
│
├── callback/                    # 平台回调接收
│   └── internal/handler/
│       ├── uu_callback.go      # UU回调
│       ├── dada_callback.go    # 达达回调
│       └── ...                 # 其他14个平台
│
├── mq/                         # MQ消费者
│   └── internal/logic/
│
├── pkg/                        # 核心业务包
│   ├── engine/                 # 四大引擎
│   │   ├── platform/          # 平台管理引擎
│   │   │   ├── manager.go
│   │   │   └── adapter/       # 16个平台适配器
│   │   │       ├── interface.go
│   │   │       ├── uu_adapter.go
│   │   │       ├── dada_adapter.go
│   │   │       └── ...
│   │   ├── dispatch/          # 发单引擎
│   │   │   ├── engine.go
│   │   │   └── executor/      # 并发/顺序执行器
│   │   ├── strategy/          # 策略引擎
│   │   │   ├── engine.go
│   │   │   └── selector/      # 平台选择策略
│   │   └── track/             # 轨迹追踪引擎
│   │       └── tracker.go
│   ├── client/                # 平台SDK
│   │   ├── uu/
│   │   ├── dada/
│   │   └── ...
│   └── utils/                 # 工具类
│
├── model/                     # 数据模型（20张表）
│   ├── dispatch_order_model.go
│   ├── dispatch_inquiry_log_model.go
│   └── ...
│
└── etc/                       # 配置文件
    ├── delivery-api.yaml
    ├── delivery-rpc.yaml
    └── callback-api.yaml
```

### 🗄️ 核心职责引擎

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

### 队列设计

```yaml
负责范围:
  ✅ 配送平台管理（适配器模式）
  ✅ 智能询价（并发调用多平台）
  ✅ 发单策略执行
  ✅ 配送跟踪（轮询/webhook）
  ✅ 订单操作执行（取消/加小费）
  ✅ 违约金获取（调用下游平台）
  
消费MQ:
  - order-event-topic (order.created → 触发询价发单)
  - upstream-operation-topic (处理取消/加小费操作)
  
调用下游API:
  - UU询价/发单/取消/加小费
  - 顺丰询价/发单/取消
  - 闪送询价/发单/取消
  
发布事件:
  - delivery-event-topic (配送状态变更)
```

## 询价发单细节

### PHP 代码
```php
// 溢出单配送平台实例，app/Services/OverFlowOrder/isv 
OverFlowDeliveryLogic::getInstance($item);

//平台调度运力配置表
config_user_store
```

### 建店逻辑
```text
UU
- openID 商户身份发单 点对点逻辑？建店逻辑
- https://open.uupt.com/#/development/guide?t=%E6%8E%A5%E5%85%A5%E6%B5%81%E7%A8%8B&index=1
- 百度地图坐标系

达达
- 开发者身份 建店发单
- https://newopen.imdada.cn/#/development/file/guide

蜂鸟
- 开发者身份 建店发单
- https://open.ele.me/documents/openApi/1652

闪送
- 门店ID 商户身份发单 点对点
- https://open.ishansong.com/documentCenter

顺丰
- shopId发单 无门店信息，用订单信息发单
- https://openic.sf-express.com/#/apidoc
```

### 自建虚拟门店

推单经纬度，创建发货虚拟门店映射发货地址

```sql
select id from store where lng = ? and lat = ? and platform_code = ?
```

## TODO 速记卡
1. [ ]  适配器模式 + 责任链模式 + 装饰器模式