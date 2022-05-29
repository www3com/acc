package r

import (
	"acc/internal/pkg/e"
	"acc/internal/pkg/logger"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) int {
	err := c.Bind(form)
	if err != nil {
		return e.INVALID_PARAMS
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return e.ERROR
	}
	if !check {
		markErrors(valid.Errors)
		return e.INVALID_PARAMS
	}

	return e.OK
}

func markErrors(errors []*validation.Error) {
	for _, err := range errors {
		logger.Info(err.Key, err.Message)
	}

	return
}
