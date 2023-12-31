package trace

import (
	"github.com/bingxindan/bxd_data_access/framework"
	"github.com/bingxindan/bxd_data_access/framework/contract"
)

type BxdTraceProvider struct {
	c framework.Container
}

// Register a new function for make a service instance
func (provider *BxdTraceProvider) Register(c framework.Container) framework.NewInstance {
	return NewBxdTraceService
}

// Boot will called when the service instantiate
func (provider *BxdTraceProvider) Boot(c framework.Container) error {
	provider.c = c
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *BxdTraceProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *BxdTraceProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.c}
}

// Name define the name for this service
func (provider *BxdTraceProvider) Name() string {
	return contract.TraceKey
}
