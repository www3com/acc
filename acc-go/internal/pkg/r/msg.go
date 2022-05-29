package r

import (
	"acc/internal/pkg/e"
)

type Msg struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *Msg) Ok() bool {
	return r.Code == e.OK
}

func (r *Msg) Fail() bool {
	return r.Code != e.OK
}
