package routers

import (
	"acc/internal/api"
	"acc/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode("release")
	r := gin.New()

	r.Use(middleware.Recovery())

	//r.StaticFS("/web", http.Dir("./"))

	root := r.Group("/api")
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

	user := r.Group("/api/user")
	user.GET("/username", api.ExistUsername)
	user.POST("/sign-in", api.SignIn)
	user.POST("/sign-up", api.SignUp)
	return r
}
