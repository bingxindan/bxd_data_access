package id

import (
	"github.com/bingxindan/bxd_data_access/framework"
	"github.com/bingxindan/bxd_data_access/framework/contract"
)

type BxdIDProvider struct {
}

// Register registe a new function for make a service instance
func (provider *BxdIDProvider) Register(c framework.Container) framework.NewInstance {
	return NewBxdIDService
}

// Boot will called when the service instantiate
func (provider *BxdIDProvider) Boot(c framework.Container) error {
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *BxdIDProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *BxdIDProvider) Params(c framework.Container) []interface{} {
	return []interface{}{}
}

// Name define the name for this service
func (provider *BxdIDProvider) Name() string {
	return contract.IDKey
}
