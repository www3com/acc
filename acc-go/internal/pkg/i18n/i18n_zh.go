package i18n

import "acc/internal/pkg/e"

var zh = map[int]string{
	e.OK:            "",
	e.ERROR:         "服务器内部错误",
	e.InvalidParams: "请求参数错误",
	e.Unauthorized:  "访问令牌已超时",
	e.Forbidden:     "无权限访问",
	e.NotFound:      "访问的资源不存在",

	e.UserDisagreement:      "必须同意协议",
	e.UserDuplicateUsername: "用户名重复",
	e.UserFreeze:            "账号被冻结",
	e.UserAuthFailed:        "账号名或密码错误",
}
