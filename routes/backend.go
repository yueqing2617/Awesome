package routes

import (
	"Awesome/app/http/helper"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func Backend() {
	facades.Route.Get("/", func(ctx http.Context) {
		password, err := facades.Hash.Make("123456")
		if err != nil {
			helper.RestfulError(ctx, err.Error())
			return
		}
		helper.RestfulSuccess(ctx, "success", http.Json{
			"password": password,
		})
	})
	facades.Route.Static("static", "./storage")
}
