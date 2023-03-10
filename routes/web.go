package routes

import (
	"Awesome/app/http/helper"
	"Awesome/app/models"
	"fmt"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func Web() {
	facades.Route.Get("/", func(ctx http.Context) {
		pms := new(models.Permission)
		list, err := pms.GetPermissionList()
		if err != nil {
			return
		}
		fmt.Println(list)
		helper.RestfulSuccess(ctx, "success", http.Json{
			"list": list,
		})
	})

}
