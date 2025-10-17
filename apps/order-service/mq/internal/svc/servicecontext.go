package svc

import (
	"chuandao-sails-core/apps/order-service/model"
	"chuandao-sails-core/apps/order-service/mq/internal/config"
	"chuandao-sails-core/apps/order-service/pkg/client"
	"chuandao-sails-core/apps/order-service/pkg/engine/lifecycle/event"
	"chuandao-sails-core/common/rocketmq"
	"chuandao-sails-core/common/snowflake"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	//Model
	OrderModel     model.PlatformPushOrderMasterModel
	StatusLogModel model.OrderStatusLogModel
	PenaltyModel   model.OrderPenaltyRecordModel
	//业务引擎
	//StateMachine  *lifecycle.StateMachine
	//PenaltyEngine *penalty.Calculator
	//外部客户端
	DeliveryClient client.DeliveryClient
	FinanceClient  client.FinanceClient
	//MQ
	MQClient       *rocketmq.Client
	EventPublisher *event.Publisher
}

func NewServiceContext(c config.Config) *ServiceContext {
	//init mysql conn
	conn := sqlx.NewMysql(c.DataSource)
	orderModel := model.NewPlatformPushOrderMasterModel(conn, c.Redis)
	statusLogModel := model.NewOrderStatusLogModel(conn, c.Redis)
	penaltyModel := model.NewOrderPenaltyRecordModel(conn, c.Redis)

	//雪花算法
	if err := snowflake.InitDefaultGenerator(c.Snowflake.WorkerId); err != nil {
		logx.Severef("order-mp snowflake.InitDefaultGenerator fail: %v", err)
	}

	//MQ client
	mqClient, err := rocketmq.NewClient(&c.RocketMQ)
	if err != nil {
		logx.Severef("order-mp Failed to create MQ client: %v", err)
	}

	//事件发生器
	eventPublisher := event.NewPublisher(mqClient)

	return &ServiceContext{
		Config:         c,
		OrderModel:     orderModel,
		StatusLogModel: statusLogModel,
		PenaltyModel:   penaltyModel,
		MQClient:       mqClient,
		EventPublisher: eventPublisher,
	}
}
