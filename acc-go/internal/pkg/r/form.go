package r

import (
	"acc/internal/pkg/e"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) error {
	err := c.Bind(form)
	if err != nil {
		return e.New(e.InvalidParams)
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return e.Wrap(err, "valid params")
	}
	if !check {
		markErrors(valid.Errors)
		return e.New(e.InvalidParams)
	}

	return nil
}

func markErrors(errors []*validation.Error) {
	for _, err := range errors {
		logrus.Info(err.Message)
	}

	return
}
