package routers

import (
	"acc/internal/api"
	"acc/internal/middleware"
	"acc/internal/pkg/logger"
	"acc/internal/pkg/setting"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	gin.SetMode(setting.ConfigSetting.Server.RunMode)
	gin.DefaultErrorWriter = logger.GetLog().Out
	r := gin.New()
	//r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/web", http.Dir("./"))

	root := r.Group("/api")
	root.Use(middleware.Auth())
	ledger := root.Group("/ledger")
	// 查询账本
	ledger.GET("/", api.ListLedger)
	// 创建账本
	ledger.POST("/", api.CreateLedger)
	// 更新账本
	ledger.PUT("/", api.UpdateLedger)
	// 删除账本
	ledger.DELETE("/", api.DeleteLedger)

	account := root.Group("/account")
	account.GET("/", api.ListAccount)

	transaction := root.Group("/transaction")
	transaction.POST("/transaction", api.CreateTransaction)

	user := r.Group("/api/user")
	user.GET("/username", api.ExistUsername)
	user.POST("/sign-in", api.SignIn)
	user.POST("/sign-up", api.SignUp)
	return r
}
