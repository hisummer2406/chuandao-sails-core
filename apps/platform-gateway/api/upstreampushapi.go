package main

import (
	"chuandao-sails-core/common/sls"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
	"os/signal"
	"syscall"

	"chuandao-sails-core/apps/platform-gateway/api/internal/config"
	"chuandao-sails-core/apps/platform-gateway/api/internal/handler"
	"chuandao-sails-core/apps/platform-gateway/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "../etc/upstream-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	//注册日志
	logx.MustSetup(c.Log)
	slsWriter, err := sls.NewWriter(c.SLS)
	if err != nil {
		logx.Errorf("Failed to initialize SLS writer: %v", err)
	} else {
		defer slsWriter.Close()
		logx.SetWriter(slsWriter)
		//测试环境
		logx.AddWriter(logx.NewWriter(os.Stdout))
	}

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	//优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
		server.Start()
	}()

	<-quit
	fmt.Println("Shutting down platform gateway...")

	//关闭RocketMQ
	if err := ctx.MQClient.Close(); err != nil {
		logx.Errorf("Failed to close MQ client: %v", err)
	}
	fmt.Println("platform gateway exited")
}
