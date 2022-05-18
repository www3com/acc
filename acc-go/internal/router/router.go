package routers

import (
	"accounting-service/internal/api"
	"accounting-service/pkg/logger"
	"accounting-service/pkg/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(setting.ServerSetting.RunMode)
	gin.DefaultErrorWriter = logger.GetLog().Out
	r := gin.New()
	//r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	//r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	//r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	root := r.Group("/api")
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

	user := root.Group("/user")
	user.GET("/username", api.ExistUsername)
	user.GET("/email", api.ExistEmail)
	user.POST("/sign-up", api.SignUp)
	return r
}
