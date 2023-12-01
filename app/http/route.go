package http

import (
	"github.com/bingxindan/bxd_data_access/app/http/module/demo"
	"github.com/bingxindan/bxd_data_access/framework/gin"
)

func Routes(r *gin.Engine) {
	r.Static("/dist/", "./dist/")

	demo.Register(r)
}
