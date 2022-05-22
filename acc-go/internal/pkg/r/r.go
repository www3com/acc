package r

import "accounting-service/internal/misc/consts"

type R[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func (r *R[T]) Ok() bool {
	return r.Code == consts.OK
}

func (r *R[T]) Fail() bool {
	return r.Code != consts.OK
}

func Ok[T any](data T) *R[T] {
	return &R[T]{Code: consts.OK, Data: data}
}

func Fail[T any](code int) *R[T] {
	return &R[T]{Code: code}
}
