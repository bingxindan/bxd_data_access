package command

import (
	"context"
	"github.com/bingxindan/bxd_data_access/framework/cobra"
	"github.com/bingxindan/bxd_data_access/framework/contract"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// app启动地址
var appAddress = ""

// initAppCommand 初始化app命令和其子命令
func initAppCommand() *cobra.Command {
	appStartCommand.Flags().StringVar(&appAddress, "address", ":8888", "设置app启动的地址，默认为:8888")

	appCommand.AddCommand(appStartCommand)
	return appCommand
}

// appStartCommand 启动一个Web服务
var appStartCommand = &cobra.Command{
	Use:   "start",
	Short: "启动一个Web服务",
	RunE: func(c *cobra.Command, args []string) error {
		// 从Command中获取服务容器
		container := c.GetContainer()
		// 从服务容器中获取kernel的服务实例
		kernelService := container.MustMake(contract.KernelKey).(contract.Kernel)
		// 从kernel服务实例中获取引擎
		core := kernelService.HttpEngine()

		// 创建一个Server服务
		server := &http.Server{
			Handler: core,
			Addr:    ":8888",
		}

		// 这个goroutine是启动服务的goroutine
		go func() {
			server.ListenAndServe()
		}()

		// 当前goroutine等待信号量
		quit := make(chan os.Signal)
		// 监控信号：SIGINT, SIGTERM, SIGQUIT
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		// 这里会阻塞当前goroutine等待信号
		<-quit

		// 调用Server.Shutdown graceful结束
		timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(timeoutCtx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}

		return nil
	},
}
