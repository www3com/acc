package api

import (
	"accounting-service/internal/misc/consts"
	"accounting-service/internal/pkg/logger"
	"accounting-service/internal/pkg/r"
	"accounting-service/internal/service"
	"github.com/gin-gonic/gin"
)

func ListAccount(c *gin.Context) {
	//defer func() {
	//	e := recover()
	//	fmt.Println(e)
	//	fmt.Println("释放数据库连接...")
	//}()

	accounts, err := service.ListAccount(ownerId, 7)
	if err != nil {
		logger.Error("Query user account e, user id: {}, details: ", ownerId, err)
		r.RenderCode(c, consts.INTERNAL_ERROR)
		return
	}
	r.RenderOk(c, accounts)
}
