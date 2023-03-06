package providers

import (
	"github.com/goravel/framework/facades"

	"agpc/app/http"
	"agpc/routes"
)

type RouteServiceProvider struct {
}

func (receiver *RouteServiceProvider) Register() {
	//Add HTTP middlewares
	kernel := http.Kernel{}
	facades.Route.GlobalMiddleware(kernel.Middleware()...)
}

func (receiver *RouteServiceProvider) Boot() {
	//Add routes
	routes.Web()
}
