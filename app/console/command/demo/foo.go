package demo

import (
	"github.com/bingxindan/bxd_data_access/framework/cobra"
	"log"
)

var FooCommand = &cobra.Command{
	Use:     "foo",
	Short:   "foo的简要说明",
	Long:    "foo的长说明",
	Aliases: []string{"fo", "f"},
	Example: "foo命令的例子",
	RunE: func(c *cobra.Command, args []string) error {
		//container := c.GetContainer()
		log.Println("execute foo command")
		return nil
	},
}

// 绑定业务的命令
func AddAppCommand(rootCmd *cobra.Command) {
	// 每秒调用一次Foo命令
	//rootCmd.
}

var Foo1Command = &cobra.Command{
	Use:     "foo1",
	Short:   "foo1的简要说明",
	Long:    "foo1的长说明",
	Aliases: []string{"fo1", "f1"},
	Example: "foo1命令的例子",
	RunE: func(c *cobra.Command, args []string) error {
		container := c.GetContainer()
		log.Println(container)
		return nil
	},
}

func InitFoo() *cobra.Command {
	FooCommand.AddCommand(Foo1Command)
	return FooCommand
}
