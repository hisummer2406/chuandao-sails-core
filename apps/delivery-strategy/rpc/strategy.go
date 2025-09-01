package main

import (
	"flag"
	"fmt"

	"chuandao-sails-core/apps/delivery-strategy/rpc/internal/config"
	v1Server "chuandao-sails-core/apps/delivery-strategy/rpc/internal/server/v1"
	"chuandao-sails-core/apps/delivery-strategy/rpc/internal/svc"
	"chuandao-sails-core/apps/delivery-strategy/rpc/strategy"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/delivery-strategy.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		strategy.RegisterV1Server(grpcServer, v1Server.NewV1Server(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
