// Package helper
// @file : request.go
// @project : AGPC(Awesome Garment Production Cloud)
// @author : 周东明（Empty）
// @contact : empty@inzj.cn
// @created at: 2023/3/6 15:59
// ----------------------------------------------------------
package helper

import (
	"encoding/json"
	"github.com/goravel/framework/facades"
	"io"
	"net/http"
	"strings"
)

type ResponseData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type Request struct {
	Url string
}

func NewRequest() *Request {
	return &Request{
		Url: facades.Config.Env("AUTH_SERVICE_HOST", "http://auth-awesome.inzj.cn").(string),
	}
}

// Get 发起GET请求
func (r *Request) Get(url string) (ResponseData, error) {
	var result ResponseData
	url = r.Url + url
	// 将参数拼接到url上
	resp, err := http.Get(url)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	// 解析响应
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Post 发起POST请求
func (r *Request) Post(url string, params map[string]string) (ResponseData, error) {
	var result ResponseData
	url = r.Url + url
	// 将params转换为url.Values
	urlValues := make(map[string][]string)
	for k, v := range params {
		urlValues[k] = []string{v}
	}
	resp, err := http.PostForm(url, urlValues)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	// 解析响应
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Put 发起PUT请求
func (r *Request) Put(url string, params map[string]string) (ResponseData, error) {
	var result ResponseData
	url = r.Url + url
	// 将params转换为url.Values
	urlValues := make(map[string][]string)
	for k, v := range params {
		urlValues[k] = []string{v}
	}
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return result, err
	}
	q := req.URL.Query()
	for k, v := range urlValues {
		q.Add(k, v[0])
	}
	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	// 解析响应
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Delete 发起DELETE请求
func (r *Request) Delete(url string, params map[string]string) (ResponseData, error) {
	var result ResponseData
	url = r.Url + url
	// 将params转换为url.Values
	urlValues := make(map[string][]string)
	for k, v := range params {
		urlValues[k] = []string{v}
	}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return result, err
	}
	q := req.URL.Query()
	for k, v := range urlValues {
		q.Add(k, v[0])
	}
	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	// 解析响应
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// JsonEncode 将map转换为json字符串
func JsonEncode(params map[string]string) io.Reader {
	// 将params转换为url.Values
	urlValues := make(map[string][]string)
	for k, v := range params {
		urlValues[k] = []string{v}
	}
	jsonStr, _ := json.Marshal(urlValues)
	return strings.NewReader(string(jsonStr))
}

// JsonDecode 将json字符串转换为map
func JsonDecode(jsonStr string) map[string]string {
	var result map[string]string
	json.Unmarshal([]byte(jsonStr), &result)
	return result
}
