package main

import (
	"github.com/bingxindan/bxd_data_access/framework/gin"
	"github.com/bingxindan/bxd_data_access/framework/middleware"
)

func registerRouter(core *gin.Engine) {
	// 静态路由+HTTP方法匹配
	core.GET("/user/login", middleware.Test3(), UserLoginController)

	// 批量通用前缀
	subjectApi := core.Group("/subject")
	{
		subjectApi.Use(middleware.Test3())
		// 动态路由
		subjectApi.DELETE("/:id", SubjectDelController)
		subjectApi.PUT("/:id", SubjectUpdateController)
		subjectApi.GET("/:id", SubjectGetController)
		subjectApi.GET("/list/all", SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.GET("/name", SubjectNameController)
		}
	}
}
