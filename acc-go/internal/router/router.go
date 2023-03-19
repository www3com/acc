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

	engine.GET("api/sign-in", api.SignIn)  // 登录
	engine.POST("api/sign-up", api.SignUp) // 注册

	root := engine.Group("/api")
	root.Use(middleware.Auth())

	root.GET("/ledger", api.ListLedger) // 查询账本

	root.GET("/accounts", api.ListAccounts) // 查询账户列表

	//accountApi := api.NewAccountApi()
	//root.GET("/account", accountApi.List)
	//root.POST("/account", accountApi.Create)

	// ------------------------------------------------------

	//root.GET("/account/expenses", accountApi.ListExpenses)
	//
	//// 查询账本
	//root.GET("/ledger", api.ListLedger)
	//// 创建账本
	//root.POST("/ledger", api.CreateLedger)
	//// 更新账本
	//root.PUT("/ledger", api.UpdateLedger)
	//// 删除账本
	//root.DELETE("/ledger", api.DeleteLedger)
	//
	//root.POST("/transaction", api.CreateTransaction)

	return engine
}
