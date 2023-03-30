// Package helper
// @file : response.go
// @project : AGPC(Awesome Garment Production Cloud)
// @author : 周东明（Empty）
// @contact : empty@inzj.cn
// @created at: 2023/3/8 11:17
// ----------------------------------------------------------
package helper

import (
	"github.com/goravel/framework/contracts/http"
	"strings"
)

// RestfulSuccess restful_success is the response of success
func RestfulSuccess(ctx http.Context, msg string, data any) {
	ctx.Response().Json(200, http.Json{
		"code":    200,
		"message": msg,
		"data":    data,
	})
	return
}

// RestfulError restful_error is the response of error
func RestfulError(ctx http.Context, msg string) {
	ctx.Response().Json(200, http.Json{
		"code":    0,
		"message": msg,
	})
	return
}

// UnAuthorized is the response of unauthorized
func UnAuthorized(ctx http.Context, msg string) {
	ctx.Response().Json(401, http.Json{
		"code":    401,
		"message": msg,
	})
	return
}

// GetRequestError 获取请求错误信息
func GetRequestError(errs map[string]map[string]string) string {
	for _, v := range errs {
		for _, v2 := range v {
			return v2
		}
	}
	return ""
}

// 用,分割字符串为数组
func StringToArray(str string) []string {
	return strings.Split(str, ",")
}
