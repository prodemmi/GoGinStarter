//go:build wireinject
// +build wireinject

package console

import (
	"GoGinStarter/app/console/commands"
	"GoGinStarter/app/console/commands/make"
	"GoGinStarter/wire"
	"github.com/spf13/cobra"
	"os"
)

var Commands = []*cobra.Command{
	(&commands.ServeCommand{}).NewCommand(),
	(&commands.HelloWorldCommand{}).NewCommand(),
	(&make.MakeCommand{}).NewCommand(),
}

var rootCmd = &cobra.Command{}

func Run() error {
	container := wire.InitializeContainer()
	for _, command := range Commands {
		rootCmd.AddCommand(command)
	}

	if len(os.Args) == 1 && container.Config.App.Debug == true {
		return Commands[0].Execute()
	}

	return rootCmd.Execute()
}
