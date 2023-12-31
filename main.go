package main

import (
	"github.com/bingxindan/bxd_data_access/app/console"
	bxdHttp "github.com/bingxindan/bxd_data_access/app/http"
	"github.com/bingxindan/bxd_data_access/framework"
	"github.com/bingxindan/bxd_data_access/framework/provider/app"
	"github.com/bingxindan/bxd_data_access/framework/provider/config"
	"github.com/bingxindan/bxd_data_access/framework/provider/distributed"
	"github.com/bingxindan/bxd_data_access/framework/provider/env"
	"github.com/bingxindan/bxd_data_access/framework/provider/id"
	"github.com/bingxindan/bxd_data_access/framework/provider/kernel"
	"github.com/bingxindan/bxd_data_access/framework/provider/log"
	"github.com/bingxindan/bxd_data_access/framework/provider/trace"
)

func main() {
	// 初始化服务容器
	container := framework.NewBxdContainer()
	// 绑定App服务提供者
	container.Bind(&app.BxdAppProvider{})
	// 后续初始化需要绑定的服务提供者...
	container.Bind(&env.BxdEnvProvider{})
	container.Bind(&distributed.LocalDistributedProvider{})
	container.Bind(&config.BxdConfigProvider{})
	container.Bind(&id.BxdIDProvider{})
	container.Bind(&trace.BxdTraceProvider{})
	container.Bind(&log.BxdLogServiceProvider{})

	// 将HTTP引擎初始化，并且作为服务提供者绑定到服务容器中
	if engine, err := bxdHttp.NewHttpEngine(); err == nil {
		container.Bind(&kernel.BxdKernelProvider{HttpEngine: engine})
	}

	// 运行root命令
	console.RunCommand(container)
}
