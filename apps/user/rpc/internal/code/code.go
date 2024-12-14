package code

import "wetalk/pkg/xcode"

var (
	PhoneIsRegister     = xcode.New(100001, "手机号已注册")
	UserNotFound        = xcode.New(100002, "用户不存在")
	PhoneNotRegister    = xcode.New(100003, "手机号未注册")
	UserPwdError        = xcode.New(100004, "密码错误")
	RegisterPasswdEmpty = xcode.New(100005, "密码不能为空")
)
