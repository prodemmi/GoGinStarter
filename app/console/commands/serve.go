package commands

import (
	"GoGinStarter/app/http/routes"
	"GoGinStarter/wire"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

type ServeCommand struct{}

func (s *ServeCommand) RunE(cmd *cobra.Command, args []string) error {
	container := wire.InitializeContainer()
	router := gin.Default()

	routes.SetupApiRoutes(router, container)

	err := router.Run(fmt.Sprintf(":%v", container.Config.App.Port))
	if err != nil {
		container.Log.Error(err.Error())
	}

	return nil
}

func (s *ServeCommand) NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Start the server",
		Long:  "This command starts the server.",
		RunE:  s.RunE,
	}
}
