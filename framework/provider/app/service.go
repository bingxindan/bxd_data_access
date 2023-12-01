package app

import (
	"errors"
	"github.com/bingxindan/bxd_data_access/framework"
)

// BxdApp 代表bxd框架的App实现
type BxdApp struct {
	container  framework.Container // 服务容器
	baseFolder string              // 基础路径
}

// NewBxdApp 初始化BxdApp
func NewBxdApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}

	// 有两个参数，一个是容器，一个是baseFolder
	container := params[0].(framework.Container)
	baseFolder := params[1].(string)

	return &BxdApp{baseFolder: baseFolder, container: container}, nil
}
