package main

import (
    "chuandao-sails-core/common/sls"
    "github.com/zeromicro/go-zero/core/logx"
	"flag"
	"fmt"

	{{.importPackages}}
)

var configFile = flag.String("f", "etc/{{.serviceName}}.yaml", "the config file")

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

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
