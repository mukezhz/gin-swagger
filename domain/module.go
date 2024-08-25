package domain

import (
	"github.com/mukezhz/gin_swag/domain/hello"
	"github.com/mukezhz/gin_swag/domain/middlewares"

	"go.uber.org/fx"
)

var Module = fx.Options(
	middlewares.Module,
	hello.Module,
	fx.Provide(
		NewDocsRoute,
	),
	fx.Invoke(RegisterDocsRoute),
)
