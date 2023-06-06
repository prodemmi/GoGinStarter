package console

import (
	"GoGinStarter/app/console/commands"
	"GoGinStarter/app/console/commands/make"
	"GoGinStarter/app/console/commands/migrate"
	"GoGinStarter/app/console/commands/seed"
	"GoGinStarter/internal/container"
	"github.com/spf13/cobra"
)

var Commands = []func(container *container.Container) *cobra.Command{
	(&commands.ServeCommand{}).NewCommand,
	(&commands.HelloWorldCommand{}).NewCommand,
	(&make.MakeCommand{}).NewCommand,
	(&migrate.MigrateCommand{}).NewCommand,
	(&migrate.RollbackCommand{}).NewCommand,
	(&seed.SeedCommand{}).NewCommand,
}

var rootCmd = &cobra.Command{}

func Run(container *container.Container) error {
	for _, command := range Commands {
		rootCmd.AddCommand(command(container))
	}
	return rootCmd.Execute()
}
