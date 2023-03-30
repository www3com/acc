package api

import (
	"acc/internal/consts"
	user2 "acc/internal/consts/userstate"
	"acc/internal/dao"
	"acc/internal/pkg/r"
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
	token, err := userService.SignIn(c.Query("username"), c.Query("password"), c.ClientIP())
	r.Render(c, token, err)
}

// SignUp 注册
func SignUp(c *gin.Context) {
	var p user
	if err := r.BindAndValid(c, &p); err != nil {
		r.RenderFail(c, err)
		return
	}

	if !p.Agreement {
		r.RenderFail(c, consts.ErrUserDisagreement)
		return
	}

	ok, err := userService.Exist(p.Username)
	if err != nil {
		r.RenderFail(c, err)
		return
	}
	if ok {
		r.RenderFail(c, consts.ErrUserDuplicateUsername)
		return
	}

	user := dao.User{
		Nickname:  p.Nickname,
		Username:  p.Username,
		Password:  p.Password,
		State:     user2.UserStateNormal,
		Agreement: 1,
	}

	err = userService.SignUp(&user)
	r.Render(c, nil, err)
}
