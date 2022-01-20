package validate

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// 自定义手机号码验证器
func ValidatorMobile(fl validator.FieldLevel) bool {
	// 获取参数数据
	mobile := fl.Field().String()
	// 使用正则表达式判断数据是否合法
	ok, _ := regexp.MatchString(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`, mobile)
	// 返回验证结果
	return ok
}
