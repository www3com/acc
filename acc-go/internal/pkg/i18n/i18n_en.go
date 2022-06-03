package i18n

import "acc/internal/pkg/e"

var en = map[int]string{
	e.OK:             "",
	e.ERROR:          "Internal server error",
	e.INVALID_PARAMS: "Request parameter error",
	e.UNAUTHORIZED:   "Access token has timed out",
	e.FORBIDDEN:      "No access rights",
	e.NOT_FOUND:      "The accessed resource does not exist",

	e.USER_DISAGREEMENT: "Must agree to the agreement",
	e.USER_NO_USERNAME:  "No such account name",
	e.USER_FREEZE:       "Account is frozen",
	e.USER_AUTH_ERROR:   "Wrong account name or password",
}
