package middleware

import (
	"Awesome/app/models"
	"errors"
	"fmt"
	"github.com/goravel/framework/auth"
	contractshttp "github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"regexp"
	"strings"
)

func AdminAuthCheck() contractshttp.Middleware {
	return func(ctx contractshttp.Context) {
		token := ctx.Request().Header("token", "")
		err := facades.Auth.Parse(ctx, token)
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
		err = facades.Auth.User(ctx, &admin)
		if admin.ID == 0 {
			ctx.Request().AbortWithStatusJson(401, contractshttp.Json{
				"Code":    401,
				"Message": "登录信息已过期，请重新登录",
			})
			return
		}

		permission := strings.ToLower(fmt.Sprintf("%s.%s", ctx.Request().Method(), ctx.Request().Path()))
		if admin.RoleName != "super_admin" {
			var role models.AdminRole
			err := facades.Orm.Query().Where("name", admin.RoleName).Find(&role)
			if err != nil {
				ctx.Request().AbortWithStatusJson(401, contractshttp.Json{
					"Code":    401,
					"Message": err.Error(),
				})
				return
			}
			if role.ID == 0 {
				ctx.Request().AbortWithStatusJson(401, contractshttp.Json{
					"Code":    401,
					"Message": "该用户组不存在或已停用，您没有权限访问",
				})
				return
			}
			if !checkAuth(&role, permission) {
				ctx.Request().AbortWithStatusJson(401, contractshttp.Json{
					"Code":    401,
					"Message": "您没有权限访问",
				})
				return
			}
		}
		ctx.Request().Next()
	}
}

func checkAuth(role *models.AdminRole, permission string) bool {
	pd := role.Permissions
	fmt.Println("asda:", pd)
	_ = pd.Scan(&pd)
	fmt.Println("asd:", pd)
	for _, v := range *pd {
		if strings.HasPrefix(permission, v) {
			return true
		} else {
			reg := regexp.MustCompile(`\{.*\}`)
			if reg.MatchString(v) {
				reg = regexp.MustCompile(`\{.*\}`)
				regStr := reg.FindString(v)
				regStr = strings.Trim(regStr, "{")
				regStr = strings.Trim(regStr, "}")
				regStr = strings.Replace(regStr, "*", ".*", -1)
				regStr = strings.Replace(regStr, "/", "\\/", -1)
				regStr = strings.Replace(regStr, ".", "\\.", -1)
				regStr = strings.Replace(regStr, "?", ".?", -1)
				regStr = fmt.Sprintf("^%s$", regStr)
				reg = regexp.MustCompile(regStr)
				fmt.Println(regStr, permission)
				if reg.MatchString(permission) {
					return true
				}
			}
		}
	}
	return false
}
