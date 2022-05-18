package api

import (
	"accounting-service/internal/model"
	"accounting-service/internal/service"
	"accounting-service/pkg/logger"
	"accounting-service/pkg/ret"
	"github.com/gin-gonic/gin"
)

type user struct {
	Nickname  string `form:"nickname" binding:"required"`
	Username  string `form:"username" binding:"required"`
	Email     string `form:"email" binding:"required"`
	Password  string `form:"password" binding:"required"`
	Agreement bool   `form:"agreement" binding:"required"`
	Captcha   int    `form:"captcha" binding:"required"`
}

type signIn struct {
	Type     string `form:"type" binding:"required"`
	LoginId  string `form:"loginId" binding:"required"`
	Password string `form:"password"`
	Captcha  int    `form:"captcha"`
}

const userName = "username"

func SignIn(c *gin.Context) {
	var signInParam signIn
	if err := c.ShouldBind(&signInParam); err != nil {
		ret.RenderError(c, err)
		return
	}

	userService := service.UserService{Username: signInParam.LoginId,
		Email:    signInParam.LoginId,
		Password: signInParam.Password,
		Captcha:  signInParam.Captcha}

	var userId int64
	var err error

	if signInParam.Type == userName {
		userId, err = userService.SignInByUsername()
	} else {
		userId, err = userService.SignInByEmail()
	}

	if err != nil {
		logger.Errorf("Sign in, type: %s, cause by: %v", signInParam.Type, err)
		ret.RenderError(c, err)
		return
	}

	if userId != 0 {
		ret.RenderOk(c, map[string]int64{"userId": userId})
	} else {
		ret.RenderCode(c, ret.USER_SIGN_IN_ERROR)
	}
}

func SignUp(c *gin.Context) {
	var userParam user
	if err := c.Bind(&userParam); err != nil {
		ret.RenderError(c, err)
		return
	}

	if !userParam.Agreement {
		ret.RenderCode(c, ret.USER_DISAGREEMENT)
		return
	}

	userService := service.UserService{Username: userParam.Username, Email: userParam.Email}

	exist, err := userService.ExistUsername()
	if err != nil {
		logger.Errorf("Check if username exist, cause by: %v ", err)
		ret.RenderError(c, err)
		return
	}

	if exist {
		ret.RenderCode(c, ret.USER_NO_USERNAME)
		return
	}

	exist, err = userService.ExistEmail()
	if err != nil {
		logger.Errorf("Check if email exist, cause by: %v ", err)
		ret.RenderError(c, err)
		return
	}

	if exist {
		ret.RenderCode(c, ret.USER_NO_EMAIL)
		return
	}

	user := model.User{
		Nickname: userParam.Nickname,
		Username: userParam.Username,
		Email:    userParam.Email,
		Password: userParam.Password,
	}
	if userParam.Agreement {
		user.Agreement = 1
	}
	id, err := userService.SignUp(&user)
	if err != nil {
		logger.Errorf("create user e, cause by: %v ", err)
		ret.RenderError(c, err)
		return
	}
	ret.RenderOk(c, map[string]int64{"userId": id})
}

func ExistUsername(c *gin.Context) {
	userService := service.UserService{Username: c.Query("username")}
	exist, err := userService.ExistUsername()
	if err != nil {
		logger.Errorf("查询用户名是否重复，错误原因: %v", err)
		ret.RenderError(c, err)
		return
	}

	ret.RenderOk(c, map[string]bool{"exist": exist})
}

func ExistEmail(c *gin.Context) {
	userService := service.UserService{Email: c.Query("email")}
	exist, err := userService.ExistEmail()
	if err != nil {
		logger.Errorf("查询电子邮箱是否重复，错误原因: %v", err)
		ret.RenderError(c, err)
		return
	}

	ret.RenderOk(c, map[string]bool{"exist": exist})
}
