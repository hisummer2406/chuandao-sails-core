package main

import (
	"chuandao-sails-core/demo/restful/queue/internal/consumer"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"chuandao-sails-core/demo/restful/queue/internal/config"
	"chuandao-sails-core/demo/restful/queue/internal/handler"
	"chuandao-sails-core/demo/restful/queue/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/queue-api.yaml", "the config file")

func main() {
	flag.Parse()

	//加载配置
	var c config.Config
	conf.MustLoad(*configFile, &c)

	//创建服务上下文
	ctx := svc.NewServiceContext(c)
	defer ctx.Stop()

	//启动消费者
	consumerManager := consumer.NewConsumerManager(ctx)
	if err := consumerManager.Start(); err != nil {
		logx.Errorf("rocketmq start consumer failed, err=%v", err)
		os.Exit(-1)
	}

	//启动健康检查
	startHealthCheck(&c)

	fmt.Printf("rocketmq starting server at %s:%d...\n", c.Host, c.Port)

	//关闭信号
	gracefulShutdown()
}

func startHealthCheck(c *config.Config) {
	// 启动健康检查服务
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 注册健康检查处理函数
	server.AddRoute(rest.Route{
		Path:   "/health",
		Method: http.MethodGet,
		Handler: func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		},
	})

	// 启动健康检查服务
	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
	logx.Infof("rocketmq health check server started at %s", addr)
	server.Start()
}

// gracefulShutdown 优雅关闭
func gracefulShutdown() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	sig := <-quit
	logx.Infof("rocketmq received signal: %v, shutting down queue service...", sig)
}
