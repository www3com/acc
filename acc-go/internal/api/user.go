package api

import (
	"accounting-service/internal/misc/consts"
	"accounting-service/internal/model"
	"accounting-service/internal/pkg/logger"
	"accounting-service/internal/pkg/r"
	"accounting-service/internal/service"
	"github.com/gin-gonic/gin"
)

type user struct {
	Nickname  string `form:"nickname" binding:"required"`
	Username  string `form:"username" binding:"required"`
	Password  string `form:"password" binding:"required"`
	Agreement bool   `form:"agreement" binding:"required"`
}

type signIn struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func SignIn(c *gin.Context) {
	var signInParam signIn
	if err := c.Bind(&signInParam); err != nil {
		r.RenderError(c, err)
		return
	}

	userService := service.UserService{Username: signInParam.Username,
		Password: signInParam.Password}

	res, err := userService.SignIn()

	if err != nil {
		logger.Errorf("Sign in error, cause by: %v", err)
		r.RenderError(c, err)
		return
	}

	if res.Ok() {
		r.RenderOk(c, map[string]int64{"userId": res.Data})
	} else {
		r.RenderCode(c, res.Code)
	}
}

func SignUp(c *gin.Context) {
	var userParam user
	if err := c.Bind(&userParam); err != nil {
		r.RenderError(c, err)
		return
	}

	if !userParam.Agreement {
		r.RenderCode(c, consts.USER_DISAGREEMENT)
		return
	}

	userService := service.UserService{Username: userParam.Username}

	exist, err := userService.ExistUsername()
	if err != nil {
		logger.Errorf("Check if username exist, cause by: %v ", err)
		r.RenderError(c, err)
		return
	}

	if exist {
		r.RenderCode(c, consts.USER_NO_USERNAME)
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
	id, err := userService.SignUp(&user)
	if err != nil {
		logger.Errorf("create user e, cause by: %v ", err)
		r.RenderError(c, err)
		return
	}
	r.RenderOk(c, map[string]int64{"userId": id})
}

func ExistUsername(c *gin.Context) {
	userService := service.UserService{Username: c.Query("username")}
	exist, err := userService.ExistUsername()
	if err != nil {
		logger.Errorf("查询用户名是否重复，错误原因: %v", err)
		r.RenderError(c, err)
		return
	}

	r.RenderOk(c, map[string]bool{"exist": exist})
}
