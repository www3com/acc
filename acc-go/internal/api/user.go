package api

import (
	"acc/internal/consts"
	"acc/internal/model"
	"acc/internal/pkg/e"
	"acc/internal/pkg/r"
	"acc/internal/service"
	"github.com/gin-gonic/gin"
)

type user struct {
	Nickname  string `form:"nickname" valid:"Required"`
	Username  string `form:"username" valid:"Required"`
	Password  string `form:"password" valid:"Required"`
	Agreement bool   `form:"agreement" valid:"Required"`
}

type signIn struct {
	Username string `form:"username" valid:"Required"`
	Password string `form:"password" valid:"Required"`
}

func SignIn(c *gin.Context) {
	var param signIn
	if err := r.BindAndValid(c, &param); err != nil {
		r.RenderFail(c, err)
		return
	}

	userService := service.UserService{
		Username: param.Username,
		Password: param.Password,
		IP:       c.ClientIP()}

	token, err := userService.SignIn()
	r.Render(c, token, err)
}

func SignUp(c *gin.Context) {
	var param user

	if err := r.BindAndValid(c, &param); err != nil {
		r.RenderFail(c, err)
		return
	}

	if !param.Agreement {
		r.RenderFail(c, e.New(e.UserDisagreement))
		return
	}

	userService := service.UserService{Username: param.Username}

	exist, err := userService.ExistUsername()

	if err != nil {
		r.RenderFail(c, err)
		return
	}

	if exist {
		r.RenderFail(c, e.UserDuplicateUsernameError)
		return
	}

	user := model.User{
		Nickname: param.Nickname,
		Username: param.Username,
		Password: param.Password,
		State:    consts.UserStateNormal,
	}
	if param.Agreement {
		user.Agreement = 1
	}
	err = userService.SignUp(&user)
	r.Render(c, nil, err)
}

func ExistUsername(c *gin.Context) {
	userService := service.UserService{Username: c.Query("username")}
	exist, err := userService.ExistUsername()
	r.Render(c, map[string]bool{"exist": exist}, err)
}
