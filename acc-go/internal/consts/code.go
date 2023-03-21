package consts

import "github.com/upbos/go-saber/e"

var (
	ErrUserDisagreement      = e.New(1000)
	ErrUserDuplicateUsername = e.New(1001)
	ErrUserFreeze            = e.New(1002)
	ErrUserAuthFailed        = e.New(1003)
	ErrAccountNotRecord      = e.New(1010)
	ErrTemplateNotFound      = e.New(2001)
	ErrExistSubAccount       = e.New(2002)
	ErrDeleteTopAccount      = e.New(2003)
)
