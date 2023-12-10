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
	// 挂载AppCommand命令
	root.AddCommand(initAppCommand())
}
