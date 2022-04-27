package routers

import (
	"accounting-service/internal/api"
	"accounting-service/pkg/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(setting.ServerSetting.RunMode)

	r := gin.New()
	//r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	//r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	//r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	//r.GET("/test", controller)

	//tpl := r.Group("/t")
	//tpl.GET("/ledgers", template.ListTemplateLedger)

	ledger := r.Group("/")
	// 查询账本
	ledger.GET("/ledger", api.ListLedger)
	// 创建账本
	ledger.POST("/ledger", api.CreateLedger)
	// 更新账本
	ledger.PUT("/ledger", api.UpdateLedger)
	// 删除账本
	ledger.DELETE("/ledger", api.DeleteLedger)
	return r
}
