package domain

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mukezhz/gin_swag/domain/hello"
	"github.com/mukezhz/gin_swag/pkg/framework"
	"github.com/mukezhz/gin_swag/pkg/infrastructure"
	"github.com/savaki/swag"
	"github.com/savaki/swag/swagger"
)

type SwagRoute struct {
	logger      framework.Logger
	helloRouter *hello.Route
	endpoints   []*swagger.Endpoint
	router      *infrastructure.Router
}

func NewDocsRoute(
	logger framework.Logger,
	helloRouter *hello.Route,
	router *infrastructure.Router,
) *SwagRoute {
	return &SwagRoute{
		logger:      logger,
		helloRouter: helloRouter,
		endpoints:   []*swagger.Endpoint{},
		router:      router,
	}
}

func (d *SwagRoute) combine(endpoints []*swagger.Endpoint) []*swagger.Endpoint {
	return append(d.endpoints, endpoints...)
}

func RegisterDocsRoute(d *SwagRoute) {
	d.logger.Info("Setting up docs routes")
	helloEndpoints := hello.RegisterRoute(d.helloRouter)
	d.logger.Infoln("Registering hello route", len(helloEndpoints))
	endpoints := d.combine(helloEndpoints)
	d.logger.Infoln("Registering docs route", len(endpoints))
	api := swag.New(
		swag.Endpoints(endpoints...),
		swag.Description("THis is the test description"),
		swag.Version("1.0.0"),
		swag.Title("Test Title"),
	)
	api.Walk(func(path string, endpoint *swagger.Endpoint) {
		h := endpoint.Handler.(func(c *gin.Context))
		path = swag.ColonPath(path)

		d.router.Handle(endpoint.Method, path, h)
	})

	enableCors := true
	d.router.GET("/swagger", gin.WrapH(api.Handler(enableCors)))
	d.router.LoadHTMLGlob("templates/*.html")
	d.router.GET("/docs", func(ctx *gin.Context) {
		ctx.Header("Content-Type", "text/html")
		d.logger.Infoln(ctx.Request.Host)
		d.logger.Infof("%#v", ctx.Request.URL)
		d.logger.Infoln(ctx.Request.TLS)
		scheme := "http://"
		if ctx.Request.TLS != nil {
			scheme = "https://"
		}
		content := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Scalar API Reference</title>
			<meta charset="utf-8" />
			<meta name="viewport" content="width=device-width, initial-scale=1" />
		</head>
		<body>
			<!-- Need a Custom Header? Check out this example https://codepen.io/scalarorg/pen/VwOXqam -->
			<script
			id="api-reference"
			type="application/json"
			data-url="%s"
			></script>
			<script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
		</body>
		</html>
		`, scheme+ctx.Request.Host+"/swagger")
		ctx.String(http.StatusOK, content)
	})

}
