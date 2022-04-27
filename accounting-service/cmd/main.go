package main

import (
	"accounting-service/internal/router"
	"accounting-service/pkg/db"
	"accounting-service/pkg/i18n"
	"accounting-service/pkg/logger"
	"accounting-service/pkg/setting"
	"accounting-service/pkg/translator"
	"fmt"
	"net/http"
	"time"
)

func init() {
	setting.Setup()
	logger.Setup()
	db.Setup()
	translator.Setup()
	i18n.Setup()
}

func main() {

	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	server := &http.Server{
		Addr:           endPoint,
		Handler:        routers.InitRouter(),
		ReadTimeout:    setting.ServerSetting.ReadTimeout * time.Second,
		WriteTimeout:   setting.ServerSetting.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	logger.Infof("start http server listening %s", endPoint)

	server.ListenAndServe()
}
