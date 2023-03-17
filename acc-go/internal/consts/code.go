package consts

import "github.com/upbos/go-saber/e"

const (
	UserDisagreement      = 1000
	UserDuplicateUsername = 1001
	UserFreeze            = 1002
	UserAuthFailed        = 1003

	AccountNotRecord = 1010

	ErrTemplateNotFound = 2001
)

var ErrUserDisagreement = e.New(UserDisagreement)

var ErrUserDuplicateUsername = e.New(UserDuplicateUsername)

var ErrUserFreeze = e.New(UserFreeze)

var ErrUserAuthFailed = e.New(UserAuthFailed)
