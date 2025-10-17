# 平台接入域（Platform）：上游平台对接

## 1. 目录结构

## 2. 微服务划分

2. platform-gateway-service（平台网关服务）

* 上游平台接入适配
* 协议转换（京东、美团、淘宝等）
* 限流与熔断
* 订单格式标准化

## 3. 对接平台

#### UU
- http://showdoc.canzan.com/web/#/296
- 密码: 6PY9ZQ8B

**注意事项**
1. 选择报价模式
2. UU 定义数据格式，主动推送

#### 顺丰
- https://showdoc.canzan.com/web/#/383
- https://www.yuque.com/jinyubuwangjiushichen/kb/guff3g6bsbhagwht?singleDoc#UoTAU

**注意事项**
1. 顺丰询价发单模式，顺丰暂未对接
2. 顺丰定义数据格式，主动推送

#### 公共
- https://showdoc.canzan.com/web/#/364

**注意事项**

1. 裹小递 / 快服务使用
2. 船到定义数据结构，厂商主动推送

## 4. PHP同步细节

### 1. 删除 `store` 逻辑，接单直接发单

**相关表：**
    - business_store_user 商户信息表
    - config_store_auto 自动发单配置表
    - config_user_store 账户信息配置表，存储发单配置，渠道店铺ID等

## 5.队列设计
```yaml
负责范围:
  ✅ 接收上游平台的HTTP回调
  ✅ 协议转换为标准格式
  ✅ 数据校验和清洗
  ✅ 发送到MQ解耦
  
处理接口:
  - POST /platform/callback/uu/create (UU推单)
  - POST /platform/callback/uu/cancel (UU取消)
  - POST /platform/callback/uu/addOnlineFee (UU加小费)
  - POST /platform/callback/sf/create (顺丰推单)
  
发布到MQ:
  - upstream-push-topic (订单创建)
  - upstream-operation-topic (订单操作: 取消/加小费)
```