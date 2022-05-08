package ret

const (

	// OK : The server successfully accepts the request from the client
	OK = 200

	// INVALID_PARAMS : The client's request has illegal parameters
	INVALID_PARAMS = 400

	// UNAUTHORIZED 请求要求身份验证
	UNAUTHORIZED = 401

	// FORBIDDEN : The client does not have access rights to the content
	FORBIDDEN = 403

	// NOT_FOUND : The server can not find the requested resource
	NOT_FOUND = 404

	// INTERNAL_ERROR : The server has an unknown e
	INTERNAL_ERROR = 500

	USER_DISAGREEMENT = 1000

	USER_NO_USERNAME = 1001

	USER_NO_EMAIL = 1002

	USER_SIGN_IN_ERROR = 1003
)
