package routers

import (
	"acc/internal/api"
	"acc/internal/middleware"
	"acc/internal/pkg/r"
	"github.com/gin-gonic/gin"
	"github.com/upbos/go-saber/e"
	"net/http"
)

func InitRouter() *gin.Engine {
	gin.SetMode("release")
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.NoRoute(func(c *gin.Context) {
		r.RenderFail(c, e.ErrNotFound)
	})

	// 静态资源
	engine.StaticFS("/web", http.Dir("./web"))

	// 注册和登录
	engine.GET("api/sign-in", api.SignIn)  // 登录
	engine.POST("api/sign-up", api.SignUp) // 注册

	// 权限校验
	root := engine.Group("/api")
	root.Use(middleware.Auth())

	// 账本
	root.GET("/ledger", api.ListLedger) // 查询账本

	// 账户
	root.GET("/accounts", api.ListAccounts)         // 查询账户列表
	root.POST("/account", api.SaveAccount)          // 创建账户或者更新账户
	root.DELETE("/account", api.DeleteAccount)      // 删除账户
	root.PUT("/account/name", api.UpdateName)       // 调整账户名称
	root.PUT("/account/remark", api.UpdateRemark)   // 调整账户描述
	root.PUT("/account/balance", api.UpdateBalance) // 调整账户余额

	return engine
}
