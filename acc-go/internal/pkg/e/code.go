package e

const (

	// OK : The server successfully accepts the request from the client
	OK = 200

	// ERROR : The server has an unknown e
	ERROR = 500

	// INVALID_PARAMS : The client's request has illegal parameters
	INVALID_PARAMS = 400

	// UNAUTHORIZED 请求要求身份验证
	UNAUTHORIZED = 401

	// FORBIDDEN : The client does not have access rights to the content
	FORBIDDEN = 403

	// NOT_FOUND : The server can not find the requested resource
	NOT_FOUND = 404

	USER_DISAGREEMENT = 1000
	USER_NO_USERNAME  = 1001
	USER_FREEZE       = 1002
	USER_AUTH_ERROR   = 1003
)

func IsOk(code int) bool {
	return code == OK
}

func IsFail(code int) bool {
	return code != OK
}
