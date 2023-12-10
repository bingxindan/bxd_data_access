package http

import "github.com/bingxindan/bxd_data_access/framework/gin"

func NewHttpEngine() (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	Routes(r)

	return r, nil
}
