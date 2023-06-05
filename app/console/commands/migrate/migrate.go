package migrate

import (
	"GoGinStarter/internal/migrator"
	"GoGinStarter/wire"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

type MigrateCommand struct{}

func (m *MigrateCommand) RunE(cmd *cobra.Command, args []string) error {
	var container = wire.InitializeContainer()
	if len(args) == 0 {
		errMessage := "migration ID failed"
		container.Log.Error(errMessage)
		return fmt.Errorf(errMessage)
	}
	migrationID := strings.Trim(string(args[0]), " ")
	newMigrator := migrator.NewMigrator(container.DB)
	if err := newMigrator.MigrateTo(migrationID); err != nil {
		container.Log.Error("Migration failed: " + err.Error())
		return err
	}
	container.Log.Error("Migration did run successfully for migrate ID " + migrationID)
	return nil
}

func (m *MigrateCommand) NewCommand() *cobra.Command {
	migrateCommand := cobra.Command{
		Use:   "migrate",
		Short: "Migrate",
		RunE:  m.RunE,
	}

	return &migrateCommand
}
