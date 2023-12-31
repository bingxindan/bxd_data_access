package command

import "github.com/bingxindan/bxd_data_access/framework/cobra"

// AppCommand 是命令行参数第一级为 app 的命令，它没有实际功能，只是打印帮助文档
var appCommand = &cobra.Command{
	Use:   "app",
	Short: "业务应用控制命令",
	RunE: func(c *cobra.Command, args []string) error {
		// 打印帮助文档
		c.Help()
		return nil
	},
}

// AddKernelCommands will add all command/* to root command
func AddKernelCommands(root *cobra.Command) {
	// app 命令
	root.AddCommand(initAppCommand())
	// env 命令
	root.AddCommand(initEnvCommand())
	// cron 命令
	root.AddCommand(initCronCommand())
	// config 命令
	root.AddCommand(initConfigCommand())
	// build 命令
	root.AddCommand(initBuildCommand())
	// go build
	root.AddCommand(goCommand)
	// dev
	root.AddCommand(initDevCommand())
	// cmd
	root.AddCommand(initCmdCommand())
	// provider
	root.AddCommand(initProviderCommand())
	// middleware
	root.AddCommand(initMiddlewareCommand())
	//
	//// swagger
	//swagger.IndexCommand.AddCommand(swagger.InitServeCommand())
	//swagger.IndexCommand.AddCommand(swagger.GenCommand)
	//root.AddCommand(swagger.IndexCommand)
	//
}
