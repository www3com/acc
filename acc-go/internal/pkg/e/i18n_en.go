package e

var en = map[int]string{
	OK:             "",
	ERROR:          "Internal server error",
	INVALID_PARAMS: "Request parameter error",
	UNAUTHORIZED:   "Access token has timed out",
	FORBIDDEN:      "No access rights",
	NOT_FOUND:      "The accessed resource does not exist",

	USER_DISAGREEMENT: "Must agree to the agreement",
	USER_NO_USERNAME:  "No such account name",
	USER_FREEZE:       "Account is frozen",
	USER_AUTH_ERROR:   "Wrong account name or password",
}
