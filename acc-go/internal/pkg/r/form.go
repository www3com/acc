package r

import (
	"acc/internal/pkg/e"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
		logrus.Info(err.Message)
	}

	return
}
