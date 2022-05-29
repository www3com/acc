package e

var zh = map[int]string{
	OK:             "",
	ERROR:          "服务器内部错误",
	INVALID_PARAMS: "请求参数错误",
	UNAUTHORIZED:   "Token已超时",
	FORBIDDEN:      "无权限访问",
	NOT_FOUND:      "访问的资源不存在",

	USER_DISAGREEMENT: "必须同意协议",
	USER_NO_USERNAME:  "没有此账号名",
	USER_FREEZE:       "账号被冻结",
	USER_AUTH_ERROR:   "账号名或密码错误",
}
