package command

import (
	"fmt"
	"github.com/bingxindan/bxd_data_access/framework"
	"github.com/bingxindan/bxd_data_access/framework/cobra"
	"github.com/bingxindan/bxd_data_access/framework/contract"
	"github.com/bingxindan/bxd_data_access/framework/util"
	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"
)

// devConfig 代表调试模式的配置信息
type devConfig struct {
	Port    string   // 调试模式最终监听的端口，默认为8070
	Backend struct { // 后端调试模式配置
		RefreshTime   int    // 调试模式后端更新时间，如果文件变更，等待3s才进行一次更新，能让频繁保存变更更为顺畅, 默认1s
		Port          string // 后端监听端口， 默认 8072
		MonitorFolder string // 监听文件夹，默认为AppFolder
	}
}

// 初始化配置文件
func initDevConfig(c framework.Container) *devConfig {
	devConfig := &devConfig{
		Port: "8087",
		Backend: struct {
			RefreshTime   int
			Port          string
			MonitorFolder string
		}{
			1,
			"8072",
			"",
		},
	}
	configer := c.MustMake(contract.ConfigKey).(contract.Config)
	if configer.IsExist("app.dev.port") {
		devConfig.Port = configer.GetString("app.dev.port")
	}
	if configer.IsExist("app.dev.backend.refresh_time") {
		devConfig.Backend.RefreshTime = configer.GetInt("app.dev.backend.refresh_time")
	}
	if configer.IsExist("app.dev.backend.port") {
		devConfig.Backend.Port = configer.GetString("app.dev.backend.port")
	}
	monitorFolder := configer.GetString("app.dev.backend.monitor_folder")
	if monitorFolder == "" {
		appService := c.MustMake(contract.AppKey).(contract.App)
		devConfig.Backend.MonitorFolder = appService.AppFolder()
	}

	return devConfig
}

// Proxy 代表serve启动的服务器代理
type Proxy struct {
	devConfig   *devConfig   // 配置文件
	proxyServer *http.Server // proxy的服务
	backendPid  int          // 当前的backend服务的pid
	frontendPid int          // 当前的frontend服务的pid
}

// NewProxy 初始化一个Proxy
func NewProxy(c framework.Container) *Proxy {
	devConfig := initDevConfig(c)
	return &Proxy{
		devConfig: devConfig,
	}
}

// 重新启动一个proxy网关
func (p *Proxy) newProxyReverseProxy(frontend, backend *url.URL) *httputil.ReverseProxy {
	if p.backendPid == 0 {
		fmt.Println("前端和后端服务都不存在")
		return nil
	}

	// 后端服务存在
	if p.frontendPid == 0 && p.backendPid != 0 {
		return httputil.NewSingleHostReverseProxy(backend)
	}

	// 两个都有进程
	// 先创建一个后端服务的directory
	director := func(req *http.Request) {
		req.URL.Scheme = backend.Scheme
		req.URL.Host = backend.Host
	}

	// 定义一个NotFoundErr
	NotFoundErr := errors.New("response is 404, need to redirect")
	return &httputil.ReverseProxy{
		Director: director, // 先转发到后端服务
		ModifyResponse: func(response *http.Response) error {
			// 如果后端服务返回了404，我们返回NotFoundErr 会进入到errorHandler中
			if response.StatusCode == 404 {
				return NotFoundErr
			}
			return nil
		},
		ErrorHandler: func(writer http.ResponseWriter, request *http.Request, err error) {
			// 判断 Error 是否为NotFoundError, 是的话则进行前端服务的转发，重新修改writer
			if errors.Is(err, NotFoundErr) {
				httputil.NewSingleHostReverseProxy(frontend).ServeHTTP(writer, request)
			}
		}}
}

// rebuildBackend 重新编译后端
func (p *Proxy) rebuildBackend() error {
	// 重新编译bxd
	cmdBuild := exec.Command("./bxd", "build", "backend")
	cmdBuild.Stdout = os.Stdout
	cmdBuild.Stderr = os.Stderr
	if err := cmdBuild.Start(); err == nil {
		err = cmdBuild.Wait()
		if err != nil {
			return err
		}
	}
	return nil
}

// restartBackend 启动后端服务
func (p *Proxy) restartBackend() error {

	// 杀死之前的进程
	if p.backendPid != 0 {
		syscall.Kill(p.backendPid, syscall.SIGKILL)
		p.backendPid = 0
	}

	// 设置随机端口，真实后端的端口
	port := p.devConfig.Backend.Port
	bxdAddress := fmt.Sprintf(":" + port)
	// 使用命令行启动后端进程
	cmd := exec.Command("./bxd", "app", "start", "--address="+bxdAddress)
	cmd.Stdout = os.NewFile(0, os.DevNull)
	cmd.Stderr = os.Stderr
	fmt.Println("启动后端服务: ", "http://127.0.0.1:"+port)
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	p.backendPid = cmd.Process.Pid
	fmt.Println("后端服务pid:", p.backendPid)
	return nil
}

// 重启后端服务, 如果frontend为nil，则没有包含后端
func (p *Proxy) startProxy(startBackend bool) error {
	var backendURL, frontendURL *url.URL
	var err error

	// 启动后端
	if startBackend {
		if err := p.restartBackend(); err != nil {
			return err
		}
	}

	// 如果已经启动过proxy了，就不要进行设置了
	if p.proxyServer != nil {
		return nil
	}

	if backendURL, err = url.Parse(fmt.Sprintf("%s%s", "http://127.0.0.1:", p.devConfig.Backend.Port)); err != nil {
		return err
	}

	// 设置反向代理
	proxyReverse := p.newProxyReverseProxy(frontendURL, backendURL)
	p.proxyServer = &http.Server{
		Addr:    "127.0.0.1:" + p.devConfig.Port,
		Handler: proxyReverse,
	}

	fmt.Println("代理服务启动:", "http://"+p.proxyServer.Addr)
	// 启动proxy服务
	err = p.proxyServer.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

// monitorBackend 监听应用文件
func (p *Proxy) monitorBackend() error {
	// 监听
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	appFolder := p.devConfig.Backend.MonitorFolder
	fmt.Println("监控文件夹：", appFolder)
	filepath.Walk(appFolder, func(path string, info os.FileInfo, err error) error {
		if info != nil && !info.IsDir() {
			return nil
		}
		if util.IsHiddenDirectory(path) {
			return nil
		}
		return watcher.Add(path)
	})

	refreshTime := p.devConfig.Backend.RefreshTime
	t := time.NewTimer(time.Duration(refreshTime) * time.Second)
	t.Stop()
	for {
		select {
		case <-t.C:
			fmt.Println("...检测到文件更新，重启服务开始...")
			if err := p.rebuildBackend(); err != nil {
				fmt.Println("重新编译失败：", err.Error())
			} else {
				if err := p.restartBackend(); err != nil {
					fmt.Println("重新启动失败：", err.Error())
				}
			}
			fmt.Println("...检测到文件更新，重启服务结束...")
			t.Stop()
		case _, ok := <-watcher.Events:
			if !ok {
				continue
			}
			t.Reset(time.Duration(refreshTime) * time.Second)
		case err, ok := <-watcher.Errors:
			if !ok {
				continue
			}
			fmt.Println("监听文件夹错误：", err.Error())
			t.Reset(time.Duration(refreshTime) * time.Second)
		}
	}
	return nil
}

// 初始化Dev命令
func initDevCommand() *cobra.Command {
	devCommand.AddCommand(devBackendCommand)
	devCommand.AddCommand(devAllCommand)
	return devCommand
}

// devCommand 为调试模式的一级命令
var devCommand = &cobra.Command{
	Use:   "dev",
	Short: "调试模式",
	RunE: func(c *cobra.Command, args []string) error {
		c.Help()
		return nil
	},
}

// devBackendCommand 启动后端调试模式
var devBackendCommand = &cobra.Command{
	Use:   "backend",
	Short: "启动后端调试模式",
	RunE: func(c *cobra.Command, args []string) error {
		proxy := NewProxy(c.GetContainer())
		go proxy.monitorBackend()
		if err := proxy.startProxy(true); err != nil {
			return err
		}
		return nil
	},
}

var devAllCommand = &cobra.Command{
	Use:   "all",
	Short: "同时启动后端调试",
	RunE: func(c *cobra.Command, args []string) error {
		proxy := NewProxy(c.GetContainer())
		go proxy.monitorBackend()
		if err := proxy.startProxy(true); err != nil {
			return err
		}
		return nil
	},
}
