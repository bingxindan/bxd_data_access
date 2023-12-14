package kernel

import (
	"github.com/bingxindan/bxd_data_access/framework/gin"
	"net/http"
)

// BxdKernelService 引擎服务
type BxdKernelService struct {
	engine *gin.Engine
}

// NewBxdKernelService 初始化web引擎服务实例
func NewBxdKernelService(params ...interface{}) (interface{}, error) {
	httpEngine := params[0].(*gin.Engine)
	return &BxdKernelService{engine: httpEngine}, nil
}

// HttpEngine 返回web引擎
func (s *BxdKernelService) HttpEngine() http.Handler {
	return s.engine
}
