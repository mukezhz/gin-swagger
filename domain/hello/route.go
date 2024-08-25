package hello

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mukezhz/gin_swag/pkg/infrastructure"
	"github.com/savaki/swag/endpoint"
	"github.com/savaki/swag/swagger"
)

type Route struct {
	router      *infrastructure.Router
	controller  *Controller
	groupRouter *gin.RouterGroup
}

func NewRoute(router *infrastructure.Router, controller *Controller) *Route {
	route := Route{router: router, controller: controller}
	route.groupRouter = route.router.Group("api/hello")
	return &route
}

func RegisterRoute(r *Route) []*swagger.Endpoint {
	endpoints := []*swagger.Endpoint{}

	rootHandler := endpoint.New(
		http.MethodGet,
		"/hello",
		"Handle Hello Root",
		endpoint.Handler(r.controller.HandleRoot),
		endpoint.Response(http.StatusOK, Model{}, "Hello Root handler Response"),
	)
	getHelloByIDHandler := endpoint.New(
		http.MethodGet,
		"/hello/:id",
		"Handle Get Hello By ID",
		endpoint.Handler(r.controller.HandleGreet),
		endpoint.Response(http.StatusOK, Greet{}, "Hello  get by id Response"),
	)

	addHelloHandler := endpoint.New(
		http.MethodPost,
		"/hello",
		"Handle Add Hello",
		endpoint.Handler(r.controller.HandleAddGreet),
		endpoint.Body(Greet{}, "Hello Request Payload", true),
		endpoint.Response(http.StatusCreated, gin.H{}, "Hello create Response"),
	)

	updateHelloHandler := endpoint.New(
		http.MethodPut,
		"/hello/:id",
		"Handle Update Hello",
		endpoint.Handler(r.controller.HandleUpdateGreet),
		endpoint.Response(http.StatusOK, gin.H{}, "Hello update Response"),
	)

	deleteHelloHandler := endpoint.New(
		http.MethodDelete,
		"/hello/:id",
		"Handle Delete Hello",
		endpoint.Handler(r.controller.HandleDeleteGreet),
		endpoint.Response(http.StatusOK, gin.H{}, "Hello Response"),
	)

	endpoints = append(endpoints, rootHandler, getHelloByIDHandler, addHelloHandler, updateHelloHandler, deleteHelloHandler)
	return endpoints
}
