package config

import (
	"github.com/bingxindan/bxd_data_access/framework"
	"github.com/bingxindan/bxd_data_access/framework/contract"
	"path/filepath"
)

type BxdConfigProvider struct {
}

// Register registe a new function for make a service instance
func (provider *BxdConfigProvider) Register(c framework.Container) framework.NewInstance {
	return NewBxdConfig
}

// Boot will called when the service instantiate
func (provider *BxdConfigProvider) Boot(c framework.Container) error {
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *BxdConfigProvider) IsDefer() bool {
	return false
}

func (provider *BxdConfigProvider) Params(c framework.Container) []interface{} {
	appService := c.MustMake(contract.AppKey).(contract.App)
	envService := c.MustMake(contract.EnvKey).(contract.Env)
	env := envService.AppEnv()
	// 配置文件夹地址
	configFolder := appService.ConfigFolder()
	envFolder := filepath.Join(configFolder, env)

	return []interface{}{c, envFolder, envService.All()}
}

func (provider *BxdConfigProvider) Name() string {
	return contract.ConfigKey
}
