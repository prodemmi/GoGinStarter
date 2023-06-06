package seed

import (
	"GoGinStarter/internal/container"
	"GoGinStarter/internal/seeder"
	"github.com/spf13/cobra"
)

type SeedCommand struct {
	seeder seeder.Seeder
}

func (s *SeedCommand) RunE(cmd *cobra.Command, args []string) error {
	return s.seeder.Seed()
}

func (s *SeedCommand) NewCommand(container *container.Container) *cobra.Command {
	s.seeder = container.Seeder
	return &cobra.Command{
		Use:   "seed",
		Short: "Seed into database",
		Long:  "Seed into database.",
		RunE:  s.RunE,
	}
}
