package api

import (
	"acc/internal/misc/consts"
	"acc/internal/model"
	"acc/internal/pkg/e"
	"acc/internal/pkg/logger"
	"acc/internal/pkg/r"
	"acc/internal/service"
	"github.com/allegro/bigcache/v3"
	"github.com/gin-gonic/gin"
	"time"
)

var cache *bigcache.BigCache

func init() {
	cache, _ = bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
}

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
	if code := r.BindAndValid(c, &signInParam); code != e.OK {
		r.RenderFail(c, code)
		return
	}

	userService := service.UserService{
		Username: signInParam.Username,
		Password: signInParam.Password}

	m := userService.SignIn()
	if m.Ok() {
		r.RenderOk(c, m.Data)
	} else {
		r.RenderFail(c, m.Code)
	}
}

func SignUp(c *gin.Context) {
	var userParam user

	if code := r.BindAndValid(c, &userParam); e.IsFail(code) {
		r.RenderFail(c, code)
		return
	}

	if !userParam.Agreement {
		r.RenderFail(c, e.USER_DISAGREEMENT)
		return
	}

	userService := service.UserService{Username: userParam.Username}

	exist, err := userService.ExistUsername()
	if err != nil {
		logger.Errorf("Check if username exist, cause by: %v ", err)
		r.RenderFail(c, e.ERROR)
		return
	}

	if exist {
		r.RenderFail(c, e.USER_NO_USERNAME)
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
		r.RenderFail(c, e.ERROR)
		return
	}
	r.RenderOk(c, map[string]int64{"jwt": id})
}

func ExistUsername(c *gin.Context) {
	userService := service.UserService{Username: c.Query("username")}
	exist, err := userService.ExistUsername()
	if err != nil {
		logger.Errorf("查询用户名是否重复，错误原因: %v", err)
		r.RenderFail(c, e.ERROR)
		return
	}

	r.RenderOk(c, map[string]bool{"exist": exist})
}
