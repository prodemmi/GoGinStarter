package migrate

import (
	"GoGinStarter/internal/container"
	"GoGinStarter/internal/migrator"
	"GoGinStarter/wire"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

type RollbackCommand struct{}

func (m *RollbackCommand) RunE(cmd *cobra.Command, args []string) error {
	var container = wire.InitializeContainer()
	if len(args) == 0 {
		errMessage := "rollback ID failed"
		container.Log.Error(errMessage)
		return fmt.Errorf(errMessage)
	}
	migrationID := strings.Trim(string(args[0]), " ")
	newMigrator := migrator.NewMigrator(container.DB)
	if err := newMigrator.RollbackTo(migrationID); err != nil {
		container.Log.Error("Rollback failed: " + err.Error())
		return err
	}
	container.Log.Info("Rollback did run successfully for Migration ID " + migrationID)

	return nil
}

func (m *RollbackCommand) NewCommand(container *container.Container) *cobra.Command {
	migrateCommand := cobra.Command{
		Use:   "rollback",
		Short: "Rollback migration using migration ID",
		RunE:  m.RunE,
	}

	return &migrateCommand
}
