package migrations

import (
	"github.com/mukezhz/gin_swag/pkg/framework"

	"go.uber.org/fx"
)

func AsMigrator(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(framework.Migration)),
		fx.ResultTags(`group:"migrations"`),
	)
}

var (
	Module = fx.Module("migrations",
		fx.Provide(
			AsMigrator(NewHelloMigration),
			fx.Annotate(
				NewMigrator,
				fx.ParamTags(`group:"migrations"`),
			),
		),
	)
)
