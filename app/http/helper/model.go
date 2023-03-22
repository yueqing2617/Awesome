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
	"encoding/json"
	"math/rand"
	"reflect"
	"time"
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

// PermissionMetaToJSON 权限元数据转换为JSON
func PermissionMetaToJSON(meta *models.Meta) string {
	if meta != nil {
		b, err := json.Marshal(meta)
		if err != nil {
			return err.Error()
		}
		return string(b)
	}
	return ""
}

// PermissionMetaToStruct 权限元数据转换为结构体
func PermissionMetaToStruct(meta string) *models.Meta {
	var result models.Meta
	if meta != "" {
		_ = json.Unmarshal([]byte(meta), &result)
	}
	return &result
}

// ToJSON 将数组转换为JSON
func ToJSON(val interface{}) string {
	b, err := json.Marshal(val)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

// ToArray 将JSON转换为数组
func ToArray(val string) interface{} {
	var result interface{}
	_ = json.Unmarshal([]byte(val), &result)
	return result
}

// ArrayToString 将数组转换为字符串,用,分隔
func ArrayToString(array []string) string {
	var result string
	arrayValue := reflect.ValueOf(array)
	for i := 0; i < arrayValue.Len(); i++ {
		result += arrayValue.Index(i).String() + ","
	}
	return result
}

// RandomString 生成随机字符串
func RandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return "Agpc_" + string(result)
}
