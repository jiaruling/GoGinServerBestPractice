package validate

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// 自定义邮箱验证器
func ValidatorEmail(fl validator.FieldLevel) bool {
	// 获取参数数据
	email := fl.Field().String()
	// 使用正则表达式判断数据是否合法
	ok, _ := regexp.MatchString(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`, email)
	// 返回验证结果
	return ok
}
