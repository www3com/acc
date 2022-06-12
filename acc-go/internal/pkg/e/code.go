package e

const (

	// OK : The server successfully accepts the request from the client
	OK = 200

	// ERROR : The server has an unknown e
	ERROR = 500

	// InvalidParams : The client's request has illegal parameters
	InvalidParams = 400

	// Unauthorized 请求要求身份验证
	Unauthorized = 401

	// Forbidden The client does not have access rights to the content
	Forbidden = 403

	// NotFound : The server can not find the requested resource
	NotFound = 404

	UserDisagreement      = 1000
	UserDuplicateUsername = 1001
	UserFreeze            = 1002
	UserAuthFailed        = 1003
)

var Error = New(ERROR)

var InvalidParamsError = New(InvalidParams)

var UnauthorizedError = New(Unauthorized)

var ForbiddenError = New(Forbidden)

var UserDisagreementError = New(UserDisagreement)

var UserDuplicateUsernameError = New(UserDuplicateUsername)

var UserFreezeError = New(UserFreeze)

var UserAuthFailedError = New(UserAuthFailed)
