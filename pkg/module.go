package pkg

import (
	"github.com/mukezhz/gin_swag/pkg/framework"
	"github.com/mukezhz/gin_swag/pkg/infrastructure"
	"github.com/mukezhz/gin_swag/pkg/middlewares"
	"github.com/mukezhz/gin_swag/pkg/services"

	"go.uber.org/fx"
)

var Module = fx.Module("pkg",
	framework.Module,
	services.Module,
	middlewares.Module,
	infrastructure.Module,
)
