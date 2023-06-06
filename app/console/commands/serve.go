package commands

import (
	"GoGinStarter/app/http/routes"
	"GoGinStarter/internal/container"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	csrf "github.com/utrack/gin-csrf"
)

type ServeCommand struct {
	container *container.Container
}

func (s *ServeCommand) RunE(cmd *cobra.Command, args []string) error {
	router := gin.Default()

	router.Use(sessions.Sessions("mysession", s.container.Session.Store))
	router.Use(csrf.Middleware(csrf.Options{
		Secret: "SHsHZ28711587148418",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}))

	routes.SetupApiRoutes(router, s.container)

	err := router.Run(fmt.Sprintf(":%v", s.container.Config.App.Port))
	if err != nil {
		s.container.Log.Error(err.Error())
	}

	return nil
}

func (s *ServeCommand) NewCommand(container *container.Container) *cobra.Command {
	s.container = container
	return &cobra.Command{
		Use:   "serve",
		Short: "Start the server",
		Long:  "This command starts the server.",
		RunE:  s.RunE,
	}
}
