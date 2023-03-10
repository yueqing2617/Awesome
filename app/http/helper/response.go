// Package helper
// @file : response.go
// @project : AGPC(Awesome Garment Production Cloud)
// @author : 周东明（Empty）
// @contact : empty@inzj.cn
// @created at: 2023/3/8 11:17
// ----------------------------------------------------------
package helper

import "github.com/goravel/framework/contracts/http"

// RestfulSuccess restful_success is the response of success
func RestfulSuccess(ctx http.Context, msg string, data any) {
	ctx.Response().Json(200, http.Json{
		"Code":    200,
		"Message": msg,
		"Data":    data,
	})
	return
}

// RestfulError restful_error is the response of error
func RestfulError(ctx http.Context, msg string) {
	ctx.Response().Json(200, http.Json{
		"Code":    0,
		"Message": msg,
	})
	return
}

// UnAuthorized is the response of unauthorized
func UnAuthorized(ctx http.Context, msg string) {
	ctx.Response().Json(401, http.Json{
		"Code":    401,
		"Message": msg,
	})
	return
}
