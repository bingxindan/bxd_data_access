package env

import (
	"github.com/bingxindan/bxd_data_access/framework"
	"github.com/bingxindan/bxd_data_access/framework/contract"
)

type BxdEnvProvider struct {
	Folder string
}

// Register registe a new function for make a service instance
func (provider *BxdEnvProvider) Register(c framework.Container) framework.NewInstance {
	return NewBxdEnv
}

// Boot will called when the service instantiate
func (provider *BxdEnvProvider) Boot(c framework.Container) error {
	app := c.MustMake(contract.AppKey).(contract.App)
	provider.Folder = app.BaseFolder()
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *BxdEnvProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *BxdEnvProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.Folder}
}

// Name define the name for this service
func (provider *BxdEnvProvider) Name() string {
	return contract.EnvKey
}
