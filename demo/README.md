## 工程目录结构

```plainText
.
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
