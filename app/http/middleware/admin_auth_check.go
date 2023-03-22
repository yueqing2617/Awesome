package middleware

import (
	"Awesome/app/models"
	"fmt"
	contractshttp "github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"regexp"
	"strings"
	"time"
)

func AdminAuthCheck() contractshttp.Middleware {
	return func(ctx contractshttp.Context) {
		admin := ctx.Value("CurrentAdmin").(*models.Admin)

		permission := strings.ToLower(fmt.Sprintf("%s.%s", ctx.Request().Method(), ctx.Request().Path()))
		if admin.Role.Name != "super_admin" {
			var role models.AdminRole
			err := facades.Orm.Query().Where("name", admin.Role.Name).Find(&role)
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

//func checkAuth(role *models.AdminRole, permission string) bool {
//	_ = role.Scan()
//	var policies []models.AdminPermissionPolicy
//	_ = facades.Orm.Query().Where("name", "in", role.Permissions).Find(&policies)
//	for _, v := range policies {
//		if strings.HasPrefix(permission, v.Path) {
//			return true
//		} else {
//			reg := regexp.MustCompile(`\{.*\}`)
//			if reg.MatchString(v.Path) {
//				reg = regexp.MustCompile(`\{.*\}`)
//				regStr := reg.FindString(v.Path)
//				regStr = strings.Trim(regStr, "{")
//				regStr = strings.Trim(regStr, "}")
//				regStr = strings.Replace(regStr, "*", ".*", -1)
//				regStr = strings.Replace(regStr, "/", "\\/", -1)
//				regStr = strings.Replace(regStr, ".", "\\.", -1)
//				regStr = strings.Replace(regStr, "?", ".?", -1)
//				regStr = fmt.Sprintf("^%s$", regStr)
//				reg = regexp.MustCompile(regStr)
//				if reg.MatchString(permission) {
//					return true
//				}
//			}
//		}
//	}
//	return false
//}

func checkAuth(role *models.AdminRole, permission string) bool {
	_ = role.Scan()
	key := fmt.Sprintf("admin_permission_%s", role.Name)
	// 判断是否存在缓存
	policies, err := facades.Cache.Remember(key, 5*time.Hour, func() interface{} {
		var policies []models.AdminPermissionPolicy
		_ = facades.Orm.Query().Where("name", "in", role.Permissions).Find(&policies)
		return policies
	})
	if err != nil {
		return false
	}
	for _, v := range policies.([]models.AdminPermissionPolicy) {
		if strings.HasPrefix(permission, v.Path) {
			return true
		} else if strings.Contains(v.Path, "{") {
			regexPath := regexp.QuoteMeta(v.Path)
			regexPath = strings.ReplaceAll(regexPath, "\\{([^/]+)\\}", "(?P<$1>[^/]+)")
			regexPath = "^" + regexPath + "$"
			reg := regexp.MustCompile(regexPath)
			if reg.MatchString(permission) {
				return true
			}
		}
	}

	return false
}
