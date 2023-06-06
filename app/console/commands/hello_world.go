package commands

import (
	"GoGinStarter/internal/container"
	"fmt"
	"github.com/spf13/cobra"
)

type HelloWorldCommand struct{}

func (s *HelloWorldCommand) RunE(cmd *cobra.Command, args []string) error {
	fmt.Println("Hello World!")
	return nil
}

func (s *HelloWorldCommand) NewCommand(container *container.Container) *cobra.Command {
	return &cobra.Command{
		Use:   "hello",
		Short: "Show hello world",
		Long:  "This command shows hello world.",
		RunE:  s.RunE,
	}
}
