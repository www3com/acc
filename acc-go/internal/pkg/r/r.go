package r

import (
	"acc/internal/pkg/e"
)

type R struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *R) Ok() bool {
	return r.Code == e.OK
}

func (r *R) Fail() bool {
	return r.Code != e.OK
}
