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

// SignIn 登录
func SignIn(c *gin.Context) {
	var p user
	if err := r.BindAndValid(c, &p); err != nil {
		r.RenderFail(c, err)
		return
	}
	fmt.Printf("%v\n", p)
	//str, err := service.UserSrv.SignIn("", "", "")
	//if err != nil {
	//	return
	//}
}
