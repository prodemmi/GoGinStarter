package console

import (
	"GoGinStarter/app/console/commands"
	"GoGinStarter/app/console/commands/make"
	"GoGinStarter/app/console/commands/migrate"
	"github.com/spf13/cobra"
)

var Commands = []*cobra.Command{
	(&commands.ServeCommand{}).NewCommand(),
	(&commands.HelloWorldCommand{}).NewCommand(),
	(&make.MakeCommand{}).NewCommand(),
	(&migrate.MigrateCommand{}).NewCommand(),
	(&migrate.RollbackCommand{}).NewCommand(),
}

var rootCmd = &cobra.Command{}

func Run() error {
	for _, command := range Commands {
		rootCmd.AddCommand(command)
	}
	return rootCmd.Execute()
}
