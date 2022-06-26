package routers

import (
	"acc/internal/api"
	"acc/internal/middleware"
	"acc/internal/pkg/e"
	"acc/internal/pkg/r"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode("release")
	engine := gin.New()
	engine.Use(middleware.Recovery())
	engine.NoRoute(func(c *gin.Context) {
		r.RenderFail(c, e.NotFoundError)
	})

	//r.StaticFS("/web", http.Dir("./"))

	user := engine.Group("/api/user")
	user.POST("/sign-in", api.SignIn)
	user.POST("/sign-up", api.SignUp)

	root := engine.Group("/api")
	root.Use(middleware.Auth())

	// 查询账本
	root.GET("/ledger", api.ListLedger)
	// 创建账本
	root.POST("/ledger", api.CreateLedger)
	// 更新账本
	root.PUT("/ledger", api.UpdateLedger)
	// 删除账本
	root.DELETE("/ledger", api.DeleteLedger)

	root.GET("/account", api.ListAccount)

	root.POST("/transaction", api.CreateTransaction)

	return engine
}
