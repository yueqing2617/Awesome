package middleware

import (
	"Awesome/app/models"
	"errors"
	"github.com/goravel/framework/auth"
	contractshttp "github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func AdminLoginCheck() contractshttp.Middleware {
	return func(ctx contractshttp.Context) {
		token := ctx.Request().Header("token", "")
		payload, err := facades.Auth.Parse(ctx, token)
		if token == "" || err != nil {
			ctx.Request().AbortWithStatusJson(401, contractshttp.Json{
				"Code":    401,
				"Message": "您还没有登录，请先登录",
			})
			return
		}
		errors.Is(err, auth.ErrorTokenExpired)
		if err != nil {
			ctx.Request().AbortWithStatusJson(401, contractshttp.Json{
				"Code":    401,
				"Message": err.Error(),
			})
			return
		}
		var admin models.Admin
		err = facades.Orm.Query().Where("id", payload.Key).With("Role").Find(&admin)
		if admin.ID == 0 {
			ctx.Request().AbortWithStatusJson(401, contractshttp.Json{
				"Code":    401,
				"Message": "登录信息已过期，请重新登录",
			})
			return
		}
		ctx.WithValue("CurrentAdmin", &admin)
		ctx.Request().Next()
	}
}
