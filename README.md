# 船到-千帆微服务架构

## 项目简介
船到千帆系统，基于Go-zero 和 PolarDB Mysql 8.0 、RocketMQ 消息队列的微服务架构方式。
系统功能是本地生活聚合配送平台，关键角色是京东外卖、美团外卖、淘宝闪购等外卖平台，在订单高峰期讲溢出单推送给本系统，本系统将订单推送给 UU跑腿、闪送、顺丰、蜂鸟等配送平台，从中赚取配送差价。
财务方面，订单平台根据订单完成情况给本系统线下打款，本系统垫付配送费给配送平台。
订单流程为：上游推送平台、系统接单、根据策略下游配送平台询价（城市、平台等维度）、根据发单策略发单（顺序、同时、价格、配送时效等） 、配送平台接单，配送完成。
订单取消情况：下游根据取消状态收取违约金； 本系统根据订单状态收取订单配送违约金。
性能要求：每天300w订单，且集中在餐饮高峰期节假日高峰期。

**项目目录如下：**
```plainText

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