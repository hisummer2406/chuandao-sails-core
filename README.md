# 船到-千帆微服务架构

## 项目简介
船到-千帆微服务架构，是一个基于微服务架构的订单系统，用于管理订单、支付、配送等业务。


**项目目录如下：**
```plainText
====== 多服务微服务架构 =======

project/
├── apps/              # 各个微服务
│   ├── user/             # 用户服务
│   │   ├── api/
│   │   ├── rpc/
│   │   └── model/
│   ├── order/            # 订单服务
│   │   ├── api/
│   │   ├── rpc/
│   │   └── model/
│   └── product/          # 商品服务
│       ├── api/
│       ├── rpc/
│       └── model/
├── common/               # 公共组件
│   ├── response/         # 统一响应格式
│   ├── middleware/       # 中间件
│   ├── utils/           # 工具函数
│   ├── constants/       # 常量定义
│   └── errorx/          # 错误处理
├── deploy/              # 部署相关
│   ├── docker-compose/
│   ├── k8s/
│   └── scripts/
├── docs/                # 文档
├── .gitignore
├── go.mod
└── go.sum

====== 单服务微服务架构 =======

service-name/
├── consumer # 队列消费服务
├── go.mod
├── internal # 工程内部可以访问的公共模块
│   └── model
├── job # cron job 服务
├── pkg # 工程外部可以访问的公共模块
├── restful # HTTP 服务
├── script # 脚本服务
└── service # gRPC 服务
```

已接入技术栈：
- 框架：Go-zero
- 服务注册发现：etcd
- 服务调用：gRPC
- 消息队列：阿里云 RocketMQ
- 数据库：阿里云 PolarDB MySQL8.0
- 缓存：Redis


- 负载均衡：Nginx
- 监控：Prometheus
- 日志：ELK
- 链路追踪：Jaeger

## 订单流程代码

## 系统架构图

## 业务架构图

## 网关
### 阿里云云原生网关
> 代理Nginx：[云原生API网关](https://help.aliyun.com/zh/api-gateway/cloud-native-api-gateway/product-overview/what-is-cloud-native-api-gateway)

> 阿里云SLS：[日志服务](https://help.aliyun.com/zh/sls/developer-reference/go-sdk/)

## 监控

## 链路追踪
> 阿里云 ARMS jaeger 方式接入 
> - [配置方式](https://www.shiguopeng.cn/posts/2023083117/)

## 日志
> 阿里云SLS [Aliyun LOG Go Producer ](https://github.com/aliyun/aliyun-log-go-sdk/tree/master/producer)
> Go-zero [第三方日志库集成](https://github.com/zeromicro/zero-examples/blob/main/logx/readme-cn.md)

## 消息队列

### RocketMQ
> https://www.aliyun.com/product/rocketmq

**1. 分布式事务消息**

分布式事务消息：最终一致性；解决生产者端本地事务与消息发送的原子性；
分布式事务：强一致性；解决跨服务、跨数据库的全局事务一致性；

- go-zero DTM: [事务消息](https://go-zero.dev/docs/tutorials/mysql/distribute/transaction)
- DTM:  [go-zero 集成](https://dtm.pub/ref/gozero.html)
- 分布式事务: [小米信息部技术团队](https://xiaomi-info.github.io/2020/01/02/distributed-transaction/)

**2. 消息队列 **
> 代替go-zero go-queue(kafka): [顺序消息](https://go-zero.dev/docs/tutorials/message-queue/kafka)

**3. 延时消息**
> 代替go-zero beanstalkd: [定时/延时消息](https://help.aliyun.com/zh/apsaramq-for-rocketmq/cloud-message-queue-rocketmq-5-x-series/developer-reference/scheduled-and-delayed-messages?spm=a2c4g.11186623.help-menu-search-29530.d_0)

### Go-zero 聚合消息
[go-zero微服务实战系列（九、极致优化秒杀性能](https://mp.weixin.qq.com/s/8VSS9WNSy4jkOSSIA4BmLw)

聚合策略是啥呢？聚合策略有两个维度分别是聚合消息条数和聚合时间，比如聚合消息达到100条我们就往Kafka发送一次，这个条数是可以配置的，那如果一直也达不到100条消息怎么办呢？通过聚合时间来兜底，这个聚合时间也是可以配置的，比如配置聚合时间为1秒钟，也就是无论目前聚合了多少条消息只要聚合时间达到1秒，那么就往Kafka发送一次数据。聚合条数和聚合时间是或的关系，也就是只要有一个条件满足就触发。