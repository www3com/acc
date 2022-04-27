package ret

type RetData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *RetData) Ok() bool {
	return r.Code == OK
}

func Ok(data interface{}) *RetData {
	return &RetData{Code: OK, Data: data}
}

func Fail(code int) *RetData {
	return &RetData{Code: code}
}
