package migrate

import (
	"GoGinStarter/internal/container"
	"GoGinStarter/internal/migrator"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

type MigrateCommand struct {
	container *container.Container
}

func (m *MigrateCommand) RunE(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		errMessage := "migration ID failed"
		m.container.Log.Error(errMessage)
		return fmt.Errorf(errMessage)
	}
	migrationID := strings.Trim(string(args[0]), " ")
	newMigrator := migrator.NewMigrator(m.container.DB)
	if err := newMigrator.MigrateTo(migrationID); err != nil {
		m.container.Log.Error("Migration failed: " + err.Error())
		return err
	}
	m.container.Log.Error("Migration did run successfully for migrate ID " + migrationID)
	return nil
}

func (m *MigrateCommand) NewCommand(container *container.Container) *cobra.Command {
	m.container = container
	migrateCommand := cobra.Command{
		Use:   "migrate",
		Short: "Migrate",
		RunE:  m.RunE,
	}

	return &migrateCommand
}
