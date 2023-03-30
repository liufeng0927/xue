package user

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	globalkey "xuedengwang/core/global"
	"xuedengwang/core/interceptor/rpcserver"

	"xuedengwang/service/user/rpc/internal/config"
	"xuedengwang/service/user/rpc/internal/server"
	"xuedengwang/service/user/rpc/internal/svc"
	"xuedengwang/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Run(context *cli.Context) error {
	var c config.Config
	conf.MustLoad(context.String(globalkey.CtxConfKey), &c)

	prefix := fmt.Sprintf("%s:", c.Mode)
	keys := globalkey.BuildKey(prefix)

	log.Printf("redis buildKey success with prefix:【%s】, keyCounts:%d, keys:%v\n", prefix, len(keys), keys)

	//migrator.AutoMigrate(&c.DB)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	//rpc log
	s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)
	defer s.Stop()

	log.Printf("Starting stat rpc server at %s...\n", c.ListenOn)
	s.Start()
	return nil
}
