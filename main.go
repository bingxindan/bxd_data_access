package main

import (
	"bxd_data_access/framework"
	"bxd_data_access/framework/middleware"
	"net/http"
)

func main() {
	core := framework.NewCore()

	//core.Use(middleware.Test1(), middleware.Test2())
	core.Use(middleware.Recovery())
	core.Use(middleware.Cost())
	//core.Use(middleware.Timeout(1*time.Second))

	registerRouter(core)

	server := &http.Server{
		// 自定义的请求核心处理函数
		Handler: core,
		// 请求监听地址
		Addr: ":8888",
	}
	server.ListenAndServe()
}
