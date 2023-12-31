package http

import (
	"github.com/bingxindan/bxd_data_access/app/http/middleware/cors"
	"github.com/bingxindan/bxd_data_access/app/http/module/demo"
	"github.com/bingxindan/bxd_data_access/framework/contract"
	"github.com/bingxindan/bxd_data_access/framework/gin"
	ginSwagger "github.com/bingxindan/bxd_data_access/framework/middleware/gin-swagger"
	"github.com/bingxindan/bxd_data_access/framework/middleware/gin-swagger/swaggerFiles"
)

func Routes(r *gin.Engine) {
	container := r.GetContainer()
	configService := container.MustMake(contract.ConfigKey).(contract.Config)

	r.Static("/dist/", "./dist/")

	// 使用cors中间件
	r.Use(cors.Default())

	// 如果配置了swagger，则显示swagger的中间件
	if configService.GetBool("app.swagger") == true {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	demo.Register(r)
}
