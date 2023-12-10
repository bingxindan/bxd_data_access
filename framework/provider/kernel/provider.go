package kernel

import (
	"github.com/bingxindan/bxd_data_access/framework"
	"github.com/bingxindan/bxd_data_access/framework/contract"
	"github.com/bingxindan/bxd_data_access/framework/gin"
)

type BxdKernelProvider struct {
	HttpEngine *gin.Engine
}

func (provider *BxdKernelProvider) Register(c framework.Container) framework.NewInstance {
	return nil
}

func (provider *BxdKernelProvider) Boot(c framework.Container) error {
	if provider.HttpEngine == nil {
		provider.HttpEngine = gin.Default()
	}
	provider.HttpEngine.SetContainer(c)
	return nil
}

func (provider *BxdKernelProvider) IsDefer() bool {
	return false
}

func (provider *BxdKernelProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.HttpEngine}
}

func (provider *BxdKernelProvider) Name() string {
	return contract.KernelKey
}
