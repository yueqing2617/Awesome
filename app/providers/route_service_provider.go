package providers

import (
	"github.com/goravel/framework/facades"

	"Awesome/app/http"
	"Awesome/routes"
)

type RouteServiceProvider struct {
}

func (receiver *RouteServiceProvider) Register() {
	//Add HTTP middlewares
	facades.Route.GlobalMiddleware(http.Kernel{}.Middleware()...)
}

func (receiver *RouteServiceProvider) Boot() {
	receiver.configureRateLimiting()

	routes.Backend()
	routes.Manager()
}

func (receiver *RouteServiceProvider) configureRateLimiting() {

}
