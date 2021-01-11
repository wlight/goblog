package requests

import (
	"github.com/thedevsaddam/govalidator"
	"goblog/app/models/user"
)

// ValidateRegistrationForm 验证表单，返回errs长度等于零即通过
func ValidateRegistrationForm(data user.User) map[string][]string {
	// 1.定制认证规则
	rules := govalidator.MapData{
		"name": []string{
			"required", "alpha_num", "between:3,20",
		},
		"email": []string{
			"required", "min:4", "max:30", "email",
		},
		"password": []string{
			"required", "min:6",
		},
		"password_comfirm": []string{
			"required",
		},
	}
	// 2. 定制错误信息
	message := govalidator.MapData{
		"name": []string{
			"required:用户名为必填项",
			"alpha_num:格式错误，只允许数字和英文",
			"between:用户名长度需在 3——20 之间",
		},
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
		"password": []string{
			"required:密码为必填项",
			"min:长度需大于 6",
		},
		"password_comfirm": []string{
			"required:确认密码为必填项",
		},
	}
	// 3.配置选项
	opts := govalidator.Options{
		Data:          &data,
		Rules:         rules,
		TagIdentifier: "valid", // Struct 标签标识符
		Messages:      message,
	}

	// 4.开始验证
	errs  := govalidator.New(opts).ValidateStruct()

	// 5.因 govalidator 不支持 password_comfirm 验证，我们自己写一个
	if data.Password != data.PasswordComfirm {
		errs["password_comfirm"] = append(errs["password_comfirm"], "两次输入密码不匹配")
	}

	return errs
}
