package api

import (
	"acc/internal/pkg/r"
	"fmt"
	"github.com/gin-gonic/gin"
)

type user struct {
	Nickname  string `form:"nickname" binding:"required"`
	Username  string `form:"username" binding:"required"`
	Password  string `form:"password" binding:"required"`
	Agreement bool   `form:"agreement" binding:"required"`
}

type UserApi struct {
}

func NewUserApi() *UserApi {
	return &UserApi{}
}

// SignIn 登录
func (api *UserApi) SignIn(c *gin.Context) {
	var p user
	if err := r.BindAndValid(c, &p); err != nil {
		r.RenderFail(c, err)
		return
	}
	fmt.Printf("%v\n", p)
}
