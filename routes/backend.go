package routes

import (
	"Awesome/app/http/helper"
	"encoding/json"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"io"
	curl "net/http"
)

func Backend() {
	facades.Route.Get("/", func(ctx http.Context) {
		url := facades.Config.Env("AUTH_SERVICE_HOST", "http://auth-awesome.inzj.cn").(string) + "/auth/heartbeat"
		// 发起请求
		form := make(map[string]string)
		form["code"] = "1000012"
		form["secret"] = "AGPC-035cc0bc40be92ab802e927b951173c0"
		resp, err := curl.Post(url, "application/json", helper.JsonEncode(form))
		if err != nil {
			helper.RestfulError(ctx, "请求失败："+err.Error())
			return
		}
		defer resp.Body.Close()
		// 解析响应
		var result helper.ResponseData
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			helper.RestfulError(ctx, "请求失败："+err.Error())
			return
		}
		err = json.Unmarshal(body, &result)
		if err != nil {
			helper.RestfulError(ctx, "请求失败："+err.Error())
			return
		}
		if result.Code != 200 {
			helper.RestfulError(ctx, "请求失败："+result.Message)
			return
		}
		helper.RestfulSuccess(ctx, "请求成功", result)
	})
	facades.Route.Static("static", "./storage")
}
