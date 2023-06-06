package migrate

import (
	"GoGinStarter/internal/container"
	"GoGinStarter/internal/migrator"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

type RollbackCommand struct {
	container *container.Container
}

func (m *RollbackCommand) RunE(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		errMessage := "rollback ID failed"
		m.container.Log.Error(errMessage)
		return fmt.Errorf(errMessage)
	}
	migrationID := strings.Trim(string(args[0]), " ")
	newMigrator := migrator.NewMigrator(m.container.DB)
	if err := newMigrator.RollbackTo(migrationID); err != nil {
		m.container.Log.Error("Rollback failed: " + err.Error())
		return err
	}
	m.container.Log.Info("Rollback did run successfully for Migration ID " + migrationID)

	return nil
}

func (m *RollbackCommand) NewCommand(container *container.Container) *cobra.Command {
	m.container = container
	migrateCommand := cobra.Command{
		Use:   "rollback",
		Short: "Rollback migration using migration ID",
		RunE:  m.RunE,
	}

	return &migrateCommand
}
