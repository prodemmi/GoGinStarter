package commands

import (
	"GoGinStarter/app/http/routes"
	"GoGinStarter/internal/container"
	"GoGinStarter/wire"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	csrf "github.com/utrack/gin-csrf"
)

type ServeCommand struct{}

func (s *ServeCommand) RunE(cmd *cobra.Command, args []string) error {
	container := wire.InitializeContainer()
	router := gin.Default()

	routes.SetupApiRoutes(router, container)

	router.Use(csrf.Middleware(csrf.Options{
		Secret: "SHsHZ28711587148418",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}))

	err := router.Run(fmt.Sprintf(":%v", container.Config.App.Port))
	if err != nil {
		container.Log.Error(err.Error())
	}

	return nil
}

func (s *ServeCommand) NewCommand(container *container.Container) *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Start the server",
		Long:  "This command starts the server.",
		RunE:  s.RunE,
	}
}
