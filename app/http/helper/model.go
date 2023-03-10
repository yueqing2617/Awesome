// Package helper
// @file : model.go
// @project : AGPC(Awesome Garment Production Cloud)
// @author : 周东明（Empty）
// @contact : empty@inzj.cn
// @created at: 2023/3/6 15:49
// ----------------------------------------------------------
package helper

import (
	"Awesome/app/models"
	"crypto/md5"
	"encoding/hex"
	"reflect"
)

// PasswordEncrypt 密码加密
func PasswordEncrypt(password string) string {
	h := md5.New()
	salt := "AGPC"
	h.Write([]byte(password + salt))
	return hex.EncodeToString(h.Sum(nil))
}

// RequestToModel 将请求参数赋值给模型
func RequestToModel(request interface{}, model interface{}) {
	// TODO: 赋值
	requestValue := reflect.ValueOf(request)
	modelValue := reflect.ValueOf(model)
	modelType := reflect.TypeOf(model)
	for i := 0; i < modelValue.Elem().NumField(); i++ {
		fieldName := modelType.Elem().Field(i).Name
		requestFieldValue := requestValue.Elem().FieldByName(fieldName)
		if requestFieldValue.IsValid() {
			modelValue.Elem().FieldByName(fieldName).Set(requestFieldValue)
		}
	}
}

// ToStringArray 将数组转换为字符串数组
func ToStringArray(array interface{}) models.StringArray {
	// TODO: 转换
	var result models.StringArray
	arrayValue := reflect.ValueOf(array)
	for i := 0; i < arrayValue.Len(); i++ {
		result = append(result, arrayValue.Index(i).String())
	}
	return result
}
