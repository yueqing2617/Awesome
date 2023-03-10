// Package helper
// @file : captcha.go
// @project : AGPC(Awesome Garment Production Cloud)
// @author : 周东明（Empty）
// @contact : empty@inzj.cn
// @created at: 2023/3/6 14:43
// ----------------------------------------------------------
package helper

import (
	"errors"
	"fmt"
	"github.com/goravel/framework/facades"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"strconv"
	"time"
)

var result = base64Captcha.NewMemoryStore(20480, 5*time.Minute)
var c_w, _ = strconv.Atoi(facades.Config.Env("CAPTCHA_WIDTH", 240).(string))
var c_h, _ = strconv.Atoi(facades.Config.Env("CAPTCHA_HEIGHT", 80).(string))
var c_n_c, _ = strconv.Atoi(facades.Config.Env("CAPTCHA_NOISE_COUNT", 4).(string))
var c_l, _ = strconv.Atoi(facades.Config.Env("CAPTCHA_LENGTH", 4).(string))
var c_s = facades.Config.Env("CAPTCHA_SOURCE", "1234567890qwertyuiopasdfghjklzxcvbnm").(string)
var c_s_z = facades.Config.Env("CAPTCHA_SOURCE_ZH", "富强,民主,文明,和谐,自由,平等,公正,法治,爱国,敬业,诚信,友善").(string)

// mathConfig 数字运算类型
func mathConfig() *base64Captcha.DriverMath {
	mathType := &base64Captcha.DriverMath{
		Height:          c_h,
		Width:           c_w,
		NoiseCount:      c_n_c,
		ShowLineOptions: base64Captcha.OptionShowHollowLine,
		BgColor: &color.RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 85,
		},
	}
	return mathType
}

// digitConfig 数字类型
func digitConfig() *base64Captcha.DriverDigit {
	digitType := &base64Captcha.DriverDigit{
		Height:   c_h,
		Width:    c_w,
		Length:   c_l,
		MaxSkew:  0.45,
		DotCount: 80,
	}
	return digitType
}

// stringConfig 字符类型
func stringConfig() *base64Captcha.DriverString {
	stringType := &base64Captcha.DriverString{
		Height:          c_h,
		Width:           c_w,
		NoiseCount:      c_n_c,
		ShowLineOptions: base64Captcha.OptionShowHollowLine | base64Captcha.OptionShowSlimeLine | base64Captcha.OptionShowSineLine,
		Length:          c_l,
		Source:          c_s,
		BgColor: &color.RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 85,
		},
	}
	return stringType
}

// chineseConfig 中文类型
func chineseConfig() *base64Captcha.DriverChinese {
	chineseType := &base64Captcha.DriverChinese{
		Height:          c_h,
		Width:           c_w,
		NoiseCount:      c_n_c,
		ShowLineOptions: base64Captcha.OptionShowHollowLine | base64Captcha.OptionShowSlimeLine | base64Captcha.OptionShowSineLine,
		Length:          c_l,
		Source:          c_s_z,
		BgColor: &color.RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 85,
		},
		Fonts: nil,
	}
	return chineseType
}

// audioConfig 音频类型
func audioConfig() *base64Captcha.DriverAudio {
	audioType := &base64Captcha.DriverAudio{
		Length:   c_l,
		Language: "zh",
	}
	return audioType
}

// MakeCaptcha 生成验证码
// opt: 1:算数 2:数字 3:字符串 4:中文 5：音频
func MakeCaptcha() (id string, b64s string, err error) {
	var driver base64Captcha.Driver
	opt, _ := strconv.Atoi(facades.Config.Env("CAPTCHA_MODE", 2).(string))
	fmt.Println(opt)
	switch opt {
	case 1:
		driver = mathConfig()
	case 2:
		driver = digitConfig()
	case 3:
		driver = stringConfig()
	case 4:
		driver = chineseConfig()
	case 5:
		driver = audioConfig()
	default:
		driver = digitConfig()
	}
	if driver == nil {
		return "", "", errors.New("captcha driver is nil")
	}
	c := base64Captcha.NewCaptcha(driver, result)
	id, b64s, err = c.Generate()
	return id, b64s, err
}

// VerifyCaptcha 验证验证码
func VerifyCaptcha(id, verifyValue string) bool {
	return result.Verify(id, verifyValue, true)
}
