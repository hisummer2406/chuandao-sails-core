# 财务域（Finance）：资金结算与对账

## 推荐目录结构

```text
finance-service/
├── api/                                    # HTTP API服务
│   ├── finance.api                        # API定义文件
│   ├── internal/
│   │   ├── config/
│   │   │   └── config.go                  # API配置结构
│   │   ├── handler/                       # HTTP处理器
│   │   │   ├── account/                   # 账户管理处理器
│   │   │   │   ├── balance_handler.go     # 余额查询处理器
│   │   │   │   ├── freeze_handler.go      # 资金冻结处理器
│   │   │   │   └── unfreeze_handler.go    # 资金解冻处理器
│   │   │   ├── transaction/               # 交易处理器
│   │   │   │   ├── create_handler.go      # 创建交易处理器
│   │   │   │   ├── query_handler.go       # 查询交易处理器
│   │   │   │   └── cancel_handler.go      # 取消交易处理器
│   │   │   ├── settlement/                # 结算处理器
│   │   │   │   ├── create_handler.go      # 创建结算单处理器
│   │   │   │   ├── approve_handler.go     # 审批结算单处理器
│   │   │   │   └── execute_handler.go     # 执行结算处理器
│   │   │   ├── penalty/                   # 违约金处理器
│   │   │   │   ├── calculate_handler.go   # 计算违约金处理器
│   │   │   │   ├── collect_handler.go     # 收取违约金处理器
│   │   │   │   └── refund_handler.go      # 退还违约金处理器
│   │   │   ├── report/                    # 报表处理器
│   │   │   │   ├── income_handler.go      # 收入报表处理器
│   │   │   │   ├── expense_handler.go     # 支出报表处理器
│   │   │   │   └── profit_handler.go      # 利润报表处理器
│   │   │   └── routes.go                  # 路由注册
│   │   ├── logic/                         # 业务逻辑层
│   │   │   ├── account/
│   │   │   │   ├── balance_logic.go       # 余额查询逻辑
│   │   │   │   ├── freeze_logic.go        # 资金冻结逻辑
│   │   │   │   └── unfreeze_logic.go      # 资金解冻逻辑
│   │   │   ├── transaction/
│   │   │   │   ├── create_logic.go        # 创建交易逻辑
│   │   │   │   ├── query_logic.go         # 查询交易逻辑
│   │   │   │   └── cancel_logic.go        # 取消交易逻辑
│   │   │   ├── settlement/
│   │   │   │   ├── create_logic.go        # 创建结算单逻辑
│   │   │   │   ├── approve_logic.go       # 审批结算单逻辑
│   │   │   │   └── execute_logic.go       # 执行结算逻辑
│   │   │   ├── penalty/
│   │   │   │   ├── calculate_logic.go     # 计算违约金逻辑
│   │   │   │   ├── collect_logic.go       # 收取违约金逻辑
│   │   │   │   └── refund_logic.go        # 退还违约金逻辑
│   │   │   └── report/
│   │   │       ├── income_logic.go        # 收入报表逻辑
│   │   │       ├── expense_logic.go       # 支出报表逻辑
│   │   │       └── profit_logic.go        # 利润报表逻辑
│   │   ├── middleware/                    # 中间件
│   │   │   ├── auth.go                   # 认证中间件
│   │   │   ├── permission.go             # 权限中间件
│   │   │   ├── audit.go                  # 审计日志中间件
│   │   │   └── risk_control.go           # 风控中间件
│   │   ├── svc/
│   │   │   └── servicecontext.go         # 服务上下文
│   │   └── types/
│   │       └── types.go                  # API类型定义
│   └── finance.go                        # API入口文件
├── rpc/                                  # RPC服务
│   ├── finance.proto                     # protobuf定义
│   ├── internal/
│   │   ├── config/
│   │   │   └── config.go                 # RPC配置
│   │   ├── logic/                        # RPC业务逻辑
│   │   │   ├── account/
│   │   │   │   ├── get_balance_logic.go   # 获取余额RPC逻辑
│   │   │   │   ├── freeze_amount_logic.go # 冻结金额RPC逻辑
│   │   │   │   └── unfreeze_amount_logic.go # 解冻金额RPC逻辑
│   │   │   ├── transaction/
│   │   │   │   ├── create_transaction_logic.go # 创建交易RPC逻辑
│   │   │   │   ├── confirm_transaction_logic.go # 确认交易RPC逻辑
│   │   │   │   └── rollback_transaction_logic.go # 回滚交易RPC逻辑
│   │   │   ├── settlement/
│   │   │   │   ├── create_settlement_logic.go # 创建结算RPC逻辑
│   │   │   │   └── batch_settlement_logic.go  # 批量结算RPC逻辑
│   │   │   └── penalty/
│   │   │       ├── process_penalty_logic.go   # 处理违约金RPC逻辑
│   │   │       └── calculate_penalty_logic.go # 计算违约金RPC逻辑
│   │   ├── server/
│   │   │   └── finance_server.go         # RPC服务器实现
│   │   └── svc/
│   │       └── servicecontext.go         # RPC服务上下文
│   ├── pb/                               # 生成的protobuf文件
│   │   └── finance.pb.go
│   └── finance.go                        # RPC入口文件
├── job/                                  # 定时任务
│   ├── internal/
│   │   ├── config/
│   │   │   └── config.go                 # 任务配置
│   │   ├── logic/
│   │   │   ├── daily_settlement_job.go   # 日结算任务
│   │   │   ├── monthly_settlement_job.go # 月结算任务
│   │   │   ├── account_reconcile_job.go  # 对账任务
│   │   │   ├── penalty_collect_job.go    # 违约金收取任务
│   │   │   ├── interest_calculate_job.go # 利息计算任务
│   │   │   └── report_generate_job.go    # 报表生成任务
│   │   └── svc/
│   │       └── servicecontext.go         # 任务服务上下文
│   └── job.go                            # 任务入口文件
├── mq/                                   # 消息队列消费者
│   ├── internal/
│   │   ├── config/
│   │   │   └── config.go                 # MQ配置
│   │   ├── logic/
│   │   │   ├── order_complete_consumer.go    # 订单完成消息消费
│   │   │   ├── penalty_consumer.go           # 违约金消息消费
│   │   │   ├── settlement_consumer.go        # 结算消息消费
│   │   │   ├── refund_consumer.go            # 退款消息消费
│   │   │   └── notification_consumer.go      # 通知消息消费
│   │   └── svc/
│   │       └── servicecontext.go         # MQ服务上下文
│   └── mq.go                             # MQ入口文件
├── model/                                # 数据模型
│   ├── account_model.go                  # 账户模型
│   ├── transaction_model.go              # 交易模型
│   ├── settlement_model.go               # 结算单模型
│   ├── penalty_record_model.go           # 违约记录模型
│   ├── account_balance_model.go          # 账户余额模型
│   ├── frozen_amount_model.go            # 冻结金额模型
│   ├── financial_flow_model.go           # 资金流水模型
│   ├── reconciliation_model.go           # 对账记录模型
│   └── vars.go                           # 模型常量
├── pkg/                                  # 业务核心包
│   ├── engine/                           # 财务核心引擎
│   │   ├── account/                      # 账户管理引擎
│   │   │   ├── manager.go               # 账户管理器
│   │   │   ├── balance/                 # 余额管理
│   │   │   │   ├── balance_manager.go   # 余额管理器
│   │   │   │   ├── calculator.go        # 余额计算器
│   │   │   │   └── validator.go         # 余额验证器
│   │   │   ├── freeze/                  # 冻结管理
│   │   │   │   ├── freeze_manager.go    # 冻结管理器
│   │   │   │   ├── freeze_calculator.go # 冻结计算器
│   │   │   │   └── freeze_validator.go  # 冻结验证器
│   │   │   └── history/                 # 历史记录
│   │   │       ├── history_manager.go   # 历史管理器
│   │   │       └── history_analyzer.go  # 历史分析器
│   │   ├── transaction/                 # 交易引擎
│   │   │   ├── engine.go               # 交易引擎主类
│   │   │   ├── processor/              # 交易处理器
│   │   │   │   ├── interface.go        # 处理器接口
│   │   │   │   ├── payment_processor.go # 付款处理器
│   │   │   │   ├── receipt_processor.go # 收款处理器
│   │   │   │   ├── transfer_processor.go # 转账处理器
│   │   │   │   └── refund_processor.go  # 退款处理器
│   │   │   ├── validator/              # 交易验证器
│   │   │   │   ├── amount_validator.go  # 金额验证器
│   │   │   │   ├── account_validator.go # 账户验证器
│   │   │   │   └── rule_validator.go    # 规则验证器
│   │   │   ├── state/                  # 交易状态机
│   │   │   │   ├── state_machine.go    # 状态机
│   │   │   │   ├── transitions.go      # 状态转换
│   │   │   │   └── validators.go       # 状态验证器
│   │   │   └── compensator/            # 补偿机制
│   │   │       ├── compensator.go      # 补偿器
│   │   │       └── rollback.go         # 回滚机制
│   │   ├── settlement/                 # 结算引擎
│   │   │   ├── engine.go              # 结算引擎主类
│   │   │   ├── calculator/            # 结算计算器
│   │   │   │   ├── commission_calculator.go # 佣金计算器
│   │   │   │   ├── cost_calculator.go       # 成本计算器
│   │   │   │   ├── profit_calculator.go     # 利润计算器
│   │   │   │   └── tax_calculator.go        # 税费计算器
│   │   │   ├── generator/             # 结算单生成器
│   │   │   │   ├── daily_generator.go      # 日结算生成器
│   │   │   │   ├── weekly_generator.go     # 周结算生成器
│   │   │   │   └── monthly_generator.go    # 月结算生成器
│   │   │   ├── approver/              # 结算审批器
│   │   │   │   ├── auto_approver.go        # 自动审批器
│   │   │   │   ├── manual_approver.go      # 人工审批器
│   │   │   │   └── risk_approver.go        # 风险审批器
│   │   │   └── executor/              # 结算执行器
│   │   │       ├── batch_executor.go       # 批量执行器
│   │   │       ├── single_executor.go      # 单笔执行器
│   │   │       └── retry_executor.go       # 重试执行器
│   │   ├── penalty/                   # 违约金引擎
│   │   │   ├── engine.go             # 违约金引擎主类
│   │   │   ├── calculator/           # 违约金计算器
│   │   │   │   ├── interface.go      # 计算器接口
│   │   │   │   ├── fixed_calculator.go      # 固定金额计算器
│   │   │   │   ├── ratio_calculator.go      # 比例计算器
│   │   │   │   ├── tiered_calculator.go     # 阶梯计算器
│   │   │   │   └── time_calculator.go       # 时间相关计算器
│   │   │   ├── rule/                 # 违约规则
│   │   │   │   ├── rule_engine.go    # 规则引擎
│   │   │   │   ├── cancel_rule.go    # 取消违约规则
│   │   │   │   ├── timeout_rule.go   # 超时违约规则
│   │   │   │   └── quality_rule.go   # 质量违约规则
│   │   │   ├── collector/            # 违约金收取器
│   │   │   │   ├── auto_collector.go       # 自动收取器
│   │   │   │   ├── manual_collector.go     # 人工收取器
│   │   │   │   └── batch_collector.go      # 批量收取器
│   │   │   └── refunder/             # 违约金退还器
│   │   │       ├── refund_calculator.go    # 退还计算器
│   │   │       └── refund_executor.go      # 退还执行器
│   │   ├── reconciliation/           # 对账引擎
│   │   │   ├── engine.go            # 对账引擎主类
│   │   │   ├── matcher/             # 对账匹配器
│   │   │   │   ├── amount_matcher.go       # 金额匹配器
│   │   │   │   ├── time_matcher.go         # 时间匹配器
│   │   │   │   └── order_matcher.go        # 订单匹配器
│   │   │   ├── differ/              # 差异检测器
│   │   │   │   ├── amount_differ.go        # 金额差异检测
│   │   │   │   ├── status_differ.go        # 状态差异检测
│   │   │   │   └── missing_differ.go       # 缺失差异检测
│   │   │   ├── resolver/            # 差异解决器
│   │   │   │   ├── auto_resolver.go        # 自动解决器
│   │   │   │   ├── manual_resolver.go      # 人工解决器
│   │   │   │   └── suggest_resolver.go     # 建议解决器
│   │   │   └── reporter/            # 对账报告器
│   │   │       ├── summary_reporter.go     # 汇总报告器
│   │   │       └── detail_reporter.go      # 详细报告器
│   │   └── risk/                    # 风控引擎
│   │       ├── monitor.go           # 风控监控器
│   │       ├── detector/            # 风险检测器
│   │       │   ├── amount_detector.go      # 金额异常检测器
│   │       │   ├── frequency_detector.go   # 频率异常检测器
│   │       │   └── pattern_detector.go     # 模式异常检测器
│   │       ├── rule/                # 风控规则
│   │       │   ├── amount_rule.go          # 金额风控规则
│   │       │   ├── time_rule.go            # 时间风控规则
│   │       │   └── behavior_rule.go        # 行为风控规则
│   │       └── action/              # 风控动作
│   │           ├── alert_action.go         # 告警动作
│   │           ├── freeze_action.go        # 冻结动作
│   │           └── review_action.go        # 审核动作
│   ├── report/                      # 报表组件
│   │   ├── generator/               # 报表生成器
│   │   │   ├── income_generator.go         # 收入报表生成器
│   │   │   ├── expense_generator.go        # 支出报表生成器
│   │   │   ├── profit_generator.go         # 利润报表生成器
│   │   │   └── flow_generator.go           # 流水报表生成器
│   │   ├── analyzer/                # 数据分析器
│   │   │   ├── trend_analyzer.go           # 趋势分析器
│   │   │   ├── compare_analyzer.go         # 对比分析器
│   │   │   └── forecast_analyzer.go        # 预测分析器
│   │   ├── exporter/                # 报表导出器
│   │   │   ├── excel_exporter.go           # Excel导出器
│   │   │   ├── pdf_exporter.go             # PDF导出器
│   │   │   └── csv_exporter.go             # CSV导出器
│   │   └── scheduler/               # 报表调度器
│   │       ├── daily_scheduler.go          # 日报调度器
│   │       ├── weekly_scheduler.go         # 周报调度器
│   │       └── monthly_scheduler.go        # 月报调度器
│   ├── payment/                     # 支付组件
│   │   ├── gateway/                 # 支付网关
│   │   │   ├── interface.go               # 网关接口
│   │   │   ├── alipay_gateway.go          # 支付宝网关
│   │   │   ├── wechat_gateway.go          # 微信支付网关
│   │   │   └── bank_gateway.go            # 银行网关
│   │   ├── processor/               # 支付处理器
│   │   │   ├── payment_processor.go       # 支付处理器
│   │   │   ├── refund_processor.go        # 退款处理器
│   │   │   └── query_processor.go         # 查询处理器
│   │   └── notifier/                # 支付通知器
│   │       ├── callback_notifier.go       # 回调通知器
│   │       └── webhook_notifier.go        # Webhook通知器
│   ├── client/                      # 外部客户端
│   │   ├── order_client.go                # 订单服务客户端
│   │   ├── delivery_client.go             # 配送服务客户端
│   │   ├── bank_client.go                 # 银行接口客户端
│   │   └── notification_client.go         # 通知服务客户端
│   ├── cache/                       # 缓存组件
│   │   ├── account_cache.go               # 账户缓存
│   │   ├── balance_cache.go               # 余额缓存
│   │   ├── transaction_cache.go           # 交易缓存
│   │   └── settlement_cache.go            # 结算缓存
│   ├── mq/                          # 消息队列组件
│   │   ├── producer.go                    # 消息生产者
│   │   ├── message/                       # 消息定义
│   │   │   ├── transaction_message.go     # 交易消息
│   │   │   ├── settlement_message.go      # 结算消息
│   │   │   ├── penalty_message.go         # 违约金消息
│   │   │   └── notification_message.go    # 通知消息
│   │   └── topic/                         # Topic管理
│   │       └── finance_topic.go           # 财务Topic
│   ├── utils/                       # 工具类
│   │   ├── decimal_utils.go               # 金额计算工具
│   │   ├── encrypt_utils.go               # 加密工具
│   │   ├── id_generator.go                # ID生成器
│   │   ├── time_utils.go                  # 时间工具
│   │   └── validation_utils.go            # 验证工具
│   ├── constants/                   # 常量定义
│   │   ├── account_constants.go           # 账户常量
│   │   ├── transaction_constants.go       # 交易常量
│   │   ├── settlement_constants.go        # 结算常量
│   │   ├── penalty_constants.go           # 违约金常量
│   │   └── error_constants.go             # 错误码常量
│   └── validator/                   # 验证器
│       ├── account_validator.go           # 账户验证器
│       ├── amount_validator.go            # 金额验证器
│       └── business_validator.go          # 业务规则验证器
├── etc/                             # 配置文件
│   ├── finance-api.yaml            # API服务配置
│   ├── finance-rpc.yaml            # RPC服务配置
│   ├── finance-job.yaml            # 定时任务配置
│   ├── finance-mq.yaml             # MQ消费者配置
│   └── payment-gateway.yaml        # 支付网关配置
├── deploy/                          # 部署文件
│   ├── dockerfile/
│   │   ├── Dockerfile.api                 # API服务镜像
│   │   ├── Dockerfile.rpc                 # RPC服务镜像
│   │   ├── Dockerfile.job                 # 定时任务镜像
│   │   └── Dockerfile.mq                  # MQ消费者镜像
│   └── k8s/
│       ├── deployment.yaml                # K8S部署配置
│       ├── service.yaml                   # K8S服务配置
│       └── configmap.yaml                 # K8S配置映射
├── docs/                            # 文档
│   ├── api.md                      # API文档
│   ├── transaction_flow.md         # 交易流程文档
│   ├── settlement_guide.md         # 结算指南
│   ├── penalty_rules.md            # 违约金规则文档
│   └── deployment.md               # 部署文档
├── scripts/                         # 脚本
│   ├── build.sh                    # 构建脚本
│   ├── deploy.sh                   # 部署脚本
│   ├── migrate.sql                 # 数据库迁移脚本
│   └── init_accounts.sql           # 账户初始化脚本
├── test/                            # 测试
│   ├── api_test.go                 # API测试
│   ├── rpc_test.go                 # RPC测试
│   ├── integration_test.go         # 集成测试
│   ├── transaction_test.go         # 交易测试
│   ├── settlement_test.go          # 结算测试
│   └── penalty_test.go             # 违约金测试
├── go.mod                           # Go模块文件
├── go.sum                           # Go依赖锁定文件
├── Makefile                         # Make构建文件
└── README.md                        # 项目说明文档

```

## 2. 微服务划分

6. finance-service（财务服务）

* 资金流水管理
* 自动对账
* 垫付资金管理
* 结算单生成