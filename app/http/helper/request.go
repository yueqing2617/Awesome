// Package helper
// @file : request.go
// @project : AGPC(Awesome Garment Production Cloud)
// @author : 周东明（Empty）
// @contact : empty@inzj.cn
// @created at: 2023/3/6 15:59
// ----------------------------------------------------------
package helper

// GetRequestError 获取请求错误信息
func GetRequestError(errs map[string]map[string]string) string {
	for _, v := range errs {
		for _, v2 := range v {
			return v2
		}
	}
	return ""
}
