package manager

import (
	"Awesome/app/http/helper"
	"Awesome/app/http/requests"
	"Awesome/app/models"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/goravel/framework/auth"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"strconv"
	"time"
)

type Common struct {
	//Dependent services
	Model models.Admin
}

func NewCommon() *Common {
	return &Common{
		//Inject services
	}
}

// Login POST /dummy
func (r *Common) Login(ctx http.Context) {
	var params requests.AdminLoginRequest
	valid, err := ctx.Request().ValidateRequest(&params)
	if err != nil {
		helper.RestfulError(ctx, "表单参数错误："+err.Error())
		return
	}
	if valid != nil {
		helper.RestfulError(ctx, "表单参数错误："+helper.GetRequestError(valid.All()))
		return
	}
	// 查询
	var admin models.Admin
	if err := facades.Orm.Query().Where("phone", params.Phone).With("Role").FirstOrFail(&admin); err != nil {
		helper.RestfulError(ctx, err.Error())
		return
	}
	if admin.ID == 0 {
		helper.RestfulError(ctx, "该手机号未注册")
		return
	}
	var errCount int
	if err := facades.Cache.Get("login_error_count_"+params.Phone, &errCount); err != nil {
		errCount = 0
	}
	// 密码错误次数限制
	if errCount >= 5 {
		helper.RestfulError(ctx, "密码错误次数过多，请稍后再试")
		return
	}
	// 密码验证
	if !facades.Hash.Check(params.Password, admin.Password) {
		errCount++
		_ = facades.Cache.Put("login_error_count_"+params.Phone, errCount, 60)
		helper.RestfulError(ctx, "密码错误,您还有"+strconv.Itoa(5-errCount)+"次机会")
		return
	}
	// 登录成功,清除错误次数
	facades.Cache.Forget("login_error_count_" + params.Phone)
	// 生成token
	fmt.Println(admin)
	token, err := facades.Auth.Login(ctx, &admin)
	if err != nil {
		helper.RestfulError(ctx, "Token generate failed, error: "+err.Error())
		return
	}
	//fmt.Println(payload, token)
	helper.RestfulSuccess(ctx, "登录成功", http.Json{
		"token": token,
		"userInfo": http.Json{
			"id":        admin.ID,
			"phone":     admin.Phone,
			"role":      admin.Role,
			"gender":    admin.Gender,
			"avatar":    admin.Avatar,
			"nickname":  admin.Nickname,
			"email":     admin.Email,
			"remark":    admin.Remark,
			"dashboard": "0",
		},
	})

}

// Logout POST /dummy
func (r *Common) Logout(ctx http.Context) {
	if err := facades.Auth.Logout(ctx); err != nil {
		helper.RestfulError(ctx, "退出登录失败")
		return
	}
	helper.RestfulSuccess(ctx, "退出登录成功", nil)
}

// Info GET /dummy
func (r *Common) Info(ctx http.Context) {
	token := ctx.Request().Header("token", "")
	if token == "" {
		helper.RestfulError(ctx, "请先登录")
		return
	}
	// 查询
	payload, err := facades.Auth.Parse(ctx, token)
	if err != nil {
		helper.RestfulError(ctx, "请先登录")
		return
	}
	if errors.Is(err, auth.ErrorTokenExpired) {
		helper.RestfulError(ctx, "登录已过期,请重新登录")
		return
	}
	if payload == nil {
		helper.RestfulError(ctx, "登录已过期,请重新登录")
		return
	}
	var admin models.Admin
	if err := facades.Orm.Query().Where("id", payload.Key).With("Role").FirstOrFail(&admin); err != nil {
		helper.RestfulError(ctx, "请先登录")
		return
	}
	helper.RestfulSuccess(ctx, "登录成功", http.Json{
		"userInfo": http.Json{
			"id":        admin.ID,
			"phone":     admin.Phone,
			"role":      admin.Role,
			"gender":    admin.Gender,
			"avatar":    admin.Avatar,
			"nickname":  admin.Nickname,
			"email":     admin.Email,
			"remark":    admin.Remark,
			"dashboard": "0",
		},
	})
}

// Captcha GET /dummy
func (r *Common) Captcha(ctx http.Context) {
	appSecret, b4s, err := helper.MakeCaptcha()
	if err != nil {
		helper.RestfulError(ctx, "验证码生成失败")
		return
	}
	// 返回图片
	helper.RestfulSuccess(ctx, "获取成功", http.Json{
		"app_secret": appSecret,
		"image":      b4s,
	})
}

// refresh POST /dummy
func (r *Common) Refresh(ctx http.Context) {
	token := ctx.Request().Header("token", "")
	if token == "" {
		helper.RestfulError(ctx, "请先登录")
		return
	}
	// 查询
	payload, err := facades.Auth.Parse(ctx, token)
	if err != nil {
		helper.RestfulError(ctx, "请先登录")
		return
	}
	if errors.Is(err, auth.ErrorTokenExpired) {
		helper.RestfulError(ctx, "登录已过期,请重新登录")
		return
	}
	if payload == nil {
		helper.RestfulError(ctx, "登录已过期,请重新登录")
		return
	}
	token, err = facades.Auth.Refresh(ctx)
	if err != nil {
		helper.RestfulError(ctx, "Token error: "+err.Error())
		return
	}
	helper.RestfulSuccess(ctx, "刷新成功", http.Json{
		"token": token,
	})
}

// Upload POST /dummy
func (r *Common) Upload(ctx http.Context) {
	fileType := ctx.Request().Form("type", "file")
	file, err := ctx.Request().File("file")
	if err != nil {
		helper.RestfulError(ctx, "文件上传失败： "+err.Error())
		return
	}
	if file == nil {
		helper.RestfulError(ctx, "文件上传失败：文件不存在")
		return
	}
	switch fileType {
	case "Image":
		if !helper.IsImage(file.GetClientOriginalName()) {
			helper.RestfulError(ctx, "文件上传失败：文件格式不正确")
			return
		}
		break
	case "Video":
		if !helper.IsVideo(file.GetClientOriginalName()) {
			helper.RestfulError(ctx, "文件上传失败：文件格式不正确")
			return
		}
		break
	case "Doc":
		if !helper.IsDoc(file.GetClientOriginalName()) {
			helper.RestfulError(ctx, "文件上传失败：文件格式不正确")
			return
		}
		break
	case "Audio":
		if !helper.IsAudio(file.GetClientOriginalName()) {
			helper.RestfulError(ctx, "文件上传失败：文件格式不正确")
			return
		}
		break
	default:
		helper.RestfulError(ctx, "文件上传失败：文件类型不正确")
		return
	}
	// 上传到临时文件目录
	path, err := file.Store("tmp")
	if err != nil {
		helper.RestfulError(ctx, "文件上传失败： "+err.Error())
		return
	}
	url, err := facades.Storage.TemporaryUrl(path, time.Now().Add(120*time.Minute))
	if err != nil {
		helper.RestfulError(ctx, "文件上传失败： "+err.Error())
		return
	}
	helper.RestfulSuccess(ctx, "文件上传成功", http.Json{
		"url": url,
	})
}

// Options GET /dummy
func (r *Common) Options(ctx http.Context) {
	dashboardGrid := []string{"welcome", "ver", "time", "progress", "echarts", "about"}
	menu, err := (new(models.AdminPermission)).GetPermissionList(0)
	if err != nil {
		helper.RestfulError(ctx, "获取失败")
		return
	}
	admin := ctx.Value("CurrentAdmin").(*models.Admin)
	var per []models.AdminPermission
	facades.Orm.Query().Model(&models.AdminPermission{}).Find(&per)
	var permissions []string
	if admin.Role.Name == "super_admin" {
		if len(per) > 0 {
			for _, v := range per {
				var meta *models.Meta
				_ = json.Unmarshal([]byte(v.MetaJSON), &meta)
				fmt.Println(meta.Type)
				if meta.Type == "button" {
					permissions = append(permissions, v.Name)
				}
			}
		}
	} else {
		var role models.AdminRole
		facades.Orm.Query().Model(&models.AdminRole{}).Where("id", admin.RoleID).First(&role)
		if role.ID == 0 {
			helper.RestfulError(ctx, "获取失败")
			return
		}
		_ = role.Scan()
		fmt.Println(role.Permissions)
		if len(role.Permissions) > 0 {
			for _, v := range per {
				var meta *models.Meta
				meta = &models.Meta{}
				_ = json.Unmarshal([]byte(v.MetaJSON), &meta)
				if meta.Type == "button" {
					for _, p := range role.Permissions {
						if p == v.Name {
							permissions = append(permissions, v.Name)
						}
					}
				}
			}
		}
	}
	helper.RestfulSuccess(ctx, "获取成功", http.Json{
		"dashboardGrid": dashboardGrid,
		"menu":          menu,
		"permissions":   permissions,
	})
}

// EditPassword POST /dummy
func (r *Common) EditPassword(ctx http.Context) {
	var params requests.AdminEditPasswordRequest
	valid, err := ctx.Request().ValidateRequest(&params)
	if err != nil {
		helper.RestfulError(ctx, "表单参数错误： "+err.Error())
		return
	}
	if valid != nil {
		helper.RestfulError(ctx, "表单参数错误： "+helper.GetRequestError(valid.All()))
		return
	}
	admin := ctx.Value("CurrentAdmin").(*models.Admin)
	if !facades.Hash.Check(admin.Password, params.OldPassword) {
		helper.RestfulError(ctx, "旧密码错误")
		return
	}
	pwd, _ := facades.Hash.Make(params.NewPassword)
	if err := facades.Orm.Query().Model(&models.Admin{}).Where("id = ?", admin.ID).Update("password", pwd); err != nil {
		helper.RestfulError(ctx, "修改失败")
		return
	}
	helper.RestfulSuccess(ctx, "修改成功", nil)
}
