package middleware

import (
	contractshttp "github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func ErrorHandler() contractshttp.Middleware {
	return func(ctx contractshttp.Context) {
		//处理全局的错误
		defer func() {
			if err := recover(); err != nil {
				facades.Log.Errorf("Error: %v", err)
				ctx.Request().AbortWithStatusJson(500, contractshttp.Json{
					"code":    500,
					"message": err,
				})
			}
		}()
		//继续执行
		ctx.Request().Next()
	}
}
