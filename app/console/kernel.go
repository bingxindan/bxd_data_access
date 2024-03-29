package console

import (
	"github.com/bingxindan/bxd_data_access/app/console/command/demo"
	"github.com/bingxindan/bxd_data_access/framework"
	"github.com/bingxindan/bxd_data_access/framework/cobra"
	"github.com/bingxindan/bxd_data_access/framework/command"
)

// RunCommand 初始化根 Command 并运行
func RunCommand(container framework.Container) error {
	// 根 Command
	var rootCmd = &cobra.Command{
		// 定义根命令的关键字
		Use: "bxd",
		// 简短介绍
		Short: "bxd 命令",
		// 根命令的详细介绍
		Long: "bxd 框架提供的命令行工具，使用这个命令行工具能很方便执行框架自带命令，也能很方便编写业务命令",
		// 根命令的执行函数
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.InitDefaultHelpFlag()
			return cmd.Help()
		},
		// 不需要出现 cobra 默认的 completion 子命令
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}

	// 为根 Command 设置服务容器
	rootCmd.SetContainer(container)
	// 绑定框架命令
	command.AddKernelCommands(rootCmd)
	// 绑定业务的命令
	AddAppCommand(rootCmd)

	// 执行 RootCommand
	return rootCmd.Execute()
}

// AddAppCommand 绑定业务命令
func AddAppCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(demo.InitFoo())
}
