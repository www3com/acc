package ret

const (

	// OK : The server successfully accepts the request from the client
	OK = 200

	// INVALID_PARAMS : The client's request has illegal parameters
	INVALID_PARAMS = 400

	// FORBIDDEN : The client does not have access rights to the content
	FORBIDDEN = 403

	// NOT_FOUND : The server can not find the requested resource
	NOT_FOUND = 404

	// TOO_MANY_REQUESTS : flow control limited.
	TOO_MANY_REQUESTS = 429

	// INTERNAL_ERROR : The server has an unknown e
	INTERNAL_ERROR = 500
)
