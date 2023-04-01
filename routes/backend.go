package routes

import (
	"Awesome/app/http/helper"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func Backend() {
	facades.Route.Get("/", func(ctx http.Context) {
		result := helper.GetEnv("APP_NAME", "Awesome1")
		helper.RestfulSuccess(ctx, "请求成功", result)
	})
	facades.Route.Static("static", "./storage")
}
