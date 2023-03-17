package routers

import (
	"acc/internal/api"
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

	// 注册用户和登录
	engine.GET("api/sign-up", api.SignUp)

	//root := engine.Group("/api")
	//root.Use(middleware.Auth())

	//user := engine.Group("/api/user")
	//user.GET("/sign-in", api)
	//user.POST("/sign-up", userApi.SignUp)
	//
	//root := engine.Group("/api")
	//root.Use(middleware.Auth())
	//
	//ledgerApi := api.NewLedgerApi()
	//// 查询账本
	//root.GET("/ledger", ledgerApi.List)
	//
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
