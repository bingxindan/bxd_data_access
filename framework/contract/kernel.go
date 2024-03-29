package contract

import "net/http"

// KernelKey 提供 kernel 服务凭证
const KernelKey = "bxd:kernel"

// Kernel 接口提供框架最核心的结构
type Kernel interface {
	// HttpEngine http.Handler结构，作为net/http框架使用, 实际上是gin.Engine
	HttpEngine() http.Handler
}
