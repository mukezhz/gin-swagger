package migrations

import (
	"github.com/mukezhz/gin_swag/pkg/framework"
)

type HelloMigration struct {
	logger framework.Logger
}

func NewHelloMigration(
	logger framework.Logger,
) *HelloMigration {
	return &HelloMigration{
		logger: logger,
	}
}

func (r *HelloMigration) Migrate() error {
	r.logger.Infoln("[Migrating...] Hello")

    // inject the DB connection and run the migration
    
	return nil
}
