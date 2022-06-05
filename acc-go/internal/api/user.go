package api

import (
	"acc/internal/consts"
	"acc/internal/model"
	"acc/internal/pkg/context"
	"acc/internal/pkg/e"
	"acc/internal/pkg/r"
	"acc/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
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
	if err == nil {
		r.RenderOk(c, token)
	} else {
		r.RenderFail(c, err)
	}
}

func SignUp(c *gin.Context) {
	var userParam user

	if err := r.BindAndValid(c, &userParam); err != nil {
		r.RenderFail(c, err)
		return
	}

	if !userParam.Agreement {
		r.RenderFail(c, e.New(e.USER_DISAGREEMENT))
		return
	}

	userService := service.UserService{Username: userParam.Username}

	exist, err := userService.ExistUsername()
	if err != nil {
		logrus.Errorf("Check if username exist, cause by: %v ", err)
		r.Render(c, http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	if exist {
		r.RenderFail(c, e.New(e.USER_NO_USERNAME))
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
	ownerId := context.GetUserId(c)
	id, err := userService.SignUp(ownerId, &user)
	if err != nil {
		logrus.Errorf("create user error, cause by: %v ", err)
		r.Render(c, http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	r.RenderOk(c, map[string]int64{"jwt": id})
}

func ExistUsername(c *gin.Context) {
	userService := service.UserService{Username: c.Query("username")}
	exist, err := userService.ExistUsername()
	if err != nil {
		logrus.Errorf("查询用户名是否重复，错误原因: %v", err)
		r.Render(c, http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	r.RenderOk(c, map[string]bool{"exist": exist})
}
