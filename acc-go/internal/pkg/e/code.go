package e

const (

	// OK : The server successfully accepts the request from the client
	OK = 200

	// ERROR : The server has an unknown e
	ERROR = 500

	// InvalidParams : The client's request has illegal parameters
	InvalidParams = 400

	// UNAUTHORIZED 请求要求身份验证
	UNAUTHORIZED = 401

	// FORBIDDEN : The client does not have access rights to the content
	FORBIDDEN = 403

	// NotFound : The server can not find the requested resource
	NotFound = 404

	UserDisagreement = 1000
	UserNoUsername   = 1001
	UserFreeze       = 1002
	UserAuthFailed   = 1003
)

func IsOk(code int) bool {
	return code == OK
}

func IsFail(code int) bool {
	return code != OK
}

var Error = New(ERROR)

var UserDisagreementError = New(UserDisagreement)

var UserNoUsernameError = New(UserNoUsername)

var UserFreezeError = New(UserFreeze)

var UserAuthFailedError = New(UserAuthFailed)
