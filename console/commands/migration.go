package commands

import (
	"github.com/mukezhz/gin_swag/migrations"
	"github.com/mukezhz/gin_swag/pkg/framework"

	"github.com/spf13/cobra"
)

type MigrateCommand struct {
}

func (s *MigrateCommand) Short() string {
	return "run migrate command"
}

func NewMigrateCommand() *MigrateCommand {
	return &MigrateCommand{}
}

func (s *MigrateCommand) Setup(cmd *cobra.Command) {
	// left empty intentionally
}

func (s *MigrateCommand) Run() framework.CommandRunner {
	return func(
		migrator *migrations.Migrator,
	) {
		if err := migrator.Exec(); err != nil {
			framework.GetLogger().Fatal(err)
		}
	}
}
