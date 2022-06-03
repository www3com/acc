package i18n

import "acc/internal/pkg/e"

var zh = map[int]string{
	e.OK:             "",
	e.ERROR:          "服务器内部错误",
	e.INVALID_PARAMS: "请求参数错误",
	e.UNAUTHORIZED:   "访问令牌已超时",
	e.FORBIDDEN:      "无权限访问",
	e.NOT_FOUND:      "访问的资源不存在",

	e.USER_DISAGREEMENT: "必须同意协议",
	e.USER_NO_USERNAME:  "没有此账号名",
	e.USER_FREEZE:       "账号被冻结",
	e.USER_AUTH_ERROR:   "账号名或密码错误",
}
