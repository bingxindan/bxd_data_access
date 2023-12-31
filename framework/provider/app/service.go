package app

import (
	"errors"
	"github.com/bingxindan/bxd_data_access/framework"
	"github.com/bingxindan/bxd_data_access/framework/util"
	"github.com/google/uuid"
	flag "github.com/spf13/pflag"
	"path/filepath"
)

// BxdApp 代表bxd框架的App实现
type BxdApp struct {
	container  framework.Container // 服务容器
	baseFolder string              // 基础路径
	appId      string              // 表示当前这个app的唯一id, 可以用于分布式锁等

	configMap map[string]string // 配置加载
}

// AppID 表示这个App的唯一ID
func (h BxdApp) AppID() string {
	return h.appId
}

// Version 实现版本
func (h BxdApp) Version() string {
	return "0.0.3"
}

// BaseFolder 表示基础目录，可以代表开发场景的目录，也可以代表运行时候的目录
func (h BxdApp) BaseFolder() string {
	if h.baseFolder != "" {
		return h.baseFolder
	}

	// 如果参数也没有，使用默认的当前路径
	return util.GetExecDirectory()
}

// ConfigFolder  表示配置文件地址
func (h BxdApp) ConfigFolder() string {
	return filepath.Join(h.BaseFolder(), "config")
}

// LogFolder 表示日志存放地址
func (h BxdApp) LogFolder() string {
	return filepath.Join(h.StorageFolder(), "log")
}

func (h BxdApp) HttpFolder() string {
	return filepath.Join(h.BaseFolder(), "http")
}

func (h BxdApp) ConsoleFolder() string {
	return filepath.Join(h.BaseFolder(), "console")
}

func (h BxdApp) StorageFolder() string {
	return filepath.Join(h.BaseFolder(), "storage")
}

// ProviderFolder 定义业务自己的服务提供者地址
func (h BxdApp) ProviderFolder() string {
	return filepath.Join(h.BaseFolder(), "provider")
}

// MiddlewareFolder 定义业务自己定义的中间件
func (h BxdApp) MiddlewareFolder() string {
	return filepath.Join(h.HttpFolder(), "middleware")
}

// CommandFolder 定义业务定义的命令
func (h BxdApp) CommandFolder() string {
	return filepath.Join(h.ConsoleFolder(), "command")
}

// RuntimeFolder 定义业务的运行中间态信息
func (h BxdApp) RuntimeFolder() string {
	return filepath.Join(h.StorageFolder(), "runtime")
}

// TestFolder 定义测试需要的信息
func (h BxdApp) TestFolder() string {
	return filepath.Join(h.BaseFolder(), "test")
}

// NewBxdApp 初始化BxdApp
func NewBxdApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}

	// 有两个参数，一个是容器，一个是baseFolder
	container := params[0].(framework.Container)
	baseFolder := params[1].(string)
	// 如果没有设置，则使用参数
	if baseFolder == "" {
		flag.StringVar(&baseFolder, "base_folder", "", "base_folder参数, 默认为当前路径")
		flag.Parse()
	}
	appId := uuid.New().String()
	configMap := map[string]string{}

	return &BxdApp{baseFolder: baseFolder, container: container, appId: appId, configMap: configMap}, nil
}

// LoadAppConfig 加载配置map
func (h *BxdApp) LoadAppConfig(kv map[string]string) {
	for key, val := range kv {
		h.configMap[key] = val
	}
}

// AppFolder 代表app目录
func (h *BxdApp) AppFolder() string {
	if val, ok := h.configMap["app_folder"]; ok {
		return val
	}
	return filepath.Join(h.BaseFolder(), "app")
}
