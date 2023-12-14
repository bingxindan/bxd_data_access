package main

import (
	"github.com/bingxindan/bxd_data_access/app/console"
	bxdHttp "github.com/bingxindan/bxd_data_access/app/http"
	"github.com/bingxindan/bxd_data_access/framework"
	"github.com/bingxindan/bxd_data_access/framework/provider/app"
	"github.com/bingxindan/bxd_data_access/framework/provider/kernel"
)

func main() {
	// 初始化服务容器
	container := framework.NewBxdContainer()
	// 绑定App服务提供者
	container.Bind(&app.BxdAppProvider{})

	// 将HTTP引擎初始化，并且作为服务提供者绑定到服务容器中
	if engine, err := bxdHttp.NewHttpEngine(); err == nil {
		container.Bind(&kernel.BxdKernelProvider{HttpEngine: engine})
	}

	// 运行root命令
	console.RunCommand(container)
}
