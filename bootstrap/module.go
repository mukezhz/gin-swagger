package bootstrap

import (
	"github.com/mukezhz/gin_swag/domain"
	"github.com/mukezhz/gin_swag/migrations"
	"github.com/mukezhz/gin_swag/pkg"
	"github.com/mukezhz/gin_swag/seeds"

	"go.uber.org/fx"
)

var CommonModules = fx.Module("common",
	fx.Options(
		pkg.Module,
		seeds.Module,
		migrations.Module,
		domain.Module,
	),
)
