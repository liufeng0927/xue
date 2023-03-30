package gateway

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"log"
	globalkey "xuedengwang/core/global"
	"xuedengwang/service/gateway/api/internal/config"
	"xuedengwang/service/gateway/api/internal/handler"
	"xuedengwang/service/gateway/api/internal/svc"
)

func Run(context *cli.Context) error {

	var c config.Config
	conf.MustLoad(context.String(globalkey.CtxConfKey), &c)

	prefix := fmt.Sprintf("%s:", c.Mode)
	keys := globalkey.BuildKey(prefix)

	log.Printf("redis buildKey success with prefix:【%s】, keyCounts:%d, keys:%v\n", prefix, len(keys), keys)

	//migrator.AutoMigrate(&c.DB)

	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	log.Printf("Starting gateway server at %s:%d...\n", c.Host, c.Port)
	server.Start()
	return nil
}
