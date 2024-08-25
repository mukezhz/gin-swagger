package migrations

import "github.com/mukezhz/gin_swag/pkg/framework"

type Migrator struct {
	migrations []framework.Migration
	logger     framework.Logger
}

func NewMigrator(
	migrations []framework.Migration,
	logger framework.Logger,
) *Migrator {
	return &Migrator{
		migrations: migrations,
		logger:     logger,
	}
}

func (m *Migrator) Exec() error {
	for _, migration := range m.migrations {
		if err := migration.Migrate(); err != nil {
			return err
		}
	}
	return nil
}
