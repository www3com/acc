package api

import (
	"acc/internal/consts"
	"acc/internal/model"
	"acc/internal/pkg/context"
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
	var signInParam signIn
	if err := r.BindAndValid(c, &signInParam); err != nil {
		r.RenderFail(c, err)
		return
	}

	userService := service.UserService{
		Username: signInParam.Username,
		Password: signInParam.Password,
		IP:       c.ClientIP()}

	token, err := userService.SignIn()
	r.Render(c, token, err)
}

func SignUp(c *gin.Context) {
	var userParam user

	if err := r.BindAndValid(c, &userParam); err != nil {
		r.RenderFail(c, err)
		return
	}

	if !userParam.Agreement {
		r.RenderFail(c, e.New(e.UserDisagreement))
		return
	}

	userService := service.UserService{Username: userParam.Username}

	if err := userService.ExistUsername(); err != nil {
		r.RenderFail(c, err)
		return
	}

	user := model.User{
		Nickname: userParam.Nickname,
		Username: userParam.Username,
		Password: userParam.Password,
		State:    consts.USER_STATE_NORMAL,
	}
	if userParam.Agreement {
		user.Agreement = 1
	}
	id, err := userService.SignUp(context.GetUserId(c), &user)
	r.Render(c, map[string]int64{"jwt": id}, err)
}

func ExistUsername(c *gin.Context) {
	userService := service.UserService{Username: c.Query("username")}
	err := userService.ExistUsername()
	exist := false
	if err == nil {
		exist = false
	} else {
		if err == e.UserNoUsernameError {
			exist = true
		}
	}

	r.Render(c, map[string]bool{"exist": exist}, err)
}
