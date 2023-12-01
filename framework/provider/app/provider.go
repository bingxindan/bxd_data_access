package app

import (
	"github.com/bingxindan/bxd_data_access/framework"
	"github.com/bingxindan/bxd_data_access/framework/contract"
)

// BxdAppProvider 提供App的具体实现方法
type BxdAppProvider struct {
	BaseFolder string
}

// Register 注册BxdApp方法
func (b *BxdAppProvider) Register(container framework.Container) framework.NewInstance {
	return NewBxdApp
}

// Boot 启动调用
func (b *BxdAppProvider) Boot(container framework.Container) error {
	return nil
}

// IsDefer 是否延迟初始化
func (b *BxdAppProvider) IsDefer() bool {
	return false
}

func (b *BxdAppProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container, b.BaseFolder}
}

func (b *BxdAppProvider) Name() string {
	return contract.AppKey
}
