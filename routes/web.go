package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func Web() {
	facades.Route.Get("/", func(ctx http.Context) {
		ctx.Response().Json(200, http.Json{
			"Hello": "Goravel",
		})
	})

}
