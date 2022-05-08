package api

import (
	"accounting-service/internal/service"
	"accounting-service/pkg/logger"
	"accounting-service/pkg/ret"
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
		ret.RenderCode(c, ret.INTERNAL_ERROR)
		return
	}
	ret.RenderOk(c, accounts)
}
