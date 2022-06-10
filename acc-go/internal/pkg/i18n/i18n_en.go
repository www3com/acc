package i18n

import "acc/internal/pkg/e"

var en = map[int]string{
	e.OK:            "",
	e.ERROR:         "Internal server error",
	e.InvalidParams: "Request parameter error",
	e.Unauthorized:  "Access token has timed out",
	e.Forbidden:     "No access rights",
	e.NotFound:      "The accessed resource does not exist",

	e.UserDisagreement: "Must agree to the agreement",
	e.UserNoUsername:   "No such account name",
	e.UserFreeze:       "Account is frozen",
	e.UserAuthFailed:   "Wrong account name or password",
}
